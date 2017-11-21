package visual_object

import (
	"fmt"
)

type Color struct {
	Red   uint64
	Green uint64
	Blue  uint64
}

func (c1 *Color) Equal(c2 *Color) (equal bool, inequalityMessage string) {
	if c1 == nil || c2 == nil {
		if c1 == nil && c2 == nil {
			return true, ""
		}
		if c1 == nil {
			return false, "Color I is nil, but nit Color II"
		}
		return false, "Color II is nil, but nit Color I"
	}

	if (*c1).Red != (*c2).Red {
		return false, fmt.Sprintf("Color I red is %d, Color II red is %d", (*c1).Red, (*c2).Red)
	}
	if (*c1).Green != (*c2).Green {
		return false, fmt.Sprintf("Color I green is %d, Color II green is %d", (*c1).Green, (*c2).Green)
	}
	if (*c1).Blue != (*c2).Blue {
		return false, fmt.Sprintf("Color I blue is %d, Color II blue is %d", (*c1).Blue, (*c2).Blue)
	}

	return true, ""
}
