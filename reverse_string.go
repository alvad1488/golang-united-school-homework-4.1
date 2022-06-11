package reverse_string

func ReverseString(input string) (output string) {

	check := []rune(input)
	length := len(check)

	if length > 0 {
		for i := 0; i < length/2; i++ {
			tmp := check[length-1-i]
			check[length-1-i] = check[i]
			check[i] = tmp
		}
	}
	output = string(check)
	return
}
