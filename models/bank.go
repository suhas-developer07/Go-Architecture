package models

type Customer struct {
	ID       string   `json:"id"` 
	Name     string   `json:"name"` 
	Email    string   `json:"email"` 
	Phone    string   `json:"phone"` 
	Address  string   `json:"address"` 
}

type CustomerRequest struct {
	Name     string   `json:"name"` 
	Email    string   `json:"email"` 
	Phone    string   `json:"phone"` 
	Address  string   `json:"address"` 
}

type CustomerResponse struct {
	ID      string     `json:"id"`
	Name    string     `json:"name"`
	Email   string     `json:"email"`
}

type BankRepo interface {
	CreateCustomer(customer *Customer) error
	GetCustomer(id string) (*Customer,error)
	UpdateCustomer(customer *Customer) error
	DeleteCustomer(id string) error
	 ListOfCustomers()([]*Customer,error)
}