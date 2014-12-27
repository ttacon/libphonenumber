package libphonenumber

import (
	"io"
	"os"
)

type MetadataLoader interface {
	LoadMetadata(filename string) io.ReadCloser
}

type metaLoader struct{}

func (m metaLoader) LoadMetadata(filename string) io.ReadCloser {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	return f
}

var DEFAULT_METADATA_LOADER = metaLoader{}
