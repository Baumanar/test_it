package interface_adapters

import (
	"bufio"
	"errors"
	"github.com/Baumanar/test_it/models"
	"os"
	"strconv"
	"strings"
)

// Parses a file and returns a field and the list of mowers
// Makes sure the file is valid
type FileParser struct{}

var MowerNotUniqueErr = errors.New("found two mowers at the same initial position")
var InstructionsNotFoundErr = errors.New("no instructions found")
var InvalidFieldSizeErr = errors.New("file should include field size in first line")
var NotInFieldErr = errors.New("a mower is not in the field")
var EmptyFileErr = errors.New("empty file")
var InvalidInstructionErr = errors.New("invalid instruction")
var InvalidOrientationErr = errors.New("invalid orientation")
var InvalidPositionErr = errors.New("invalid position")

func (f FileParser) Parse(fileName string) (*models.Field, []*models.Mower, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	size, err := f.getFieldSize(scanner)
	if err != nil {
		return nil, nil, err
	}
	field := models.Field{
		MowerPositions: map[models.Position]bool{},
		MaxPos:         *size,
	}

	mowers := make([]*models.Mower, 0)
	positionSet := make(map[models.Position]bool)
	// Create new mower from file
	// Check if it's initial position is unique
	// And that its position is in the field
	for scanner.Scan() {
		mower, err := f.getMowerProgram(scanner)
		if err != nil {
			return nil, nil, err
		}
		if positionSet[mower.Position] {
			return nil, nil, MowerNotUniqueErr
		}
		positionSet[mower.Position] = true
		if mower.Position.IsAbovePos(field.MaxPos) || !mower.Position.IsPositive() {
			return nil, nil, NotInFieldErr
		}
		mowers = append(mowers, mower)
	}
	err = file.Close()
	if err != nil {
		return nil, nil, err
	}
	return &field, mowers, nil
}

// Parses first line to get the field size
func (FileParser) getFieldSize(scanner *bufio.Scanner) (*models.Position, error) {
	if scanner.Scan() {
		sizeTxt := strings.Split(scanner.Text(), " ")
		if len(sizeTxt) != 2 {
			return nil, InvalidFieldSizeErr
		}
		maxX, err := strconv.Atoi(sizeTxt[0])
		if err != nil {
			return nil, InvalidPositionErr
		}
		maxy, err := strconv.Atoi(sizeTxt[1])
		if err != nil {
			return nil, InvalidPositionErr
		}
		return models.NewPosition(maxX, maxy), nil
	}
	return nil, EmptyFileErr
}

// Parses other lines to the the mowers initial position and instructions
func (FileParser) getMowerProgram(scanner *bufio.Scanner) (*models.Mower, error) {
	sizeTxt := strings.Split(scanner.Text(), " ")
	if len(sizeTxt) != 3 {
		return nil, InvalidPositionErr
	}
	maxX, err := strconv.Atoi(sizeTxt[0])
	if err != nil {
		return nil, InvalidPositionErr
	}
	maxy, err := strconv.Atoi(sizeTxt[1])
	if err != nil {
		return nil, InvalidPositionErr
	}
	if !models.Orientation(sizeTxt[2]).IsValid() {
		return nil, InvalidOrientationErr
	}
	orientation := models.Orientation(sizeTxt[2])
	if !scanner.Scan() {
		return nil, InstructionsNotFoundErr
	}
	instructionsTxt := scanner.Text()
	instructions := make([]models.Instruction, len(instructionsTxt))
	for idx, char := range instructionsTxt {
		if !models.Instruction(char).IsValid() {
			return nil, InvalidInstructionErr
		}
		instructions[idx] = models.Instruction(char)
	}
	mower := models.Mower{
		Orientation:  orientation,
		Position:     *models.NewPosition(maxX, maxy),
		Instructions: instructions,
	}
	return &mower, nil
}
