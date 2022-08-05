package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

type Record struct {
	Id      string   `json:"_id"`
	Output  string   `json:"output"`
	Aliases []string `json:"aliases"`
}

type Bot struct {
	BotName  string
	Commands []Record
}

var ErrDbLocationUnset = errors.New("DB_LOCATION environment variable not set")

func botLocation() (string, error) {
	value := os.Getenv("DB_LOCATION")
	if len(value) == 0 {
		return "", ErrDbLocationUnset
	}

	return value, nil
}

func dbFileLocation(name string) string {
	basePath, _ := botLocation()
	return path.Join(basePath, name, "commands.db")
}

func hasDbFile(name string) bool {
	commandsFile := dbFileLocation(name)

	if _, err := os.Stat(commandsFile); err != nil {
		return !os.IsNotExist(err)
	}

	return true
}

func LoadDb(name string) ([]Record, error) {
	commandsFile := dbFileLocation(name)

	content, err := ioutil.ReadFile(commandsFile)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]
	ret := make([]Record, len(lines))

	for idx, line := range lines {
		json.Unmarshal([]byte(line), &(ret[idx]))
	}

	return ret, nil
}

func AvailableBots() (*mapset.Set[string], error) {
	ret := mapset.NewSet[string]()

	basePath, err := botLocation()
	if err != nil {
		return nil, err
	}

	names, err := ioutil.ReadDir(basePath)
	if err != nil {
		return nil, err
	}

	for _, file := range names {
		if hasDbFile(file.Name()) {
			ret.Add(file.Name())
		}
	}

	return &ret, nil
}
