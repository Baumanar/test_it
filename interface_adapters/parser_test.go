package interface_adapters

import (
	"github.com/Baumanar/test_it/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fileParser_Parse(t *testing.T) {

	tests := []struct {
		name        string
		fileName    string
		wantField   *models.Field
		wantMowers  []*models.Mower
		wantLengths []int
		wantErr     error
	}{
		{
			name:     "test ok",
			fileName: "../fixtures/example.txt",
			wantField: &models.Field{
				MaxPos:         models.Position{X: 5, Y: 5},
				MowerPositions: make(map[models.Position]bool),
			},
			wantMowers: []*models.Mower{
				{
					Orientation: models.North,
					Position:    models.Position{X: 1, Y: 2},
					Instructions: []models.Instruction{models.Left, models.Forward, models.Left, models.Forward, models.Left,
						models.Forward, models.Left, models.Forward, models.Forward},
				},
				{
					Orientation: models.East,
					Position:    models.Position{X: 3, Y: 3},
					Instructions: []models.Instruction{models.Forward, models.Forward, models.Right, models.Forward,
						models.Forward, models.Right, models.Forward, models.Right, models.Right, models.Forward},
				},
			},
			wantLengths: []int{9, 10},
		},
		{
			name:        "test two mowers same pos",
			fileName:    "../fixtures/two_mowers_same_pos.txt",
			wantField:   nil,
			wantMowers:  nil,
			wantLengths: nil,
			wantErr:     MowerNotUniqueErr,
		},
		{
			name:        "test mower not in field",
			fileName:    "../fixtures/mower_not_in_field.txt",
			wantField:   nil,
			wantMowers:  nil,
			wantLengths: nil,
			wantErr:     NotInFieldErr,
		},
		{
			name:        "test mower not in field negative",
			fileName:    "../fixtures/mower_not_in_field2.txt",
			wantField:   nil,
			wantMowers:  nil,
			wantLengths: nil,
			wantErr:     NotInFieldErr,
		},
		{
			name:        "invalid position",
			fileName:    "../fixtures/invalid.txt",
			wantField:   nil,
			wantMowers:  nil,
			wantLengths: nil,
			wantErr:     InvalidPositionErr,
		},
		{
			name:        "invalid position in mower",
			fileName:    "../fixtures/invalid.txt",
			wantField:   nil,
			wantMowers:  nil,
			wantLengths: nil,
			wantErr:     InvalidPositionErr,
		},
		{
			name:        "invalid position in field",
			fileName:    "../fixtures/invalid2.txt",
			wantField:   nil,
			wantMowers:  nil,
			wantLengths: nil,
			wantErr:     InvalidPositionErr,
		},
		{
			name:        "invalid orientation for mower",
			fileName:    "../fixtures/invalid3.txt",
			wantField:   nil,
			wantMowers:  nil,
			wantLengths: nil,
			wantErr:     InvalidOrientationErr,
		},
		{
			name:        "empty file",
			fileName:    "../fixtures/empty.txt",
			wantField:   nil,
			wantMowers:  nil,
			wantLengths: nil,
			wantErr:     EmptyFileErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FileParser{}
			gotField, gotMowers, err := f.Parse(tt.fileName)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantField.MaxPos, gotField.MaxPos)
				for i, prog := range gotMowers {
					assert.Equal(t, tt.wantLengths[i], len(prog.Instructions))
				}
			}
		})
	}
}
