package hangmanweb

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func OpenScoreboard() []Scoreboard {
	var line int

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
		if line == 15 {
			break
		}
		line++
		username := record[0]
		category := record[1]
		points, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, Scoreboard{Username: username, Category: category, Points: points})
		for i := 0; i < len(data); i++ {
			for j := i + 1; j < len(data); j++ {
				if data[i].Points < data[j].Points {
					data[i], data[j] = data[j], data[i]
				}
			}
		}
	}

	return data
}
