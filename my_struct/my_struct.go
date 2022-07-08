package my_struct

type User struct {
	UserId int `json:"user_id"`
}
type Order struct {
	OrderId     int    `json:"order_id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
	UserId      int    `json:"user_id"`
}
type UserOrder struct {
	UserId int     `json:"user_id"`
	Orders []Order `json:"orders"`
}
type UserArray []UserOrder
