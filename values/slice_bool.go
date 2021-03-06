package values

// Bool does the best to convert the value whose index is i to bool.
func (s Slice) Bool(i int) (v bool, err error) {
	if len(s) <= i {
		err = ErrOutOfLen
		return
	}

	return !IsZero(s[i]), nil
}

// IsBool returns true when the type of the ith value is bool; or false.
func (s Slice) IsBool(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].(bool)
	return ok
}
