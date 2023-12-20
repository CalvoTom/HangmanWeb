package hangmanweb

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func Open() []Scoreboard {
	file, err := os.Open("scoreboard.csv")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
	}

	var data []Scoreboard
	for _, record := range records {
		username := record[0]
		category := record[1]
		points, _ := strconv.Atoi(record[2][1:])
		data = append(data, Scoreboard{Username: username, Category: category, Points: points})
	}

	return data
}
