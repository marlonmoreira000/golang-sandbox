package main

import "fmt"

type Point struct {
	x float32
	y float32
}

type Vehicle struct {
	velocity float32
	point Point
}

type Spaceship interface {
	fly()
	land()
	position() Point
}

func (this *Vehicle) fly() {
	this.velocity = 10
}

func (this *Vehicle) land() {
	this.velocity = 0
}

func (this Vehicle) position() {
}

func main() {
	v := Vehicle{ velocity: 0, point: Point{x: 0, y: 0}}
	v.fly()
	fmt.Println(v.velocity)
	v.land()
	fmt.Println(v.velocity)
}

