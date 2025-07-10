package repository

import (
	"database/sql"
	"fmt"

	"github.com/suhas-developer07/Bank-service/models"
)


type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository{
	return &PostgresRepository{db:db}
}

func (r *PostgresRepository) Init() error{
	query := `CREATE TABLE IF NOT EXISTS customers (
	id UUID PRIMARY KEY,	
	name VARCHAR(255),
	email VARCHAR(255),
	phone VARCHAR(20),
	address TEXT
	);
	`
	_,err := r.db.Exec(query)
	return err
}

func (r *PostgresRepository) CreateCustomer(customer *models.Customer) error{
    query := `INSERT INTO customers (id,name,email,phone,address) VALUES ($1,$2,$3,$4,$5)`

	_,err := r.db.Exec(query,customer.ID,customer.Name,customer.Email,customer.Phone,customer.Address)

	return err
}

func (r *PostgresRepository) GetCustomer(id string) (*models.Customer,error){
   
	query:= `SELECT id,name,email,phone,address FROM customers WHERE id = $1`

	var customer models.Customer

	err := r.db.QueryRow(query,id).Scan(
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.Phone,
		&customer.Address,
	)

	if err!=nil{
		return nil,err
	}

	return &customer,nil

}

func (r *PostgresRepository) UpdateCustomer(customer *models.Customer) error{

	query := `UPDATE customers SET name=$1, email=$2, phone=$3, address=$4 WHERE id = $5`

	_,err := r.db.Exec(query,customer.Name,customer.Email,customer.Phone,customer.Address,customer.ID)

	return err
}

func (r *PostgresRepository) DeleteCustomer(id string) error{

	query:= `DELETE FROM customers WHERE id=$1`

	result,err:= r.db.Exec(query,id)

	if err!=nil {
		return err
	}

	rowsAffected,err := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("task with ID %s not found",id)
	}

	return  err
}

func (r *PostgresRepository) ListOfCustomers()([]*models.Customer,error){

	query := `SELECT * FROM customers`

	rows,err := r.db.Query(query)

	if err!=nil{
		return nil ,err
	}

	var customers []*models.Customer

	for rows.Next(){
		var customer models.Customer

		if err := rows.Scan(&customer.ID,&customer.Name,&customer.Email,&customer.Phone,&customer.Address); err!=nil{
			return nil , err
		}

		customers = append(customers, &customer)
	}

	return customers,nil

}