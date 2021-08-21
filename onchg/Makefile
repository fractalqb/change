.PRECIOUS: %.go-m4

all: builtin.go multihook.go

%.go: %.go-m4 %.go-m4~
	diff3 -m $@ $*.go-m4 $*.go-m4~ > $@~
	mv $@~ $@
	cp $*.go-m4~ $*.go-m4
	gofmt -w -s $@

%.go-m4~: %.m4
	m4 $< > $@
