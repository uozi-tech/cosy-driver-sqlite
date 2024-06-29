package sqlite

import (
	"fmt"
	"git.uozi.org/uozi/cosy/settings"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path"
)

func Open(dir string, dbs *settings.DataBase) gorm.Dialector {
	return sqlite.Open(path.Join(dir, fmt.Sprintf("%s.db", dbs.Name)))
}
