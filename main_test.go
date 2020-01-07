package main

import (
	"testing"
)

func Test_toplivo(t *testing.T) {

	tests := []struct {
		name string
		fuelCons int
		fuelVol int
		want int
	}{
		// TODO: Add test cases.
		{"One hundred km", 15,15,100},
		{"Two hundred km", 13,26,200},
		{"50 km", 10,5,50},
	}
	for _, tt := range tests {
		got := fuel(tt.fuelCons, tt.fuelVol)
		if got != tt.want{
			t.Error("for toplivo test:",tt.name,"got:",got,"want:",tt.want)
		}

	}
}