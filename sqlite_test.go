package sqlite

import (
	"git.uozi.org/uozi/cosy/settings"
	"gorm.io/driver/sqlite"
	"testing"
)

func TestOpen(t *testing.T) {
	dbs := &settings.DataBase{
		Name: "cosy",
	}

	dialector := Open("", dbs)

	d, ok := dialector.(*sqlite.Dialector)
	if !ok {
		t.Fatal("dialector is not *Dialector")
	}
	if d.DSN == "" {
		t.Error("dialector.DSN is empty")
	}
}
