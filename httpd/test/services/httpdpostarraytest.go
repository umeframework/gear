package services

import (
	"github.com/umeframework/gear/httpd"
	"fmt"
)

// Test for simple ordered arguments
// [POST] /postArray
// REQUEST BODY:
//	[1, "A"]
func postArray(context httpd.HttpRequestContext, n int, text string) string {
	return fmt.Sprintf("%v-%v", n, text)
}

// Test for ordered arguments as slice
// [POST] /postArray2
// REQUEST BODY:
//	[5, 6, 7, 8]
func postArray2(context httpd.HttpRequestContext, args []int) int {
	var sum = 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

// Test for simple ordered arguments
// [POST] /service3
// REQUEST BODY:
//	["A", 1, "B", 2.5]
func postArray3(context httpd.HttpRequestContext, args []interface{}) string {
	var text = ""
	for _, arg := range args {
		text += fmt.Sprintf("%v", arg)
	}
	return text
}



func init() {
	httpd.NewServicePoint("/postArray", nil, postArray)
	httpd.NewServicePoint("/postArray2", nil, postArray2)
	httpd.NewServicePoint("/postArray3", nil, postArray3)
}
