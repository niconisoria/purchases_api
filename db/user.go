package db

import "workshop/models"

type Users struct{}

func (dbp *Users) Save(user models.User) (models.User, error) {
	sqlStatement := `
	INSERT INTO users (name, last_name, dni)
	VALUES ($1, $2, $3) RETURNING user_id`
	err := db.QueryRow(sqlStatement, user.Name, user.LastName, user.DNI).Scan(&user.ID)
	return user, err
}

func (dbp *Users) GetAll() []models.User {
	var users = []models.User{}
	rows, err := db.Query("SELECT user_id, name, last_name, dni FROM users LIMIT $1", 1000)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user = models.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.LastName, &user.DNI)
		users = append(users, user)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return users
}

func (dbp *Users) GetByID(key string) (models.User, error) {
	sqlStatement := "SELECT user_id, name, last_name, dni FROM users WHERE user_id=$1;"
	var user = models.User{}
	row := db.QueryRow(sqlStatement, key)
	err := row.Scan(&user.ID, &user.Name, &user.LastName, &user.DNI)
	return user, err
}

func (dbp *Users) Update(key string, user models.User) error {
	sqlStatement := `UPDATE users SET name = $2, last_name = $3, dni = $4 WHERE user_id = $1;`
	_, err := db.Exec(sqlStatement, key, user.Name, user.LastName, user.DNI)
	return err
}

func (dbp *Users) Delete(key string) error {
	sqlStatement := `DELETE FROM users WHERE user_id = $1;`
	_, err := db.Exec(sqlStatement, key)
	return err
}
