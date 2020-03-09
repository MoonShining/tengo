package stdlib_test

import (
	"github.com/d5/tengo/objects"
	"testing"
)

func TestBigEndianBinary(t *testing.T) {
	bigOne := []byte{
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x01),
	}
	bigNegativeOne := []byte{
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
	}
	bigMax := []byte{
		byte(0x7f),
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
	}
	bigMin := []byte{
		byte(0x80),
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x00),
	}



	module(t, `binary`).call("big_endian_bytes_to_long", &objects.Bytes{Value: bigOne}).expect(1)
	module(t, `binary`).call("big_endian_bytes_to_long", &objects.Bytes{Value: bigNegativeOne}).expect(-1)
	module(t, `binary`).call("big_endian_bytes_to_long", &objects.Bytes{Value: bigMax}).expect(9223372036854775807)
	module(t, `binary`).call("big_endian_bytes_to_long", &objects.Bytes{Value: bigMin}).expect(-9223372036854775808)
}

func TestLittleEndianBinary(t *testing.T) {
	littleOne := []byte{
		byte(0x01),
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x00),
	}
	littleNegativeOne := []byte{
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
	}
	littleMax := []byte{
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0xff),
		byte(0x7f),
	}
	littleMin := []byte{
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x00),
		byte(0x80),
	}
	module(t, `binary`).call("little_endian_bytes_to_long", &objects.Bytes{Value: littleOne}).expect(1)
	module(t, `binary`).call("little_endian_bytes_to_long", &objects.Bytes{Value: littleNegativeOne}).expect(-1)
	module(t, `binary`).call("little_endian_bytes_to_long", &objects.Bytes{Value: littleMax}).expect(9223372036854775807)
	module(t, `binary`).call("little_endian_bytes_to_long", &objects.Bytes{Value: littleMin}).expect(-9223372036854775808)
}