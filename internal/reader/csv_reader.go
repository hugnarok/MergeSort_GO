package reader

import (
	"encoding/csv"
	"os"
	"strconv"
)

func LoadRatings(filePath string, sampleSize int) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // pula o cabeçalho

	ratings := []int{}
	for i := 0; i < sampleSize; {
		record, err := reader.Read()
		if err != nil {
			break
		}
		ratingFloat, _ := strconv.ParseFloat(record[2], 64)
		ratingInt := int(ratingFloat * 10) // evita float e facilita ordenação
		ratings = append(ratings, ratingInt)
		i++
	}
	return ratings, nil
}
