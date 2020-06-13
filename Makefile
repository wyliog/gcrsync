BUILD_VERSION   := $(version)
BUILD_TIME      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD)

all:
	gox -osarch="darwin/amd64 linux/386 linux/amd64" \
        -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}"

release: all
	ghr -u wyliog -t $(GITHUB_RELEASE_TOKEN) -replace -recreate --debug $(version) dist

clean:
	rm -rf dist

install:
	go install 

.PHONY : all release clean install

.EXPORT_ALL_VARIABLES:

GO111MODULE = on
