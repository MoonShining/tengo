package stdlib

import (
	"encoding/binary"
	"github.com/d5/tengo/objects"
)

var binaryModule = map[string]objects.Object{
	"big_endian_bytes_to_long": &objects.UserFunction{Name:"big_endian_bytes_to_long", Value: bigEndianBytesToLong},
	"little_endian_bytes_to_long": &objects.UserFunction{Name:"little_endian_bytes_to_long", Value: littleEndianBytesToLong},
}

func bigEndianBytesToLong(args ...objects.Object)(objects.Object, error){
	numArgs := len(args)
	if numArgs != 1 {
		return nil, objects.ErrWrongNumArguments
	}
	b, ok := args[0].(*objects.Bytes)
	if !ok {
		return nil, objects.ErrInvalidArgumentType{
			Name:     "bigEndianBytesToLong",
			Expected: "bytes",
			Found:    args[0].TypeName(),
		}
	}
	if len(b.Value) != 8 {
		return nil, objects.ErrInvalidArgumentType{
			Name:     "bigEndianBytesToLong",
			Expected: "8 bytes",
			Found:    args[0].TypeName(),
		}
	}

	num := int64(binary.BigEndian.Uint64(b.Value))
	return &objects.Int{Value: num}, nil
}

func littleEndianBytesToLong(args ...objects.Object)(objects.Object, error){
	numArgs := len(args)
	if numArgs != 1 {
		return nil, objects.ErrWrongNumArguments
	}
	b, ok := args[0].(*objects.Bytes)
	if !ok {
		return nil, objects.ErrInvalidArgumentType{
			Name:     "bigEndianBytesToLong",
			Expected: "bytes",
			Found:    args[0].TypeName(),
		}
	}
	if len(b.Value) != 8 {
		return nil, objects.ErrInvalidArgumentType{
			Name:     "bigEndianBytesToLong",
			Expected: "8 bytes",
			Found:    args[0].TypeName(),
		}
	}

	num := int64(binary.LittleEndian.Uint64(b.Value))
	return &objects.Int{Value: num}, nil
}