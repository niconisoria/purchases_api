package db

import (
	"errors"
	"fmt"
	"workshop/config"
	"workshop/models"

	"syreclabs.com/go/faker"
)

var localdb map[string]interface{}

type DBPurchases struct{}

func (dbp *DBPurchases) Save(key string, value models.Purchase) error {
	if _, ok := localdb[key]; ok {
		return errors.New(fmt.Sprintf("Purchase whit key %v already exist.", key))
	}
	localdb[key] = value
	return nil
}

func (dbp *DBPurchases) GetAll() map[string]interface{} {
	return localdb
}

func (dbp *DBPurchases) GetById(key string) (interface{}, error) {
	if _, ok := localdb[key]; ok {
		return localdb[key], nil
	} else {
		return nil, errors.New(fmt.Sprintf("Purchase with key %v not found ", key))
	}
}

func (dbp *DBPurchases) Update(key string, purchase models.Purchase) (interface{}, error) {
	if _, ok := localdb[key]; ok {
		localdb[key] = purchase
		return localdb[key], nil
	} else {
		return nil, errors.New(fmt.Sprintf("Purchase with key %v not found ", key))
	}
}

func (dbp *DBPurchases) Delete(key string) string {
	delete(localdb, key)
	if _, ok := localdb[key]; !ok {
		return fmt.Sprintf("Purchase id:" + key + " deleted ")
	} else {
		return fmt.Sprintf("Purchase id:" + key + " cant be deleted ")
	}
}

func init() {
	fmt.Println("--------------------- INIT db ---------------------")
	localdb = map[string]interface{}{}

	for i := 1; i <= 10; i++ {
		id := fmt.Sprintf("%v", i)
		purchase := models.Purchase{
			Amount: faker.Commerce().Price(),
			ID:     id,
			Image:  faker.Internet().Url(),
			Title:  faker.Commerce().ProductName(),
		}
		switch i % 3 {
		case 0:
			purchase.Status = config.FINISHED
		case 1:
			purchase.Status = config.NEW
		case 2:
			purchase.Status = config.CANCELLED
		}
		localdb[id] = purchase
	}

}
