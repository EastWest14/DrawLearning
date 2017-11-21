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
		//Equal
		{nil, nil, true, ""},
		{NewVisualObject(&Color{Red: 255}), NewVisualObject(&Color{Red: 255}), true, ""},
		//Not equal
		{nil, NewVisualObject(&Color{}), false, "Visual Object I is nil, but not Visual Object II"},
		{NewVisualObject(nil), NewVisualObject(&Color{}), false, "Color not equal"},
		{NewVisualObject(&Color{Red: 255}), NewVisualObject(&Color{Red: 254}), false, "Visual objects' colors not equal: Color I red is 255, Color II red is 254"},
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
