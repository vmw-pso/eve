package mer

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/vmw-pso/eve/sde"
)

type Stats struct {
	TotalKills    int `json:"totalKills"`
	HighsecKills  int `json:"highsecKills"`
	LowsecKills   int `json:"lowsecKills"`
	NullsecKills  int `json:"nullsecKills"`
	WormholeKills int `json:"wormholeKills"`
	AbyssalKills  int `json:"abyssalKills"`
	VoidKills     int `json:"voidKills"`
}

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

func (m *mer) Analyze() (Stats, error) {
	fmt.Println("Analyzing...")
	solarsystems, _ := sde.NewFromJSON("solarsystems.json")
	highsec := solarsystems.HighsecSystems()
	lowsec := solarsystems.LowsecSystems()
	nullsec := solarsystems.NullsecSystems()
	wormholes := solarsystems.JSpaceSystems()
	abyssal := solarsystems.AbyssalSystems()
	// void := solarsystems.VoidSystems()
	stats := Stats{
		TotalKills:    len(m.rows),
		HighsecKills:  m.Kills(&highsec),
		LowsecKills:   m.Kills(&lowsec),
		NullsecKills:  m.Kills(&nullsec),
		WormholeKills: m.Kills(&wormholes),
		AbyssalKills:  m.Kills(&abyssal),
	}
	return stats, nil
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
