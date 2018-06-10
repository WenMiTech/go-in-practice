package pkga

import (
	"fmt"
	"package/pkgb"
)

func TestA() {
	pkgb.FuncFromPkga()
	fmt.Println("test a")
}
