package sde

import (
	"archive/zip"
	"bufio"
	"os"
)

type SDE struct {
	Filename string
}

func (sde *SDE) WriteFileStructure(outfile string) error {
	zf, err := zip.OpenReader(sde.Filename)
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
