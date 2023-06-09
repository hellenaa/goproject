package promotions

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgxpool"
)

func addMany(db *pgxpool.Pool, data []Promotion) {
	_, err := db.CopyFrom(
		context.Background(),
		pgx.Identifier{"promotions"},
		[]string{"promotion_id", "price", "expiration_date"},
		pgx.CopyFromSlice(len(data), func(i int) ([]interface{}, error) {
			return []interface{}{data[i].PromotionId, data[i].Price, data[i].ExpirationDate}, nil
		}),
	)
	if err != nil {
		fmt.Println("err", err)
		return
	}
}

func deleteAll(db *pgxpool.Pool) {
	query := `truncate table promotions`
	_, err := db.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getById(db *pgxpool.Pool, id int64) (Promotion, error) {
	query := `select * from promotions where id = $1::int`
	rows, err := db.Query(context.Background(), query, id)
	if err != nil {
		fmt.Println(err)
		return Promotion{}, errors.New("something went wrong")
	}
	defer rows.Close()

	var promotions []Promotion
	for rows.Next() {
		var p Promotion
		err := rows.Scan(&p.Id, &p.PromotionId, &p.Price, &p.ExpirationDate)
		if err != nil {
			fmt.Println(err)
			return Promotion{}, errors.New("something went wrong")
		}
		promotions = append(promotions, p)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return Promotion{}, errors.New("something went wrong")
	}

	if len(promotions) == 0 {
		return Promotion{}, errors.New("promotion not found")
	}

	return promotions[0], nil
}
