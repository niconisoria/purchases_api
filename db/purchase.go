package db

import (
	"errors"
	"purchases-api/models"
	"strconv"
)

type Purchases struct{}

func (dbp *Purchases) Save(purchase models.Purchase, user models.User) (models.Purchase, error) {
	user, err := DBUsers.GetByID(strconv.FormatInt(user.ID, 10))
	if err != nil {
		return purchase, errors.New("User doesn't exists")
	}
	purchase.User = user
	sqlStatement := `
	INSERT INTO purchases (purchase_id, image, title, status, amount, user_id)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING creation_date`
	err = db.QueryRow(sqlStatement, purchase.ID, purchase.Image, purchase.Title, purchase.Status, purchase.Amount, user.ID).Scan(&purchase.CreationDate)
	return purchase, err
}

func (dbp *Purchases) GetAll() []models.Purchase {
	var purchases = []models.Purchase{}
	rows, err := db.Query("SELECT purchase_id, image, title, status, amount, creation_date, user_id FROM purchases LIMIT $1", 1000)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var purchase = models.Purchase{}
		err = rows.Scan(&purchase.ID, &purchase.Image, &purchase.Title, &purchase.Status, &purchase.Amount, &purchase.CreationDate, &purchase.User.ID)
		if err != nil {
			panic(err)
		}
		user, err := DBUsers.GetByID(strconv.FormatInt(purchase.User.ID, 10))
		if err != nil {
			panic(err)
		}
		purchase.User = user
		purchases = append(purchases, purchase)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return purchases
}

func (dbp *Purchases) GetByID(key string) (models.Purchase, error) {
	sqlStatement := "SELECT purchase_id, image, title, status, amount, creation_date, user_id FROM purchases WHERE purchase_id=$1;"
	var purchase = models.Purchase{}
	row := db.QueryRow(sqlStatement, key)
	err := row.Scan(&purchase.ID, &purchase.Image, &purchase.Title, &purchase.Status, &purchase.Amount, &purchase.CreationDate, &purchase.User.ID)
	user, err := DBUsers.GetByID(strconv.FormatInt(purchase.User.ID, 10))
	purchase.User = user
	return purchase, err
}

func (dbp *Purchases) Update(key string, purchase models.Purchase) error {
	sqlStatement := `UPDATE purchases SET amount = $2, title = $3, image = $4, status = $5 WHERE purchase_id = $1;`
	_, err := db.Exec(sqlStatement, key, purchase.Amount, purchase.Title, purchase.Image, purchase.Status)
	return err
}

func (dbp *Purchases) Delete(key string) error {
	sqlStatement := `DELETE FROM purchases WHERE purchase_id = $1;`
	_, err := db.Exec(sqlStatement, key)
	return err
}
