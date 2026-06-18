package main

import (
	"fmt"
	"log"
)

type Todo struct {
	completed bool
}

type TodoList struct {
	tList map[string]Todo
}

func (t TodoList) MarkComplete(entry string) {
	en, ok := t.tList[entry]
	if !ok {
		log.Println(entry+" is not in the todo list", en)
		return
	}
	t.tList[entry] = Todo{
		completed: true,
	}
}

func (t TodoList) MarkIncomplete(entry string) {
	en, ok := t.tList[entry]
	if !ok {
		log.Println(entry+" is not in the todo list", en)
	}
	t.tList[entry] = Todo{
		completed: false,
	}
}

func main() {
	todo := TodoList{
		tList: map[string]Todo{
			"Get food": {
				completed: false,
			},
			"Exercise": {
				completed: true,
			},
			"Program": {
				completed: false,
			},
			"Go Sledding": {
				completed: true,
			},
			"Eat pie": {
				completed: false,
			},
		},
	}

	todo.MarkComplete("Get food")
	todo.MarkIncomplete("Laugh")
	fmt.Println(todo)
}
