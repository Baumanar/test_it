package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMowerMove(t *testing.T) {
	tests := []struct {
		name        string
		field       *Field
		instruction Instruction
		orientation Orientation
		initPos     Position
		wantPos     Position
	}{
		{
			name: "test move north",
			field: &Field{
				MowerPositions: map[Position]bool{},
				MaxPos:         Position{5, 5},
			},
			instruction: Forward,
			orientation: North,
			initPos:     Position{1, 1},
			wantPos:     Position{1, 2},
		},
		{
			name: "test move west",
			field: &Field{
				MowerPositions: map[Position]bool{},
				MaxPos:         Position{5, 5},
			},
			instruction: Forward,
			orientation: West,
			initPos:     Position{1, 1},
			wantPos:     Position{0, 1},
		},
		{
			name: "test move south",
			field: &Field{
				MowerPositions: map[Position]bool{},
				MaxPos:         Position{5, 5},
			},
			instruction: Forward,
			orientation: South,
			initPos:     Position{1, 1},
			wantPos:     Position{1, 0},
		},
		{
			name: "test move east",
			field: &Field{
				MowerPositions: map[Position]bool{},
				MaxPos:         Position{5, 5},
			},
			instruction: Forward,
			orientation: East,
			initPos:     Position{1, 1},
			wantPos:     Position{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mower := &Mower{
				Orientation: tt.orientation,
				Position:    tt.initPos,
			}
			tt.field.MowerPositions[tt.initPos] = true
			err := mower.execInstruction(tt.instruction, tt.field)
			assert.NoError(t, err)
			assert.Equal(t, mower.Position, tt.wantPos)
			_, ok := tt.field.MowerPositions[tt.wantPos]
			assert.True(t, ok)
			_, ok = tt.field.MowerPositions[tt.initPos]
			assert.False(t, ok)
		})
	}
}

func TestMowerMoveMaxPos(t *testing.T) {
	tests := []struct {
		name        string
		field       *Field
		instruction Instruction
		orientation Orientation
		initPos     Position
	}{
		{
			name: "test move north",
			field: &Field{
				MowerPositions: map[Position]bool{},
				MaxPos:         Position{5, 5},
			},
			instruction: Forward,
			orientation: North,
			initPos:     Position{2, 5},
		},
		{
			name: "test move west",
			field: &Field{
				MowerPositions: map[Position]bool{},
				MaxPos:         Position{5, 5},
			},
			instruction: Forward,
			orientation: West,
			initPos:     Position{0, 2},
		},
		{
			name: "test move south",
			field: &Field{
				MowerPositions: map[Position]bool{},
				MaxPos:         Position{5, 5},
			},
			instruction: Forward,
			orientation: South,
			initPos:     Position{2, 0},
		},
		{
			name: "test move east",
			field: &Field{
				MowerPositions: map[Position]bool{},
				MaxPos:         Position{5, 5},
			},
			instruction: Forward,
			orientation: East,
			initPos:     Position{5, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mower := &Mower{
				Orientation: tt.orientation,
				Position:    tt.initPos,
			}
			tt.field.MowerPositions[tt.initPos] = true
			err := mower.execInstruction(tt.instruction, tt.field)
			assert.NoError(t, err)
			assert.Equal(t, mower.Position, tt.initPos)
			_, ok := tt.field.MowerPositions[tt.initPos]
			assert.True(t, ok)
		})
	}
}

func TestMowerResolve(t *testing.T) {
	tests := []struct {
		name         string
		field        *Field
		instructions []Instruction
		orientation  Orientation
		initPos      Position
		endString    string
	}{
		{
			name: "test move example 1",
			field: &Field{
				MowerPositions: map[Position]bool{},
				MaxPos:         Position{5, 5},
			},
			instructions: []Instruction{Left, Forward, Left, Forward, Left,
				Forward, Left, Forward, Forward},
			orientation: North,
			initPos:     Position{1, 2},
			endString:   "1 3 N",
		},
		{
			name: "test move example 2",
			field: &Field{
				MowerPositions: map[Position]bool{},
				MaxPos:         Position{5, 5},
			},
			instructions: []Instruction{Forward, Forward, Right, Forward,
				Forward, Right, Forward, Right, Right, Forward},
			orientation: East,
			initPos:     Position{3, 3},
			endString:   "5 1 E",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mower := &Mower{
				Orientation:  tt.orientation,
				Position:     tt.initPos,
				Instructions: tt.instructions,
			}
			tt.field.MowerPositions[tt.initPos] = true
			err := mower.Resolve(tt.field)
			assert.NoError(t, err)
			assert.Equal(t, tt.endString, mower.ToString())
		})
	}
}
