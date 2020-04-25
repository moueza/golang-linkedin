package stringutil

 
//uppercase
func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length //order important
}

//order not important
func FullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l //=because already in signature
	length = len(full)
	return
}
