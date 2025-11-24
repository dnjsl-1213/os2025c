package main

import "fmt"

type Kilometers float64
type Meters float64
type Miless float64

//변환 메서드(Km -> Mile,Meter -> Mile)
func (km Kilometers) ToMiless() Miless {
	return Miless(km / 1.609)
}
func (m Meters) ToMiless() Miless {
	return Miless(m / 1609)
}

func main() {
	kmph := Kilometers(151)
	fmt.Printf("%0.3f Kilometers per hour %0.2f mile per hour\n", kmph, kmph.ToMiless())
	meter := Meters(151000)
	fmt.Printf("%0.3f Miless equals %0.2f miles\n", meter, meter.ToMiless())
}
