package main

import (
	"context"
	"encoding/hex"
	"example/hello/http_with_ent/ent"
	"log"

	"github.com/spf13/viper"
)

type contextKey string

const userContextKey = contextKey("user")

var pasetoKey []byte

func getPasetoKey() []byte {
	pasetoKeyHex := viper.GetString("PASETO_KEY")
	pasetoKey, err := hex.DecodeString(pasetoKeyHex)
	if err != nil {
		log.Panic("failed to generate PASETO key: ", err)
	}
	return pasetoKey
}

// GetUserFromContext retrieves the User from the request context.
func GetUserFromContext(ctx context.Context) *ent.User {
	if u, ok := ctx.Value(userContextKey).(*ent.User); ok {
		return u
	}
	return nil
}
