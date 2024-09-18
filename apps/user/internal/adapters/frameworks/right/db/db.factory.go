package db

import (
	"context"
	"errors"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/adapters/frameworks/right/db/mongodb"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/ports"
)

func CreateDbByType(ctx context.Context, dbType DbType, url string, dbName string) (ports.DbPort, error) {
	ctx = context.WithValue(ctx, "database", dbName)
	switch dbType {
	case Mongo:
		return mongodb.NewMongoApi(ctx, url), nil
	}

	return nil, errors.New("database type is incorrect")
}
