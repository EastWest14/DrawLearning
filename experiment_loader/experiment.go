package experiment_loader

import (
	"DrawLearning/utils"
	"encoding/xml"
	"fmt"
)

type ExperimentDescriptor struct {
	XMLName            xml.Name `xml:"ExperimentDescriptor"`
	Name               string   `xml:"Name"`
	NumberOfGeneration int      `xml:"NumberOfGenerations"`
}

func LoadExperimentDescriptor(filepath string) (*ExperimentDescriptor, error) {
	content, err := utils.LoadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Error loading experiment descriptor from file: %s", err.Error())
	}
	return unmarshalDescriptor(content)
}

func unmarshalDescriptor(content []byte) (*ExperimentDescriptor, error) {
	descr := &ExperimentDescriptor{}
	err := xml.Unmarshal(content, &descr)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling experiment descriptor: %s", err.Error())
	}
	return descr, nil
}
