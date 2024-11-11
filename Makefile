PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))

define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go install $(2) ;\
}
endef

MOCKERY := $(PWD)/bin/mockery
$(MOCKERY):
	$(call go-get-tool,$(MOCKERY),github.com/vektra/mockery/v2)

.PHONY: generate
generate: $(MOCKERY)
	PATH=$(PWD)/bin:$$PATH go generate ./...

test: generate
	go test -v ./p
