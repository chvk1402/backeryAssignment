package structs

type Price struct {
	Pack   int     `json:"pack"`
	QtySet int     `json:"quantity"`
	Price  float32 `json:"price(in $)"`
}

type OrderReq struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type OrderResp struct {
	Code       string  `json:"code"`
	TotalPrice float32 `json:"total_price"`
	Packs      []Price `json:"packs"`
}
