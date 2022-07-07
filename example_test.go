package json_parser_test

import (
	"fmt"

	json_parser "aa.bb.cc/util/json"
)

func ExampleParse() {
	fmt.Println(json_parser.Parse(""))
}
