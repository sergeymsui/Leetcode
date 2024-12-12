package main

import (
	"fmt"
	"log"
)

func getCar(type_name string, year int, color int, engine string) (ICar, error) {

	if type_name == "race" {
		return newRaceCar(year, color, engine), nil
	}

	return nil, fmt.Errorf("undefined type for %s", type_name)
}

func main() {

	car, err := getCar("race", 2020, 1, "Doodge v12")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(car.getYear(), car.getColor(), car.getEngine())
}
