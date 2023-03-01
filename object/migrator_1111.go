package object

import (
	"errors"
	"github.com/xorm-io/xorm/migrate"
)

type Migrator_1111 struct{}

func (*Migrator_1111) IsMigrationNeeded() bool {
	return true
}

func (*Migrator_1111) DoMigration() *migrate.Migration {
	panic(errors.New("panic migration"))
	panic(errors.New("panic migration"))
}
