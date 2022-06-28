package common

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var RandomState string

func SetupConfig() *oauth2.Config {
	godotenv.Load(".env")
	fmt.Println(os.Getenv("GOOGLE_CLIENT_ID"))
	conf := &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	RandomState="random"
	return conf
}
