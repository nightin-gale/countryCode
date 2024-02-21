package countryCode

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

var nameToCode = make(map[string]string)
var codeToName = make(map[string]string)
var initialised = false

const version = "v0.0.4"

func getGoPackagePath() string {
	var goPath string
	if os.Getenv("GOPATH") != "" {
		goPath = os.Getenv("GOPATH")
	} else {
		goPath = os.Getenv("HOME") + "/go/"
	}

	// shall be something like: /home/username/go/pkg/mod/github.com/nightin-gale/countryCode@v0.0.2/
	return fmt.Sprintf("%spkg/mod/github.com/nightin-gale/country!code@%s/", goPath, version)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = '|'
	csvReader.TrimLeadingSpace = true
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func saveCsvFile(filePath string, record [][]string) error {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Fail to create" + filePath)
	}
	defer f.Close()

	csvWriter := csv.NewWriter(f)
	csvWriter.Comma = '|'
	defer csvWriter.Flush()
	err = csvWriter.WriteAll(record)
	return err
}

func GetCountryCode(name string) (string, error) {
	if !initialised {
		initialise()
		initialised = true
	}
	key := strings.ToUpper(name)
	res, ok := nameToCode[key]
	if ok {
		return res, nil
	} else {
		return "", fmt.Errorf("Country Code for %s not found", key)
	}
}

func GetCountryName(code string) (string, error) {
	if !initialised {
		initialise()
		initialised = true
	}
	if len(code) != 2 {
		return "", fmt.Errorf("Country code should be 2 characters")
	}

	key := strings.ToUpper(code)
	res, ok := codeToName[key]
	if ok {
		return res, nil
	} else {
		return "UNKNOWN", nil
	}
}

func initialise() {
	path := getGoPackagePath()
	countryCode := readCsvFile(path + "countryCode.csv")
	moreCountryName := readCsvFile(path + "moreCountryName.csv")
	for _, row := range countryCode {
		nameToCode[row[0]] = row[1]
		codeToName[row[1]] = row[0]
	}
	for _, row := range moreCountryName {
		nameToCode[row[0]] = row[1]
	}
}
