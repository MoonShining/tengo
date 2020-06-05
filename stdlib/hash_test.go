package stdlib_test

import "testing"

func TestMd5(t *testing.T) {
	module(t, "hash").call("md5", "abc").expect("900150983cd24fb0d6963f7d28e17f72")
	module(t, "hash").call("md5", "").expect("d41d8cd98f00b204e9800998ecf8427e")
	module(t, "hash").call("md5", 1).expect("c4ca4238a0b923820dcc509a6f75849b")
}
