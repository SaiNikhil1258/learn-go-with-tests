package iteration

func Repeat(char string, count int) string {
	var res string
	for i := 0; i < count; i++ {
		res += char
	}
	return res
}
