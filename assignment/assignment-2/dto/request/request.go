package request

import "time"

type OrderRequest struct {
	OrderedAt    time.Time     `json:"ordered_at" form:"ordered_at" example:"2024-03-10T19:00:00+07:00"`
	CustomerName string        `json:"customer_name" form:"customer_name" example:"Fildza"`
	Items        []ItemRequest `json:"items" form:"items"`
}

type UpdateOrderRequest struct {
	OrderedAt    time.Time     `json:"ordered_at" form:"ordered_at" example:"2024-03-10T19:00:00+07:00"`
	CustomerName string        `json:"customer_name" form:"customer_name" example:"Fildza"`
	Items        []UpdateItemRequest `json:"items" form:"items"`

}

type ItemRequest struct {
	ItemCode    string `json:"item_code" form:"item_code" example:"ABC123"`
	Description string `json:"description" form:"description" example:"Item A"`
	Quantity    int    `json:"quantity" form:"quantity" example:"7"`
}

type UpdateItemRequest struct {
	ID          string `json:"id" form:"id" example:"ce135609-1426-4011-bb95-8c475c45fb33"`
	ItemCode    string `json:"item_code" form:"item_code" example:"ABC123"`
	Description string `json:"description" form:"description" example:"Item A"`
	Quantity    int    `json:"quantity" form:"quantity" example:"7"`
}
