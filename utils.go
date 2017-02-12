package main

import (
	"database/sql"
	"gopkg.in/gorp.v1"
	"log"
	"fmt"
		"io"
		"io/ioutil"
		"os"
"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:98muzunu@/tests")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	checkErr(err, "Create tables failed")
	// _, err = dbmap.Exec("ALTER TABLE post DROP text;")
	// _, err = dbmap.Exec("ALTER TABLE post ADD text VARCHAR(65535) AFTER title;")
	// _, err = dbmap.Exec("ALTER TABLE post DROP tags;")
	// _, err = dbmap.Exec("ALTER TABLE post ADD tags VARCHAR(65535) AFTER truncated;")
	// _, err = dbmap.Exec("ALTER TABLE post DROP summary;")
	// _, err = dbmap.Exec("ALTER TABLE post ADD summary VARCHAR(65535) AFTER tags;")
	// checkErr(err, "Setting mode failed")
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func loadData(p string) ([]byte, error) {
	if p == "" {
		return nil, fmt.Errorf("No path specified")
	}

	var rdr io.Reader
	if p == "-" {
		rdr = os.Stdin
	} else {
		if f, err := os.Open(p); err == nil {
			rdr = f
			defer f.Close()
		} else {
			return nil, err
		}
	}
	return ioutil.ReadAll(rdr)
}
