package experiment_loader

import (
	"DrawLearning/utils"
	"encoding/xml"
	"fmt"
)

type ExperimentDescriptor struct {
	XMLName xml.Name `xml:"Person"`
	Name    string   `xml:"Name"`
}

func LoadExperimentDescriptor(filepath string) (*ExperimentDescriptor, error) {
	content, err := utils.LoadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Error loading experiment descriptor from file: %s", err.Error())
	}

	descr := &ExperimentDescriptor{}
	err = xml.Unmarshal(content, &descr)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling experiment descriptor: %s", err.Error())
	}
	fmt.Println(descr.Name)
	return nil, nil
}
