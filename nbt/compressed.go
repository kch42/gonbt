package nbt

import (
	"compress/gzip"
	"compress/zlib"
	"io"
)

// Some helpers for reading / writing compressed NBT data, since NBT data is often compressed.

// ReadGzipdNamedTag reads a gzip compressed named tag. See ReadNamedTags for more info.
func ReadGzipdNamedTag(r io.Reader) (Tag, string, error) {
	decomp, err := gzip.NewReader(r)
	if err != nil {
		return Tag{}, "", err
	}

	return ReadNamedTag(decomp)
}

// WriteGzipdNamedTag writes a gzip compressed named tag. See WriteNamedTag for more info.
func WriteGzipdNamedTag(w io.Writer, name string, tag Tag) (outerr error) {
	comp := gzip.NewWriter(w)
	defer func(){
		err := comp.Close()
		if outerr != nil {
			outerr = err
		}
	}()
	outerr = WriteNamedTag(comp, name, tag)
	return
}

// ReadZlibdNamedTag reads a zlib compressed named tag. See ReadNamedTags for more info.
func ReadZlibdNamedTag(r io.Reader) (Tag, string, error) {
	decomp, err := zlib.NewReader(r)
	if err != nil {
		return Tag{}, "", err
	}

	return ReadNamedTag(decomp)
}

// WriteZlibdNamedTag writes a zlib compressed named tag. See WriteNamedTag for more info.
func WriteZlibdNamedTag(w io.Writer, name string, tag Tag) (outerr error) {
	comp := zlib.NewWriter(w)
	defer func(){
		err := comp.Close()
		if outerr != nil {
			outerr = err
		}
	}()
	return WriteNamedTag(comp, name, tag)
}
