TOPDIR=$(shell cd "`dirname $(lastword $(MAKEFILE_LIST))`" && pwd)
PKGDIR=$(TOPDIR)/pkg

PACKAGES=dbi

all: $(PACKAGES)
	for package in $(PACKAGES); do \
		$(MAKE) -C "$(PKGDIR)/$$package"; \
	done

clean:
	for package in $(PACKAGES); do \
		$(MAKE) -C "$(PKGDIR)/$$package" clean; \
	done

.PHONY: all clean

