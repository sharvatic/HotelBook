package firebase

import (
    "context"
    "log"
    firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
    "google.golang.org/api/option"
)

var App *firebase.App
var AuthClient *auth.Client


func InitFirebaseApp() {
	var err error
    opt := option.WithCredentialsFile("./serviceAccount.json")
    app, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
        log.Fatalf("error initializing app: %v\n", err)
    }
    App = app;

	AuthClient, err = App.Auth(context.Background())
	if err != nil {
        log.Fatalf("error getting Auth client: %v\n", err)
    }
}
