package main

import (
	"log"
	"os"

	pro "github.com/ErlanRG/bluecpprint/internal"
)

func main() {
	if err := pro.CheckArgs(); err != nil {
		log.Fatal(err)
	}

	currentWorkingDir, err := os.Getwd()
	if err != nil {
		log.Printf("Could not get current working directory: %v", err)
		os.Exit(1)
	}

	p := pro.Project{
		ProjectName:  os.Args[1],
		AbsolutePath: currentWorkingDir,
	}

	if err := p.CreateProjectStructure(); err != nil {
		log.Fatalf("Error creating project structure: %v", err)
	}

    log.Printf("Project %s created successfully", p.ProjectName)
}
