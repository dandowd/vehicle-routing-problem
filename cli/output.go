package cli

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"vehicle-routing-problem/entities"
)

var Logger = log.New(NewFileWriter("./out.txt"), "", log.LstdFlags|log.Lshortfile)

func PrintRoutes(drivers []*entities.Driver) {
	for _, driver := range drivers {
		fmt.Println(driver.GetPath())
	}
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
	if err := os.WriteFile(f.path, bytes, fs.FileMode(0644)); err != nil {
		fmt.Println("Error writing to file ", err)
	}

	return 0, nil
}
