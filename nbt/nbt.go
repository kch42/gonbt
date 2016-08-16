// Package nbt provides functions for reading and writing NBT encoded data.
package nbt

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/silvasur/kagus"
	"io"
	"strconv"
)

// Tag holds the data of an NBT tag. Type is a TAG_* value.
//
// The type of Payload depends on Type:
//
// 	TAG_Byte       -- byte
// 	TAG_Short      -- int16
// 	TAG_Int        -- int32
// 	TAG_Long       -- int64
// 	TAG_Float      -- float32
// 	TAG_Double     -- float64
// 	TAG_Byte_Array -- []byte
// 	TAG_String     -- string
// 	TAG_List       -- TagList
// 	TAG_Compound   -- TagCompound
// 	TAG_Int_Array  -- []int32
type Tag struct {
	Type    TagType
	Payload interface{}
}

// TagList is the payload of a TAG_List Tag. The type of the Elems depends on Type (see docu of Tag).
type TagList struct {
	Type  TagType
	Elems []interface{}
}

func (t Tag) String() string {
	s := t.Type.String()
	switch t.Type {
	case TAG_Byte:
		s += fmt.Sprintf(": 0x%02x", t.Payload.(byte))
	case TAG_Short:
		s += fmt.Sprintf(": %d", t.Payload.(int16))
	case TAG_Int:
		s += fmt.Sprintf(": %d", t.Payload.(int32))
	case TAG_Long:
		s += fmt.Sprintf(": %d", t.Payload.(int64))
	case TAG_Float:
		s += fmt.Sprintf(": %f", t.Payload.(float32))
	case TAG_Double:
		s += fmt.Sprintf(": %f", t.Payload.(float64))
	case TAG_Byte_Array:
		s += "\n" + kagus.Indent(hex.Dump(t.Payload.([]byte)), "  ")
	case TAG_String:
		s += ": " + strconv.Quote(t.Payload.(string))
	case TAG_List:
		l := t.Payload.(TagList)
		s += " of " + l.Type.String() + ":"
		for _, elem := range l.Elems {
			s += "\n" + kagus.Indent(Tag{l.Type, elem}.String(), "  ")
		}
	case TAG_Compound:
		s += ":"
		comp := t.Payload.(TagCompound)
		for k, v := range comp {
			s += "\n" + kagus.Indent(strconv.Quote(k)+"  ->"+kagus.Indent(v.String(), "  "), "  ")
		}
		return s
	case TAG_Int_Array:
		l := t.Payload.([]int32)
		s += ": "
		sep := ""
		for _, elem := range l {
			s += sep + fmt.Sprintf("%d", elem)
			sep = ", "
		}
	}

	return s
}

// TagCompund is the payload of a TAG_Compound. Initialize with make.
type TagCompound map[string]Tag

func readTagData(r io.Reader, tt TagType) (interface{}, error) {
	switch tt {
	case TAG_End:
	case TAG_Byte:
		var v uint8
		err := binary.Read(r, binary.BigEndian, &v)
		return v, err
	case TAG_Short:
		var v int16
		err := binary.Read(r, binary.BigEndian, &v)
		return v, err
	case TAG_Int:
		var v int32
		err := binary.Read(r, binary.BigEndian, &v)
		return v, err
	case TAG_Long:
		var v int64
		err := binary.Read(r, binary.BigEndian, &v)
		return v, err
	case TAG_Float:
		var v float32
		err := binary.Read(r, binary.BigEndian, &v)
		return v, err
	case TAG_Double:
		var v float64
		err := binary.Read(r, binary.BigEndian, &v)
		return v, err
	case TAG_Byte_Array:
		var l int32
		if err := binary.Read(r, binary.BigEndian, &l); err != nil {
			return nil, err
		}
		if l < 0 {
			return nil, errors.New("Byte array has negative length?")
		}

		data := make([]byte, l)
		_, err := io.ReadFull(r, data)
		return data, err
	case TAG_String:
		var l int16
		if err := binary.Read(r, binary.BigEndian, &l); err != nil {
			return nil, err
		}
		if l < 0 {
			return nil, errors.New("String has negative length?")
		}

		data := make([]byte, l)
		_, err := io.ReadFull(r, data)
		return string(data), err
	case TAG_List:
		_ltt, err := kagus.ReadByte(r)
		if err != nil {
			return nil, err
		}
		ltt := TagType(_ltt)

		var l int32
		if err := binary.Read(r, binary.BigEndian, &l); err != nil {
			return nil, err
		}
		if l < 0 {
			return nil, errors.New("List has negative length?")
		}

		tl := TagList{Type: ltt, Elems: make([]interface{}, l)}
		for i := 0; i < int(l); i++ {
			if tl.Elems[i], err = readTagData(r, ltt); err != nil {
				return nil, err
			}
		}
		return tl, nil
	case TAG_Compound:
		comp := make(TagCompound)
		for {
			tag, name, err := ReadNamedTag(r)
			if err != nil {
				return nil, err
			}
			if tag.Type == TAG_End {
				break
			}
			comp[name] = tag
		}
		return comp, nil
	case TAG_Int_Array:
		var l int32
		if err := binary.Read(r, binary.BigEndian, &l); err != nil {
			return nil, err
		}
		if l < 0 {
			return nil, errors.New("Int Array has negative length?")
		}

		data := make([]int32, l)
		for i := 0; i < int(l); i++ {
			var e int32
			if err := binary.Read(r, binary.BigEndian, &e); err != nil {
				return nil, err
			}
			data[i] = e
		}
		return data, nil
	}

	return nil, errors.New("Unknown tag type")
}

// ReadNamedTag reads a named Tag from an io.Reader. It returns the Tag, the tags Name and an error.
func ReadNamedTag(r io.Reader) (Tag, string, error) {
	_tt, err := kagus.ReadByte(r)
	if err != nil {
		return Tag{}, "", err
	}
	tt := TagType(_tt)

	if tt == TAG_End {
		return Tag{Type: tt}, "", nil
	}

	name, err := readTagData(r, TAG_String)
	if err != nil {
		return Tag{}, "", err
	}

	td, err := readTagData(r, tt)
	return Tag{Type: tt, Payload: td}, name.(string), err
}

func writeByte(w io.Writer, b byte) error {
	_, err := w.Write([]byte{b})
	return err
}

func writeTagData(w io.Writer, tt TagType, data interface{}) error {
	switch tt {
	case TAG_End:
		return nil
	case TAG_Byte:
		return writeByte(w, data.(byte))
	case TAG_Short:
		return binary.Write(w, binary.BigEndian, data.(int16))
	case TAG_Int:
		return binary.Write(w, binary.BigEndian, data.(int32))
	case TAG_Long:
		return binary.Write(w, binary.BigEndian, data.(int64))
	case TAG_Float:
		return binary.Write(w, binary.BigEndian, data.(float32))
	case TAG_Double:
		return binary.Write(w, binary.BigEndian, data.(float64))
	case TAG_Byte_Array:
		slice := data.([]byte)
		if err := binary.Write(w, binary.BigEndian, int32(len(slice))); err != nil {
			return err
		}
		_, err := w.Write(slice)
		return err
	case TAG_String:
		strEnc := []byte(data.(string))
		if err := binary.Write(w, binary.BigEndian, int16(len(strEnc))); err != nil {
			return err
		}
		_, err := w.Write(strEnc)
		return err
	case TAG_List:
		list := data.(TagList)
		if err := writeByte(w, byte(list.Type)); err != nil {
			return err
		}

		if err := binary.Write(w, binary.BigEndian, int32(len(list.Elems))); err != nil {
			return err
		}

		for _, el := range list.Elems {
			if err := writeTagData(w, list.Type, el); err != nil {
				return err
			}
		}
		return nil
	case TAG_Compound:
		comp := data.(TagCompound)
		for name, tag := range comp {
			if err := WriteNamedTag(w, name, tag); err != nil {
				return err
			}
		}
		return writeByte(w, TAG_End)
	case TAG_Int_Array:
		slice := data.([]int32)
		if err := binary.Write(w, binary.BigEndian, int32(len(slice))); err != nil {
			return err
		}

		for _, el := range slice {
			if err := binary.Write(w, binary.BigEndian, el); err != nil {
				return err
			}
		}

		return nil
	}

	return errors.New("Unknown tage type")
}

// WriteNamedTag writes a named Tag to an io.Writer.
func WriteNamedTag(w io.Writer, name string, tag Tag) error {
	if err := writeByte(w, byte(tag.Type)); err != nil {
		return err
	}

	if err := writeTagData(w, TAG_String, name); err != nil {
		return err
	}

	return writeTagData(w, tag.Type, tag.Payload)
}
