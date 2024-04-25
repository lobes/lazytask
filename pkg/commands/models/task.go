package models

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/google/uuid"
)

type Priority string

const (
	None  Priority = "tell'em he's dreamin'"
	Some  Priority = "maybe not today... but some day"
	Lots  Priority = "the most important thing"
	Heaps Priority = "the actual most important thing"
	All   Priority = "awareness"
)

const (
	New   string = "limbo"
	Later string = "maybe not today... but some day"
	Next  string = "batter up..."
	Now   string = "getting after it"
	Done  string = "made the thing do the thing"
)

type Task struct {
	Uuid        string
	Description string
	Tags        []string
	Priority    Priority
	State       string
	DueOn       time.Time
	Parents     []string
}

func (t Task) Error() string {
	return fmt.Sprintf(`
    "Task Error: [%s] %s 
    Priority: %s
    State: %s`, t.Uuid, t.Description, t.Priority, t.State)

}

func ReadTask() []string {
	var fromDisk Task
	var readTasks []string
	if _, err := toml.DecodeFile("sisiphus/tasks/template.toml", &fromDisk); err != nil {
		fmt.Printf("Error decoding task from toml: %s\n", err)
	}

	readTasks = append(readTasks, fromDisk.Description)

	return readTasks
}

func (t *Task) SetState(newState string) error {
	// TODO: expand to the whole task state machine
	if t.State == Done && newState != Done {
		return fmt.Errorf("invalid transition from %s to %s", Done, newState)
	}
	t.State = newState
	return nil
}

func (t Task) CreateTask(data Task) error {
	// encodes it as <Uuid>.toml
	// err if a file with that name already exists
	return Task{
		Uuid:        uuid.NewString(),
		Description: data.Description,
	}
}

// func UpdateTask()
// is passed in a Task
// encodes it as <Uuid>.toml and over-writes any existing file if there

// func DeleteTask()
// is passed in a Uuid
