package service

import (
	"fmt"
	"github.com/Baumanar/test_it/interface_adapters"
	"github.com/Baumanar/test_it/models"
	"strings"
	"sync"
)


// check benchmarks: discuss on shared map

// mowersService is the service that does the global job,
// parses the files, initializes mowers and runs them concurrently
type mowersService struct {
	fileParser interface_adapters.FileParser
	field      *models.Field
	mowers     []*models.Mower
}

func NewMowersService() *mowersService {
	return &mowersService{
		fileParser: interface_adapters.FileParser{},
	}
}

// Initialize the field with the mowers initial positions
func (m *mowersService) Init(filename string) error {
	field, mowers, err := m.fileParser.Parse(filename)
	if err != nil {
		return err
	}
	m.field = field
	for _, mower := range mowers {
		mower.Init(m.field)
	}
	m.mowers = mowers
	return nil
}

// Run the mowers programs concurrently
// And store the end positions of the mowers
func (m *mowersService) Run() error {
	var wg sync.WaitGroup
	for idx, currMower := range m.mowers {
		wg.Add(1)
		go func(mower *models.Mower, idx int) {
			_ = mower.Resolve(m.field)
			wg.Done()
		}(currMower, idx)
	}
	wg.Wait()
	return nil
}

// Formats end positions of all mowers to a string
func (m mowersService) ResToString() string {
	var resStr strings.Builder
	for _, mower := range m.mowers {
		resStr.WriteString(fmt.Sprintf("%s\n", mower.ToString()))
	}
	return resStr.String()
}
