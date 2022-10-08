package helpers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ProductCategory struct {
	ProductID          int            `json:"product_id"`
	ProductName        string         `json:"product_name"`
	ProductPrice       int            `json:"product_price"`
	ProductImageLink   sql.NullString `json:"product_image_link,omitempty"`
	ProductDescription sql.NullString `json:"product_description,omitempty"`
	CategoryID         int            `json:"category_id,omitempty"`
	CategoryName       string         `json:"category_name"`
}

func SuccessResponseJSON(w http.ResponseWriter, msg string, data interface{}) {
	response := JsonResponse{Success: true, Message: msg, Data: data}
	responseByte, _ := json.Marshal(response)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseByte)
}

func ErrorResponseJSON(w http.ResponseWriter, msg string, statusCode int) {
	response := JsonResponse{
		Success: false,
		Message: msg,
		Data:    nil,
	}

	responseByte, _ := json.Marshal(response)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseByte)
}
