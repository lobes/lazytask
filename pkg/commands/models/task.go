package models

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

type Uuid string

type Priority string

const (
	None  Priority = "tell'em he's dreamin'"
	Some  Priority = "maybe not today... but some day"
	Lots  Priority = "the most important thing"
	Heaps Priority = "the actual most important thing"
	All   Priority = "awareness"
)

type Status string

const (
	New   Status = "limbo"
	Later Status = "maybe not today... but some day"
	Next  Status = "batter up..."
	Now   Status = "getting after it"
	Done  Status = "made the thing do the thing"
)

// Task represents a task
type Task struct {
	Uuid        Uuid
	Description string
	Tags        []string
	Priority    Priority
	Status      Status
	DueOn       time.Time
	Parents     []Uuid
}

func readTask() {
	var testicles Task
	if _, err := toml.DecodeFile("sisiphus/tasks/template.toml", &testicles); err != nil {
		fmt.Printf("Error decoding task from toml: %s\n", err)
	}
	fmt.Printf("description: %s\n", testicles.Description)
}

// func createTask()
// is passed in a Task
// generates a uuid
// assigns it to the field uuid
// encodes it as <uuid>.toml
// err if a file with that name already exists

// func updateTask()
// is passed in a Task
// encodes it as <uuid>.toml and over-writes any existing file if there

// func deleteTask()
// is passed in a Uuid
