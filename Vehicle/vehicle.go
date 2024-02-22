package main

import "fmt"

type Vehicle interface {
	Colour() string
	Type() string
	Start()
	Stop()
}

type Honda struct {
	name       string
	colour     string
	engineType string
	state      bool
}

func (h *Honda) Colour() string {
	return h.colour
}

func (h *Honda) Type() string {
	return h.engineType
}

func (h *Honda) Start() {
	if h.state {
		fmt.Println("already running")
	} else {
		h.state = true
		fmt.Println("running")
	}
}

func (h *Honda) Stop() {
	if h.state {
		h.state = false
		fmt.Println("stopped")
	} else {
		fmt.Println("already stopped")
	}
}

func main() {
	fmt.Println("Welcome to my Vehicle Repair Shop")
	car := Honda{colour: "blue", engineType: "petrol", name: "Honda City", state: false}

	fmt.Printf("The colour is %s\n", car.Colour())
	fmt.Printf("The engine type is %s\n", car.Type())
	car.Start()
	car.Start()
	car.Stop()
	car.Stop()
	fmt.Println("Seems like your car is working!")
}
