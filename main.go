package main

func main() {

}
func distance(fuelCons,fuelVol int)  int{
	const km  = 100
	total := 0
	total = fuelVol*km/fuelCons
	return total
}
