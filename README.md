# New Mall API

## Description

This is a small API for e-commerce.

Build for learning purpose. 

## Tech Stack

- Golang 
- Gin Framework
- MySQL 8.0 for database
- Gorm for ORM
- Redis for cache
- JWT for authentication
- Makefile
- Docker and Docker Compose for deployment

## Build and run

### Docker

- Requirements: Docker, Docker Compose

```bash
cp config-compose.yaml config.yaml # copy the config file

make comepose-up # start the docker compose
```

### Local
- Default config is for docker compose, if you want to run it locally, please change the config in `config-local.yaml`, for database, you can use docker compose to start a MySQL container.

```bash
cp config-local.yaml config.yaml # copy the config file

docker-compose up -d mysql # start a MySQL container

make # build and run the project
```

## Project Structure

```
.
├── cmd
│   └── main.go
├── internal
│   ├── config          # config package
│   ├── controllers     # controller handlers for routes
│   ├── database        # database for mysql setup and seeding
│   ├── global          # global store DB, CONFIG
│   ├── middleware      # middleware for routes
│   ├── models          # models for database
│   ├── repositories    # repositories handle database operations
│   ├── routes          # setup routes for gin
│   ├── services        # services handle business logic
│   └── types           # types for models (request, response)
├── pkg
│   ├── common          # common constants, errors
│   └── component       # app components like token, upload, hash, email
└── static                    # static files
    └── upload
```

## API

Import the postman collection in `static/postman` to test the API.

## Features

- [x] Authentication
  - [x] JWT
  - [x] Refresh token
  - [x] Logout
  - [ ] 3rd party login (Google, Facebook)
- [x] User
  - [x] Register
  - [x] Login
  - [x] Get user info
  - [x] Update user info
  - [x] Delete user
- [x] Product
  - [x] Create product
  - [x] Get product list
  - [x] Get product detail
  - [x] Update product
  - [x] Delete product
- [x] Category
  - [x] Create category
  - [x] Get category list
  - [x] Update category
  - [x] Delete category
- [x] Order
  - [x] Create order from cart
  - [x] Get order list
  - [x] Get order detail
  - [ ] Update order status
  - [x] Delete order
- [x] Cart
  - [x] Add product to cart
  - [x] Get cart item list
  - [x] Update cart item
  - [x] Delete cart
- [x] Address
  - [x] Create address
  - [x] Get address list for user
  - [x] Update address
  - [x] Delete address
- [ ] Payment

## TODO

- [ ] Add Validation for request
- [ ] Add Role and Permission for user (rbac)
- [ ] Add unit test
- [ ] Implement database transaction
- [ ] Implement upload/download provider (AWS S3 or Cloudinary)
- [ ] Add integration test
- [ ] Add CI/CD
- [ ] Add more features
- [ ] Refactor code to clean architecture
- [ ] Deploy to cloud (AWS)
- [ ] Build frontend (React/Vue/Android/iOS)

## Reference

- [Gin](https://gin-gonic.com/docs/)
- [Gorm](https://gorm.io/docs/)
- [Golang](https://golang.org/doc/)
- [Docker](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Makefile](https://www.gnu.org/software/make/manual/make.html)
- [JWT](https://jwt.io/)
- [OpenAI](https://openai.com/)

## License

[MIT](https://choosealicense.com/licenses/mit/)
```
