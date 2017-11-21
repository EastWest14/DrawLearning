package visual_object_loader

import (
	vis_obj "DrawLearning/visual_object"
	"testing"
)

func TestLoadVisualObject(t *testing.T) {
	nonexistPath := "./nonexist.xml"
	visObj, err := LoadVisualObject(nonexistPath)
	if err == nil {
		t.Errorf("Expected non-nil error for nonexisting filepath, got nil")
	}
	if equal, ineqMessage := visObj.Equal(nil); !equal {
		t.Errorf("Expected visual object descriptor for non-existing filepath to be nil, got inequality message: %s", ineqMessage)
	}

	badFormatPath := "../test_files/visual_object_descriptors/invalid.xml"
	visObj, err = LoadVisualObject(badFormatPath)
	if err == nil {
		t.Errorf("Expected non-nil error for invalid format filepath, got nil")
	}
	if equal, ineqMessage := visObj.Equal(nil); !equal {
		t.Errorf("Expected visual object descriptor for invalid format filepath to be nil, got inequality message: %s", ineqMessage)
	}

	validFormatPath := "../test_files/visual_object_descriptors/valid_visual_object.xml"
	visObj, err = LoadVisualObject(validFormatPath)
	if err != nil {
		t.Errorf("Expected a nil error for valid format filepath, got %s", err.Error())
	}
	expectedVisObj := vis_obj.NewVisualObject(&vis_obj.Color{Red: 123, Green: 234, Blue: 99})
	if equal, ineqMessage := visObj.Equal(expectedVisObj); !equal {
		t.Errorf("Expected visual object descriptor for invalid format filepath to be nil, got inequality message: %s", ineqMessage)
	}
}

const (
	VALID_FORMAT = `<?xml version = "1.0" encoding = "utf-8"?>
<VisualObject>
   <Color>
   	<Red>123</Red>
   	<Green>234</Green>
   	<Blue>99</Blue>
   </Color>
</VisualObject>`

	VALID_FORMAT_II = `<?xml version = "1.0" encoding = "utf-8"?>
<VisualObject>
   <Color>
   	<Green>210</Green>
   </Color>
</VisualObject>`

	INVALID_FORMAT = `<?xml version = "1.0" encoding = "utf-8"?>
<Dog>
   <Name>Dog_Name</Name>
</Dog>`

	INVALID_FORMAT_II = `Not_XML`
)

func TestUnmarshalVisualObject(t *testing.T) {
	cases := []struct {
		content       []byte
		expectedError bool
		expectedColor *Color_
	}{
		//Error cases
		{[]byte(INVALID_FORMAT), true, nil},
		{[]byte(INVALID_FORMAT_II), true, nil},
		{[]byte{}, true, nil},
		{nil, true, nil},
		//Valid cases
		{[]byte(VALID_FORMAT), false, &Color_{Red: 123, Green: 234, Blue: 99}},
		{[]byte(VALID_FORMAT_II), false, &Color_{Green: 210}},
	}

	for i, aCase := range cases {
		visObj, err := unmarshalVisualObject(aCase.content)
		if (err != nil) != aCase.expectedError {
			t.Errorf("Error in case %d. Expected error %v != actual error %v", i, aCase.expectedError, (err != nil))
		}
		if err != nil {
			if visObj != nil {
				t.Errorf("Error in case %d. If error is not nil, Visual Object should be nil. Got: %v", i, visObj)
			}
			continue
		}
		if visObj.Color == nil || aCase.expectedColor == nil {
			if visObj.Color == nil && aCase.expectedColor != nil {
				t.Errorf("Error in case %d. Expected color %v, got nil", i, aCase.expectedColor)
				continue
			}
			if visObj.Color != nil && aCase.expectedColor == nil {
				t.Errorf("Error in case %d. Expected nil color, got %v", i, visObj.Color)
				continue
			}
		}
		if *visObj.Color != *aCase.expectedColor {
			t.Errorf("Error in case %d. Expected color %v, got %v", i, *aCase.expectedColor, *visObj.Color)
		}
	}
}

func TestConvertToDomainVisObject(t *testing.T) {
	visObjOne := vis_obj.NewVisualObject(&vis_obj.Color{Red: 123})
	visObjTwo := vis_obj.NewVisualObject(&vis_obj.Color{Red: 1, Green: 2, Blue: 3})

	cases := []struct {
		loadedObject              *VisualObject_
		expectedDomainObject      *vis_obj.VisualObject
		expectedEqual             bool
		expectedInequalityMessage string
	}{
		//Control case:
		{&VisualObject_{}, nil, false, "Visual Object II is nil, but not Visual Object I"},
		//Actual cases:
		{nil, nil, true, ""},
		{&VisualObject_{}, vis_obj.NewVisualObject(nil), true, ""},
		{&VisualObject_{Color: &Color_{Red: 123}}, visObjOne, true, ""},
		{&VisualObject_{Color: &Color_{Red: 1, Green: 2, Blue: 3}}, visObjTwo, true, ""},
	}

	for i, aCase := range cases {
		domainVisObj := convertToDomainVisObject(aCase.loadedObject)
		equal, inequalityMessage := domainVisObj.Equal(aCase.expectedDomainObject)
		if equal != aCase.expectedEqual {
			t.Errorf("Error in case %d. Expected equal %v, got %v", i, aCase.expectedEqual, equal)
		}
		if inequalityMessage != aCase.expectedInequalityMessage {
			t.Errorf("Error in case %d. Expected inequality message %s, got %s", i, aCase.expectedInequalityMessage, inequalityMessage)
		}
	}
}
