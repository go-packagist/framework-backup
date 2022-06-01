package foundation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplication(t *testing.T) {
}

func TestApplication_Version(t *testing.T) {
	app := NewApplication("./")

	assert.Equal(t, VERSION, app.Version())
}

func TestApplication_Register(t *testing.T) {
	app := NewApplication("./")

	app.Register(NewTestProvider(app))
	app.Register(NewTestProvider(app)) // 验证重复导入

	assert.Equal(t, []Provider{
		NewTestProvider(app),
	}, app.GetProviders())
}

func TestApplication_Bind(t *testing.T) {
	app := NewApplication("./")

	app.Register(NewTestProvider(app))

	// 测试容器的单例（Singleton）效果
	testService := app.MustMake("test").(*TestService)
	testService.WriteContent("aaa")
	assert.Equal(t, "aaa", testService.ReadContent())

	testService2 := app.MustMake("test").(*TestService)
	testService2.WriteContent("bbb")

	assert.Equal(t, "bbb", testService2.ReadContent())
	assert.Equal(t, "bbb", testService.ReadContent())
	assert.Equal(t, "bbb", app.MustMake("test").(*TestService).ReadContent())

	// 测试容器的Bind（no Shared）效果
	testService3 := app.MustMake("test2").(*TestService)
	testService3.WriteContent("aaa")
	assert.Equal(t, "aaa", testService3.ReadContent())

	testService4 := app.MustMake("test2").(*TestService)
	testService4.WriteContent("bbb")
	assert.Equal(t, "bbb", testService4.ReadContent())
	assert.Equal(t, "aaa", testService3.ReadContent())
	assert.Equal(t, "", app.MustMake("test2").(*TestService).ReadContent())
}

func TestApplication_AppInstance(t *testing.T) {
	app := NewApplication("./")

	app.Register(NewTestProvider(app))

	// GetInstance
	GetInstance().MustMake("test").(*TestService).WriteContent("aaa")
	assert.Equal(t, "aaa", GetInstance().MustMake("test").(*TestService).ReadContent())

	// App
	App().MustMake("test").(*TestService).WriteContent("bbb")
	assert.Equal(t, "bbb", App().MustMake("test").(*TestService).ReadContent())

	// Instance
	Instance().MustMake("test").(*TestService).WriteContent("ccc")
	assert.Equal(t, "ccc", Instance().MustMake("test").(*TestService).ReadContent())
}

func TestApplication_Instance(t *testing.T) {
	app := NewApplication("./")

	// map[string]string
	app.Instance("config", map[string]string{
		"key":  "value",
		"key1": "value1",
	})
	assert.Equal(t, "value", app.MustMake("config").(map[string]string)["key"])

	// string
	app.Instance("path.base", "dirname")
	assert.Equal(t, "dirname", app.MustMake("path.base").(string))

	// struct
	type TestStruct struct {
		Name string
	}
	app.Instance("test", &TestStruct{"test"})
	assert.Equal(t, "test", app.MustMake("test").(*TestStruct).Name)

	// func
	app.Instance("func", func() string {
		return "func"
	})
	assert.Equal(t, "func", app.MustMake("func").(func() string)())
}
