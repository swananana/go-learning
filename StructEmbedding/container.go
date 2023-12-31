package main
import "fmt"
type base struct {
    num int
}
func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}
type container struct {
    base
    str string
}

func (c container) define() string {
	return fmt.Sprintf("container with base=%v and str=%v", c.base, c.str)
}
