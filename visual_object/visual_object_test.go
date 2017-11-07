package visual_object

import (
	"testing"
)

func TestVisualObjectsEqual(t *testing.T) {
	cases := []struct {
		objectOne                 *VisualObject
		objectTwo                 *VisualObject
		expectedEqual             bool
		expectedInequalityMessage string
	}{
		{nil, NewVisualObject(&Color{true}), false, "Visual Object I is nil, but not Visual Object II"},
		{NewVisualObject(&Color{false}), NewVisualObject(&Color{true}), false, "Color not equal"},
		{nil, nil, true, ""},
		{NewVisualObject(&Color{true}), NewVisualObject(&Color{true}), true, ""},
	}

	for i, aCase := range cases {
		equal, inequalityMessage := aCase.objectOne.Equal(aCase.objectTwo)
		if equal != aCase.expectedEqual {
			t.Errorf("Error in case %d. Expected equal %v, got %v", i, aCase.expectedEqual, equal)
		}
		if inequalityMessage != aCase.expectedInequalityMessage {
			t.Errorf("Error in case %d. Expected inequalty message %s, got %s", i, aCase.expectedInequalityMessage, inequalityMessage)
		}
	}
}
