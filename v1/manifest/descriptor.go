package manifest

import (
	"github.com/grokify/gocharts/v2/data/table"
	digest "github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type DescriptorEdit struct {
	Descriptor v1.Descriptor
}

func NewDescriptor() v1.Descriptor {
	return v1.Descriptor{
		MediaType:   v1.MediaTypeImageManifest,
		Annotations: map[string]string{},
	}
}

func NewDescriptorFromTableRow(cols, row []string) (v1.Descriptor, error) {
	edit := NewDescriptorEdit()
	err := edit.UpsertTableRow(cols, row)
	return edit.Descriptor, err
}

func NewDescriptorEdit() DescriptorEdit {
	return DescriptorEdit{
		Descriptor: NewDescriptor(),
	}
}

func (e *DescriptorEdit) init() {
	if e.Descriptor.Annotations == nil {
		e.Descriptor.Annotations = map[string]string{}
	}
}

func (e *DescriptorEdit) UpsertDigest(v string) {
	if v == "" {
		return
	} else {
		e.Descriptor.Digest = digest.Digest(v)
	}
}

func (e *DescriptorEdit) UpsertSize(v int64) {
	e.Descriptor.Size = v
}

func (e *DescriptorEdit) UpsertTitle(v string) {
	if v == "" {
		return
	} else {
		e.init()
		e.Descriptor.Annotations[v1.AnnotationTitle] = v
	}
}

func (e *DescriptorEdit) UpsertVersion(v string) {
	if v == "" {
		return
	} else {
		e.init()
		e.Descriptor.Annotations[v1.AnnotationVersion] = v
	}
}

func (e *DescriptorEdit) UpsertTableRow(cols, row []string) error {
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
