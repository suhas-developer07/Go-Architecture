package repository

import (
	"database/sql"

	"github.com/suhas-developer07/Bank-service/models"
)


type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository{
	return &PostgresRepository{db:db}
}

func Init(r *PostgresRepository) error{
	query := `CREATE TABLE IF NOT EXIST customers (
	id UUID PRIMARY KEY,
	name VARCHAR(255),
	email VARCHAR(255).
	phone VARCHAR(20),
	address TEXT
	);
	`
	_,err := r.db.Exec(query)
	return err
}

func (r *PostgresRepository) CreateCustomer(customer *models.Customer) error{
	return nil
}

func (r *PostgresRepository) GetCustomer(id string) (*models.Customer,error){
	return &models.Customer{},nil
}

func (r *PostgresRepository) UpdateCustomer(customer *models.Customer) error{
	return nil
}

func (r *PostgresRepository) DeleteCustomer(id string) error{
	return  nil
}

func (r *PostgresRepository) ListOfCustomers()([]*models.Customer,error){
   return nil
}