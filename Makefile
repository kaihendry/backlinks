INFILES = $(shell find . -name "*.mdwn")
OUTFILES = $(INFILES:.mdwn=.html)
LINKFILES = $(INFILES:.mdwn=.links)

all: $(OUTFILES)

%.links: %.mdwn
	./links $< > $@

%.html: %.mdwn $(LINKFILES)
	cmark $< > $@
	./backlink $< | cmark >> $@

clean:
	rm -fv $(OUTFILES) $(LINKFILES)

PHONY: all clean
