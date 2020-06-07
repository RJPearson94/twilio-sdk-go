package utils

//  Returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

//  Returns a pointer to the bool value passed in.
func Bool(v bool) *bool {
	return &v
}

//  Returns a pointer to the int value passed in.
func Int(v int) *int {
	return &v
}
