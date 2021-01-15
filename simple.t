#!/bin/sh

test_description="backlinks"

. $SHARNESS_TEST_SRCDIR/sharness.sh

test_expect_success "pageA is rendered" "
	redo clean &&
    redo all &&
    diff test/pageA.html pageA.html &&
    diff test/pageB.html pageB.html &&
    diff test/pageC.html pageC.html
"

test_done
