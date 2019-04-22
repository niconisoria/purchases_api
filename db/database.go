package db

import (
	"errors"
	"fmt"
	"workshop/config"
	"workshop/models"

	"syreclabs.com/go/faker"
)

var localdb map[string]interface{}
var users = map[int]models.User{}

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
	var message = " purchase " + key + " doesn't exist"
	if _, ok := localdb[key]; ok {
		resultado := models.Purchase{}
		resultado = localdb[key].(models.Purchase)
		if resultado.Status == config.FINISHED {
			return fmt.Sprintf("Purchase id:" + key + " has final status, cant be deleted ")
		}
		delete(localdb, key)
		if _, ok := localdb[key]; !ok {
			message = fmt.Sprintf("Purchase id:" + key + " deleted ")
		} else {
			message = fmt.Sprintf("Purchase id:" + key + " cant be deleted ")
		}
	}
	return message
}

func init() {
	fmt.Println("--------------------- INIT db ---------------------")
	localdb = map[string]interface{}{}
	users = map[int]models.User{}
	for i := 1; i <= 3; i++ {
		u := models.User{
			DNI:      faker.RandomInt(20000000, 40000000),
			ID:       int64(i),
			Name:     faker.Name().FirstName(),
			LastName: faker.Name().LastName(),
		}
		users[i] = u
	}

	for i := 1; i <= 10; i++ {
		id := fmt.Sprintf("%v", i)
		purchase := models.Purchase{
			Amount: faker.Commerce().Price(),
			ID:     id,
			Image:  fmt.Sprintf("https://loremflickr.com/320/240?random=%v", i),
			Title:  faker.Commerce().ProductName(),
		}
		switch i % 3 {
		case 0:
			purchase.Status = config.FINISHED
			purchase.User = users[1]
		case 1:
			purchase.Status = config.NEW
			purchase.User = users[2]
		case 2:
			purchase.Status = config.CANCELLED
			purchase.User = users[3]
		}
		localdb[id] = purchase
	}

}
