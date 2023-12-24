package iteration 

const counter = 4
func Repeat(char string, n int ) string {
	var repeated string
	counter := n
	for i := 0; i < counter; i++ {
		repeated += char
	}
	return repeated
}