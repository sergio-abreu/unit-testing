package v2

import (
	"fmt"
	"path"
	"time"
)

// Arquitetura Hexagonal

type IFileSystem interface {
	GetFiles(directory string) []string
	WriteText(filename, text string) error
	ReadAllLines(filename string) ([]string, error)
}

func NewAuditManager(maxEntriesPerFile int, directoryName string, fileSystem IFileSystem) AuditManager {
	return AuditManager{maxEntriesPerFile: maxEntriesPerFile, directoryName: directoryName, fileSystem: fileSystem}
}

type AuditManager struct {
	maxEntriesPerFile int
	directoryName     string
	fileSystem        IFileSystem
}

func (a AuditManager) AddRecord(visitorName string, timeOfVisit time.Time) error {
	filePaths := a.fileSystem.GetFiles(a.directoryName)
	newRecord := fmt.Sprintf("%s; %s", visitorName, timeOfVisit.Format(time.RFC3339))
	if len(filePaths) == 0 {
		filename := path.Join(a.directoryName, "audit-1.txt")
		return a.fileSystem.WriteText(filename, newRecord)
	}
	currentIndex := len(filePaths)
	currentFile := filePaths[currentIndex-1]
	lines, err := a.fileSystem.ReadAllLines(currentFile)
	if err != nil {
		return err
	}
	if len(lines) < a.maxEntriesPerFile {
		return a.fileSystem.WriteText(currentFile, newRecord)
	}
	newIndex := currentIndex + 1
	newFile := fmt.Sprintf("/audit-%d.txt", newIndex)
	filename := path.Join(a.directoryName, newFile)
	return a.fileSystem.WriteText(filename, newRecord)
}
