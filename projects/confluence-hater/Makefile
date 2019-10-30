PKGPATH=github.com/lovemew67/go-misc/projects/confluence-hater/command
SOURCE=./...
DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
REVISION=$(shell git rev-list -1 HEAD)
TAG=$(shell git tag -l --points-at HEAD)
ifeq ($(TAG),)
TAG=$(REVISION)
endif

test:
	@echo "start unit tests & vet"
	go vet $(SOURCE)
	go test -mod=vendor -race -cover $(SOURCE)

build:
	GO111MODULE=on go install -mod=vendor -i -v -ldflags "-s -X $(PKGPATH).gitCommit=$(REVISION) -X $(PKGPATH).appVersion=$(TAG) -X $(PKGPATH).buildDate=$(DATE)" $(SOURCE)

vendor:
	rm -f go.sum
	GO111MODULE=on go build -v $(SOURCE) 
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor
