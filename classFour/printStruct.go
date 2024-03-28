package main

import "fmt"

type Funcionario struct {
	name   string
	age    int
	job    string
	salary float32
}

func main() {
	funcionarioUm := Funcionario{
		name:   "Alexia",
		age:    21,
		job:    "Dev",
		salary: 9000,
	}
	funcionarioDois := Funcionario{
		name:   "Camila",
		age:    21,
		job:    "Dev",
		salary: 9000,
	}
	funcionarioTres := Funcionario{
		name:   "Paula",
		age:    21,
		job:    "Dev",
		salary: 9000,
	}
	fmt.Printf("%+v\n", funcionarioUm)
	fmt.Printf("%+v\n", funcionarioDois.job)
	fmt.Printf("%+v", funcionarioTres.job)

}
