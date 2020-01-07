ARCH := amd64
OS := linux darwin

VERSION=$(shell ./version.sh)

VERSION_FLAG=\"github.com/tjinjin/ssm-edit/cmd.version=$(shell ./version.sh)\"

clean:
	rm -rf build/*

build-all: clean
	gox -os="$(OS)" -arch="$(ARCH)" -ldflags "-X $(VERSION_FLAG)" -output="build/{{.OS}}_{{.Arch}}/{{.Dir}}"
	cd build/ && find . -name "$(VERSION)"

build:
	go build -ldflags "-X $(VERSION_FLAG)" -o ./dist/ssm-edit

package: build-all
	cd build/darwin_amd64 && tar czf ssm-edit_$(VERSION)_darwin_amd64.tar.gz ssm-edit
	cd build/linux_amd64 && tar czf ssm-edit_$(VERSION)_linux_amd64.tar.gz ssm-edit
	find build -type f -name *.gz -exec mv {} pkg \;

release: package
	ghr -u tjinjin -r ssm-edit -n "$(VERSION)" $(VERSION) pkg/
