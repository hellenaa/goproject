package promotions

type Promotion struct {
	Id             int     `json:"id" gorm:"primary_key"`
	PromotionId    string  `json:"promotion_id"`
	Price          float64 `json:"price"`
	ExpirationDate string  `json:"expiration_date"`
}
