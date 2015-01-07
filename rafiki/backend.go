package rafiki

import (
	"crypto/sha256"
    "crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
    "log"
    "os"
	"github.com/bndr/gotabulate"
	"github.com/codegangsta/cli"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(c *cli.Context) *sql.DB {

	fname := c.String("db")

	if _, err := os.Stat(fname); os.IsNotExist(err) {
		log.Print("db doesnt exit")
		PromptToCreateDB()
	}

	db := createDBConn(fname)

	return db

}

func PromptToCreateDB() {

	msg := "No DB Specified, Y/N to create a new one"
	var i string
	fmt.Println(msg)
	fmt.Scan(&i)
	if i == "y" {
		CreateDB()
	} else {
		os.Exit(1)
	}

}

func GetKeyName() string {

	msg := "Which key?"
	var i string
	fmt.Println(msg)
	fmt.Scan(&i)

	return i

}

func CreateDB() error {

	// Create Database File
	//
	db, err := sql.Open("sqlite3", "./rafiki.db")
	if err != nil {
		return err
	}

	// Generate Schema for DB
	//
	sqlStmt := `
        create table files (id integer not null primary key,
                          identity text,
                          type text,
                          filename text,
                          file blob);
        `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	// Create password table
	//
	sqlStmt = `CREATE TABLE password (
                hashed_password text NOT NULL);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	return nil

}

func ListKeys(db *sql.DB, fileType string) error {

	new := [][]string{}


    var query string
    if fileType != "" {
        query = fmt.Sprintf("select id, identity, filename, type from files WHERE type = %s", fileType)
    } else {
        query = fmt.Sprintf("select id, identity, filename, type from files order by type")
    }

    rows, err := db.Query(query)

	if err != nil {
		return err
	}

	for rows.Next() {
		var id string
		var identity string
        var filename string
        var ftype string
		rows.Scan(&id, &identity, &filename, &ftype)
		new = append(new, []string{id, identity, filename, ftype})
	}
	rows.Close()

	tabulate := gotabulate.Create(new)
	tabulate.SetHeaders([]string{"ID", "CommonName", "Filename", "File Type"})

	if len(new) > 0 {
		fmt.Println(tabulate.Render("grid"))
	} else {
		fmt.Println("No Keys to Print\n")
	}

	return nil
}




func checkDB(fname string) (password string, err error) {

	if _, err := os.Stat(fname); os.IsNotExist(err) {
		log.Print("db doesnt exit")
		PromptToCreateDB()
		//password, err := setPassword()
		return password, err
	} else {
		//password, err := checkPassword()
		return password, err
	}
	return password, nil

}

func InsertKey(db *sql.DB, identity string, fileType string, fileContents string, fileName string) error {

	_, err := db.Exec("INSERT INTO files (identity, type, filename, file) VALUES (?, ?, ?, ?)", identity, fileType, fileName, fileContents)
	if err != nil {
		return err
	}

	return nil

}

func DeleteKey(db *sql.DB, kId string) error {

	_, err := db.Exec("DELETE FROM files WHERE id = ?", kId)
	if err != nil {
		return err
	}

	return nil

}

func SelectKey(db *sql.DB, id string) string {

	rows, err := db.Query("SELECT file from files WHERE id = ? ", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var csr string = ""
	for rows.Next() {
		rows.Scan(&csr)
	}

	return csr

}

func InsertPassword(db *sql.DB, password string) error {

	_, err := db.Exec("INSERT INTO password (hashed_password) VALUES (?)", password)
	if err != nil {
		return err
	}

	return nil

}

func SelectPassword(db *sql.DB) (string, error) {

	rows, err := db.Query("SELECT hashed_password from password LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pass string
	for rows.Next() {
		rows.Scan(&pass)
	}

	return pass, nil

}

func CheckIsPasswordSet(db *sql.DB) (string, error) {

	var count string

	err := db.QueryRow("SELECT COUNT(hashed_password) from password").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	//defer rows.Close()

	//var pass string
	//for rows.Next() {
	//    rows.Scan(&pass)
	//}

	return count, nil

}

func createDBConn(fname string) *sql.DB {

	db, err := sql.Open("sqlite3", fname)
	if err != nil {
		log.Print(err)
	}

	return db
}

func hashStringToSHA256(input string) string {

	hash := sha256.New()
	hash.Write([]byte(input))
	chkSum := hash.Sum(nil)
	return hex.EncodeToString(chkSum)

}

func md5String(input string) string {

    hash := md5.New()
    hash.Write([]byte(input))
    return hex.EncodeToString(hash.Sum(nil))

}


func formatMd5(input string) string {

    i := 0
    final := ""

    for _, c := range input {
           
        final = final + string(c)
        i ++

        if i == len(input) {
            break
        }

        if i%2 == 0 {
            final = final + ":"
        }
    }

    return final
}



