package visual_object_loader

import (
	"DrawLearning/utils"
	"encoding/xml"
	"fmt"
)

type VisualObject_ struct {
	XMLName xml.Name `xml:"VisualObject"`
	Color   bool     `xml:"Color"`
}

func LoadVisualObject(filepath string) (*VisualObject_, error) {
	content, err := utils.LoadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Error loading visual object from file: %s", err.Error())
	}
	return unmarshalVisualObject(content)
}

func unmarshalVisualObject(content []byte) (*VisualObject_, error) {
	visObj := &VisualObject_{}
	err := xml.Unmarshal(content, &visObj)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling visual object: %s", err.Error())
	}
	return visObj, nil
}
