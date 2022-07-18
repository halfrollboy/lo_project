package dto

type Dto struct {
	Id int
}

type OrderApiDto struct {
	OrderUID          string         `json:"order_uid,omitempty"`
	TrackNumber       string         `json:"track_number,omitempty"`
	Entry             string         `json:"entry,omitempty"`
	Delivery          DeliveryApiDto `json:"delivery"`
	Payment           PaymentApiDto  `json:"payment"`
	Items             []ItemsApiDto  `json:"items"`
	Locale            string         `json:"locale,omitempty"`
	InternalSignature string         `json:"internal_signature,omitempty"`
	CustomerID        string         `json:"customer_id,omitempty"`
	DeliveryService   string         `json:"delivery_service,omitempty"`
	Shardkey          string         `json:"shardkey,omitempty"`
	SmID              int            `json:"sm_id,omitempty"`
	OofShard          string         `json:"oof_shard"`
}

type DeliveryApiDto struct {
	Name    string `json:"name,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Zip     string `json:"zip,omitempty"`
	City    string `json:"city,omitempty"`
	Address string `json:"address,omitempty"`
	Region  string `json:"region,omitempty"`
	Email   string `json:"email,omitempty"`
}

type PaymentApiDto struct {
	Transaction  string `json:"transaction,omitempty"`
	RequestID    string `json:"request_id,omitempty"`
	Currency     string `json:"currency,omitempty"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type ItemsApiDto struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}
