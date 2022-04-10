
# Shopping Cart API

This project is an basic web API for shopping operations

### Stack
* Golang - Gin Web Framework
* MySQL
* Docker



## Deployment

#### Requirements
* go
* docker
* docker-compose

To run MySQL Database
```bash
docker-compose -d docker-compose.yml up -d
```

To run the web server
```bash
go run main.go 
```




## API Reference
Project has Swagger integration and it can be reached on
[Local Swagger Doc](http://localhost:8080/swagger/index.html)
after running the project

### Authantication and Authorization
Project JWT for both Authorization and Authantication.

#### Auth Endpoints
```http
POST /user/login
POST /user
```

#### No Token Needed Endpoints
```http
GET    /category/
GET    /product/
```
#### User Spesific Endpoints
```http
GET    /cart/
POST   /cart/item
PATCH  /cart/item

GET    /order
POST   /order
DELETE /order
```
#### Access Restricted Endpoints (Admin Only)

```http
POST   /category
POST   /category/upload

POST   /product
DELETE /product
PATCH  /product
```





## Demo

After starting the project there initial data to try endpoints

### User

* User{Username: user, Password: user, IsAdmin: false}
* User{Username: admin, Password: admin, IsAdmin: true}

### Category

* Category{Name: CAT1, Desc: Category 1}
* Category{Name: CAT2, Desc: Category 2}



## Roadmap

- Add unit tests
- Seperate development config and production config
- Add detailed error messages
- Add initial data
- Add web server to docker compose






## Author

- [@yunusgok](https://www.github.com/yunusgok)


## License

[MIT](https://choosealicense.com/licenses/mit/)

