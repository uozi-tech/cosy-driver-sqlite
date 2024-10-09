package sqlite

import (
	"fmt"
	"github.com/uozi-tech/cosy/settings"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path"
)

func Open(dir string, dbs *settings.DataBase) gorm.Dialector {
	return sqlite.Open(path.Join(dir, fmt.Sprintf("%s.db", dbs.Name)))
}
