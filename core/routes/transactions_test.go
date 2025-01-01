package routes

import (
	"api/db"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	db.ConnectDB()

	id, err := CreateTransaction(Transaction{
		Title:         "Test_transaccion",
		Amount:        1010,
		CategoryID:    1,
		SubCategoryID: 1,
		Currency:      "USD",
		PaymentMethod: "Cash",
		ExchangeRate:  1,
		Notes:         "",
	})

	if err != nil {
		t.Errorf("Error creating transaction: %v", err)
	}

	if id == 0 {
		t.Error("Expected non-zero ID")
	}
}
