INFILES     := $(shell find . -name "*.mdwn")
OUTFILES    := $(INFILES:.mdwn=.html)
LINKFILES   := $(INFILES:.mdwn=.whatlinkshere)
LINKPATTERN := $(INFILES:.mdwn=.w%e)

.PHONY: all clean
.PRECIOUS: $(LINKFILES)

all: $(OUTFILES)

# These need to be all made before the HTML is processed
$(LINKPATTERN): $(INFILES)
	@echo Creating backlinks
	@rm -f $(LINKFILES)
	@touch $(LINKFILES)
	@for m in $^; do go run backlinks.go $$m; done

%.html: %.mdwn %.whatlinkshere
	@echo Deps $^
	@cmark $^ > $@

clean:
	rm -f $(LINKFILES) $(OUTFILES)
