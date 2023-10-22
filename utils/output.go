package utils

import (
	"fmt"
	"log"
	"os"
	"vehicle-routing-problem/entities"
)

func NewFileLogger(path string) *log.Logger {
	return log.New(NewFileWriter(path), "", 0)
}

func PrintRoutes(drivers []*entities.Driver) {
	fmt.Print(FormatPath(drivers))
}

func FormatPath(drivers []*entities.Driver) string {
	var path string
	for _, driver := range drivers {
		driverLoads := "["
		for _, load := range driver.GetPath() {
			driverLoads += fmt.Sprintf("%v,", load.LoadNumber)
		}

		driverLoads = driverLoads[:len(driverLoads)-1]
		driverLoads += "]\n"

		path += driverLoads
	}

	return path
}

func FormatDrivers(drivers []*entities.Driver) {
	fmt.Print(FormatPath(drivers))
}

type FileWriter struct {
	path string
}

func NewFileWriter(path string) *FileWriter {
	return &FileWriter{path: path}
}

func (f *FileWriter) Write(bytes []byte) (int, error) {
	file, err := os.OpenFile(f.path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	content := string(bytes)
	// Write the content to the file.
	_, err = file.WriteString(content)
	if err != nil {
		return 0, err
	}

	return 0, nil
}
