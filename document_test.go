package goapidoc

import (
	"testing"
)

func TestSetGet(t *testing.T) {
	t.Cleanup(func() {
		SetDocument("", "", nil)
		SetOperations()
		SetDefinitions()
	})
	t.Run("Set and get in document.go", func(t *testing.T) {
		SetDocument("", "", nil)
		SetHost("localhost:12334")
		SetBasePath("v1")
		SetInfo(NewInfo("", "", "").
			Title("Test api").
			Desc("This is a test api").
			Version("v0.0.1").
			TermsOfService("ToS").
			License(NewLicense("", "").
				Name("MIT").
				Url("https://opensource.org/licenses/MIT")).
			Contact(NewContact("", "", "").
				Name("Aoi-hosizora").
				Url("https://github.com/Aoi-hosizora").
				Email("a970335605@hotmail.com")))
		SetOption(NewOption().
			Schemas("http").
			AddSchemas("https", "ws", "wss").
			Consumes("application/json").
			AddConsumes("multipart/form-data", "application/protobuf").
			Produces("application/json").
			AddProduces("application/xml", "application/protobuf").
			Tags(NewTag("", "").
				Name("Authorization").
				Desc("auth-controller")).
			AddTags(NewTag("User", "user-controller"),
				NewTag("Resource", "resource-controller")).
			Securities(NewSecurity("", "").
				Title("jwt").
				Type(APIKEY).
				Desc("A apiKey security called jwt").
				InLoc(HEADER).
				Name("Authorization")).
			AddSecurities(NewBasicSecurity("basic"),
				NewApiKeySecurity("another_jwt", HEADER, "X-JWT")))
		AddOperations(NewOperation("", "", ""))
		SetOperations(NewOperation("", "", ""),
			NewOperation("", "", ""))
		AddDefinitions(NewDefinition("", ""))
		SetDefinitions(NewDefinition("", ""),
			NewDefinition("", ""))

		if GetHost() != "localhost:12334" {
			failNow(t, "SetHost has a wrong behavior")
		}
		if GetBasePath() != "v1" {
			failNow(t, "SetBasePath has a wrong behavior")
		}
		if GetInfo().GetTitle() != "Test api" {
			failNow(t, "Info.Title has a wrong behavior")
		}
		if GetInfo().GetDesc() != "This is a test api" {
			failNow(t, "Info.Desc has a wrong behavior")
		}
		if GetInfo().GetVersion() != "v0.0.1" {
			failNow(t, "Info.Version has a wrong behavior")
		}
		if GetInfo().GetTermsOfService() != "ToS" {
			failNow(t, "Info.TermsOfService has a wrong behavior")
		}
		if GetInfo().GetLicense().GetName() != "MIT" {
			failNow(t, "Info.License.Name has a wrong behavior")
		}
		if GetInfo().GetLicense().GetUrl() != "https://opensource.org/licenses/MIT" {
			failNow(t, "Info.License.Url has a wrong behavior")
		}
		if GetInfo().GetContact().GetName() != "Aoi-hosizora" {
			failNow(t, "Info.Contact.Name has a wrong behavior")
		}
		if GetInfo().GetContact().GetUrl() != "https://github.com/Aoi-hosizora" {
			failNow(t, "Info.Contact.Url has a wrong behavior")
		}
		if GetInfo().GetContact().GetEmail() != "a970335605@hotmail.com" {
			failNow(t, "Info.Contact.Email has a wrong behavior")
		}
		if s := GetOption().GetSchemas(); s[0] != "http" || s[1] != "https" || s[2] != "ws" || s[3] != "wss" {
			failNow(t, "Option.Schemas or Option.AddSchemas has a wrong behavior")
		}
		if c := GetOption().GetConsumes(); c[0] != "application/json" || c[1] != "multipart/form-data" || c[2] != "application/protobuf" {
			failNow(t, "Option.Consumes or Option.AddConsumes has a wrong behavior")
		}
		if p := GetOption().GetProduces(); p[0] != "application/json" || p[1] != "application/xml" || p[2] != "application/protobuf" {
			failNow(t, "Option.Produces or Option.AddProduces has a wrong behavior")
		}
		if a := GetOption().GetTags(); a[0].GetName() != "Authorization" || a[0].GetDesc() != "auth-controller" ||
			a[1].GetName() != "User" || a[1].GetDesc() != "user-controller" || a[2].GetName() != "Resource" || a[2].GetDesc() != "resource-controller" {
			failNow(t, "Option.Tags or Option.AddTags or Tags.XXX has a wrong behavior")
		}
		if s := GetOption().GetSecurities(); s[0].GetTitle() != "jwt" || s[0].GetType() != "apiKey" || s[0].GetDesc() != "A apiKey security called jwt" || s[0].GetInLoc() != "header" || s[0].GetName() != "Authorization" ||
			s[1].GetTitle() != "basic" || s[1].GetType() != "basic" || s[2].GetTitle() != "another_jwt" || s[2].GetType() != "apiKey" || s[2].GetInLoc() != "header" || s[2].GetName() != "X-JWT" {
			failNow(t, "Option.Securities or Option.AddSecurities or Security.XXX has a wrong behavior")
		}
		if len(GetOperations()) != 2 {
			failNow(t, "AddOperations or SetOperations has a wrong behavior")
		}
		if len(GetDefinitions()) != 2 {
			failNow(t, "AddDefinitions or SetDefinitions has a wrong behavior")
		}
	})

	t.Run("Set and get in operation.go", func(t *testing.T) {
		SetOperations()
		AddOperations(NewGetOperation("/user", "Get all users"),
			NewPutOperation("/user/{id}", "Replace user"),
			NewPostOperation("/user", "Add new user"),
			NewDeleteOperation("/user/{id}", "Delete user"),
			NewOptionsOperation("/user", "Show options of route"),
			NewHeadOperation("/user", "Show head of route"),
			NewPatchOperation("/user/{id}", "Update user"))
		AddOperations(NewOperation("", "", "").
			Method(GET).
			Route("/user/{id}").
			Summary("Get user information").
			Desc("Query a specific user with id's information").
			OperationId("-user-:id").
			Schemas("http").
			AddSchemas("https", "ws", "wss").
			Consumes("application/json").
			AddConsumes("multipart/form-data", "application/protobuf").
			Produces("application/json").
			AddProduces("application/xml", "application/protobuf").
			Tags("User").
			AddTags("Authorized", "Information").
			Securities("jwt").
			AddSecurities("basic", "another_jwt").
			Deprecated(true).
			Responses(NewResponse(404, "Result")).
			AddResponses(NewResponse(0, "").
				Code(200).
				Type("_Result<User>").
				Desc("200 OK").
				Examples(map[string]string{JSON: `{"code": "200", "message": "success", "data": {"id": 1, "name": "user1"}}`}).
				AddExample(XML, `<Result> <code>200</code> <message>success</message> </Result>`).
				Headers(NewHeader("X-RateLimit-Remaining", "integer#int64", "Request rate limit remaining")).
				AddHeaders(NewHeader("", "", "").
					Name("X-RateLimit-Limit").
					Type("integer#int64").
					Desc("Request rate limit size"))).
			Params(NewPathParam("id", "integer#int64", true, "user id"),
				NewQueryParam("need_details", "boolean", false, "is need details")).
			AddParams(NewFormParam("key", "string", false, "fake form key"), // only for test
				NewBodyParam("param", "Fake", false, "fake body"), // only for test
				NewHeaderParam("X-NEED-DETAILS", "boolean", false, "a duplicate parameter of need_details query")).
			AddParams(NewParam("", "", "", false, "").
				Name("X-TEST").
				InLoc(HEADER).
				Type("string").
				Required(true).
				Desc("A test header").
				AllowEmpty(true).
				Default("test###").
				Example("a###").
				Pattern("^.+###$").
				Enum("a###", "b###", "test###").
				MinLength(0).
				MaxLength(0).
				LengthRange(4, 8),
				NewParam("X-TEST2", HEADER, "integer[]", false, "Another test header").
					MinItems(0).
					MaxItems(0).
					ItemsRange(3, 18).
					UniqueItems(true).
					CollectionFormat(CSV).
					ItemOption(NewItemOption()),
				NewParam("X-TEST3", HEADER, "integer", false, "More than another test header").
					Minimum(0).
					Maximum(0).
					ValueRange(1.1, 3.3).
					ExclusiveMax(true).
					ExclusiveMin(true).
					MultipleOf(1.1)))

		if len(GetOperations()) != 8 {
			failNow(t, "AddOperations or SetOperations has a wrong behavior")
		}
		if o := GetOperations(); o[0].GetRoute() != "/user" || o[0].GetMethod() != "get" || o[0].GetSummary() != "Get all users" ||
			o[1].GetRoute() != "/user/{id}" || o[1].GetMethod() != "put" || o[1].GetSummary() != "Replace user" ||
			o[2].GetRoute() != "/user" || o[2].GetMethod() != "post" || o[2].GetSummary() != "Add new user" ||
			o[3].GetRoute() != "/user/{id}" || o[3].GetMethod() != "delete" || o[3].GetSummary() != "Delete user" ||
			o[4].GetRoute() != "/user" || o[4].GetMethod() != "options" || o[4].GetSummary() != "Show options of route" ||
			o[5].GetRoute() != "/user" || o[5].GetMethod() != "head" || o[5].GetSummary() != "Show head of route" ||
			o[6].GetRoute() != "/user/{id}" || o[6].GetMethod() != "patch" || o[6].GetSummary() != "Update user" {
			failNow(t, "NewXXXOperation has a wrong behavior")
		}
		op := GetOperations()[7]
		if op.GetMethod() != "get" {
			failNow(t, "Operation.Method has a wrong behavior")
		}
		if op.GetRoute() != "/user/{id}" {
			failNow(t, "Operation.Route has a wrong behavior")
		}
		if op.GetSummary() != "Get user information" {
			failNow(t, "Operation.Summary has a wrong behavior")
		}
		if op.GetDesc() != "Query a specific user with id's information" {
			failNow(t, "Operation.Desc has a wrong behavior")
		}
		if op.GetOperationId() != "-user-:id" {
			failNow(t, "Operation.OperationId has a wrong behavior")
		}
		if s := op.GetSchemas(); s[0] != "http" || s[1] != "https" || s[2] != "ws" || s[3] != "wss" {
			failNow(t, "Operation.Schemas or Operation.AddSchemas has a wrong behavior")
		}
		if c := op.GetConsumes(); c[0] != "application/json" || c[1] != "multipart/form-data" || c[2] != "application/protobuf" {
			failNow(t, "Operation.Consumes or Option.AddConsumes has a wrong behavior")
		}
		if p := op.GetProduces(); p[0] != "application/json" || p[1] != "application/xml" || p[2] != "application/protobuf" {
			failNow(t, "Operation.Produces or Option.AddProduces has a wrong behavior")
		}
		if a := op.GetTags(); a[0] != "User" || a[1] != "Authorized" || a[2] != "Information" {
			failNow(t, "Operation.Tags or Operation.AddTags has a wrong behavior")
		}
		if s := op.GetSecurities(); s[0] != "jwt" || s[1] != "basic" || s[2] != "another_jwt" {
			failNow(t, "Operation.Securities or Operation.AddSecurities has a wrong behavior")
		}
		if op.GetDeprecated() != true {
			failNow(t, "Operation.Deprecated has a wrong behavior")
		}
		if r := op.GetResponses()[0]; r.GetCode() != 404 || r.GetType() != "Result" {
			failNow(t, "NewResponse has a wrong behavior")
		}
		resp := op.GetResponses()[1]
		if resp.GetCode() != 200 {
			failNow(t, "Response.Code has a wrong behavior")
		}
		if resp.GetType() != "_Result<User>" {
			failNow(t, "Response.Type has a wrong behavior")
		}
		if resp.GetDesc() != "200 OK" {
			failNow(t, "Response.Desc has a wrong behavior")
		}
		if e := resp.GetExamples(); e["application/json"] != `{"code": "200", "message": "success", "data": {"id": 1, "name": "user1"}}` ||
			e["text/xml"] != `<Result> <code>200</code> <message>success</message> </Result>` {
			failNow(t, "Response.Example or Response.AddExample has a wrong behavior")
		}
		if h := resp.GetHeaders()[0]; h.GetName() != "X-RateLimit-Remaining" || h.GetType() != "integer#int64" || h.GetDesc() != "Request rate limit remaining" {
			failNow(t, "NewHeader has a wrong behavior")
		}
		header := resp.GetHeaders()[1]
		if header.GetName() != "X-RateLimit-Limit" {
			failNow(t, "Header.Name has a wrong behavior")
		}
		if header.GetType() != "integer#int64" {
			failNow(t, "Header.Type has a wrong behavior")
		}
		if header.GetDesc() != "Request rate limit size" {
			failNow(t, "Header.Desc has a wrong behavior")
		}
		if len(op.GetParams()) != 8 {
			failNow(t, "Operation.Params or Operation.AddParams has a wrong behavior")
		}
		if p := op.GetParams(); p[0].GetName() != "id" || p[0].GetType() != "integer#int64" || p[0].GetInLoc() != "path" || p[0].GetRequired() != true || p[0].GetDesc() != "user id" ||
			p[1].GetName() != "need_details" || p[1].GetType() != "boolean" || p[1].GetInLoc() != "query" || p[1].GetRequired() != false || p[1].GetDesc() != "is need details" ||
			p[2].GetName() != "key" || p[2].GetType() != "string" || p[2].GetInLoc() != "formData" || p[2].GetRequired() != false || p[2].GetDesc() != "fake form key" ||
			p[3].GetName() != "param" || p[3].GetType() != "Fake" || p[3].GetInLoc() != "body" || p[3].GetRequired() != false || p[3].GetDesc() != "fake body" ||
			p[4].GetName() != "X-NEED-DETAILS" || p[4].GetType() != "boolean" || p[4].GetInLoc() != "header" || p[4].GetRequired() != false || p[4].GetDesc() != "a duplicate parameter of need_details query" {
			failNow(t, "NewXXXParam has a wrong behavior")
		}
		param := op.GetParams()[5]
		if param.GetName() != "X-TEST" {
			failNow(t, "Param.Name has a wrong behavior")
		}
		if param.GetInLoc() != "header" {
			failNow(t, "Param.InLoc has a wrong behavior")
		}
		if param.GetType() != "string" {
			failNow(t, "Param.Type has a wrong behavior")
		}
		if param.GetRequired() != true {
			failNow(t, "Param.Required has a wrong behavior")
		}
		if param.GetDesc() != "A test header" {
			failNow(t, "Param.Desc has a wrong behavior")
		}
		if param.GetAllowEmpty() != true {
			failNow(t, "Param.AllowEmpty has a wrong behavior")
		}
		if param.GetDefault() != "test###" {
			failNow(t, "Param.Default has a wrong behavior")
		}
		if param.GetExample() != "a###" {
			failNow(t, "Param.Example has a wrong behavior")
		}
		if param.GetPattern() != "^.+###$" {
			failNow(t, "Param.Pattern has a wrong behavior")
		}
		if e := param.GetEnum(); e[0] != "a###" || e[1] != "b###" || e[2] != "test###" {
			failNow(t, "Param.Enum has a wrong behavior")
		}
		if *param.GetMinLength() != 4 {
			failNow(t, "Param.MinLength or Param.LengthRange has a wrong behavior")
		}
		if *param.GetMaxLength() != 8 {
			failNow(t, "Param.MaxLength or Param.LengthRange has a wrong behavior")
		}
		param = op.GetParams()[6]
		if *param.GetMinItems() != 3 {
			failNow(t, "Param.MinItems or Param.ItemsRange has a wrong behavior")
		}
		if *param.GetMaxItems() != 18 {
			failNow(t, "Param.MaxItems or Param.ItemsRange has a wrong behavior")
		}
		if param.GetUniqueItems() != true {
			failNow(t, "Param.UniqueItems has a wrong behavior")
		}
		if param.GetCollectionFormat() != "csv" {
			failNow(t, "Param.CollectionFormat has a wrong behavior")
		}
		if param.GetItemOption() == nil {
			failNow(t, "Param.ItemOption has a wrong behavior")
		}
		param = op.GetParams()[7]
		if *param.GetMinimum() != 1.1 {
			failNow(t, "Param.Minimum or Param.ValueRange has a wrong behavior")
		}
		if *param.GetMaximum() != 3.3 {
			failNow(t, "Param.Maximum or Param.ValueRange has a wrong behavior")
		}
		if param.GetExclusiveMax() != true {
			failNow(t, "Param.ExclusiveMax has a wrong behavior")
		}
		if param.GetExclusiveMin() != true {
			failNow(t, "Param.ExclusiveMin has a wrong behavior")
		}
		if param.GetMultipleOf() != 1.1 {
			failNow(t, "Param.MultipleOf has a wrong behavior")
		}
	})

	t.Run("Set and get in definition.go", func(t *testing.T) {
		SetDefinitions()
		AddDefinitions(NewDefinition("Result", "A global response"),
			NewDefinition("_Result", "A global response with generics").
				Generics("T", "E").
				Properties(NewProperty("data", "T", true, "response data"),
					NewProperty("error", "E", false, "response error")))
		AddDefinitions(NewDefinition("", "").
			Name("Page").
			Desc("A global paged response").
			Generics("T").
			Properties(NewProperty("total", "integer#int32", true, "total data"),
				NewProperty("page", "integer#int32", true, "current page")).
			AddProperties(
				NewProperty("limit", "integer#int32", true, "page size"),
				NewProperty("", "", false, "").
					Name("data").
					Type("T").
					Required(true).
					Desc("paged data").
					AllowEmpty(true).
					Default([]string{}).
					Example([]string{"hello", "world"}).
					Pattern("^.+$"). // followings are only for test
					Enum("hello", "world", "").
					MinLength(0).
					MaxLength(0).
					LengthRange(4, 8).MinItems(0).
					MaxItems(0).
					ItemsRange(3, 18).
					UniqueItems(true).
					CollectionFormat(CSV).
					Minimum(0).
					Maximum(0).
					ValueRange(1.1, 3.3).
					ExclusiveMax(true).
					ExclusiveMin(true).
					MultipleOf(1.1).
					ItemOption(NewItemOption(). // only for test
						AllowEmpty(true).
						Default([]string{}).
						Example([]string{"hello", "world"}).
						Pattern("^.+$").
						Enum("hello", "world", "").
						MinLength(0).
						MaxLength(0).
						LengthRange(4, 8).
						MinItems(0).
						MaxItems(0).
						ItemsRange(3, 18).
						UniqueItems(true).
						CollectionFormat(CSV).
						Minimum(0).
						Maximum(0).
						ValueRange(1.1, 3.3).
						ExclusiveMax(true).
						ExclusiveMin(true).
						MultipleOf(1.1).
						ItemOption(NewItemOption()))))
		if len(GetDefinitions()) != 3 {
			failNow(t, "AddDefinitions or SetDefinitions has a wrong behavior")
		}
		if d := GetDefinitions(); d[0].GetName() != "Result" || d[0].GetDesc() != "A global response" ||
			d[1].GetName() != "_Result" || d[1].GetDesc() != "A global response with generics" {
			failNow(t, "NewDefinition has a wrong behavior")
		}
		if g := GetDefinitions()[1].GetGenerics(); g[0] != "T" || g[1] != "E" {
			failNow(t, "Definition.Generics has a wrong behavior")
		}
		if p := GetDefinitions()[1].GetProperties(); p[0].GetName() != "data" || p[0].GetType() != "T" || p[0].GetRequired() != true || p[0].GetDesc() != "response data" ||
			p[1].GetName() != "error" || p[1].GetType() != "E" || p[1].GetRequired() != false || p[1].GetDesc() != "response error" {
			failNow(t, "Definition.NewProperty has a wrong behavior")
		}
		def := GetDefinitions()[2]
		if def.GetName() != "Page" {
			failNow(t, "Definition.Name has a wrong behavior")
		}
		if def.GetDesc() != "A global paged response" {
			failNow(t, "Definition.Desc has a wrong behavior")
		}
		if def.GetGenerics()[0] != "T" {
			failNow(t, "Definition.Generics has a wrong behavior")
		}
		if len(def.GetProperties()) != 4 {
			failNow(t, "Definition.Properties or Definition.AddProperties has a wrong behavior")
		}
		prop := def.GetProperties()[3]
		if prop.GetName() != "data" {
			failNow(t, "Property.Name has a wrong behavior")
		}
		if prop.GetType() != "T" {
			failNow(t, "Property.Type has a wrong behavior")
		}
		if prop.GetRequired() != true {
			failNow(t, "Property.Required has a wrong behavior")
		}
		if prop.GetDesc() != "paged data" {
			failNow(t, "Property.Desc has a wrong behavior")
		}
		if prop.GetAllowEmpty() != true {
			failNow(t, "Property.AllowEmpty has a wrong behavior")
		}
		if prop.GetDefault() == nil {
			failNow(t, "Property.Default has a wrong behavior")
		}
		if prop.GetExample() == nil {
			failNow(t, "Property.Example has a wrong behavior")
		}
		if prop.GetPattern() != "^.+$" {
			failNow(t, "Property.Pattern has a wrong behavior")
		}
		if e := prop.GetEnum(); e[0] != "hello" || e[1] != "world" || e[2] != "" {
			failNow(t, "Property.Enum has a wrong behavior")
		}
		if *prop.GetMinLength() != 4 {
			failNow(t, "Property.MinLength or Property.LengthRange has a wrong behavior")
		}
		if *prop.GetMaxLength() != 8 {
			failNow(t, "Property.MaxLength or Property.LengthRange has a wrong behavior")
		}
		if *prop.GetMinItems() != 3 {
			failNow(t, "Property.MinItems or Property.ItemsRange has a wrong behavior")
		}
		if *prop.GetMaxItems() != 18 {
			failNow(t, "Property.MaxItems or Property.ItemsRange has a wrong behavior")
		}
		if prop.GetUniqueItems() != true {
			failNow(t, "Property.UniqueItems has a wrong behavior")
		}
		if prop.GetCollectionFormat() != "csv" {
			failNow(t, "Property.CollectionFormat has a wrong behavior")
		}
		if *prop.GetMinimum() != 1.1 {
			failNow(t, "Property.Minimum or Property.ValueRange has a wrong behavior")
		}
		if *prop.GetMaximum() != 3.3 {
			failNow(t, "Property.Maximum or Property.ValueRange has a wrong behavior")
		}
		if prop.GetExclusiveMax() != true {
			failNow(t, "Property.ExclusiveMax has a wrong behavior")
		}
		if prop.GetExclusiveMin() != true {
			failNow(t, "Property.ExclusiveMin has a wrong behavior")
		}
		if prop.GetMultipleOf() != 1.1 {
			failNow(t, "Property.MultipleOf has a wrong behavior")
		}
		opt := prop.GetItemOption()
		if opt.GetAllowEmpty() != true {
			failNow(t, "ItemOption.AllowEmpty has a wrong behavior")
		}
		if opt.GetDefault() == nil {
			failNow(t, "ItemOption.Default has a wrong behavior")
		}
		if opt.GetExample() == nil {
			failNow(t, "ItemOption.Example has a wrong behavior")
		}
		if opt.GetPattern() != "^.+$" {
			failNow(t, "ItemOption.Pattern has a wrong behavior")
		}
		if e := opt.GetEnum(); e[0] != "hello" || e[1] != "world" || e[2] != "" {
			failNow(t, "ItemOption.Enum has a wrong behavior")
		}
		if *opt.GetMinLength() != 4 {
			failNow(t, "ItemOption.MinLength or ItemOption.LengthRange has a wrong behavior")
		}
		if *opt.GetMaxLength() != 8 {
			failNow(t, "ItemOption.MaxLength or ItemOption.LengthRange has a wrong behavior")
		}
		if *opt.GetMinItems() != 3 {
			failNow(t, "ItemOption.MinItems or ItemOption.ItemsRange has a wrong behavior")
		}
		if *opt.GetMaxItems() != 18 {
			failNow(t, "ItemOption.MaxItems or ItemOption.ItemsRange has a wrong behavior")
		}
		if opt.GetUniqueItems() != true {
			failNow(t, "ItemOption.UniqueItems has a wrong behavior")
		}
		if opt.GetCollectionFormat() != "csv" {
			failNow(t, "ItemOption.CollectionFormat has a wrong behavior")
		}
		if *opt.GetMinimum() != 1.1 {
			failNow(t, "ItemOption.Minimum or ItemOption.ValueRange has a wrong behavior")
		}
		if *opt.GetMaximum() != 3.3 {
			failNow(t, "ItemOption.Maximum or ItemOption.ValueRange has a wrong behavior")
		}
		if opt.GetExclusiveMax() != true {
			failNow(t, "ItemOption.ExclusiveMax has a wrong behavior")
		}
		if opt.GetExclusiveMin() != true {
			failNow(t, "ItemOption.ExclusiveMin has a wrong behavior")
		}
		if opt.GetMultipleOf() != 1.1 {
			failNow(t, "ItemOption.MultipleOf has a wrong behavior")
		}
		if opt.GetItemOption() == nil || opt.GetItemOption().GetItemOption() != nil {
			failNow(t, "ItemOption.ItemOption has a wrong behavior")
		}
	})
}
