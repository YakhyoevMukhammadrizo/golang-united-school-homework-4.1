package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	byte_str := []rune(input)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	output = string(byte_str)
	return output
}
