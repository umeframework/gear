package services

import (
	"github.com/umeframework/gear/httpd"
	"fmt"
)

// Test for simple post DTO
// [POST] /postDto
// REQUEST BODY:
//	{"id": 100, "name": "jackson"}
type postDtoInDTO struct {
	Id int
	Name string
}

func postDto(context httpd.HttpRequestContext, inDTO postDtoInDTO) string {
	return fmt.Sprintf("id: %d, name: %s", inDTO.Id, inDTO.Name)
}

// Test for simple post DTO as map
// [POST] /postDto2
// REQUEST BODY:
//	{"id": 100, "name": "jackson"}

func postDto2(context httpd.HttpRequestContext, inDTO map[string]interface{}) string {
	return fmt.Sprintf("%v", inDTO)
}


// Test for combination of DTO and query params
// [POST] /postDto3?id=100&name=tom
// REQUEST BODY:
//	{"name": "jerry", "year": 2018}
type postDto3InDTO struct {
	Id int
	Name string
	Year int
}

func postDto3(context httpd.HttpRequestContext, inDTO postDto3InDTO) string {
	return fmt.Sprintf("id: %d, name: %s, year: %d", inDTO.Id, inDTO.Name, inDTO.Year)
}


func init() {
	httpd.NewServicePoint("/postDto", nil, postDto)
	httpd.NewServicePoint("/postDto2", nil, postDto2)
	httpd.NewServicePoint("/postDto3", nil, postDto3)
}
