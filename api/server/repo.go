package main

import "fmt"

var currentId int
var children Children
var currentChoreId int //LOL I SUCK
var chores Chores

// Give us some seed data
func init() {
	RepoCreateChild(Child{Name: "Amy",Points: 4})
	RepoCreateChild(Child{Name: "Bill"})
	RepoCreateChore(Chore{Name: "Chore1", Points: 6})
}

func RepoFindChild(id int) Child {
	for _, c := range children {
		if c.Id == id {
			return c
		}
	}
	return Child{}
}

func RepoCreateChild(c Child) Child {
	currentId += 1
	c.Id = currentId
	children = append(children, c)
	return c
}

func RepoDestroyChild(id int) error {
	for i, c := range children {
		if c.Id == id {
			children = append(children[:i], children[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Child with id of %d to delete", id)
}

func RepoCreateChore(c Chore) Chore {
	currentChoreId += 1
	c.Id = currentChoreId
	chores = append(chores, c)
	return c
}

func RepoFindChore(id int) Chore {
	for _, c := range chores {
		if c.Id == id {
			return c
		}
	}
	
	// return empty Todo if not found
	return Chore{}
}

func RepoAddChoreToChild(childId int, choreId int) {
	for _, c := range children {
		if c.Id == childId {
			c.Chores = append(c.Chores, choreId)
		}
	}
}
	
