package promotions

import (
	"fmt"
	"strconv"
)

func formatCSVData(records [][]string) ([]Promotion, error) {
	var formattedRecords []Promotion

	for _, record := range records {
		price, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			fmt.Println(err)
			return nil, nil
		}

		record := Promotion{
			PromotionId:    record[0],
			Price:          price,
			ExpirationDate: record[2],
		}

		formattedRecords = append(formattedRecords, record)
	}

	return formattedRecords, nil
}
