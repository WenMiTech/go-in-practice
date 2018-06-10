package main

//run go install
// cd /gopath/src/package (dirname)
// go install

import (
	"package/pkga" //project(gopath)/package/pkga(dirname as package name)/*.go
	"package/pkgb"
)

func main() {
	pkga.TestA()
	pkgb.FuncFromPkga()
}
