# build vars
TAG_NAME := $(shell test -d .git && git describe --abbrev=0 --tags)
SHA := $(shell test -d .git && git rev-parse --short HEAD)
COMMIT := $(SHA)
# hide commit for releases
ifeq ($(RELEASE),1)
    COMMIT :=
endif
VERSION := $(if $(TAG_NAME),$(TAG_NAME),$(SHA))
BUILD_DATE := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
BUILD_TAGS := -tags=release
LD_FLAGS := -X github.com/thommyho/robotui/server.Version=$(VERSION) -X github.com/thommyho/robotui/server.Commit=$(COMMIT) -s -w
BUILD_ARGS := -ldflags='$(LD_FLAGS)'

# docker
DOCKER_IMAGE := robotui/robotui
PLATFORM := linux/amd64,linux/arm64,linux/arm/v6

# gokrazy image
IMAGE_FILE := robotui_$(TAG_NAME).image
IMAGE_OPTIONS := -hostname robotui -http_port 8080 github.com/gokrazy/serial-busybox github.com/gokrazy/breakglass github.com/gokrazy/mkfs github.com/gokrazy/wifi github.com/thommyho/robotui

# deb
PACKAGES = ./release

# asn1-patch
GOROOT := $(shell go env GOROOT)
CURRDIR := $(shell pwd)

default:: ui build

all:: clean install install-ui ui assets lint test-ui lint-ui test build

clean::
	rm -rf dist/

install::
	go install $$(go list -f '{{join .Imports " "}}' tools.go)

install-ui::
	npm ci

ui::
	npm run build

assets::
	go generate ./...

docs::
	go generate github.com/thommyho/robotui/util/templates/...

lint::
	golangci-lint run

lint-ui::
	npm run lint

test-ui::
	npm test

toml:
	go run packaging/toml.go

test::
	@echo "Running testsuite"
	CGO_ENABLED=0 go test $(BUILD_TAGS),test ./...

porcelain::
	gofmt -w -l $$(find . -name '*.go')
	go mod tidy
	test -z "$$(git status --porcelain)" || (git status; git diff; false)

build::
	@echo Version: $(VERSION) $(SHA) $(BUILD_DATE)
	CGO_ENABLED=0 go build -v $(BUILD_TAGS) $(BUILD_ARGS)

snapshot:
	goreleaser --snapshot --skip-publish --clean

release::
	goreleaser --clean

docker::
	@echo Version: $(VERSION) $(SHA) $(BUILD_DATE)
	docker buildx build --platform $(PLATFORM) --tag $(DOCKER_IMAGE):testing .

publish-testing::
	@echo Version: $(VERSION) $(SHA) $(BUILD_DATE)
	docker buildx build --platform $(PLATFORM) --tag $(DOCKER_IMAGE):testing --push .

publish-nightly::
	@echo Version: $(VERSION) $(SHA) $(BUILD_DATE)
	docker buildx build --platform $(PLATFORM) --tag $(DOCKER_IMAGE):nightly --push .

publish-release::
	@echo Version: $(VERSION) $(SHA) $(BUILD_DATE)
	docker buildx build --build-arg RELEASE=1 --platform $(PLATFORM) --tag $(DOCKER_IMAGE):latest --tag $(DOCKER_IMAGE):$(VERSION) --push .

apt-nightly::
	$(foreach file, $(wildcard $(PACKAGES)/*.deb), \
		cloudsmith push deb robotui/unstable/any-distro/any-version $(file); \
	)

apt-release::
	$(foreach file, $(wildcard $(PACKAGES)/*.deb), \
		cloudsmith push deb robotui/stable/any-distro/any-version $(file); \
	)

# gokrazy image
gokrazy::
	go install github.com/gokrazy/tools/cmd/gokr-packer@main
	mkdir -p flags/github.com/gokrazy/breakglass
	echo "-forward=private-network" > flags/github.com/gokrazy/breakglass/flags.txt
	mkdir -p env/github.com/thommyho/robotui
	echo "robotui_NETWORK_PORT=80" > env/github.com/thommyho/robotui/env.txt
	echo "robotui_DATABASE_DSN=/perm/robotui.db" >> env/github.com/thommyho/robotui/env.txt
	mkdir -p buildflags/github.com/thommyho/robotui
	echo "$(BUILD_TAGS),gokrazy" > buildflags/github.com/thommyho/robotui/buildflags.txt
	echo "-ldflags=$(LD_FLAGS)" >> buildflags/github.com/thommyho/robotui/buildflags.txt
	gokr-packer -hostname robotui -http_port 8080 -overwrite=$(IMAGE_FILE) -target_storage_bytes=1258299392 $(IMAGE_OPTIONS)
	# gzip -f $(IMAGE_FILE)

gokrazy-run::
	MACHINE=arm64 IMAGE_FILE=$(IMAGE_FILE) ./packaging/gokrazy/run.sh

gokrazy-update::
	gokr-packer -update yes $(IMAGE_OPTIONS)

soc::
	@echo Version: $(VERSION) $(SHA) $(BUILD_DATE)
	go build $(BUILD_TAGS) $(BUILD_ARGS) github.com/thommyho/robotui/cmd/soc

# patch asn1.go to allow Elli buggy certificates to be accepted with EEBUS
patch-asn1-sudo::
	# echo $(GOROOT)
	cat $(GOROOT)/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go | grep -C 1 "out = true"
	sudo patch -N -t -d $(GOROOT)/src/vendor/golang.org/x/crypto/cryptobyte -i $(CURRDIR)/patch/asn1.diff
	cat $(GOROOT)/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go | grep -C 1 "out = true"

patch-asn1::
	# echo $(GOROOT)
	cat $(GOROOT)/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go | grep -C 1 "out = true"
	patch -N -t -d $(GOROOT)/src/vendor/golang.org/x/crypto/cryptobyte -i $(CURRDIR)/patch/asn1.diff
	cat $(GOROOT)/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go | grep -C 1 "out = true"
