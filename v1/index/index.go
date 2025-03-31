// manifest provides utilities for managing Open Container Initiative image manifests
// described here: https://github.com/opencontainers/image-spec/blob/main/image-index.md and
// here: https://pkg.go.dev/github.com/opencontainers/image-spec@v1.1.1/specs-go/v1#Index .
package index

import (
	"os"

	"github.com/grokify/gocharts/v2/data/table"
	"github.com/grokify/goopencontainers/v1/manifest"
	"github.com/grokify/mogo/encoding/jsonutil"
	specs "github.com/opencontainers/image-spec/specs-go"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

func NewImageIndexFromTable(t *table.Table) (v1.Index, error) {
	index := v1.Index{
		Versioned: specs.Versioned{
			SchemaVersion: 2,
		},
		MediaType: v1.MediaTypeImageIndex,
		Manifests: []v1.Descriptor{},
	}

	for _, r := range t.Rows {
		if m, err := manifest.NewDescriptorFromTableRow(t.Columns, r); err != nil {
			return index, err
		} else {
			index.Manifests = append(index.Manifests, m)
		}
	}
	return index, nil
}

func WriteJSON(filename string, idx v1.Index, prefix, indent string, perm os.FileMode) error {
	if b, err := jsonutil.MarshalSimple(idx, prefix, indent); err != nil {
		return err
	} else {
		return os.WriteFile(filename, b, perm)
	}
}
