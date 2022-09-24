package mer

import (
	"encoding/json"
	"os"

	"github.com/gocarina/gocsv"
)

type Row struct {
	VictimCorporationID   int     `csv:"victimCorporationID" json:"victimCorporationId"`
	VictimCorporationName string  `csv:"victimCorp" json:"victimCorporationName"`
	VictimAllianceName    string  `csv:"victimAlliance" json:"victimAllianceName"`
	KillerCorporationID   int     `csv:"finalCorporationID" json:"killerCorporationId"`
	KillerCorporationName string  `csv:"finalCorp" json:"killerCorporationName"`
	KillerAllianceName    string  `csv:"finalAlliance" json:"killerAllianceName"`
	VictimShipTypeID      int     `csv:"destroyedShipTypeID" json:"victimShipTypeId"`
	VictimShipTypeName    string  `csv:"destroyedShipType" json:"victimShipTypeName"`
	VictimShipGroupName   string  `csv:"destroyedShipGroup" json:"victimShipGroupName"`
	KillmailTime          string  `csv:"killTime" json:"killmailTime"`
	SolarSystemID         int     `csv:"solarSystemID" json:"solarSystemId"`
	SolarSystemName       string  `csv:"solarSystemName" json:"solarSystemName"`
	RegionID              int     `csv:"regionID" json:"regionId"`
	RegionName            string  `csv:"regionName" json:"regionName"`
	TotalValue            int     `csv:"iskLost" json:"totalValue"`
	DestroyedValue        int     `csv:"iskDestroyed" json:"destroyedValue"`
	BountyClaimed         float64 `csv:"bountyClaimed" json:"bountyClaimed"`
}

type mer struct {
	filename string
	rows     []Row
}

func New(filename string) *mer {
	return &mer{
		filename: filename,
		rows:     []Row{},
	}
}

func (m *mer) SaveCSV(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return gocsv.MarshalFile(&m.rows, file)
}

func (m *mer) SaveJSON(filename string) error {
	data, err := json.Marshal(m.rows)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, os.ModePerm)
}
