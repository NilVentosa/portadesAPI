# portadesAPI

This started as a way to learn how to do a basic API with GO. It is used by [El joc de les portades](https://eljocdelesportades.xyz).

### How to run
To run the API execute the following commands:

    go build cmd/portadesAPI.go
    ./portadesAPI portades.db

### Endpoints

+ [x] `GET /random`returns a random Portada
+ [ ] `GET /result/id`returns the result of the specified Portada
+ [x] `GET /portada/id` returns specific Portada
+ [ ] `POST /portada?{...}`Adds a new Portada

### TODO

+ [ ] Result is hashed
+ [ ] Add some kind of authentification
+ [x] Use Sqlite db instead of csv file
