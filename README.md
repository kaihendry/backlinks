Issue with Makefile is that for example `pageC.whatlinkshere` isn't properly
generated, it's empty.

Workaround is `for i in *.mdwn; do go run backlinks.go $i; done`

# Other issue

If I edit pageC.mdwn and append `[page B](pageB.html)`.

I expect only:

1. pageC.mdwn -> pageC.html - pageC has changed regenerate
2. pageC.mdwn -> pageB.whatlinkshere - pageB.whatlinkshere is updated since there is a new link from pageC
3. pageB.mdwn -> pageB.html - pageB needs be re-rendered

To be run
