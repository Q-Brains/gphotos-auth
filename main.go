package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/Q-Brains/gphotos"
)

func main() {
	id := os.Getenv("CLIENT_ID")
	if reflect.DeepEqual(id, "") {
		fmt.Println("Invalid Request: CLIENT_ID not found.")
		return
	}

	secret := os.Getenv("CLIENT_SECRET")
	if reflect.DeepEqual(secret, "") {
		fmt.Println("Invalid Request: CLIENT_SECRET not found.")
		return
	}

	scopes := parseScopes(os.Getenv("AUTH_SCOPES"))
	if len(scopes) <= 0 {
		fmt.Println("Invalid Request: AUTH_SCOPES not found.")
		return
	}

	conf := gphotos.Auth.OAuth2Config(id, secret, scopes)
	authURL := gphotos.Auth.OAuth2CreateURL(conf, "gphotos-auth")
	fmt.Println("Authorization URL: " + authURL)
}

func parseScopes(str string) gphotos.AuthorizationScopes {
	var scopes gphotos.AuthorizationScopes

	for _, s := range strings.Split(str, " ") {
		switch s {
		case "read":
			scopes = append(scopes, gphotos.Readonly)
		case "append":
			scopes = append(scopes, gphotos.Appendonly)
		case "access":
			scopes = append(scopes, gphotos.Appcreateddata)
		case "read/append":
			scopes = append(scopes, gphotos.ReadAndAppend)
		case "sharing":
			scopes = append(scopes, gphotos.Sharing)
		default:
			fmt.Println("Invalid Argments: Unexpected scope `" + s + "`")
			return nil
		}
	}

	return scopes
}
