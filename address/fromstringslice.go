package address

//FromStringSlice turns a length-7 string slice into an Address.
func FromStringSlice(s []string) (Address, bool) {
	if len(s) != 7 {
		return Address{}, false
	}
	return Address{s[0], s[1], s[2], s[3], s[4], s[5], s[6]}, true
}
