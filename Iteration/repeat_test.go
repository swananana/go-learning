package iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 4)
	expected := "aaaa"

	if actual != expected {
		t.Errorf("expected '%s' but got '%s'", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
}
}

func ExampleRepeat() {
	repeated := Repeat("a", 4)
	fmt.Println(repeated)
	// Output: aaaa
}