# change
Easy tracking of value changes in Go programs – ranging from simple to full-featured GoF observable

[![Go Reference](https://pkg.go.dev/badge/github.com/fractalqb/change.svg)](https://pkg.go.dev/github.com/fractalqb/change)
[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](http://www.gnu.org/licenses/agpl-3.0)
![Go Action](https://github.com/fractalqb/change/actions/workflows/go.yml/badge.svg)

```go
import "github.com/fractalqb/change"
```

## A Story of `change`
This project was started for a very specific need: An event-driven
program that collects events to to build and constantly refine the
data of a database. When this was working well, I wanted a Web UI to
have an easy look at some specific queries about the data. – No need
to mention that the data has a reasonably rich domain model.

Then came the crucial step: I wanted—tired of pressing F5—the Web UI
to update with the incoming events. That was the moment when I was
looking for some observer-like solutions to put into my program. But
as a hobby-project that runs on rather small hardware, I also wanted
it to stay modest in its use of resources. No need to fire and
instantaneously process events on each and every updated value. It
would be sufficient to collect the information which entities
*changed* during the processing of an event, and then only send the
entities of the current Web view via WebSocket. – The thing that now
is the `chgv` (“changing values”) sub-package was born. It has simply
no memory overhead for the change-detectable values and, as you can
see from the benchmarks, minimal time overhead compared to Go's
builtin types. But the story is not yet complete…

The DB and its data grew, DB indexes were optimised
read-caches[<sup>1</sup>](#fn1) took away read-load from the DB server
and change detection also reduced DB writes. The program ran nearly
imperceptible along others on the small machine. In theory the world
could have been fine!  … If it weren't for this pathological core
entity in the domain model, which is important for processing almost
all events.

- This entity has a permanent external key which, unfortunately, is
  not part of all events. But the entity has to be created, event if
  that external key is not yet known.

- No problem, I use an internal surrogate key. But how are subsequent
  events correlated to new instances. Luckily…
  
- …the entity also has a `name` that is almost unique. Far more than
  99.9% of all instances will have a unique name. A name which may
  change over time in very rare cases. – Not exactly a dream if you
  aim for deterministic processing of the data.
  
Don't get me wrong! I do not complain. This entity makes one of the
more interesting challenges in the business logic of the project. But
runtime statistics show that all three keys, the external key, the
surrogate key and the name, are similarly important for lookup during
event processing.

| Lookup by    | Count    | Hit Rate |
|--------------|----------|----------|
| External ID  |    62436 |    33.9% |
| Name         |    55761 |    29.7% |
| Surrogate ID |    20114 |    65.4% |

Now, combine that with a cache where you can lookup the same memory
object by either external key, surrogate key or name!  It will turn
out that one can make that work if you invalidate the by-name-lookup
*each time* the name of an instance *changes*. This goes beyond the
capabilities of the `chgv` package. And—I think I don't have to tell
you—spreading the code with `if name changed then invalidate
by-name-lookup` is a no-go.

OK, back to observer libraries. Many things I found for Go told me in
their docs, that Go `chan` is an important part of their concept. But
my problem is not about concurrency but only about [GoF
observer](https://en.wikipedia.org/wiki/Observer_pattern)[<sup>2</sup>](#fn2). Invalidating
the cache is fast enough, no need to do that concurrently. What I
wanted, was something similar to JavaFX Properties and *I* didn't find
it. So, rolled my own… Why not? 'Twas for my own fun.

But still I couldn't resist to think about efficiency. A value has to
know all its observers. Even when there is no observer, this will
introduce a constant amount of extra bytes to any single value.

- A slice of observers will add 3 machine words. E.g. on a 64-bit
  machine this would blow up all register types, 1–8 byte, to the size
  of 32 byte. A `string` would end up with 40 byte.
  
- A function pointer would add 1 extra machine word. Again, on a
  64-bit machine this would blow up register types to 16 byte and
  `string` to 24 byte. That's better!
  
The problem with function pointers is that you cannot transparently
resister/unregister several functions as observers. One only can
set/unset/reset their pointers. So, which way to go?  …  Both!  'Twas
for my own fun, you remember? The result is:

- `onchg`: Implements the function pointer idea. The name resembles
  the call to the hook function “on change”.

- `obsv`: Implements the observer pattern, still with only 1 machine
  word per-value memory overhead when not used as observable. The name
  is short for “observable value”.
  
For my original project it is still sufficient to use the
`onchg.String` for the `name` attribute of the beloved pathological
entity.

## Benchmark
As the benchmark shows there is a significant runtime penalty for
having the luxury of a full observable implementation. The benchmark
measures the time to access an `int` value and then to set it again –
either with change detection or without it, i.e. using the value as
nothing more than being a value.

Noramlized against the reference of a bare Go `int` we have the
factors:

| Package | Benchmark                   | ~Factor |
|---------|-----------------------------|---------|
| chgv    | BenchmarkIntReference-4     |       1 |
| chgv    | BenchmarkInt_noDetect-4     |       6 |
| chgv    | BenchmarkInt_withDetect-4   |       6 |
| onchg   | BenchmarkInt_noDetect-4     |       8 |
| onchg   | BenchmarkInt_withDetect-4   |      23 |
| obsv    | BenchmarkInt_noDetect-4     |      10 |
| obsv    | BenchmarkInt_withDetect-4   |     381 |
| obsv    | BenchmarkInt_withObserver-4 |     441 |


**Here's the benchmark output:**

```
goos: linux
goarch: amd64
pkg: github.com/fractalqb/change/chgv
cpu: Intel(R) Core(TM) i7-5500U CPU @ 2.40GHz
BenchmarkIntReference-4     	1000000000	         0.3374 ns/op
BenchmarkInt_noDetect-4     	588945325	         2.023 ns/op
BenchmarkInt_withDetect-4   	587903571	         2.020 ns/op
PASS
ok  	github.com/fractalqb/change/chgv	3.197s
goos: linux
goarch: amd64
pkg: github.com/fractalqb/change/onchg
cpu: Intel(R) Core(TM) i7-5500U CPU @ 2.40GHz
BenchmarkInt_noDetect-4     	439199455	         2.703 ns/op
BenchmarkInt_withDetect-4   	156135638	         7.592 ns/op
PASS
ok  	github.com/fractalqb/change/onchg	3.431s
goos: linux
goarch: amd64
pkg: github.com/fractalqb/change/obsv
cpu: Intel(R) Core(TM) i7-5500U CPU @ 2.40GHz
BenchmarkInt_noDetect-4       	341003700	         3.374 ns/op
BenchmarkInt_withDetect-4     	 8994292	       128.7 ns/op
BenchmarkInt_withObserver-4   	 8412120	       148.7 ns/op
PASS
ok  	github.com/fractalqb/change/obsv	4.412s
```

## Footnotes
1. <span id="fn1">Yes, that caches kill horizontal scalability. But,
   remember, this thing has to run with minimal resources, not on an
   elastic cluster.</span>

2. <span id="fn2">Having a good GoF Observer lib would make it easy to
   create wrappers around the receiving or the sending end of a `chan`
   to make them Observers or Observable. The other way around is not
   that simple. To me, a clear case of [abstraction
   inversion](https://en.wikipedia.org/wiki/Abstraction_inversion)</span>.
