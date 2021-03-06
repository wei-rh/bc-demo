module github.com/wei-rh/bc-demo

go 1.15

require (
	github.com/mr-tron/base58 v1.2.0
	github.com/syndtr/goleveldb v1.0.0
	github.com/tyler-smith/go-bip32 v1.0.0
	github.com/tyler-smith/go-bip39 v1.1.0
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
)

replace ./golang.org/x/sys => ./github.com/golang/sys

replace ./golang.orgx/crypto => ./github.com/golang/crypto
