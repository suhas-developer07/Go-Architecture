package routes

import (
	"github.com/gorilla/mux"
	"github.com/suhas-developer07/Bank-service/controllers"
	"github.com/suhas-developer07/Bank-service/models"
)




func MountRoutes(repo models.BankRepo)*mux.Router{
   handler := controllers.NewBankHandler(repo)
   
   router := mux.NewRouter()

   router.HandleFunc("/customers",handler.CreateCustomerHandler).Methods("POST")
   router.HandleFunc("/customers/{id}",handler.GetCustomerHandler).Methods("GET")
   router.HandleFunc("/customers/{id}",handler.UpdateCustomerHandler).Methods("PUT")
   router.HandleFunc("/customers/{id}",handler.DeleteCustomerHandler).Methods("DELETE")
   router.HandleFunc("/customers",handler.ListOfCustomers).Methods("GET")
   
   return router
}