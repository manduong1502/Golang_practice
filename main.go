package main

import "fmt"


type Todo struct {
	task     string
	complete bool
}

func addTask(task string, tasks chan string) {
	tasks <- task // Gửi công việc vào kênh
}

func showTodo(todos []Todo) {
	for i, todo := range todos {
		status := "not completed" 
		if todo.complete {
			status = "complete" 
		}
		fmt.Printf("%d: %s - %s\n", i+1, todo.task, status) 
	}
}

func main() {
	name := "Golang Learner"
	fmt.Printf("Welcome, %s, to todo list app \n", name)


	todos := []Todo{
		{task: "Buy groceries", complete: false},
		{task: "Complete homework", complete: true},
		{task: "Read Golang docs", complete: false},
	}

	showTodo(todos)


	if len(todos) == 0 {
		fmt.Println("Your todo list is empty!")
	} else {
		fmt.Println("You have tasks to do!")
	}

	tasks := make(chan string)

	go addTask("Buy groceries", tasks)
	go addTask("Complete homework", tasks)

	for i := 0; i < 2; i++ {
		task := <-tasks
		fmt.Printf("Task added: %s\n", task)
	}
}
