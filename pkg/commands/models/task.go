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
	Status      string
	DueOn       time.Time
	Parents     []string
}

func (t Task) Error() string {
	return fmt.Sprintf(`
    "Task Error: [%s] %s 
    Priority: %s
    Status: %s`, t.Uuid, t.Description, t.Priority, t.Status)

}

func ReadTask() {
	var testicles Task
	if _, err := toml.DecodeFile("sisiphus/tasks/template.toml", &testicles); err != nil {
		fmt.Printf("Error decoding task from toml: %s\n", err)
	}
	fmt.Printf("description: %s\n", testicles.Description)
}

func (t *Task) SetStatus(newStatus string) error {
	// Here, you could enforce rules. For example:
	// - Disallow moving away from StatusDone
	// - Define valid transitions if needed
	if t.Status == Done && newStatus != Done {
		return fmt.Errorf("invalid transition from %s to %s", Done, newStatus)
	}
	t.Status = newStatus
	return nil
}

func (t Task) CreateTask(data Task) error {
	return Task{
		Uuid:        uuid.NewString(),
		Description: data.Description,
	}
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
