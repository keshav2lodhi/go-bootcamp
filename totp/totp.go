package main

import (
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/spf13/viper"
)

func main() {

	viper.SetDefault("TOTP_KEY", "HJYHCMKROMXS6W2CM5TWGO2RKAQXORRR") //dcost-nightlybuild
	// viper.SetDefault("TOTP_KEY", "PE7CIKTJJRBWSXRRHZDDQPROJVYV4MTP") //kubet-nightlybuild

	totpToken, err := totp.GenerateCode(viper.GetString("TOTP_KEY"), time.Now().Add(time.Second*5))
	if err != nil {
		fmt.Println("Error while generating code ", err)
	}

	fmt.Println("totpToken ", totpToken)
}
