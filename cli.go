package main

import (
	"fmt"
	"log"
	"strconv"
)

type Todo struct {
	completed bool
}

type TodoList struct {
	tList map[string]Todo
}

// Needs to be a pointer because it needs to change the data
// Without the function removing what we just did in a copy
func (t *TodoList) AddTask(entry string) {
	//Handle empty map addition
	if t.tList == nil {
		t.tList = make(map[string]Todo)
	}
	//User should be able to add the same task more than once
	t.tList[entry] = Todo{
		completed: false,
	}
}

func (t *TodoList) RemoveTask(entry string) {
	//If it exists, delete it
	_, ok := t.tList[entry]
	if ok {
		delete(t.tList, entry)
	}
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
	todo := TodoList{}

	// for i := 0; i < 100; i++ {
	// 	todo.AddTask("Adding Task " + strconv.Itoa(i))
	// }

	for i := 0; i < 100; i++ {
		todo.RemoveTask("Adding Task " + strconv.Itoa(i))
	}

	fmt.Println(todo)
}
