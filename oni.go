package main

type Oni struct {
	x float64
	y float64
}

func (o *Oni) Update() {
	o.x -= 2
	if o.x < -50 {
		o.x = float64(screenWidth) + 50
	}
}
