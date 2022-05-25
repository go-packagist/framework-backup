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

	// // Test Service
	// testService := app.GetService("test").(*providers.TestService)
	// assert.Equal(t, app, testService.Application())
	//
	// testService.WriteContent("123123")
	// assert.Equal(t, "123123", testService.ReadContent())
	//
	// testService.WriteContent("234567")
	// assert.Equal(t, "234567", testService.ReadContent())
	//
	// app.GetService("test").(*providers.TestService).WriteContent("123123")
	// fmt.Println(app.GetService("test").(*providers.TestService).ReadContent())
	// fmt.Println(testService.ReadContent())
	//
	// fmt.Println(app)
}

func TestApplication_Bind(t *testing.T) {
	app := NewApplication("./")

	app.Register(NewTestProvider(app))

	// 测试容器的单例（Singleton）效果
	testService := app.Make("test").(*TestService)
	testService.WriteContent("aaa")
	assert.Equal(t, "aaa", testService.ReadContent())

	testService2 := app.Make("test").(*TestService)
	testService2.WriteContent("bbb")

	assert.Equal(t, "bbb", testService2.ReadContent())
	assert.Equal(t, "bbb", testService.ReadContent())
	assert.Equal(t, "bbb", app.Make("test").(*TestService).ReadContent())

	// 测试容器的Bind（no Shared）效果
	testService3 := app.Make("test2").(*TestService)
	testService3.WriteContent("aaa")
	assert.Equal(t, "aaa", testService3.ReadContent())

	testService4 := app.Make("test2").(*TestService)
	testService4.WriteContent("bbb")
	assert.Equal(t, "bbb", testService4.ReadContent())
	assert.Equal(t, "aaa", testService3.ReadContent())
	assert.Equal(t, "", app.Make("test2").(*TestService).ReadContent())
}
