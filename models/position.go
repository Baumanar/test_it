package models



type Instruction string
type Orientation string


const (
	North   Orientation = "N"
	West    Orientation = "W"
	South   Orientation = "S"
	East    Orientation = "E"

	Forward Instruction = "F"
	Left    Instruction = "L"
	Right   Instruction = "R"
)

func (i Instruction) IsValid() bool {
	switch i {
	case Forward, Left, Right:
		return true
	}
	return false
}

func (o Orientation) IsValid() bool {
	switch o {
	case North, West, South, East:
		return true
	}
	return false
}


type Position struct {
	X int
	Y int
}

func (p Position) IsAbovePos(position Position) bool {
	if p.Y > position.Y || p.X > position.X {
		return true
	}
	return false
}

func (p Position) IsPositive() bool {
	if p.Y < 0 || p.X < 0 {
		return false
	}
	return true
}

func NewPosition(x, y int) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}



var rotationsLeft = map[Orientation]Orientation{
	North: West,
	West:  South,
	South: East,
	East:  North,
}

var rotationsRight = map[Orientation]Orientation{
	North: East,
	East:  South,
	South: West,
	West:  North,
}
