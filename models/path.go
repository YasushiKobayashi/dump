package models

import "fmt"

type (
	// Path
	// not contain extension to FileName
	Path struct {
		FilePath string
		FileName string
	}
)

func (p *Path) GetFileName() string {
	return fmt.Sprintf("%s.sql", p.FileName)
}

func (p *Path) GetFullFilePath() string {
	return fmt.Sprintf("%s/%s", p.FilePath, p.GetFileName())
}
