#!/bin/sh -e

test_description="backlinks"

. $SHARNESS_TEST_SRCDIR/sharness.sh

test_expect_success "Success is reported like this" "
    echo hello world | grep hello
"

test_expect_success "Just heello" "
    echo hello world
"

test_expect_success "Commands are chained this way" "
    test x = 'x' &&
    test 2 -gt 1 &&
    echo success
"

test_expect_success "clean up" "
    redo clean
"

test_expect_success "pages are rendered" "
redo clean &&
redo all &&
diff test/pageA.html pageA.html &&
diff test/pageB.html pageB.html &&
diff test/pageC.html pageC.html
"

test_done
