package main

func main() {

}
func distance(fuelConsumption,fuelVolume int)  int{
	const km  = 100
	total := 0
	total = fuelVolume*km/fuelConsumption
	return total
}
