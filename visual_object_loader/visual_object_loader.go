package visual_object_loader

import (
	"DrawLearning/utils"
	vis_obj "DrawLearning/visual_object"
	"encoding/xml"
	"fmt"
)

type VisualObject_ struct {
	XMLName xml.Name `xml:"VisualObject"`
	Color   *Color_  `xml:"Color"`
}

type Color_ struct {
	Red   uint64 `xml:"Red"`
	Green uint64 `xml:"Green"`
	Blue  uint64 `xml:"Blue"`
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

	return vis_obj.NewVisualObject(&vis_obj.Color{
		Red:   visObj.Color.Red,
		Green: visObj.Color.Green,
		Blue:  visObj.Color.Blue,
	})
}
