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

	badFormatPath := "../test_files/experiment_descriptors/invalid.xml"
	descr, err = LoadExperimentDescriptor(badFormatPath)
	if err == nil {
		t.Errorf("Expected non-nil error for invalid format filepath, got nil")
	}
	if descr != nil {
		t.Errorf("Expected descriptor for invalid format filepath to be nil, got %v", descr)
	}

	validFormatPath := "../test_files/experiment_descriptors/valid_experiment_descriptor.xml"
	descr, err = LoadExperimentDescriptor(validFormatPath)
	if err != nil {
		t.Errorf("Expected a nil error for valid format filepath, got %s", err.Error())
	}
	if descr == nil {
		t.Errorf("Expected a non-nil descriptor for valid format filepath, got nil")
		return
	}
	if descr.Name != "Name_Value" {
		t.Errorf("Expected name to be %s, got %s", "Name_Value", descr.Name)
	}
}

const (
	VALID_FORMAT = `<?xml version = "1.0" encoding = "utf-8"?>
<ExperimentDescriptor>
   <Name>Name_Value</Name>
   <NumberOfGenerations>1</NumberOfGenerations>
</ExperimentDescriptor>`

	VALID_FORMAT_II = `<?xml version = "1.0" encoding = "utf-8"?>
<ExperimentDescriptor>
   <Name>Name_Value_2</Name>
   <NumberOfGenerations>13</NumberOfGenerations>
</ExperimentDescriptor>`

	INVALID_FORMAT = `<?xml version = "1.0" encoding = "utf-8"?>
<Dog>
   <Name>Dog_Name</Name>
</Dog>`

	INVALID_FORMAT_II = `Not_XML`
)

func TestUnmarshalDescriptor(t *testing.T) {
	cases := []struct {
		content                  []byte
		expectedError            bool
		expectedName             string
		expectedNumOfGenerations int
	}{
		//Error cases
		{[]byte(INVALID_FORMAT), true, "", 0},
		{[]byte(INVALID_FORMAT_II), true, "", 0},
		{[]byte{}, true, "", 0},
		{nil, true, "", 0},
		//Valid cases
		{[]byte(VALID_FORMAT), false, "Name_Value", 1},
		{[]byte(VALID_FORMAT_II), false, "Name_Value_2", 13},
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
		if descr.NumberOfGeneration != aCase.expectedNumOfGenerations {
			t.Errorf("Error in case %d. Expected number of generations %d, got %d", i, aCase.expectedNumOfGenerations, descr.NumberOfGeneration)
		}
	}
}
