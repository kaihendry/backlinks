INFILES     := $(shell find . -name "*.mdwn")
OUTFILES    := $(INFILES:.mdwn=.html)
LINKFILES   := $(INFILES:.mdwn=.bl)

.PHONY: all clean test
.PRECIOUS: $(LINKFILES)

all: $(OUTFILES)

# These need to be all made before the HTML is processed
%.bl: $(INFILES)
	@echo Creating backlinks
	@touch $(LINKFILES)
	@for m in $^; do go run backlinks.go $$m; done

%.html: %.mdwn %.bl
	@echo First $(firstword $^)
	@echo Last $(lastword $^)
	@cmark $(firstword $^) > $@
	@echo "<h1>Backlinks</h1>" >> $@
	@sort -u < $(lastword $^) | cmark >> $@

test:
	@for i in *.html; do diff -u $$i test/$$i; done

clean:
	rm -f *.bl *.html
