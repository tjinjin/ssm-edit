ARCH := amd64
OS := linux darwin

VERSION=$(shell ./version.sh)

VERSION_FLAG=\"github.com/tjinjin/ssm-edit/cmd.version=$(shell ./version.sh)\"

build-all:
	gox -os="$(OS)" -arch="$(ARCH)" -ldflags "-X $(VERSION_FLAG)" -output="build/{{.OS}}_{{.Arch}}/{{.Dir}}"

build:
	go build -ldflags "-X $(VERSION_FLAG)" -o ./dist/ssm-edit
