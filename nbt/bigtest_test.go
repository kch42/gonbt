package nbt

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func bigtest() []byte {
	// Generated with go-bindata using the bigtest NBT file.
	return []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xed, 0x54,
		0xcf, 0x4f, 0x1a, 0x41, 0x14, 0x7e, 0xc2, 0x02, 0xcb, 0x96, 0x82, 0xb1,
		0xc4, 0x10, 0x63, 0xcc, 0xab, 0xb5, 0x84, 0xa5, 0xdb, 0xcd, 0x42, 0x11,
		0x89, 0xb1, 0x88, 0x16, 0x2c, 0x9a, 0x0d, 0x1a, 0xd8, 0xa8, 0x31, 0x86,
		0xb8, 0x2b, 0xc3, 0x82, 0x2e, 0xbb, 0x66, 0x77, 0xb0, 0xf1, 0xd4, 0x4b,
		0x7b, 0x6c, 0x7a, 0xeb, 0x3f, 0xd3, 0x23, 0x7f, 0x43, 0xcf, 0xbd, 0xf6,
		0xbf, 0xa0, 0xc3, 0x2f, 0x7b, 0x69, 0xcf, 0xbd, 0xf0, 0x32, 0xc9, 0xf7,
		0xe6, 0xbd, 0x6f, 0xe6, 0x7b, 0x6f, 0x26, 0x79, 0x02, 0x04, 0x54, 0x72,
		0x4f, 0x2c, 0x0e, 0x78, 0xcb, 0xb1, 0x4d, 0x8d, 0x78, 0xf4, 0xe3, 0x70,
		0x62, 0x3e, 0x08, 0x7b, 0x1d, 0xc7, 0xa5, 0x93, 0x18, 0x0f, 0x82, 0x47,
		0xdd, 0xee, 0x84, 0x02, 0x62, 0xb5, 0xa2, 0xaa, 0xc7, 0x78, 0x76, 0x5c,
		0x57, 0xcb, 0xa8, 0x55, 0x0f, 0x1b, 0xc8, 0xd6, 0x1e, 0x6a, 0x95, 0x86,
		0x86, 0x0d, 0xad, 0x7e, 0x58, 0x7b, 0x8f, 0x83, 0xcf, 0x83, 0x4f, 0x83,
		0x6f, 0xcf, 0x03, 0x10, 0x6e, 0x5b, 0x8e, 0x3e, 0xbe, 0xa5, 0x38, 0x4c,
		0x64, 0xfd, 0x10, 0xea, 0xda, 0x74, 0xa6, 0x23, 0x40, 0xdc, 0x66, 0x2e,
		0x69, 0xe1, 0xb5, 0xd3, 0xbb, 0x73, 0xfa, 0x76, 0x0b, 0x29, 0xdb, 0x0b,
		0xe0, 0xef, 0xe8, 0x3d, 0x1e, 0x38, 0x5b, 0xef, 0x11, 0x08, 0x56, 0xf5,
		0xde, 0x5d, 0xdf, 0x0b, 0x40, 0xe0, 0x5e, 0xb7, 0xfa, 0x64, 0xb7, 0x04,
		0x00, 0x8c, 0x41, 0x4c, 0x73, 0xc6, 0x08, 0x55, 0x4c, 0xd3, 0x20, 0x2e,
		0x7d, 0xa4, 0xc0, 0xc8, 0xc2, 0x10, 0xb3, 0xba, 0xde, 0x58, 0x0b, 0x53,
		0xa3, 0xee, 0x44, 0x8e, 0x45, 0x03, 0x30, 0xb1, 0x27, 0x53, 0x8c, 0x4c,
		0xf1, 0xe9, 0x14, 0xa3, 0x53, 0x8c, 0x85, 0xe1, 0xd9, 0x9f, 0xe3, 0xb3,
		0xf2, 0x44, 0x81, 0xa5, 0x7c, 0x33, 0xdd, 0xd8, 0xbb, 0xc7, 0xaa, 0x75,
		0x13, 0x5f, 0x28, 0x1c, 0x08, 0xd7, 0x2e, 0xd1, 0x59, 0x3f, 0xaf, 0x1d,
		0x1b, 0x60, 0x21, 0x59, 0xdf, 0xfa, 0xf1, 0x05, 0xfe, 0xc1, 0xce, 0xfc,
		0x9d, 0xbd, 0x00, 0xbc, 0xf1, 0x40, 0xc9, 0xf8, 0x85, 0x42, 0x40, 0x46,
		0xfe, 0x9e, 0xeb, 0xea, 0x0f, 0x93, 0x3a, 0x68, 0x87, 0x60, 0xbb, 0xeb,
		0x32, 0x37, 0xa3, 0x28, 0x0a, 0x8e, 0xbb, 0xf5, 0xd0, 0x69, 0x63, 0xca,
		0x4e, 0xdb, 0xe9, 0xec, 0xe6, 0xe6, 0x2b, 0x3b, 0xbd, 0x25, 0xbe, 0x64,
		0x49, 0x09, 0x3d, 0xaa, 0xbb, 0x94, 0xfd, 0x18, 0x7e, 0xe8, 0xd2, 0x0e,
		0xda, 0x6f, 0x15, 0x4c, 0xb1, 0x68, 0x3e, 0x2b, 0xe1, 0x9b, 0x9c, 0x84,
		0x99, 0xbc, 0x84, 0x05, 0x09, 0x65, 0x59, 0x16, 0x45, 0x00, 0xff, 0x2f,
		0x28, 0xae, 0x2f, 0xf2, 0xc2, 0xb2, 0xa4, 0x2e, 0x1d, 0x20, 0x77, 0x5a,
		0x3b, 0xb9, 0x8c, 0xca, 0xe7, 0x29, 0xdf, 0x51, 0x41, 0xc9, 0x16, 0xb5,
		0xc5, 0x6d, 0xa1, 0x2a, 0xad, 0x2c, 0xc5, 0x31, 0x7f, 0xba, 0x7a, 0x92,
		0x8e, 0x5e, 0x9d, 0x5f, 0xf8, 0x12, 0x05, 0x23, 0x1b, 0xd1, 0xf6, 0xb7,
		0x77, 0xaa, 0xcd, 0x95, 0x72, 0xbc, 0x9e, 0xdf, 0x58, 0x5d, 0x4b, 0x97,
		0xae, 0x92, 0x17, 0xb9, 0x44, 0xd0, 0x80, 0xc8, 0xfa, 0x3e, 0xbf, 0xb3,
		0xdc, 0x54, 0xcb, 0x07, 0x75, 0x6e, 0xa3, 0xb6, 0x76, 0x59, 0x92, 0x93,
		0xa9, 0xdc, 0x51, 0x50, 0x99, 0x6b, 0xcc, 0x35, 0xe6, 0x1a, 0xff, 0x57,
		0x23, 0x08, 0x42, 0xcb, 0xe9, 0x1b, 0xd6, 0x78, 0xc2, 0xec, 0xfe, 0xfc,
		0x7a, 0xfb, 0x7d, 0x78, 0xd3, 0x84, 0xdf, 0xd4, 0xf2, 0xa4, 0xfb, 0x08,
		0x06, 0x00, 0x00,
	}
}

func byteArrayTestSeries(n int) byte { return byte((n*n*255 + n*7) % 100) }

func getKey(t *testing.T, comp TagCompound, key string, tt TagType) (tag Tag, ok bool) {
	tag, ok = comp[key]
	if !ok {
		t.Errorf("Missing key %#v.", key)
	}

	if tag.Type != tt {
		t.Errorf("Value for key %#v has wrong TagType. Want: %s, have: %s", key, tt, tag.Type)
		ok = false
	}

	return
}

func tagPayloadCompare(t *testing.T, name string, tag Tag, _want interface{}) {
	switch tag.Type {
	case TAG_Byte:
		want := _want.(byte)
		have := tag.Payload.(byte)
		if want != have {
			t.Errorf("Wrong payload for %s. Want: %#v, have: %#v", name, want, have)
		}
	case TAG_Short:
		want := _want.(int16)
		have := tag.Payload.(int16)
		if want != have {
			t.Errorf("Wrong payload for %s. Want: %#v, have: %#v", name, want, have)
		}
	case TAG_Int:
		want := _want.(int32)
		have := tag.Payload.(int32)
		if want != have {
			t.Errorf("Wrong payload for %s. Want: %#v, have: %#v", name, want, have)
		}
	case TAG_Long:
		want := _want.(int64)
		have := tag.Payload.(int64)
		if want != have {
			t.Errorf("Wrong payload for %s. Want: %#v, have: %#v", name, want, have)
		}
	case TAG_Float:
		want := _want.(float32)
		have := tag.Payload.(float32)
		if want != have {
			t.Errorf("Wrong payload for %s. Want: %#v, have: %#v", name, want, have)
		}
	case TAG_Double:
		want := _want.(float64)
		have := tag.Payload.(float64)
		if want != have {
			t.Errorf("Wrong payload for %s. Want: %#v, have: %#v", name, want, have)
		}
	case TAG_String:
		want := _want.(string)
		have := tag.Payload.(string)
		if want != have {
			t.Errorf("Wrong payload for %s. Want: %#v, have: %#v", name, want, have)
		}
	default:
		panic(fmt.Errorf("tagPayloadCompare doesn't know how to compare payloads for TagType %s!", tag.Type))
	}
}

func nestedCompoundHelper(t *testing.T, comp TagCompound, id, name string, value float32) {
	if tag, ok := getKey(t, comp, id, TAG_Compound); ok {
		nested := tag.Payload.(TagCompound)
		if tag, ok := getKey(t, nested, "name", TAG_String); ok {
			tagPayloadCompare(t, "name", tag, name)
		}
		if tag, ok := getKey(t, nested, "value", TAG_Float); ok {
			tagPayloadCompare(t, "value", tag, value)
		}
	}
}

func TestBigtest(t *testing.T) {
	testBigtest(bytes.NewReader(bigtest()), t)
}

func testBigtest(r io.Reader, t *testing.T) {
	root, name, err := ReadGzipdNamedTag(r)
	if err != nil {
		t.Fatalf("Could not read NBT data: %s", err)
	}

	if name != "Level" {
		t.Errorf("Wrong name. Want \"Level\", have %#v", name)
	}

	if root.Type != TAG_Compound {
		t.Fatalf("Root tag must be a TAG_Compound, have %s", root.Type)
	}
	rootcomp := root.Payload.(TagCompound)

	if tag, ok := getKey(t, rootcomp, "shortTest", TAG_Short); ok {
		tagPayloadCompare(t, "shortTest", tag, int16(32767))
	}
	if tag, ok := getKey(t, rootcomp, "longTest", TAG_Long); ok {
		tagPayloadCompare(t, "longTest", tag, int64(9223372036854775807))
	}
	if tag, ok := getKey(t, rootcomp, "floatTest", TAG_Float); ok {
		tagPayloadCompare(t, "floatTest", tag, float32(0.49823147))
	}
	if tag, ok := getKey(t, rootcomp, "stringTest", TAG_String); ok {
		tagPayloadCompare(t, "stringTest", tag, "HELLO WORLD THIS IS A TEST STRING ÅÄÖ!")
	}
	if tag, ok := getKey(t, rootcomp, "intTest", TAG_Int); ok {
		tagPayloadCompare(t, "intTest", tag, int32(2147483647))
	}
	if tag, ok := getKey(t, rootcomp, "byteTest", TAG_Byte); ok {
		tagPayloadCompare(t, "byteTest", tag, byte(127))
	}
	if tag, ok := getKey(t, rootcomp, "doubleTest", TAG_Double); ok {
		tagPayloadCompare(t, "doubleTest", tag, float64(0.4931287132182315))
	}

	if tag, ok := getKey(t, rootcomp, "listTest (compound)", TAG_List); ok {
		l := tag.Payload.(TagList)
		if l.Type != TAG_Compound {
			t.Errorf("\"listTest (compound)\" is list of type %s, expected list of type TAG_Compound", l.Type)
		} else if len(l.Elems) != 2 {
			t.Errorf("\"listTest (compound)\" has length %d, expected 2", len(l.Elems))
		} else {
			for i, elem := range l.Elems {
				comp := elem.(TagCompound)
				if tag, ok := getKey(t, comp, "name", TAG_String); ok {
					tagPayloadCompare(t, fmt.Sprintf("listTest (compound) [%d] name", i), tag, fmt.Sprintf("Compound tag #%d", i))
				}
				if tag, ok := getKey(t, comp, "created-on", TAG_Long); ok {
					tagPayloadCompare(t, fmt.Sprintf("listTest (compound) [%d] created-on", i), tag, int64(1264099775885))
				}
			}
		}
	}

	if tag, ok := getKey(t, rootcomp, "listTest (long)", TAG_List); ok {
		l := tag.Payload.(TagList)
		if l.Type != TAG_Long {
			t.Errorf("\"listTest (long)\" is list of type %s, expected list of type TAG_Long", l.Type)
		} else if len(l.Elems) != 5 {
			t.Errorf("\"listTest (long)\" has length %d, expected 5", len(l.Elems))
		} else {
			for i, want := range []int64{11, 12, 13, 14, 15} {
				if have := l.Elems[i].(int64); have != want {
					t.Errorf("listTest (long) [%d] = %d, expected %d", i, have, want)
				}
			}
		}
	}

	if tag, ok := getKey(t, rootcomp, "nested compound test", TAG_Compound); ok {
		comp := tag.Payload.(TagCompound)
		nestedCompoundHelper(t, comp, "ham", "Hampus", 0.75)
		nestedCompoundHelper(t, comp, "egg", "Eggbert", 0.5)
	}

	if tag, ok := getKey(t, rootcomp, "byteArrayTest (the first 1000 values of (n*n*255+n*7)%100, starting with n=0 (0, 62, 34, 16, 8, ...))", TAG_Byte_Array); ok {
		data := tag.Payload.([]byte)
		if len(data) != 1000 {
			t.Errorf("byteArrayTest data has length %d, expected 1000", len(data))
		} else {
			for n := 0; n < 1000; n++ {
				want := byteArrayTestSeries(n)
				if data[n] != want {
					t.Errorf("Wrong byteArrayTest data at index %d: 0x%02x", n, data[n])
					break
				}
			}
		}
	}
}

func makeNested(name string, value float32) Tag {
	comp := make(TagCompound)
	comp["name"] = NewStringTag(name)
	comp["value"] = NewFloatTag(value)
	return Tag{TAG_Compound, comp}
}

func TestCreateBigtest(t *testing.T) {
	rootcomp := make(TagCompound)

	rootcomp["shortTest"] = NewShortTag(32767)
	rootcomp["longTest"] = NewLongTag(9223372036854775807)
	rootcomp["floatTest"] = NewFloatTag(0.49823147)
	rootcomp["stringTest"] = NewStringTag("HELLO WORLD THIS IS A TEST STRING ÅÄÖ!")
	rootcomp["intTest"] = NewIntTag(2147483647)
	rootcomp["byteTest"] = NewByteTag(127)
	rootcomp["doubleTest"] = NewDoubleTag(0.4931287132182315)

	comp := make(TagCompound)
	comp["ham"] = makeNested("Hampus", 0.75)
	comp["egg"] = makeNested("Eggbert", 0.5)
	rootcomp["nested compound test"] = Tag{TAG_Compound, comp}

	listlong := make([]interface{}, 5)
	for i := 0; i < 5; i++ {
		listlong[i] = int64(i + 11)
	}
	rootcomp["listTest (long)"] = Tag{TAG_List, TagList{TAG_Long, listlong}}

	listcomp := make([]interface{}, 2)
	for i := 0; i < 2; i++ {
		comp := make(TagCompound)
		comp["name"] = NewStringTag(fmt.Sprintf("Compound tag #%d", i))
		comp["created-on"] = NewLongTag(1264099775885)
		listcomp[i] = comp
	}
	rootcomp["listTest (compound)"] = Tag{TAG_List, TagList{TAG_Compound, listcomp}}

	data := make([]byte, 1000)
	for n := 0; n < 1000; n++ {
		data[n] = byteArrayTestSeries(n)
	}
	rootcomp["byteArrayTest (the first 1000 values of (n*n*255+n*7)%100, starting with n=0 (0, 62, 34, 16, 8, ...))"] = NewByteArrayTag(data)

	tag := Tag{TAG_Compound, rootcomp}
	buf := new(bytes.Buffer)

	if err := WriteGzipdNamedTag(buf, "Level", tag); err != nil {
		t.Fatalf("Could not write NBT data: %s", err)
	}

	testBigtest(buf, t)
}
