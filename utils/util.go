package utils

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
)

func Verify_User(client *twitter.Client) error {

	user, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})
	if err != nil {
		return err
	}
	fmt.Printf("Account : @%s (%s)\n", user.ScreenName, user.Name)
	return nil
}
