package values

// Interface returns the interface of the ith value.
func (s Slice) Interface(i int) (v interface{}, err error) {
	if len(s) <= i {
		return nil, ErrOutOfLen
	}

	return s[i], nil
}

// Slice does the best to convert the value whose index is i to Slice.
func (s Slice) Slice(i int) (v Slice, err error) {
	if len(s) <= i {
		return nil, ErrOutOfLen
	}
	return toSlice(s[i])
}

// IsSlice returns true when the type of the ith value is Slice or []interface{};
// or false.
func (s Slice) IsSlice(i int) bool {
	if len(s) <= i {
		return false
	}
	return isSlice(s[i])
}

// SMap does the best to convert the value whose index is i to SMap.
func (s Slice) SMap(i int) (v SMap, err error) {
	if len(s) <= i {
		return nil, ErrOutOfLen
	}
	return toSMap(s[i])
}

// IsSMap returns true when the type of the ith value is SMap or
// map[string]interface{}; or false.
func (s Slice) IsSMap(i int) bool {
	if len(s) <= i {
		return false
	}
	return isSMap(s[i])
}
