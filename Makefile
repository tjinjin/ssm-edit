ARCH := amd64
OS := linux darwin

VERSION=$(shell ./version.sh)

VERSION_FLAG=\"github.com/tjinjin/ssm-edit/cmd.version=$(shell ./version.sh)\"

build-all:
	gox -os="$(OS)" -arch="$(ARCH)" -ldflags="-w -s -X $(VERSION_FLAG)" -output="assets/{{.OS}}_{{.Arch}}/{{.Dir}}"

build:
	go build -ldflags="-w -s -X $(VERSION_FLAG)" -o ./build/ssm-edit

package:
	./create.sh
