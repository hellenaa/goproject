package promotions

import (
	"../files"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"math"
	"mime/multipart"
)

func SaveCSVUseCase(db *pgxpool.Pool, file multipart.File) {
	records, err := files.ReadCSV(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	formattedRecords, err := formatCSVData(records)
	if err != nil {
		fmt.Println(err)
		return
	}

	deleteAll(db)

	limit := 2000000
	start := 0
	end := limit
	batchSize := 0
	sliceLen := len(formattedRecords)
	for {
		if batchSize == int(math.Ceil(float64(sliceLen)/float64(limit))) {
			break
		}
		if start > sliceLen {
			start = sliceLen
		}

		if end > sliceLen {
			end = sliceLen
		}

		go addMany(db, formattedRecords[start:end])

		batchSize++
		end = end + limit
		start = start + limit
	}
}
