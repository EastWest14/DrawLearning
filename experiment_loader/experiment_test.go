package experiment_loader

import (
	"testing"
)

func TestLoadExperimentDescriptor(t *testing.T) {
	nonexistPath := "./nonexist.xml"
	descr, err := LoadExperimentDescriptor(nonexistPath)
	if err == nil {
		t.Errorf("Expected non-nil error for nonexisting filepath, got nil")
	}
	if descr != nil {
		t.Errorf("Expected descriptor for non-existing filepath to be nil, got %v", descr)
	}

	badFormatPath := "../test_files/invalid.xml"
	descr, err = LoadExperimentDescriptor(badFormatPath)
	if err == nil {
		t.Errorf("Expected non-nil error for invalid format filepath, got nil")
	}
	if descr != nil {
		t.Errorf("Expected descriptor for invalid format filepath to be nil, got %v", descr)
	}

	validFormatPath := "../test_files/SampleExperiment.xml"
	descr, err = LoadExperimentDescriptor(validFormatPath)
	if err != nil {
		t.Errorf("Expected a nil error for valid format filepath, got %s", err.Error())
	}
	if descr == nil {
		t.Errorf("Expected a non-nil descriptor for valid format filepath, got nil", descr)
	}
	if descr.Name != "Name_Value" {
		t.Errorf("Expected name to be %s, got %s", "Name_Value", descr.Name)
	}
}

const (
	VALID_FORMAT = `<?xml version = "1.0" encoding = "utf-8"?>
<Person>
   <Name>Name_Value</Name>
</Person>`

	VALID_FORMAT_II = `<?xml version = "1.0" encoding = "utf-8"?>
<Person>
   <Name>Name_Value_2</Name>
</Person>`

	INVALID_FORMAT = `<?xml version = "1.0" encoding = "utf-8"?>
<Dog>
   <Name>Dog_Name</Name>
</Dog>`

	INVALID_FORMAT_II = `Not_XML`
)

func TestUnmarshalDescriptor(t *testing.T) {
	cases := []struct {
		content       []byte
		expectedError bool
		expectedName  string
	}{
		//Error cases
		{[]byte(INVALID_FORMAT), true, ""},
		{[]byte(INVALID_FORMAT_II), true, ""},
		{[]byte{}, true, ""},
		{nil, true, ""},
		//Non-error cases
		{[]byte(VALID_FORMAT), false, "Name_Value"},
		{[]byte(VALID_FORMAT_II), false, "Name_Value_2"},
	}

	for i, aCase := range cases {
		descr, err := unmarshalDescriptor(aCase.content)
		if (err != nil) != aCase.expectedError {
			t.Errorf("Error in case %d. Expected error %v != actual error %v", i, aCase.expectedError, (err != nil))
		}
		if err != nil {
			if descr != nil {
				t.Errorf("Error in case %d. If error is not nil, descriptor should be nil. Got: %v", i, descr)
			}
			continue
		}
		if descr.Name != aCase.expectedName {
			t.Errorf("Error in case %d. Expected name %s, got %s", i, aCase.expectedName, descr.Name)
		}
	}
}
