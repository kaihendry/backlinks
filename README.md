# Second run of `make`

	backlinks$ make
	rm -rf pageA.backlinks/pageA.backlinks/
	mkdir -p pageA.backlinks/pageA.backlinks/
	cp pageA.backlinks/pageA.mdwn pageA.backlinks/pageA.backlinks/
	cd pageA.backlinks/pageA.backlinks/; go run ../backlinks.go pageA.backlinks/pageA.mdwn
	stat ../backlinks.go: no such file or directory
	make: *** [pageA.backlinks/pageA.backlinks/.done] Error 1
