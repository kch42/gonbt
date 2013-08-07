package nbt

import (
	"reflect"
)

func NewByteTag(v byte) Tag        { return Tag{TAG_Byte, v} }
func NewShortTag(v int16) Tag      { return Tag{TAG_Short, v} }
func NewIntTag(v int32) Tag        { return Tag{TAG_Int, v} }
func NewLongTag(v int64) Tag       { return Tag{TAG_Long, v} }
func NewFloatTag(v float32) Tag    { return Tag{TAG_Float, v} }
func NewDoubleTag(v float64) Tag   { return Tag{TAG_Double, v} }
func NewByteArrayTag(v []byte) Tag { return Tag{TAG_Byte_Array, v} }
func NewStringTag(v string) Tag    { return Tag{TAG_String, v} }
func NewIntArrayTag(v []int32) Tag { return Tag{TAG_Int_Array, v} }

// NewCompoundTag creates a new Tag with type TAG_Compound. Usually it is more convenient to make the TagCompound payload and then manually construct the Tag value, though.
func NewCompoundTag() Tag { return Tag{TAG_Compound, make(TagCompound)} }

// NewListTag creates a new Tag of type TAG_List with tag elems of type ltt.
//
// l must either be of type []interface{}, where the elements are payloads for ltt tags OR of type []T, where T is the payload type for ltt tags.
//
// YOU are responsible for this, this function will not check the correctness.
// When given wrong data, this function will either panic, or writing the NBT data will later fail.
func NewListTag(ltt TagType, l interface{}) Tag {
	var elems []interface{}
	if is, ok := l.([]interface{}); ok {
		elems = is
	} else {
		val := reflect.ValueOf(l)

		n := val.Len()
		elems = make([]interface{}, n)

		for i := 0; i < n; i++ {
			elems[i] = val.Index(i).Interface()
		}
	}

	return Tag{TAG_List, TagList{ltt, elems}}
}
