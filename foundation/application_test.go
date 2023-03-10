package foundation

import (
	"github.com/go-packagist/framework/container"
	"github.com/go-packagist/framework/contracts/provider"
	"github.com/go-packagist/framework/version"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplication(t *testing.T) {
}

func TestApplication_Version(t *testing.T) {
	app := NewApplication()

	assert.Equal(t, version.Version, app.Version())
}

func TestApplication_Register(t *testing.T) {
	app := NewApplication()

	app.Register(NewTestProvider(app.Container))
	app.Register(NewTestProvider(app.Container)) // 验证重复导入

	assert.Equal(t, []provider.Provider{
		NewTestProvider(app.Container),
	}, app.Container.GetProviders())
}

func TestApplication_Bind(t *testing.T) {
	app := NewApplication()

	app.Register(NewTestProvider(app.Container))

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
	app := NewApplication()

	app.Register(NewTestProvider(app.Container))

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

func TestApplication_Boot(t *testing.T) {
	app := NewApplication()

	testProvider := NewTestProvider(app.Container)

	app.Register(testProvider)

	assert.False(t, testProvider.(*TestProvider).Booted)

	app.Boot()

	assert.True(t, testProvider.(*TestProvider).Booted)
}

func TestApplication_Instance(t *testing.T) {
	app := NewApplication()

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

type TestProvider struct {
	container *container.Container
	Booted    bool
}

var _ provider.Provider = (*TestProvider)(nil)

func NewTestProvider(c *container.Container) provider.Provider {
	return &TestProvider{
		container: c,
		Booted:    false,
	}
}

func (p *TestProvider) Register() {
	p.container.Singleton("test", func(c *container.Container) interface{} {
		return NewTestService(c)
	})

	p.container.Bind("test2", func(c *container.Container) interface{} {
		return NewTestService(c)
	}, false)
}

func (p *TestProvider) Boot() {
	p.Booted = true
}

// TestService is a test service
type TestService struct {
	c       *container.Container
	content string
}

func NewTestService(c *container.Container) *TestService {
	return &TestService{
		c:       c,
		content: "",
	}
}

func (s *TestService) Container() *container.Container {
	return s.c
}

func (s *TestService) WriteContent(content string) *TestService {
	s.content = content

	return s
}

func (s *TestService) ReadContent() string {
	return s.content
}
