package db

import (
	"database/sql"
	"fmt"
	"workshop/models"

	_ "github.com/lib/pq"
	"syreclabs.com/go/faker"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "niconisoria"
	password = ""
	dbname   = ""
)

var db = initDB()
var DBUsers = Users{}
var DBPurchases = Purchases{}

func initDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}

func init() {
	for i := 1; i <= 3; i++ {
		u := models.User{
			DNI:      faker.RandomInt(20000000, 40000000),
			Name:     faker.Name().FirstName(),
			LastName: faker.Name().LastName(),
		}
		DBUsers.Save(u)
	}

	for i := 1; i <= 3; i++ {
		user := models.User{
			ID: int64(i),
		}
		purchase := models.Purchase{
			User:   user,
			Amount: faker.Commerce().Price(),
			Image:  fmt.Sprintf("https://loremflickr.com/320/240?random=%v", i),
			Title:  faker.Commerce().ProductName(),
		}
		purchase.GenerateID()
	}

}
