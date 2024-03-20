package response

import "time"

type OrderResponse struct {
	ID           string         `json:"id"`
	CustomerName string         `json:"customer_name"`
	OrderedAt    time.Time      `json:"ordered_at"`
	UserID       string         `json:"user_id"`
	Items        []ItemResponse `json:"items"`
}

type ItemResponse struct {
	ID          string `json:"id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
