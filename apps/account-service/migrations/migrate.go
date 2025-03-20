package main

import (
	entityaccounts "account-system/apps/account-service/src/accounts/infrastructure/persistence/orm/entities"
	entityusers "account-system/apps/account-service/src/users/infrastructure/persistence/orm/entities"
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
)

func migrate() {
	stmts, err := gormschema.New("postgres").Load(&entityaccounts.AccountEntity{}, &entityusers.UserEntity{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}

func main() {
	migrate()
}
