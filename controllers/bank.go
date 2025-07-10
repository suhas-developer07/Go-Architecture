package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/suhas-developer07/Bank-service/models"
)

type BankHandler struct {
	repo  models.BankRepo
}

func NewBankHandler(repo models.BankRepo) *BankHandler{
	return &BankHandler{repo :repo}
}

func (h *BankHandler) CreateCustomerHandler(w http.ResponseWriter,r *http.Request){
	var Payload models.CustomerRequest

	if err := json.NewDecoder(r.Body).Decode(&Payload); err!=nil{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"true",
			"message":"Invalid request body format"})
	}

	customer := models.Customer{
		ID: uuid.New().String(),
		Name: Payload.Name,
		Email: Payload.Email,
		Phone: Payload.Phone,
		Address: Payload.Address,
	}

	if err := h.repo.CreateCustomer(&customer); err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"true",
			"message":err.Error()})
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"error":"false",
		"message":"Customer cretaed successfully",
	})
}

func (h *BankHandler) GetCustomerHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	customerId := vars["id"]

	if _,err := uuid.Parse(customerId); err!=nil {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"true",
			"message":err.Error(),
		})
		return
	}
	
    customer,err := h.repo.GetCustomer(customerId)

	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"true",
			"message":err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}
func (h *BankHandler) UpdateCustomerHandler(w http.ResponseWriter, r *http.Request){
	var Payload models.CustomerRequest
	vars := mux.Vars(r)
	customerId :=vars["id"]

	if _,err := uuid.Parse(customerId); err!=nil {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"true",
			"message":err.Error(),
		})
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&Payload); err!=nil{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"true",
			"message":"Invalid body formate",
		})
		return
	}

	Customer := models.Customer{
		ID:      customerId,
		Name:    Payload.Name,
		Email:   Payload.Email,
		Phone:   Payload.Phone,
		Address: Payload.Address,
	}

	 err := h.repo.UpdateCustomer(&Customer)

	 if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"true",
			"message":err.Error(),
		})
		return
	 }

	 w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"false",
			"message":"Customer updated successfully",
		})
}
func (h *BankHandler) DeleteCustomerHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	customerId := vars["id"]

	if _,err := uuid.Parse(customerId); err!=nil {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"true",
			"message":err.Error(),
		})
		return
	}

	err := h.repo.DeleteCustomer(customerId)

	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"true",
			"message":err.Error(),
		})
		return
	}
    
	 w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"false",
			"message":"Customer Deleted successfully",
		})
}
func (h *BankHandler) ListOfCustomers(w http.ResponseWriter, r *http.Request){

	customers,err := h.repo.ListOfCustomers(); 

	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"true",
			"message":err.Error(),
		})
		return
	}
    
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}
