package mer

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

func NewAnalyzer(inputFilename string) *merNew {
	fmt.Println("Creating new analyzer")
	return &merNew{
		inputFilename: inputFilename,
		rows:          []NewRow{},
	}
}

func (m *mer) Analyze() error {
	// solarsystems, _ := sde.NewFromJSON("solarsystems.json")
	if err := m.loadMER(); err != nil {
		return err
	}
	fmt.Println(len(m.rows))
	return nil
}

func (m *mer) loadMER() error {
	file, err := os.OpenFile(m.filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	err = gocsv.Unmarshal(file, &m.rows)
	if err != nil {
		return err
	}
	return nil
}
