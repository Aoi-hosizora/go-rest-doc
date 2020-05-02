package restdoc

import (
	"fmt"
	"log"
	"testing"
)

func TestGenerateYaml(t *testing.T) {
	SetDocument(
		"localhost:10086", "/",
		NewInfo("test-api", "a demo description", "1.0").
			SetTermsOfService("http://xxx.yyy.zzz").
			SetLicense(NewLicense("MIT", "http://xxx.yyy.zzz")).
			SetContact(NewContact("author", "http://xxx.yyy.zzz", "xxx@yyy.zzz")),
	)
	SetTags(
		NewTag("ping", "ping-controller"),
		NewTag("user", "user-controller"),
	)
	SetSecurities(
		NewSecurity("jwt", HEADER, "Authorization"),
	)

	AddPaths(
		NewPath(GET, "/api/v1/ping", "ping").
			SetDescription("ping the server").
			SetTags("ping").
			SetConsumes(JSON).
			SetProduces(JSON).
			SetResponses(
				NewResponse(200).SetDescription("success").SetExamples(map[string]string{JSON: "{\n\t\"ping\": \"pong\"\n}"}),
			),
		NewPath(GET, "/api/v1/user", "get user").
			SetDescription("get user from database").
			SetTags("user").
			SetConsumes(JSON).
			SetProduces(JSON).
			SetSecurities("jwt").
			SetParams(
				NewParam("page", QUERY, INTEGER, false, "current page").SetDefault(1),
				NewParam("total", QUERY, INTEGER, false, "page size").SetDefault(10),
				NewParam("order", QUERY, STRING, false, "order string").SetDefault(""),
			).
			SetResponses(
				NewResponse(200).SetSchema(NewSchemaRef("Result<Page<User>>")),
			),
		NewPath(PUT, "/api/v1/user/{id}", "update user (ugly api)").
			SetDescription("update user to database").
			SetTags("user").
			SetConsumes(JSON).
			SetProduces(JSON).
			SetSecurities("jwt").
			SetParams(
				NewParam("id", PATH, INTEGER, true, "user id"),
				NewParam("body", BODY, OBJECT, true, "request body").SetSchema(NewSchemaRef("User")),
			).
			SetResponses(
				NewResponse(200).SetDescription("success").SetSchema(NewSchemaRef("Result")).SetHeaders(NewHeader("Content-Type", STRING, "demo")),
				NewResponse(404).SetDescription("not found"),
				NewResponse(400).SetDescription("bad request").SetSchema(NewSchema(STRING, true)).SetExamples(map[string]string{JSON: "bad request"}),
			),
		NewPath(HEAD, "/api/v1/test", "test path").
			SetParams(
				NewParam("arr", QUERY, ARRAY, true, "test").SetItems(NewItems(INTEGER).SetFormat(INT64).SetItems(NewItems(INTEGER))),
				NewParam("ref", QUERY, ARRAY, true, "test").SetItems(NewItemsRef("User")),
				NewParam("enum", QUERY, STRING, true, "test").SetEnum("male", "female"),
			),
	)

	AddDefinitions(
		NewDefinition("Result", "global response").SetProperties(
			NewProperty("code", INTEGER, true, "status code"),
			NewProperty("message", STRING, true, "status message"),
		),
		NewDefinition("User", "user response").SetProperties(
			NewProperty("id", INTEGER, true, "user id"),
			NewProperty("name", STRING, true, "user name"),
			NewProperty("profile", STRING, false, "user profile").SetAllowEmptyValue(true),
			NewProperty("gender", STRING, true, "user gender").SetEnum("male", "female"),
			NewProperty("create_at", STRING, true, "user register time").SetFormat(DATETIME),
			NewProperty("birthday", STRING, true, "user birthday").SetFormat(DATE),
			NewProperty("scores", ARRAY, true, "user scores").SetItems(NewItems(NUMBER)),
		),
		NewDefinition("Page<User>", "user response").SetProperties(
			NewProperty("page", INTEGER, true, "current page"),
			NewProperty("total", INTEGER, true, "data count"),
			NewProperty("limit", INTEGER, true, "page size"),
			NewArrayProperty("data", NewItemsRef("User"), true, "page data"),
		),
		NewDefinition("Result<Page<User>>", "user response").SetProperties(
			NewProperty("code", INTEGER, true, "status code"),
			NewProperty("message", STRING, true, "status message"),
			NewObjectProperty("data", "Page<User>", true, "result data"),
		),
	)

	// doc, _ := yaml.Marshal(appendKvs(buildDocument(_document), map[string]interface{}{"swagger": "2.0"}))
	// fmt.Println(string(doc))

	doc, _ := jsonMarshal(appendKvs(buildDocument(_document), map[string]interface{}{"swagger": "2.0"}))
	fmt.Println(string(doc))

	err := GenerateYaml("./docs/api.yaml", map[string]interface{}{"swagger": "2.0"})
	log.Println(err)

	err = GenerateJson("./docs/api.json", map[string]interface{}{"swagger": "2.0"})
	log.Println(err)
}
