package controllers

import (
	"net/http"

	"github.com/suhas-developer07/Bank-service/models"
)

type BankHandler struct {
	repo  models.BankRepo
}

func NewBankHandler(repo models.BankRepo) *BankHandler{
	return &BankHandler{repo :repo}
}

func (h *BankHandler) CreateCustomerHandler(w http.ResponseWriter,r *http.Request){

}

func (h *BankHandler) GetCustomerHandler(w http.ResponseWriter, r *http.Request){

}
func (h *BankHandler) UpdateCustomerHandler(w http.ResponseWriter, r *http.Request){

}
func (h *BankHandler) DeleteCustomerHandler(w http.ResponseWriter, r *http.Request){

}
func (h *BankHandler) ListOfCustomers(w http.ResponseWriter, r *http.Request){

}
