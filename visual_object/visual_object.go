package visual_object

type VisualObject struct {
	objectColor *Color
}

func NewVisualObject(objectColor *Color) *VisualObject {
	return &VisualObject{
		objectColor: objectColor,
	}
}

func (v1 *VisualObject) Equal(v2 *VisualObject) (equal bool, inequalityMessage string) {
	if v1 == nil || v2 == nil {
		if v1 == nil && v2 == nil {
			return true, ""
		}
		if v1 == nil {
			return false, "Visual Object I is nil, but not Visual Object II"
		}
		return false, "Visual Object II is nil, but not Visual Object I"
	}

	if v1.objectColor == nil || v2.objectColor == nil {
		if v1.objectColor == nil && v2.objectColor == nil {
			return true, ""
		}
		return false, "Color not equal"
	}

	if equal, inequalityMsg := v1.objectColor.Equal(v2.objectColor); !equal {
		return false, "Visual objects' colors not equal: " + inequalityMsg
	}
	return true, ""
}
