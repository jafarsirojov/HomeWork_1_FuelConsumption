package main

import (
	"testing"
)

func Test_fuel(t *testing.T) {

	tests := []struct {
		name string
		fuelConsumption int
		fuelVolume int
		want int
	}{
		// TODO: Add test cases.
		{"One hundred km", 15,15,100},
		{"Two hundred km", 13,26,200},
		{"50 km", 10,5,50},
	}
	for _, tt := range tests {
		got := distance(tt.fuelConsumption, tt.fuelVolume)
		if got != tt.want{
			t.Error("for fuel test:",tt.name,"got:",got,"want:",tt.want)
		}

	}
}