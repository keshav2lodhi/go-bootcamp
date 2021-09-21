package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
)

type Fields struct {
	PublicURL   string
	TenantKey   string
	AccessToken string
	Aud         []string
	Sub         string
}

func main() {

	jsonData := []byte(`
	{
    "PublicURL": "",
    "TenantKey": "cidaas-nightlybuild-dev",
	"AccessToken": "xyz",
	"Sub": "abc"
	}`)
	validate(jsonData)
}

func validate(jsonData []byte) {
	var fields Fields
	err := json.Unmarshal(jsonData, &fields)
	if err != nil {
		fmt.Printf("fact could not be processed. Unable to unmarshal message (%v)", string(jsonData))
		return
	} else if fields.PublicURL == "" || fields.Sub == "" {
		log.Error().Bool("internal_use", true).Err(err).Msg(fmt.Sprintf("The mandatory fields are missing for fact (%v)", fields))
		return
	}
	fmt.Printf("fact data %s", fields.PublicURL)
}
