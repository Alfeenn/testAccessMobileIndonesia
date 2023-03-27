package goroutinedata

type DataTest struct {
	User  *User
	Order *Order
}

type User struct {
	Username string
	Email    string
}

type Order struct {
	OrderId    string
	TotalOrder int64
}
