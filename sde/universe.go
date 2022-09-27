package sde

import (
	"archive/zip"
	"bufio"
	"os"
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
