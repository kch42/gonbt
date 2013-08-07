package nbt

import (
	"compress/gzip"
	"compress/zlib"
	"io"
)

// Some helpers for reading / writing compressed NBT data, since NBT data is often compressed.

func ReadGzipdNamedTag(r io.Reader) (Tag, string, error) {
	decomp, err := gzip.NewReader(r)
	if err != nil {
		return Tag{}, "", err
	}

	return ReadNamedTag(decomp)
}

func WriteGzipdNamedTag(w io.Writer, name string, tag Tag) error {
	comp := gzip.NewWriter(w)
	return WriteNamedTag(comp, name, tag)
}

func ReadZlibdNamedTag(r io.Reader) (Tag, string, error) {
	decomp, err := zlib.NewReader(r)
	if err != nil {
		return Tag{}, "", err
	}

	return ReadNamedTag(decomp)
}

func WriteZlibdNamedTag(w io.Writer, name string, tag Tag) error {
	comp := zlib.NewWriter(w)
	return WriteNamedTag(comp, name, tag)
}
