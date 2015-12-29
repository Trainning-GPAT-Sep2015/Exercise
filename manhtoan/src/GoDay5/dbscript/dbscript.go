package dbscript

import (
	"fmt"
	"strings"

	"github.com/dancannon/gorethink"

	"GoDay5/api/rethinkdb"
)

type RethinkScript struct {
	re     *rethinkdb.Instance
	dbname string
}

func NewRethinkScript(re *rethinkdb.Instance, dbname string) *RethinkScript {
	if len(strings.TrimSpace(dbname)) == 0 {
		fmt.Println("Empty database name")
	}
	return &RethinkScript{
		re:     re,
		dbname: dbname,
	}
}

func (this *RethinkScript) Setup() error {
	dbname := this.dbname
	err := this.re.Exec(gorethink.DB(dbname).Info())

	if err == nil {
		fmt.Printf("Database `%v` exists, skip creating.", dbname)
		return err
	}

	if !strings.Contains(err.Error(), "does not exist") {
		fmt.Printf("Error querying database `%v`: %v", dbname, err)
		return err
	}

	fmt.Printf("Creating database `%v`", dbname)
	err = this.re.Exec(gorethink.DBCreate(dbname))
	if err != nil {
		return err
	}

	// err = this.createTableUser()
	return err
}

// func (this *RethinkScript) createTableUser() error {
// 	err := this.re.Exec(this.re.DB().TableCreate(stores.UserTable))
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("Create index: username")
// 	err = this.re.Exec(this.re.Table(stores.UserTable).IndexCreate("username"))
// 	if err != nil {
// 		return err
// 	}

// 	err = this.re.Exec(this.re.Table(stores.UserTable).IndexWait())
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
