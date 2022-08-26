package main

import (
	"fmt"
	"gorgonia.org/gorgonia"
	"log"
)

func main() {
	simpleAddition()
}

func simpleNeuralNetwork() {

}

func simpleAddition() {
	//creating an empty expression graph
	g := gorgonia.NewGraph()

	//creating nodes to associate to the expression graph
	var x, y, z *gorgonia.Node
	// error variable
	var err error

	// define the expression
	//creating variables with the appropriate node
	x = gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("x"))
	y = gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("y"))
	//this is going to do the addition of both the variables
	if z, err = gorgonia.Add(x, y); err != nil {
		log.Fatal(err)
	}

	// create a VM to run the program on
	machine := gorgonia.NewTapeMachine(g)
	defer machine.Close()

	// set initial values then run
	//assigning values to x and y
	gorgonia.Let(x, 2.0)
	gorgonia.Let(y, 2.5)

	// by default, LispMachine performs forward mode and backwards mode execution
	m := gorgonia.NewLispMachine(g)
	defer m.Close()
	if err = m.RunAll(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("z: %v\n", z.Value())

	//Autodiff showcases automatic differentiation
	if xgrad, err := x.Grad(); err == nil {
		fmt.Printf("dz/dx: %v\n", xgrad)
	}

	if ygrad, err := y.Grad(); err == nil {
		fmt.Printf("dz/dy: %v\n", ygrad)
	}
}
