include $(GOROOT)/src/Make.inc

TARG=dbi
GOFILES=$(shell find . -name '*.go' -not -name '*_test.go' -not -name '_testmain.go')

include $(GOROOT)/src/Make.pkg

