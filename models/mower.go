package models

import (
	"fmt"
	"sync"
)

// A Field keep information about each mower's position
// It is used by mowers to update their current position on the field
// The Field ensures two mowers don't collide: each mower uses the
// field's shared map to inform their position on the field
type Field struct {
	sync.Mutex
	MowerPositions map[Position]bool
	MaxPos         Position
}

func (m *Mower) checkAndUpdatePosition(oldPos, newPos Position, field *Field) {
	field.Lock()
	defer field.Unlock()
	_, ok := field.MowerPositions[newPos]
	if ok {
		return
	}
	delete(field.MowerPositions, oldPos)
	field.MowerPositions[newPos] = true
	m.Position = newPos
}

// A Mower has an orientation, a position and the list of
// instructions if will follow
type Mower struct {
	Orientation  Orientation
	Position     Position
	Instructions []Instruction
}

// Execute all instructions
func (m *Mower) Resolve(field *Field) error {
	for _, inst := range m.Instructions {
		err := m.execInstruction(inst, field)
		if err != nil {
			return err
		}
	}
	return nil
}

// Save mower's position in the field
func (m *Mower) Init(field *Field) {
	field.MowerPositions[m.Position] = true
}

// Execute one instruction
func (m *Mower) execInstruction(instruction Instruction, field *Field) error {
	if !instruction.IsValid() {
		return InvalidInstructionErr
	}
	if instruction == Left || instruction == Right {
		m.rotate(instruction)
		return nil
	}
	m.moveForward(field)
	return nil
}

// Format mower's position to a string
func (m *Mower) ToString() string {
	return fmt.Sprintf("%d %d %s", m.Position.X, m.Position.Y, m.Orientation)
}

// Update the mower position
// Before updating it, check
func (m *Mower) moveForward(field *Field) {

	var newPos Position
	switch m.Orientation {
	case North:
		newPos = *NewPosition(m.Position.X, m.Position.Y+1)
		if m.Position.Y < field.MaxPos.Y {
			m.checkAndUpdatePosition(m.Position, newPos, field)
		}
	case West:
		newPos = *NewPosition(m.Position.X-1, m.Position.Y)
		if m.Position.X > 0 {
			m.checkAndUpdatePosition(m.Position, newPos, field)
		}
	case South:
		newPos = *NewPosition(m.Position.X, m.Position.Y-1)
		if m.Position.Y > 0 {
			m.checkAndUpdatePosition(m.Position, newPos, field)
		}
	case East:
		newPos = *NewPosition(m.Position.X+1, m.Position.Y)
		if m.Position.X < field.MaxPos.X {
			m.checkAndUpdatePosition(m.Position, newPos, field)
		}
	}
}

func (m *Mower) rotate(instruction Instruction) {
	if instruction == Left {
		m.Orientation = rotationsLeft[m.Orientation]
	} else if instruction == Right {
		m.Orientation = rotationsRight[m.Orientation]
	}
}
