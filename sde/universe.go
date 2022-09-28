package sde

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type sde struct {
	filename     string // filepath to the SDE .zip file
	solarSystems []*SolarSystem
}

func New(filename string) (*sde, error) {
	s := &sde{
		filename:     filename,
		solarSystems: []*SolarSystem{},
	}
	if err := s.loadSolarSystems(); err != nil {
		return nil, err
	}
	return s, nil
}

// WriteFileStructure writes out the fullpath to each document in the SDE zipfile
func (s *sde) WriteFileStructure(outfile string) error {
	zf, err := zip.OpenReader(s.filename)
	if err != nil {
		return err
	}
	defer zf.Close()

	file, err := os.OpenFile(outfile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, document := range zf.File {
		writer.WriteString(document.Name + "\n")
	}
	return nil
}

func (s *sde) WriteSolarSystemsJSON(filename string) error {
	data, err := json.Marshal(s.solarSystems)
	if err != nil {
		return err
	}

	if err = os.WriteFile(filename, data, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func (s *sde) loadSolarSystems() error {
	zf, err := zip.OpenReader(s.filename)
	if err != nil {
		return err
	}
	defer zf.Close()

	for _, file := range zf.File {
		if strings.Contains(file.Name, "solarsystem.staticdata") {
			data, err := file.Open()
			if err != nil {
				return err
			}

			content, err := io.ReadAll(data)
			if err != nil {
				return err
			}

			var solarSystem SolarSystem
			if err = yaml.Unmarshal(content, &solarSystem); err != nil {
				if err != nil {
					return err
				}
			}

			solarSystem.SolarSystemTypeName = s.solarSystemTypeFromFile(file.Name)
			solarSystem.SolarSystemName = s.systemNameFromFile(file.Name)
			solarSystem.RegionName = s.regionNameFromFile(file.Name)
			solarSystem.ConstellationName = s.constellationNameFromFile(file.Name)

			s.solarSystems = append(s.solarSystems, &solarSystem)
		}
	}
	return nil
}

func (s *sde) solarSystemTypeFromFile(filename string) string {
	if !strings.Contains(filename, "solarsystem.staticdata") {
		return ""
	}
	parts := strings.Split(filename, "/")
	if len(parts) != 8 {
		return ""
	}
	return parts[3] // return "eve", "abyssal", "wormhole", "void"
}

func (s *sde) regionNameFromFile(filename string) string {
	if !strings.Contains(filename, "solarsystem.staticdata") {
		return ""
	}
	parts := strings.Split(filename, "/")
	if len(parts) != 8 {
		return ""
	}
	return parts[4]
}

func (s *sde) constellationNameFromFile(filename string) string {
	if !strings.Contains(filename, "solarsystem.staticdata") {
		return ""
	}
	parts := strings.Split(filename, "/")
	if len(parts) != 8 {
		return ""
	}
	return parts[5]
}

func (s *sde) systemNameFromFile(filename string) string {
	if !strings.Contains(filename, "solarsystem.staticdata") {
		return ""
	}
	parts := strings.Split(filename, "/")
	if len(parts) != 8 {
		return ""
	}
	return parts[6]
}
