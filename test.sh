#!/bin/sh

test_description="backlinks"

. ${SHARNESS_TEST_SRCDIR}/sharness.sh

test_expect_success "pageA is rendered" "
    redo pageA.html
    diff test/pageA.html pageA.html
"

