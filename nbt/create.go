package nbt

func NewByteTag(v byte) Tag        { return Tag{TAG_Byte, v} }
func NewShortTag(v int16) Tag      { return Tag{TAG_Short, v} }
func NewIntTag(v int32) Tag        { return Tag{TAG_Int, v} }
func NewLongTag(v int64) Tag       { return Tag{TAG_Long, v} }
func NewFloatTag(v float32) Tag    { return Tag{TAG_Float, v} }
func NewDoubleTag(v float64) Tag   { return Tag{TAG_Double, v} }
func NewByteArrayTag(v []byte) Tag { return Tag{TAG_Byte_Array, v} }
func NewStringTag(v string) Tag    { return Tag{TAG_String, v} }
func NewIntArrayTag(v []int32) Tag { return Tag{TAG_Int_Array, v} }

// func NewCompoundTag() Tag {return Tag{TAG_Compound, make(TagCompound)}}
// func NewListTag
