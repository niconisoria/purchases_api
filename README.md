# Purchases API
*Golang API for users and purchases*

# Developers:

- **Nisoria, Nicolás Maximiliano** - nicolas.nisoria@gmail.com
- **Ceriana, Emanuel** - emanuelceriana@gmail.com

# Tasks:

1) The complete user module (services, model, controllers and the database logic) was created.
2) The reason why the endPoint didn't allow to add new purchases was that the func IsValid() on Users model ever return false.
3) We improved the security in purchase module adding in the routes the middleware "onlyAdmin".

# Extra:

We implemented a postgres database within the project and deployed on Heroku using the "pq" library since the project hadn't any database.

# Curl (Example)

curl -v -X POST -H "role : admin" "https://purchases-api.herokuapp.com/purchases" -d '{ "image": "https://loremflickr.com/320/240?random=1",  "user_id": 3, "title": "Awesome Concrete Computer", "status": "new", "amount": 123.5 }'
