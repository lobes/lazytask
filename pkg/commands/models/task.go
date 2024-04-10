package models

import (
		"fmt"
		"time"
		"github.com/BurntSushi/toml" 
)

var taskToml = `		
uuid = "some-generated-uuid" 
description = "iam invincible!"
priority = "super-high"
tags = ["boris","fist-pump"]
status = "slugheads"
due_on = 1777-07-17 07:17:17Z
parents = [
  "parent1-uuid",
  "parent2-uuid",
] 
`

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
	if _, err := toml.DecodeFile("sisiphus/tasks/template.toml", &testicles); err != nil {
		fmt.Printf("Error decoding task from toml: %s\n", err)
	}
}
