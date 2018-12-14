INFILES = $(shell find . -name "*.mdwn")
OUTFILES = $(INFILES:.mdwn=.html)

all: $(OUTFILES)

%.links: %.mdwn
	./links $< > $@

%.html: %.mdwn %.links
	cmark $< > $@
	./backlink $< | cmark >> $@

clean:
	rm -fv $(OUTFILES)

PHONY: all clean
