INFILES   := $(shell find . -name "*.mdwn")
OUTFILES  := $(INFILES:.mdwn=.html)
DONEFILES := $(patsubst %.mdwn,%.backlinks/.done,$(INFILES))

.PHONY: all clean

ifeq ($(STEP),)
all $(OUTFILES): $(DONEFILES)
	$(MAKE) STEP=2 $@

%.backlinks/.done: %.mdwn
	rm -rf $(dir $@)
	mkdir -p $(dir $@)
	cp $< $(dir $@)
	cd $(dir $@); go run ../backlinks.go $<
	touch $@

clean:
	rm -rf *.backlinks $(OUTFILES)
else
all: $(OUTFILES)

.SECONDEXPANSION:
%.html: %.mdwn $$(wildcard *.backlinks/$$*.whatlinkshere)
	@echo Deps $^
	@cmark $^ > $@
endif
