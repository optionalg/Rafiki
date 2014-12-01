Rafiki
=========

Rafiki is a simple SSL cert storage system written in Golang.

![rafiki](https://raw.githubusercontent.com/adamar/rafiki/master/doc/rafiki.gif)


Rafiki is a CLI tool for securely storing SSL cert files in a local SQLite3 Database. Imported files are first encrypted using GPG and then stored
in the database along with an identifying key (ie. CommonName from CSRs). 

The database will be created when Rafiki is run for the first time and can be re-located and referenced by Rafiki using the --db flag. 


Usage
--------------

Import a CSR file
```sh
./Rafiki csr import --file=/loc/of/file.csr
```

List CSRs
```sh
./Rafiki csr list
```

Export CSR to a file
```sh
./Rafiki csr export --file=/loc/of/newfile.csr
```


Dependencies
-------------

- SQLite v3+
- Go 1.3+



Development Misc
-------------

Useful Testing Tools

Print Public Key Fingerprint
```sh
ssh-keygen -lf /path/to/key.pub
```

Print CSR Info
openssl req -in domain.com.csr -text -noout

Show CSR Public Key
openssl req -in domain.com.csr -noout -pubkey

Show an RSA Key's thumbprint
openssl rsa -noout -modulus -in your-private.key | openssl sha1


ToDo
-----------
- Write Tests
- Add more error checking
- Add key delete function


Useful Resources
------------

http://redkestrel.co.uk/articles/CSR-FAQ/

* [CSR FAQ] - Certificate Signing Request FAQ



[CSR FAQ]:http://redkestrel.co.uk/articles/CSR-FAQ/


