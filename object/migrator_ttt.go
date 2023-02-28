package object

import (
	"errors"

	"github.com/xorm-io/xorm/migrate"
)

type Migrator_ttt struct{}

func (*Migrator_ttt) IsMigrationNeeded() bool {
	return true
}

func (*Migrator_ttt) DoMigration() *migrate.Migration {
	panic(errors.New("aaaaaaaaaa"))
}
