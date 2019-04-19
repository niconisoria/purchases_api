package services

import (
	"encoding/json"
	"workshop/config"
	"workshop/db"
	"workshop/models"
	"workshop/tools"
	"workshop/utils"
)

var database = db.DBPurchases{}

func CreatePurchase(purchase models.Purchase) (models.Purchase, error) {
	if !tools.ValidateString(purchase.ID) {
		purchase.GenerateID()
	}

	purchase.Status = config.NEW
	return purchase, database.Save(purchase.ID, purchase)
}

func GetAllPurchases() map[string]models.Purchase {
	result := map[string]models.Purchase{}
	for k, purchase := range database.GetAll() {
		if p, ok := purchase.(models.Purchase); ok {
			result[k] = p
		}
	}
	return result
}

func GetPurchaseByID(key string) (interface{}, error) {
	if purchase, error := database.GetById(key); error != nil {
		return nil, error
	} else {
		return purchase, nil
	}
}

func UpdatePurchase(key string, purchase models.Purchase) (interface{}, error) {
	savedPurchase, err := database.GetById(key)
	if err != nil {
		return nil, err
	}
	currentPurchase := models.Purchase{}
	if err := json.Unmarshal(utils.InterfaceToBytes(savedPurchase), &currentPurchase); err != nil {
		return nil, err
	}
	if purchase.Amount != currentPurchase.Amount && purchase.Amount > 0 {
		currentPurchase.Amount = purchase.Amount
	}
	if purchase.Title != currentPurchase.Title && purchase.Title != "" {
		currentPurchase.Title = purchase.Title
	}
	if purchase.Image != currentPurchase.Image && purchase.Image != "" {
		currentPurchase.Image = purchase.Image
	}
	if purchase.Status != currentPurchase.Status && purchase.Status != "" {
		currentPurchase.Status = purchase.Status
	}

	return database.Update(key, currentPurchase)
}

func DeletePurchase(key string) string {
	return database.Delete(key)
}
