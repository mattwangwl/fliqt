package main

import (
	"context"
	"fliqt/internal/api"
	"fliqt/internal/migration"
)

func main() {
	ctx := context.Background()
	migration.New().Migrate(ctx)
	api.New().Run()
}
