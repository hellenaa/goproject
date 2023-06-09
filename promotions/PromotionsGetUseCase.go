package promotions

import (
	"github.com/jackc/pgx/pgxpool"
)

func GetUseCase(db *pgxpool.Pool, id int64) (Promotion, error) {
	return getById(db, id)
}
