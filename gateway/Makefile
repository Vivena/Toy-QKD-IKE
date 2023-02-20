ALPINE_VER 		?= 	3.15
GO_VER 			= 	1.17.8

BASE_VERSION 	=	O.1.0
EXTRA_VERSION 	?= 	$(shell git rev-parse --short HEAD)
PROJECT_VERSION	=	$(BASE_VERSION)-snapshot-$(EXTRA_VERSION)
ARCH			=	$(shell go env GOARCH)

GO_TAGS 		?=
PKGNAME 		= 	github.com/Vivena/Toy-QKD-IKE/gateway
BUILD_DIR 		?= 	build

RELEASE_IMAGES 	= 	server
EXES_LIST 		= 	server 

pkgmap.server 	:= 	$(PKGNAME)/cmd/server

.DEFAULT_GOAL 	:= 	all

include gotools.mk

.PHONY: all
all: native

# -------------------------------------------------------

.PHONY: gotools
gotools: gotools-install

.PHONY: native
native: $(EXES_LIST)

.PHONY: $(EXES_LIST)
$(EXES_LIST): %: $(BUILD_DIR)/bin/%

$(BUILD_DIR)/bin/%:
	@echo "Building $@"
	@mkdir -p $(@D)
	GOBIN=$(abspath $(@D)) go install $(pkgmap.$(@F))
	@touch $@

.PHONY: tools
tools: $(TOOLS_EXES)

# -------------------------------------------------------


.PHONY: unit-test
unit-test: unit-test-clean 
	./scripts/run-unit-tests.sh


# -------------------------------------------------------

.PHONY: clean
clean: release-clean
	-@rm -rf $(BUILD_DIR)

.PHONY: release-clean
release-clean: $(RELEASE_PLATFORMS:%=%-release-clean)
%-release-clean:
	-@rm -rf release/$*
