package main

import "fmt"

type Person struct{
	Name string
	Birthday string
}

type Worker struct{
	Person
	wage string
	score int
}

type Department struct{
	workers [] Worker
	Name string
}

func (person *Person) print ()  {
	fmt.Println(person.Name + " " + person.Birthday)
}

func (department *Department) print () {
	fmt.Println(department.Name)
}

func (department *Department) addWorker(worker Worker) {
	department.workers = append(department.workers, worker)
}

func (department *Department) getBestScoreWorker() (Worker, int) {
	max := 0
	var bestWorker Worker

	for i:= 0; i < len(department.workers); i++ {
		if department.workers[i].score > max {
			max = department.workers[i].score
			bestWorker = department.workers[i]
		}
	}

	return bestWorker, max
}

func main () {
	var workers [] Worker

	dep := Department{
		workers :workers,
		Name :"IT",
	}

	dep.addWorker(Worker{Person{"Axel", "10.10.1991"}, "200", 100})
	dep.addWorker(Worker{Person{"Müller", "10.10.1991"}, "100", 60})

	bestWorker, score := dep.getBestScoreWorker()

	fmt.Println("Best worker:", bestWorker.Name)
	fmt.Println("His score:", score)
}