package input

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"vehicle-routing-problem/entities"
)

func ParseFile(filepath string) []*entities.Load {
  file, err := os.Open(filepath)

  if err != nil {
    panic(err)
  }

  scanner := bufio.NewScanner(file)

  var loads []*entities.Load
  for scanner.Scan() {
    
  }
  file.Close()

  return loads
}

func parseLine(line string) *entities.Load {
  split := strings.Split(line, " ")
  
  loadNumber, err := strconv.ParseInt(split[0], 10, 64)
  if err != nil {
    panic(err)
  }

  pickupX, pickupY := parseLoadPoint(split[1])

  dropoffX, dropoffY := parseLoadPoint(split[2])
  
  return &entities.Load{LoadNumber: int(loadNumber), Pickup: entities.Point{X: pickupX, Y: pickupY}, Dropoff: entities.Point{X: dropoffX, Y: dropoffY}} 
}

func parseLoadPoint(point string) (float64, float64) {
  split := strings.Split(point[1 : len(point) - 1], ",")

  x, err := strconv.ParseFloat(split[0], 64)
  if err != nil {
    panic(err)
  }

  y, err := strconv.ParseFloat(split[1], 64)
  if err != nil {
    panic(err)
  }

  return x, y
}
