Rafiki
=========

Rafiki is a SSL cert storage system written in Golang.


Usage
--------------

```sh
./rafiki csr import
```



Misc
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




Useful Resources
------------

http://redkestrel.co.uk/articles/CSR-FAQ/

* [CSR FAQ] - Certificate Signing Request FAQ



[CSR FAQ]:http://redkestrel.co.uk/articles/CSR-FAQ/


