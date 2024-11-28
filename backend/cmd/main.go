package main

import (
    "log"
    "github.com/sharvatic/BookMyHotel/database"
    "github.com/sharvatic/BookMyHotel/firebase"
    "github.com/sharvatic/BookMyHotel/routes"
)


func main() {
    database.Connect()
    firebase.InitFirebaseApp()
    router := routes.SetupRouter()
    log.Fatal(router.Run(":8090"))
}
