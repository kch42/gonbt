package nbt

import (
	"errors"
)

// Errors for TagCompound.Get* functions
var (
	NotFound  = errors.New("Key not found in TagCompound")
	WrongType = errors.New("Tag has wrong type.")
)

func (tc TagCompound) GetByte(key string) (byte, error) {
	t, ok := tc[key]
	if !ok {
		return 0, NotFound
	}
	if t.Type != TAG_Byte {
		return 0, WrongType
	}
	return t.Payload.(byte), nil
}
func (tc TagCompound) GetShort(key string) (int16, error) {
	t, ok := tc[key]
	if !ok {
		return 0, NotFound
	}
	if t.Type != TAG_Short {
		return 0, WrongType
	}
	return t.Payload.(int16), nil
}
func (tc TagCompound) GetInt(key string) (int32, error) {
	t, ok := tc[key]
	if !ok {
		return 0, NotFound
	}
	if t.Type != TAG_Int {
		return 0, WrongType
	}
	return t.Payload.(int32), nil
}
func (tc TagCompound) GetLong(key string) (int64, error) {
	t, ok := tc[key]
	if !ok {
		return 0, NotFound
	}
	if t.Type != TAG_Long {
		return 0, WrongType
	}
	return t.Payload.(int64), nil
}
func (tc TagCompound) GetFloat(key string) (float32, error) {
	t, ok := tc[key]
	if !ok {
		return 0, NotFound
	}
	if t.Type != TAG_Float {
		return 0, WrongType
	}
	return t.Payload.(float32), nil
}
func (tc TagCompound) GetDouble(key string) (float64, error) {
	t, ok := tc[key]
	if !ok {
		return 0, NotFound
	}
	if t.Type != TAG_Double {
		return 0, WrongType
	}
	return t.Payload.(float64), nil
}
func (tc TagCompound) GetByteArray(key string) ([]byte, error) {
	t, ok := tc[key]
	if !ok {
		return nil, NotFound
	}
	if t.Type != TAG_Byte_Array {
		return nil, WrongType
	}
	return t.Payload.([]byte), nil
}
func (tc TagCompound) GetString(key string) (string, error) {
	t, ok := tc[key]
	if !ok {
		return "", NotFound
	}
	if t.Type != TAG_String {
		return "", WrongType
	}
	return t.Payload.(string), nil
}
func (tc TagCompound) GetList(key string) (TagList, error) {
	t, ok := tc[key]
	if !ok {
		return TagList{}, NotFound
	}
	if t.Type != TAG_List {
		return TagList{}, WrongType
	}
	return t.Payload.(TagList), nil
}
func (tc TagCompound) GetCompound(key string) (TagCompound, error) {
	t, ok := tc[key]
	if !ok {
		return nil, NotFound
	}
	if t.Type != TAG_Compound {
		return nil, WrongType
	}
	return t.Payload.(TagCompound), nil
}
func (tc TagCompound) GetIntArray(key string) ([]int32, error) {
	t, ok := tc[key]
	if !ok {
		return nil, NotFound
	}
	if t.Type != TAG_Int_Array {
		return nil, WrongType
	}
	return t.Payload.([]int32), nil
}
