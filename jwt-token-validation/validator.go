package main

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
)

type MetaData struct {
	PublicURL   string
	TenantKey   string
	AccessToken string
	Aud         []string
	Sub         string
}

func main() {

	// jwksURL := "https://companyx.okta.com/oauth2/v1/keys"
	jwksURL := "https://nightlybuild.cidaas.de/.well-known/jwks.json"

	keySet, _ := jwk.Fetch(jwksURL)
	var accessToken = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjM5NjgzODBkLThjZTItNDIzMS04MzA4LTAyOTE1NmQxN2RkYSJ9.eyJ1YV9oYXNoIjoiMGI0MzNhMTc4YTYxYWI2MzllMWRhMzI1M2Y3NmZkOGQiLCJzaWQiOiJhMjVmODkzZS1kOGFkLTRmM2YtODg0Ni1kZWVlODQxNzdmZTgiLCJzdWIiOiJBTk9OWU1PVVMiLCJhdWQiOiI1MjE1NmQ1My0zY2JmLTQzZTctOTZlYi0yNTI4MDNkNGU5Y2EiLCJpYXQiOjE2MzE3MjU3NTcsImF1dGhfdGltZSI6MTYzMTcyNTc1NywiaXNzIjoiaHR0cHM6Ly9uaWdodGx5YnVpbGQuY2lkYWFzLmRlIiwianRpIjoiOTQwMDYxOTktMDNhZi00M2U3LTkwYTAtZWQ1OTkwMmU2NDBmIiwic2NvcGVzIjpbImNpZGFhczphZG1pbl9yZWFkIiwiY2lkYWFzOmFkbWluX3dyaXRlIiwiY2lkYWFzOmFkbWluX2RlbGV0ZSIsImNpZGFhczpyZWdpc3RlciIsImNpZGFhczp0b2tlbl9jcmVhdGUiLCJjaWRhYXM6cmVtb3ZlX3Nlc3Npb24iLCJjaWRhYXM6c2Vzc2lvbl93cml0ZSIsImNpZGFhczpzZXNzaW9uX2RlbGV0ZSIsImNpZGFhczp3ZWJob29rX3JlYWQiLCJjaWRhYXM6d2ViaG9va193cml0ZSIsImNpZGFhczp3ZWJob29rX2RlbGV0ZSIsImNpZGFhczpzZWVkaW5nIl0sImV4cCI6MTYzMTgxMjE1N30.YTOfLBDbu4im3bZvz1ZfpKXJP2PKa8jmQE7rLvbNaLrrFf7cvCVPGNRTYkL1RTrmDsrZ12SHcZUnEW0fbqvf4up6KqbgD71CrxOhDqg4TdUs-60hBFvxWj_Jftps1bgUpTcYoe9wTqzIhvwnYRb2cm8uHf2LSA7GUsHRVZP9Dac"
	token, err := verify(accessToken, keySet)
	if err != nil {
		fmt.Printf("Gor an error while verifiying access token: %v\n", err)
	}

	// Check if the token is valid.
	if !token.Valid {
		fmt.Println("The token is not valid.")
	}

	// fmt.Printf("Got access token: %v\n", token)

	claims := token.Claims.(jwt.MapClaims)
	// for key, value := range claims {
	// 	fmt.Printf("%s\t%v\n", key, value)
	// }
	metaData := MetaData{
		PublicURL:   "https://nightlybuild.cidaas.de",
		Sub:         claims["sub"].(string),
		AccessToken: accessToken,
	}
	fmt.Printf("MetaData : %v\n", metaData)
}

func verify(tokenString string, keySet *jwk.Set) (*jwt.Token, error) {
	tkn, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwa.RS256.String() { 
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("kid header not found")
		}
		keys := keySet.LookupKeyID(kid)
		if len(keys) == 0 {
			return nil, fmt.Errorf("key %v not found", kid)
		}
		var raw interface{}
		return raw, keys[0].Raw(&raw)
	})
	return tkn, err
}
