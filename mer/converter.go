package mer

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

type NewRow struct {
	KillmailTime          string  `csv:"kill_datetime" json:"killmailTime"`
	KillerCorporationID   int     `csv:"killer_corporation_id" json:"killerCorporationId"`
	KillerCorporationName string  `csv:"killer_corporation_name" json:"killerCorporationName"`
	KillerAllianceName    string  `csv:"killer_alliance_name" json:"killerAllianceName"`
	VictimCorporationID   int     `csv:"victim_corporation_id" json:"victimCorporationId"`
	VictimCorporationName string  `csv:"victim_corporation_name" json:"victimCorporationName"`
	VictimAllianceName    string  `csv:"victim_alliance_name" json:"victimAllianceName"`
	VictimShipTypeID      int     `csv:"victim_ship_type_id" json:"victimShipTypeId"`
	VictimShipTypeName    string  `csv:"victim_ship_type_name" json:"victimShipTypeName"`
	VictimShipGroupName   string  `csv:"victim_ship_group_name" json:"victimShipGroupName"`
	SolarSystemID         int     `csv:"solarsystem_id" json:"solarSystemId"`
	SolarSystemName       string  `csv:"solarsystem_name" json:"solarSystemName"`
	RegionID              int     `csv:"region_id" json:"regionId"`
	RegionName            string  `csv:"region_name" json:"regionName"`
	TotalValue            int     `csv:"isk_lost" json:"totalValue"`
	DestroyedValue        int     `csv:"isk_destroyed" json:"destroyedValue"`
	BountyClaimed         float64 `csv:"bounty_claimed" json:"bountyClaimed"`
}

type merNew struct {
	inputFilename string
	rows          []NewRow
}

func NewConverter(inputFilename string) *merNew {
	fmt.Println("Creating new converter")
	return &merNew{
		inputFilename: inputFilename,
		rows:          []NewRow{},
	}
}

func (mn *merNew) Convert() error {
	input, err := os.OpenFile(mn.inputFilename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer input.Close()
	fmt.Printf("%s opened\n", mn.inputFilename)

	if err := gocsv.UnmarshalFile(input, &mn.rows); err != nil {
		return err
	}
	fmt.Println("content unmarshalled")

	mer := New("")
	for _, row := range mn.rows {
		item := Row{}
		item.VictimCorporationID = row.VictimCorporationID
		item.VictimCorporationName = row.VictimCorporationName
		item.VictimAllianceName = row.VictimAllianceName
		item.KillerCorporationID = row.KillerCorporationID
		item.KillerCorporationName = row.KillerCorporationName
		item.KillerAllianceName = row.KillerAllianceName
		item.VictimShipTypeID = row.VictimShipTypeID
		item.VictimShipTypeName = row.VictimShipTypeName
		item.VictimShipGroupName = row.VictimShipGroupName
		item.KillmailTime = row.KillmailTime
		item.SolarSystemID = row.SolarSystemID
		item.SolarSystemName = row.SolarSystemName
		item.RegionID = row.RegionID
		item.RegionName = row.RegionName
		item.TotalValue = row.TotalValue
		item.DestroyedValue = row.DestroyedValue
		item.BountyClaimed = row.BountyClaimed
		mer.rows = append(mer.rows, item)
	}

	csvFilename := fmt.Sprintf("%s_cnvrtd.csv", strings.Split(mn.inputFilename, ".")[0])
	if err = mer.SaveCSV(csvFilename); err != nil {
		return err
	}

	jsonFilename := fmt.Sprintf("%s.json", strings.Split(mn.inputFilename, ".")[0])
	if err = mer.SaveJSON(jsonFilename); err != nil {
		return err
	}

	return nil
}
