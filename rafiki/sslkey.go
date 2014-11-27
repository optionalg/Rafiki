package rafiki

import (
    _ "github.com/mattn/go-sqlite3"
	"github.com/codegangsta/cli"
)


var SSLCommand = cli.Command{
    Name:        "ssl",
    Usage:       "ssl",
    Description: "example SSL blah",
    Subcommands: []cli.Command{
        {
            Name:  "export",
            Usage: "Export a CSR Cert from the DB",
            Flags: []cli.Flag{
                FileLoc,
                DBLoc,
            },
            Action: func(c *cli.Context) {

               raf := NewRafikiInit(c)
               raf.Export()

            },
        },
        {
            Name:  "import",
            Usage: "Import an SSL Cert into the DB",
            Flags: []cli.Flag{
                FileLoc,
                DBLoc,
            },
            Action: func(c *cli.Context) {

               raf := NewRafikiInit(c)
               raf.Import("csr")

            }, 
        },
        {
            Name:  "delete",
            Usage: "Delete an SSL Cert from the DB",
            Flags: []cli.Flag{
                FileLoc,
                DBLoc,
            },
            Action: func(c *cli.Context) {

               raf := NewRafikiInit(c)
               raf.Delete()

            },
        },
        {
            Name:  "list",
            Usage: "List all SSL Certs in the DB",
            Flags: []cli.Flag{
                DBLoc,
            },
            Action: func(c *cli.Context) {

               raf := NewRafikiInit(c)
               raf.List("sslkey")
               
            },
        },
    },
}
