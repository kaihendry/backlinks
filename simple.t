#!/bin/bash

test_description="backlinks"

. $SHARNESS_TEST_SRCDIR/sharness.sh

testdir=$(dirname $(pwd))

test_expect_success "clean up" "
cd $testdir
echo Test directory: $testdir
redo clean
"

test_expect_success "pageA rendered" "
cd $testdir
redo clean
redo all
diff test/pageA.html pageA.html
"

test_expect_success "pageB rendered" "
cd $testdir
redo clean
redo all
diff test/pageB.html pageB.html
"

test_expect_success "pageC rendered" "
cd $testdir
redo clean
redo all
diff test/pageC.html pageC.html
"

test_done
