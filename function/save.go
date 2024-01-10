package hangmanweb

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func SaveScoreboard(data []Scoreboard) {
	file, err := os.OpenFile("scoreboard.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, item := range data {
		record := []string{item.Username, item.Category, strconv.Itoa(item.Points)}
		err := writer.Write(record)
		if err != nil {
			log.Println(err)
		}
	}
}

func SaveAccount(data []Account) {
	file, err := os.OpenFile("account.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, item := range data {
		record := []string{item.Username, item.Email, item.Password}
		err := writer.Write(record)
		if err != nil {
			log.Println(err)
		}
	}
}
