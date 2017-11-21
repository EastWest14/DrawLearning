package visual_object

import (
	"testing"
)

func TestColorEqual(t *testing.T) {
	cases := []struct {
		colorOne              *Color
		colorTwo              *Color
		expectedEqual         bool
		expectedInequalityMsg string
	}{
		//Equal
		{nil, nil, true, ""},
		{&Color{}, &Color{}, true, ""},
		{&Color{Red: 255, Green: 255, Blue: 255}, &Color{Red: 255, Green: 255, Blue: 255}, true, ""},
		{&Color{Red: 0, Green: 0, Blue: 0}, &Color{}, true, ""},
		//Not equal
		{nil, &Color{}, false, "Color I is nil, but nit Color II"},
		{&Color{Red: 255, Green: 255, Blue: 255}, nil, false, "Color II is nil, but nit Color I"},
		{&Color{Red: 255, Green: 255, Blue: 255}, &Color{Red: 255, Green: 255, Blue: 0}, false, "Color I blue is 255, Color II blue is 0"},
		{&Color{Red: 254, Green: 255, Blue: 255}, &Color{Red: 255, Green: 255, Blue: 255}, false, "Color I red is 254, Color II red is 255"},
		{&Color{Red: 255, Green: 100, Blue: 255}, &Color{Red: 255, Green: 0, Blue: 255}, false, "Color I green is 100, Color II green is 0"},
	}

	for i, aCase := range cases {
		equal, inequalityMessage := aCase.colorOne.Equal(aCase.colorTwo)
		if equal != aCase.expectedEqual {
			t.Errorf("Error in case %d. Expected equal %v, got %v", i, aCase.expectedEqual, equal)
		}
		if inequalityMessage != aCase.expectedInequalityMsg {
			t.Errorf("Error in case %d. Expected inequality message %s, got %s", i, aCase.expectedInequalityMsg, inequalityMessage)
		}
	}
}
