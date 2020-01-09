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
		{"The volume and fuel consumption are equal", 15,15,100},
		{"The volume and fuel consumption are equal", 7,7,100},
		{"Fuel volume is twice as much fuel consumption", 13,26,200},
		{"Fuel will run out soon", 17,1,5},
		{"No fuel", 8,0,0},
		{"No fuel", 3,0,0},
	}
	for _, tt := range tests {
		got := distance(tt.fuelConsumption, tt.fuelVolume)
		if got != tt.want{
			t.Error("for fuel test:",tt.name,"got:",got,"want:",tt.want)
		}

	}
}