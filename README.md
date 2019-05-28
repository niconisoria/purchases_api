# purchases_api
*Golang API for users and purchases*

# Developers:

- **Nisoria, Nicolás Maximiliano** - nicolas.nisoria@gmail.com
- **Ceriana, Emanuel** - emanuelceriana@gmail.com

# Tasks:

1) It was created the complete user module (services, model, controllers and the database logic).
2) The reason why the endPoint didn't allow to add new purchases was that the func IsValid() on Users model ever return false.
3) We improved the security in purchase module adding in the routes the middleware "onlyAdmin".

# Extra:

We implemented a postgres database deployed on heroku using the "pq" library because the api didn't have a database.

# Curl (Example)

curl -v -X POST -H "role : admin" "https://purchases-api.herokuapp.com/purchases" -d '{ "image": "url",  "user_id": 565, "title": "Título 1", "status": "aprobado", "amount": 123.5 }'
