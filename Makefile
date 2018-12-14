INFILES = $(shell find . -name "*.mdwn")
OUTFILES = $(INFILES:.mdwn=.html)

all: $(OUTFILES)

%.html: %.mdwn
	cmark $< >> $@
	./backlinks $< | cmark >> $@

clean:
	rm -f $(OUTFILES)

PHONY: all clean
