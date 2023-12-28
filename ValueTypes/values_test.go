package main

import "testing"

func TestValueTypes(t *testing.T) {
	valueTests := []struct {
		name    string
		value   interface{}
		want string
	}{
		{name: "int", value: 1, want: "int"},
		{name: "float64", value: 0.5, want: "float64"},
		
	}

	for _, tt := range valueTests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindType(tt.value)
			if got != tt.want {
				t.Errorf("%#v got %s want %s", tt.name, got, tt.want)
			}
		})
	}

}