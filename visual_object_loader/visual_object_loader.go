package visual_object_loader

import (
	"DrawLearning/utils"
	vis_obj "DrawLearning/visual_object"
	"encoding/xml"
	"fmt"
)

type VisualObject_ struct {
	XMLName xml.Name `xml:"VisualObject"`
	Color   *bool    `xml:"Color"`
}

func LoadVisualObject(filepath string) (*vis_obj.VisualObject, error) {
	content, err := utils.LoadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Error loading visual object from file: %s", err.Error())
	}
	loadedObj, err := unmarshalVisualObject(content)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling visual object into struct: %s", err.Error())
	}
	return convertToDomainVisObject(loadedObj), nil
}

func unmarshalVisualObject(content []byte) (*VisualObject_, error) {
	visObj := &VisualObject_{}
	err := xml.Unmarshal(content, &visObj)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling visual object: %s", err.Error())
	}
	return visObj, nil
}

func convertToDomainVisObject(visObj *VisualObject_) *vis_obj.VisualObject {
	if visObj == nil {
		return nil
	}
	if visObj.Color == nil {
		return vis_obj.NewVisualObject(nil)
	}

	return vis_obj.NewVisualObject(&vis_obj.Color{*visObj.Color})
}
