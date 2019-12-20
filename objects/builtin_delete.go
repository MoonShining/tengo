package objects

// map delete
func builtinDelete(args ...Object) (Object, error) {
	if len(args) != 2 {
		return nil, ErrWrongNumArguments
	}

	omap, ok := args[0].(*Map)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "delete",
			Expected: "map",
			Found:    args[0].TypeName(),
		}
	}

	key, ok := args[1].(*String)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "delete",
			Expected: "string",
			Found:    args[1].TypeName(),
		}
	}

	delete(omap.Value, key.Value)
	return omap, nil
}
