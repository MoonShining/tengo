package stdlib

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/d5/tengo/objects"
)

var hashModule = map[string]objects.Object{
	"md5": &objects.UserFunction{Value: md5Func},
}

func md5Func(args ...objects.Object) (ret objects.Object,err error) {
	if len(args) != 1 {
		err = objects.ErrWrongNumArguments
		return
	}
	s, ok := objects.ToString(args[0])
	if !ok {
		err = objects.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	d := md5.Sum([]byte(s))
	ret = &objects.String{Value: hex.EncodeToString(d[:])}
	return
}