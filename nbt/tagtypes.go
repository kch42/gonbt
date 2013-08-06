package nbt

// Valid TagType values.
const (
	TAG_End = iota
	TAG_Byte
	TAG_Short
	TAG_Int
	TAG_Long
	TAG_Float
	TAG_Double
	TAG_Byte_Array
	TAG_String
	TAG_List
	TAG_Compound
	TAG_Int_Array
)

// TagType describes the type of a NBT tag. Valid values are the TAG_* constants.
type TagType byte

func (tt TagType) String() string {
	switch tt {
	case TAG_End:
		return "TAG_End"
	case TAG_Byte:
		return "TAG_Byte"
	case TAG_Short:
		return "TAG_Short"
	case TAG_Int:
		return "TAG_Int"
	case TAG_Long:
		return "TAG_Long"
	case TAG_Float:
		return "TAG_Float"
	case TAG_Double:
		return "TAG_Double"
	case TAG_Byte_Array:
		return "TAG_Byte_Array"
	case TAG_String:
		return "TAG_String"
	case TAG_List:
		return "TAG_List"
	case TAG_Compound:
		return "TAG_Compound"
	case TAG_Int_Array:
		return "TAG_Int_Array"
	default:
		return "TAG_Unknown"
	}
}
