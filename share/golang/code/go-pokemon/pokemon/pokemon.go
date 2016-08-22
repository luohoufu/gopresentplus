package pokemon

type Type int

const (
	Normal Type = iota
	Water
	Fire
	Grass
	Bug
	Steel
	Fairy
	numTypes int = iota
)

type weaknessMap map[Type][]Type

var wm = weaknessMap{
	Normal: []Type{Steel},
	Fire:   []Type{Normal},
	Grass:  []Type{Fire, Bug},
}

func (t Type) WeakAgainst(att Type) bool {
	w, ok := wm[t]
	if !ok {
		return false
	}

	for _, v := range w {
		if v == att {
			return true
		}
	}

	return false
}

func (t Type) String() string {
	switch t {
	case Normal:
		return "Normal"
	case Water:
		return "Water"
	case Fire:
		return "Fire"
	case Bug:
		return "Bug"
	case Steel:
		return "Steel"
	case Fairy:
		return "Fairy"
	default:
		return ""
	}
	return ""
}
