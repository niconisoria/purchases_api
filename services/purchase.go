package services

import (
	"fmt"
	"workshop/config"
	"workshop/db"
	"workshop/models"
	"workshop/tools"
)

var dbPurchases = db.DBPurchases

func CreatePurchase(purchase models.Purchase, user models.User) (models.Purchase, error) {
	if !tools.ValidateString(purchase.ID) {
		purchase.GenerateID()
	}

	purchase.Status = config.NEW
	return dbPurchases.Save(purchase, user)
}

func GetAllPurchases() []models.Purchase {
	return dbPurchases.GetAll()
}

func GetPurchaseByID(key string) (interface{}, error) {
	purchase, error := dbPurchases.GetByID(key)
	if error != nil {
		return nil, error
	}
	return purchase, nil
}

func UpdatePurchase(key string, purchase models.Purchase) (interface{}, error) {
	currentPurchase, err := dbPurchases.GetByID(key)
	if err != nil {
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

	return currentPurchase, dbPurchases.Update(key, currentPurchase)
}

func DeletePurchase(key string) string {
	resultado, err := dbPurchases.GetByID(key)
	if err != nil {
		return fmt.Sprintf("Purchase id: " + key + " doesn't exist")
	}
	if resultado.Status == config.FINISHED {
		return fmt.Sprintf("Purchase id: " + key + " has final status, cant be deleted")
	}
	err = dbPurchases.Delete(key)
	if err != nil {
		return fmt.Sprintf("Purchase id: " + key + " can't be deleted")
	}

	return fmt.Sprintf("Purchase id: " + key + " deleted")
}
