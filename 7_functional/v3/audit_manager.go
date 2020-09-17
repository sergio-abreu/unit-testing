package v3

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	v1 "sergio/unit-testing/7_functional/v1"
	"time"
)

// Arquitetura Funcional

type FileUpdate struct {
	Filename string
	Content  string
}

type FileContent struct {
	Filename string
	Lines    []string
}

func NewAuditManager(maxEntriesPerFile int) AuditManager {
	return AuditManager{maxEntriesPerFile: maxEntriesPerFile}
}

type AuditManager struct {
	maxEntriesPerFile int
}

func (a AuditManager) AddRecord(files []FileContent, visitorName string, timeOfVisit time.Time) FileUpdate {
	newRecord := fmt.Sprintf("%s; %s", visitorName, timeOfVisit.Format(time.RFC3339))
	if len(files) == 0 {
		return FileUpdate{Filename: "audit-1.txt", Content: newRecord}
	}
	currentIndex := len(files)
	currentFile := files[currentIndex-1]
	if len(currentFile.Lines) < a.maxEntriesPerFile {
		return FileUpdate{Filename: currentFile.Filename, Content: newRecord}
	}
	newIndex := currentIndex + 1
	newFilename := fmt.Sprintf("audit-%d.txt", newIndex)
	return FileUpdate{Filename: newFilename, Content: newRecord}
}

type Persister struct{}

func (p Persister) ReadDirectory(directoryName string) (fileContent []FileContent) {
	_ = filepath.Walk(directoryName, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			file, err := os.Open(info.Name())
			if err != nil {
				return err
			}
			defer file.Close()
			lines, err := v1.ReadAllLines(file)
			if err != nil {
				return err
			}
			fileContent = append(fileContent, FileContent{
				Filename: info.Name(),
				Lines:    lines,
			})
		}
		return nil
	})
	return
}

func (p Persister) ApplyUpdate(directoryName string, update FileUpdate) error {
	filePath := path.Join(directoryName, update.Filename)
	return ioutil.WriteFile(filePath, []byte(update.Content), os.ModeAppend)
}

func NewApplicationService(directoryName string, maxEntriesPerFile int) ApplicationService {
	return ApplicationService{directoryName: directoryName, auditManager: NewAuditManager(maxEntriesPerFile), persister: Persister{}}
}

type ApplicationService struct {
	directoryName string
	auditManager  AuditManager
	persister     Persister
}

func (a ApplicationService) AddRecord(visitorName string, timeOfVisit time.Time) error {
	files := a.persister.ReadDirectory(a.directoryName)
	update := a.auditManager.AddRecord(files, visitorName, timeOfVisit)
	return a.persister.ApplyUpdate(a.directoryName, update)
}