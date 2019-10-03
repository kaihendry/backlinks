INFILES     := $(shell find . -name "*.mdwn")
OUTFILES    := $(INFILES:.mdwn=.html)
LINKFILES   := $(INFILES:.mdwn=.bl)

.PHONY: all clean
.PRECIOUS: $(LINKFILES)

all: $(OUTFILES)

# These need to be all made before the HTML is processed
%.bl: $(INFILES)
	@echo Creating backlinks
	@touch $(LINKFILES)
	@for m in $^; do go run backlinks.go $$m; done

%.html: %.mdwn %.bl
	@echo Deps $^
	@cmark $^ > $@

test:

clean:
	rm -f *.bl *.html
