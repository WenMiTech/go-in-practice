package pkgb

import (
	"fmt"
)

func FuncFromPkga() {
	//from the same package,so no need to import
	//and call with a package name
	WithSamePackage()
	fmt.Println("FuncFromPkga")
}
