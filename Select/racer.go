package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {
	select {
	case <-ping(a):
		fmt.Printf("here")
		return a, nil
	case <-ping(b):
		fmt.Printf("here2")
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan error {
	ch := make(chan error)
	go func() {
		_, ok := http.Get(url)
		ch <- ok
		fmt.Printf(url, ok)
		close(ch)
	}()
	return ch
}

func main() {
	Racer("https://pkg.go.dev/builtin#close", "https://open.spotify.com/lyrics")
}
