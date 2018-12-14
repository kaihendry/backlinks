INFILES = $(shell find . -name "*.mdwn")
OUTFILES = $(INFILES:.mdwn=.html)
LINKFILES = $(INFILES:.mdwn=.whatlinkshere)

all: $(OUTFILES)

# These need to be all made before the HTML is processed
$(LINKFILES): $(INFILES)
	@echo Creating backlinks $@
	@touch $@
	@go run backlinks.go $<

%.html: %.mdwn %.whatlinkshere
	@echo Deps $^
	@cmark $^ > $@

clean:
	rm -fv $(OUTFILES) $(LINKFILES)

PHONY: all clean
