package rafiki

import (
	"crypto/x509"
	//"encoding/hex"
	"encoding/pem"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"os"
)

func ExportCSR(c *cli.Context) {

	_, err := checkDB(c.String("db"))
    if err != nil {
        log.Print(err)
    }

	conn := createDBConn(c.String("db"))
	defer conn.Close()

	//key, err := startUp()
	//log.Print(err)

	keyname := GetKeyName()
    log.Print(keyname)

	ciphertext := SelectKey(conn, keyname)

	//cleartext, err := DecryptString(key, ciphertext)
    
    err = ioutil.WriteFile(c.String("file"), []byte(ciphertext), 0644)
    if err != nil {
        panic(err)
    }

}

func ImportCSR(c *cli.Context) {

	password, _ := checkDB(c.String("db"))
    log.Print(password)
	conn := createDBConn(c.String("db"))

	defer conn.Close()

	//password, err := startUp()
	//log.Print(err)

	err := CheckFileFlag(c)
	ErrHandler(err)

	buf, err := ioutil.ReadFile(c.String("f"))
	ErrHandler(err)

	block, _ := pem.Decode(buf)

	CertificateRequest, err := x509.ParseCertificateRequest(block.Bytes) //Requires Go 1.3+
	ErrHandler(err)

	CSRName := CertificateRequest.Subject

	log.Print(CSRName.CommonName)
	//log.Print(CertificateRequest.SignatureAlgorithm)

	//log.Print(string(hex.Dump(CertificateRequest.Signature)))
	//log.Print(string(hex.EncodeToString(CertificateRequest.Signature)))

	//key := []byte(password)

	ciphertext, err := EncryptString([]byte(password), string(block.Bytes))
	//ErrHandler(err)

    log.Print(ciphertext)
	InsertKey(conn, string(CSRName.CommonName), ciphertext)

}

func DeleteCSR(c *cli.Context) {

	log.Print("csr delete")

}

func ListCSR(c *cli.Context) {

	log.Print("csr list")
	checkDB(c.String("db"))
	conn := createDBConn(c.String("db"))
	defer conn.Close()

	ListKeys(conn)

}

func checkCSRFileSet(value bool) {
	if value == false {
		log.Print("No CSR file specified")
		os.Exit(1)
	}

	//if _, err := os.Stat(filename); os.IsNotExist(err) {
	//     log.Print("File not found")
	//     os.Exit(1)
	// }

}
