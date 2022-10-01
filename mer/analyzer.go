package mer

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

func NewAnalyzer(filename string) (*mer, error) {
	fmt.Println("Creating new analyzer")
	mer := &mer{
		filename: filename,
		rows:     []Row{},
	}
	if err := mer.loadKillDump(); err != nil {
		return nil, err
	}

	return mer, nil
}

func (m *mer) Analyze() error {
	// solarsystems, _ := sde.NewFromJSON("solarsystems.json")
	fmt.Println(len(m.rows))
	return nil
}

func (m *mer) loadKillDump() error {
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
