package models

import (
		"fmt"
		"time"
		"github.com/BurntSushi/toml" 
)

// Task: A task
type Task struct {
	Uuid        string
	Description string
	Tags        []string
	Priority    string
	Status      string
	DueOn       time.Time
	Parents     []string
}

func printTask() {
	var testicles Task
	if _, err := toml.DecodeFile("sisiphus/tasks/template.toml", &testicles) 
		err != nil {
			fmt.Printf("Error decoding task from toml: %s\n", err)
		}
	fmt.Printf("description: %s\n", testicles.Description)
}
