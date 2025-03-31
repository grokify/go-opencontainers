// manifest provides utilities for managing Open Container Initiative image manifests
// described here: https://github.com/opencontainers/image-spec/blob/main/manifest.md and
// here: https://pkg.go.dev/github.com/opencontainers/image-spec@v1.1.1/specs-go/v1#Manifest .
package manifest

import (
	"github.com/grokify/gocharts/v2/data/table"
	digest "github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type ManifestEdit struct {
	Manifest v1.Manifest
}

func NewManifest() v1.Manifest {
	return v1.Manifest{
		MediaType:   v1.MediaTypeImageManifest,
		Annotations: map[string]string{},
		Config:      v1.Descriptor{},
	}
}

func NewManifestFromTableRow(cols, row []string) (v1.Manifest, error) {
	me := NewManifestEdit()
	err := me.UpsertTableRow(cols, row)
	return me.Manifest, err
}

func NewManifestEdit() ManifestEdit {
	return ManifestEdit{
		Manifest: NewManifest(),
	}
}

func (e *ManifestEdit) init() {
	if e.Manifest.Annotations == nil {
		e.Manifest.Annotations = map[string]string{}
	}
}

func (e *ManifestEdit) UpsertDigest(v string) {
	if v == "" {
		return
	} else {
		e.Manifest.Config.Digest = digest.Digest(v)
	}
}

func (e *ManifestEdit) UpsertSize(v int64) {
	e.Manifest.Config.Size = v
}

func (e *ManifestEdit) UpsertTitle(v string) {
	if v == "" {
		return
	} else {
		e.init()
		e.Manifest.Annotations[v1.AnnotationTitle] = v
	}
}

func (e *ManifestEdit) UpsertVersion(v string) {
	if v == "" {
		return
	} else {
		e.init()
		e.Manifest.Annotations[v1.AnnotationVersion] = v
	}
}

func (e *ManifestEdit) UpsertTableRow(cols, row []string) error {
	cls := table.Columns(cols)
	if v := cls.MustCellString(FieldDigest, row); v != "" {
		e.UpsertDigest(v)
	}
	if v := cls.MustCellIntOrDefault(FieldSize, row, 0); v > 0 {
		e.UpsertSize(int64(v))
	}
	if v := cls.MustCellString(FieldTitle, row); v != "" {
		e.UpsertTitle(v)
	}
	if v := cls.MustCellString(FieldVersion, row); v != "" {
		e.UpsertVersion(v)
	}
	return nil
}
