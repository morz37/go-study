package Facade

import (
	"log"
	"strings"
)

type PedalGas struct {
}

func (p *PedalGas) Ride() string{
	return "Car rides"
}

type PedalBrake struct {
}

func (p *PedalBrake) Stop() string{
	return "The car stopped"
}

type SteeringWheel struct {
}

func (s *SteeringWheel) Turn() string {
	return "Car turn"
}

type Car struct {
	pedalGas 	*PedalGas
	pedalBrake 	*PedalBrake
	steeringWheel *SteeringWheel
}

func (c *Car) ToDo () string {
	result := []string{
		c.pedalGas.Ride(),
		c.pedalBrake.Stop(),
		c.steeringWheel.Turn(),
	}
	return strings.Join(result, "\n")
}

func NewCar() *Car {
	return &Car{
		pedalGas: &PedalGas{},
		pedalBrake:  &PedalBrake{},
		steeringWheel: &SteeringWheel{},
	}
}


func main () {
	car := NewCar()
	result := car.ToDo()

	log.Println(result)
}
