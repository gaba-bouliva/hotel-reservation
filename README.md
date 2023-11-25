# Hotel Reservation Backend

## Project outline
* Users -> book hotel rooms
* Admins -> validates reservations or bookings
* Authentication and authorization -> JWT tokens
* Hotels -> JSON CRUD API 
* Rooms -> JSON CRUD API
* Scripts -> database management (migration, seeding...) 

## Dependencies
#### Mongodb driver
Documentation
```
https://mongodb.com/docs/drivers/go/current/quick-start
```

Installing mongodb client
```
go get go.mongodb.org/mongo-driver/mongo
```

### gofiber
Documentation
```
https://gofiber.io
```

gofiber Installation
```
go get github.com/gofiber/fiber/v2
```

## Docker
### Installing mongodb as a Docker container
```
docker run --name mongodb -d mongo:latest -p 27017:27017
```
