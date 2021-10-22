package service

import (
	"testing"

	"github.com/Baumanar/test_it/interface_adapters"
	"github.com/Baumanar/test_it/models"
	"github.com/stretchr/testify/assert"
)

func TestMowersRunnerRun(t *testing.T) {
	tests := []struct {
		name      string
		filename  string
		wantErr   bool
		endString string
	}{
		{
			name:      "test ok",
			filename:  "../fixtures/example.txt",
			endString: "1 3 N\n5 1 E\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mowersService{
				fileParser: interface_adapters.FileParser{},
			}

			err := m.Init(tt.filename)

			assert.NoError(t, err)
			err = m.Run()
			assert.NoError(t, err)
			assert.Equal(t, m.ResToString(), tt.endString)
		})
	}
}

func TestMowersRunnerRunCrossing(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{
			// A test where two mowers are on the same line. They should not
			// go in contact, but because of concurrency, we don't know which
			// one is going to be first
			name:     "test crossing mowers",
			filename: "../fixtures/crossing_paths.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mowersService{
				fileParser: interface_adapters.FileParser{},
			}
			err := m.Init(tt.filename)
			assert.NoError(t, err)
			err = m.Run()
			assert.NoError(t, err)
			// Check they did not cross
			assert.Greater(t, m.mowers[1].Position.X, m.mowers[0].Position.X)
			assert.Equal(t, m.mowers[0].Orientation, models.East)
			assert.Equal(t, m.mowers[1].Orientation, models.West)
		})
	}
}
