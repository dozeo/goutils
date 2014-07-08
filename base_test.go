package goutils

import (
	"fmt"

	. "github.com/smartystreets/goconvey/convey"
)

func MyConvey(items ...interface{}) {
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println("-- " + items[0].(string))
	fmt.Println("-------------------------------------------------------------------------")
	Convey(items...)
}
