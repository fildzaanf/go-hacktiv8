package request

import "time"

type OrderRequest struct {
	OrderedAt    time.Time     `json:"ordered_at" form:"ordered_at" example:"2024-03-10T19:00:00+07:00"`
	CustomerName string        `json:"customer_name" form:"customer_name"`
	Items        []ItemRequest `json:"items" form:"items"`
}

type UpdateOrderRequest struct {
	OrderedAt    time.Time     `json:"ordered_at" form:"ordered_at" example:"2024-03-10T19:00:00+07:00"`
	CustomerName string        `json:"customer_name" form:"customer_name"`
	Items        []UpdateItemRequest `json:"items" form:"items"`

}

type ItemRequest struct {
	ItemCode    string `json:"item_code" form:"item_code"`
	Description string `json:"description" form:"description"`
	Quantity    int    `json:"quantity" form:"quantity"`
}

type UpdateItemRequest struct {
	ID          string `json:"id" form:"id"`
	ItemCode    string `json:"item_code" form:"item_code"`
	Description string `json:"description" form:"description"`
	Quantity    int    `json:"quantity" form:"quantity"`
}
