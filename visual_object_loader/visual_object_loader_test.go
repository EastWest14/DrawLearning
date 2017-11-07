package visual_object_loader

import (
	"testing"
)

func TestLoadVisualObject(t *testing.T) {
	nonexistPath := "./nonexist.xml"
	visObj, err := LoadVisualObject(nonexistPath)
	if err == nil {
		t.Errorf("Expected non-nil error for nonexisting filepath, got nil")
	}
	if visObj != nil {
		t.Errorf("Expected visual object descriptor for non-existing filepath to be nil, got %v", visObj)
	}

	badFormatPath := "../test_files/visual_object_descriptors/invalid.xml"
	visObj, err = LoadVisualObject(badFormatPath)
	if err == nil {
		t.Errorf("Expected non-nil error for invalid format filepath, got nil")
	}
	if visObj != nil {
		t.Errorf("Expected visual object descriptor for invalid format filepath to be nil, got %v", visObj)
	}

	validFormatPath := "../test_files/visual_object_descriptors/valid_visual_object.xml"
	visObj, err = LoadVisualObject(validFormatPath)
	if err != nil {
		t.Errorf("Expected a nil error for valid format filepath, got %s", err.Error())
	}
	if visObj == nil {
		t.Errorf("Expected a non-nil descriptor for valid format filepath, got nil")
		return
	}
	if visObj.Color != true {
		t.Errorf("Expected color to be %v, got %v", true, visObj.Color)
	}
}

const (
	VALID_FORMAT = `<?xml version = "1.0" encoding = "utf-8"?>
<VisualObject>
   <Color>true</Color>
</VisualObject>`

	VALID_FORMAT_II = `<?xml version = "1.0" encoding = "utf-8"?>
<VisualObject>
   <Color>false</Color>
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
		expectedColor bool
	}{
		//Error cases
		{[]byte(INVALID_FORMAT), true, false},
		{[]byte(INVALID_FORMAT_II), true, false},
		{[]byte{}, true, false},
		{nil, true, false},
		//Valid cases
		{[]byte(VALID_FORMAT), false, true},
		{[]byte(VALID_FORMAT_II), false, false},
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
		if visObj.Color != aCase.expectedColor {
			t.Errorf("Error in case %d. Expected color %v, got %v", i, aCase.expectedColor, visObj.Color)
		}
	}
}
