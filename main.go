package main

import (
	"DrawLearning/experiment_loader"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Running")
	err := processArguments()
	if err != nil {
		panic(err)
	}
}

func processArguments() error {
	args := os.Args
	if len(args) < 2 {
		return fmt.Errorf("Expected %d command-line arguments, got %d arguments", 1, len(args)-1)
	}
	experimentDescriptorPath := args[1]
	fmt.Println(experimentDescriptorPath)

	_, err := experiment_loader.LoadExperimentDescriptor(experimentDescriptorPath)
	if err != nil {
		return err
	}
	return nil
}
