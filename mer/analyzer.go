package mer

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/vmw-pso/eve/sde"
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
	solarsystems, _ := sde.NewFromJSON("solarsystems.json")
	highsec := solarsystems.HighsecSystems()
	lowsec := solarsystems.LowsecSystems()
	nullsec := solarsystems.NullsecSystems()
	wormholes := solarsystems.JSpaceSystems()
	abyssal := solarsystems.AbyssalSystems()
	fmt.Printf("Total kills: %d\n", len(m.rows))
	fmt.Printf("Highsec Kills: %d\n", m.Kills(&highsec))
	fmt.Printf("Lowsec Kills: %d\n", m.Kills(&lowsec))
	fmt.Printf("Nullsec Kills: %d\n", m.Kills(&nullsec))
	fmt.Printf("Wormhole Kills: %d\n", m.Kills(&wormholes))
	fmt.Printf("Abyssal Kills: %d\n", m.Kills(&abyssal))
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

func (m *mer) Kills(solarsystems *sde.Cluster) int {
	totalKills := 0
	for _, row := range m.rows {
		for _, system := range solarsystems.SolarSystems {
			if row.SolarSystemID == system.SolarSystemID {
				totalKills++
			}
		}
	}
	return totalKills
}
