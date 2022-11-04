package test

import (
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
    db, mock, err := sqlmock.New()
    if err != nil {
        return nil, nil, err
    }

    gdb, err := gorm.Open("postgres", db)
    if err != nil {
        return nil, nil, err
    }
    return gdb, mock, nil
}
func TestCreate(t *testing.T) {
    db, mock, err := getDBMock()
    if err != nil {
        t.Fatal(err)
    }
    defer db.Close()
    db.LogMode(true)

    r := Repository{DB: db}

    id := "2222"
    name := "BBBB"

    // Mock設定
    mock.ExpectQuery(regexp.QuoteMeta(
        `INSERT INTO "users" ("id","name") VALUES ($1,$2)
         RETURNING "users"."id"`)).
        WithArgs(id, name).
        WillReturnRows(
            sqlmock.NewRows([]string{"id"}).AddRow(id))

    // 実行
    err = r.Create(id, name)
    if err != nil {
        t.Fatal(err)
    }
}