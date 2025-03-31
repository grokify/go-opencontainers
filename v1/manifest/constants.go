// manifest provides utilities for managing Open Container Initiative image manifests
// described here: https://github.com/opencontainers/image-spec/blob/main/manifest.md and
// here: https://pkg.go.dev/github.com/opencontainers/image-spec@v1.1.1/specs-go/v1#Manifest .
package manifest

const (
	FieldDigest       = "digest"
	FieldPlatformOS   = "platform_os"
	FieldPlatformArch = "platform_arch"
	FieldSize         = "size"
	FieldTitle        = "title"
	FieldVersion      = "version"
)
