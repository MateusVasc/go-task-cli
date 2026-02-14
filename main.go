package main

import (
	"fmt"
	"math/rand"

	"github.com/MateusVasc/go-task-cli/domain"
)

func randBool() bool {
	isDone := rand.Intn(2)

	return isDone == 1
}

func main() {
	fmt.Println("Go CLI running")

	ids := [5]int{1, 2, 3, 4, 5}
	titles := [5]string{"Do this", "Do that", "Do those", "Do these", "Do all"}

	tasks := make([]domain.Task, 5)

	for i := 0; i < 5; i++ {
		tasks[i].ID = ids[i]
		tasks[i].Title = titles[i]
		tasks[i].Done = randBool()
	}

	fmt.Println(tasks)
}
