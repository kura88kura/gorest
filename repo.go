package main

import "fmt"

var currentId int

var todos Todos
var users Users

func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
	RepoCreateUser(User{UserName: "Thomas John"})
	RepoCreateUser(User{UserName: "Bashua Doris"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}

func RepoFindUser(id int) User {
	for _, t := range users {
		if t.Id == id {
			return t
		}
	}

	return User{}
}

func RepoCreateUser(t User) User {
	currentId += 1
	t.Id = currentId
	users = append(users, t)
	return t
}

func RepoDestroyUser(id int) error {
	for i, t := range users {
		if t.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}