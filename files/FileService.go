package files

import (
	"encoding/csv"
	"fmt"
	"io"
)

func ReadCSV(file io.Reader) ([][]string, error) {
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return records, nil
}
