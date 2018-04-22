package services

import (
	"github.com/umeframework/gear/httpd"
	"fmt"
)

// Test for simple query params & return dto
// [GET] /testGet?x=hello&y=world
type testGetInDTO struct {
	X, Y string // Must be in capital (exportable)
}

type testGetOutDTO struct {
	X string // can NOT be exported to json
	Y string // can be exported to json
	CombinedLength int // can be exported to json
}

func testGet(context httpd.HttpRequestContext, inDTO testGetInDTO) testGetOutDTO {
	return testGetOutDTO{
		inDTO.X, inDTO.Y, len(inDTO.X + " - " + inDTO.Y),
	}
}

// Test for different kind of query params
// [GET] /testGet2?id=100&name=tom
type testGet2InDTO struct {
	Id int
	Name string
}

func testGet2(context httpd.HttpRequestContext, inDTO testGet2InDTO) string {
	return fmt.Sprintf("id = %d, name = %s", inDTO.Id, inDTO.Name)
}


// TODO: implement for path params
// Test for combination of query params and path params
// [GET] /testGet3/2018?id=100&name=test
type testGet3InDTO struct {
	Id int
	Name string
	Year int
}

func testGet3(context httpd.HttpRequestContext, inDTO testGet3InDTO) string {
	return fmt.Sprintf("id = %d, name = %s, year = %d", inDTO.Id, inDTO.Name, inDTO.Year)
}

func init() {
	httpd.NewServicePoint("/testGet", nil, testGet)
	httpd.NewServicePoint("/testGet2", nil, testGet2)
	httpd.NewServicePoint("/testGet3", nil, testGet3)
}
