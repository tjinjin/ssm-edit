build:
	gox -osarch="linux/amd64 linux/arm64 darwin/amd64" -output="build/{{.OS}}_{{.Arch}}/{{.Dir}}"
