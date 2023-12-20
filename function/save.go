package hangmanweb

import (
	"encoding/csv"
	"log"
	"os"
)

func Save(data []Scoreboard) {
	file, err := os.OpenFile("scoreboard.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, item := range data {
		record := []string{item.Username, string(item.Points)}
		err := writer.Write(record)
		if err != nil {
			log.Println(err)
		}
	}
}
