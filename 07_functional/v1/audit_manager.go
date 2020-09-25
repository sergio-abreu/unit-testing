package v1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func NewAuditManager(maxEntriesPerFile int, directoryName string) AuditManager {
	return AuditManager{maxEntriesPerFile: maxEntriesPerFile, directoryName: directoryName}
}

type AuditManager struct {
	maxEntriesPerFile int
	directoryName     string
}

func (a AuditManager) AddRecord(visitorName string, timeOfVisit time.Time) error {
	filePaths := GetFiles(a.directoryName)
	newRecord := fmt.Sprintf("%s; %s", visitorName, timeOfVisit.Format(time.RFC3339))
	if len(filePaths) == 0 {
		file, err := os.OpenFile(fmt.Sprintf("%s/audit-1.txt", a.directoryName), os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = file.WriteString(newRecord)
		return err
	}
	currentIndex := len(filePaths)
	currentFile := filePaths[currentIndex-1]
	file, err := os.OpenFile(currentFile, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	lines, err := ReadAllLines(file)
	if err != nil {
		return err
	}
	if len(lines) < a.maxEntriesPerFile {
		_, err := file.WriteString(newRecord)
		return err
	}
	newFile, err := os.OpenFile(fmt.Sprintf("%s/audit-%d.txt", a.directoryName, currentIndex+1), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer newFile.Close()
	_, err = newFile.WriteString(newRecord)
	return err
}

func GetFiles(directoryName string) (filePaths []string) {
	_ = filepath.Walk(directoryName, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			filePaths = append(filePaths, info.Name())
		}
		return nil
	})
	return
}

func ReadAllLines(file *os.File) ([]string, error) {
	reader := bufio.NewReader(file)
	var reads []string
	for {
		read, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		reads = append(reads, read)
	}
	return reads, nil

}
