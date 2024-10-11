package main

type RaceCar struct {
	year   int
	color  int
	engine string
}

func (c *RaceCar) getYear() int {
	return c.year
}

func (c *RaceCar) getColor() int {
	return c.color
}
func (c *RaceCar) getEngine() string {
	return c.engine
}

func newRaceCar(y, c int, e string) ICar {
	return &RaceCar{
		year:   y,
		color:  c,
		engine: e,
	}
}
