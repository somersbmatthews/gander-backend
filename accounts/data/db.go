package accounts

import (
	"fmt"

	"github.com/go-pg/pg"
)

func init() {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "user",
		Password: "pass",
		Database: "gander_accounts",
	})

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
   		panic(err)
	}
}
