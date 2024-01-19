package Select

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	durationA := Duration(a)
	durationB := Duration(b)

	if durationA < durationB {
		return a
	}

	return b

}

func Duration(url string) (duration time.Duration) {
	start := time.Now()
	_, _ = http.Get(url)
	return time.Since(start)
}
