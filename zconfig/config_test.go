package zconfig

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const TestConfigPath = "config.json"

type TestDTO struct {
	Name string
	Age  int
}

func TestSaveAnGetString(t *testing.T) {
	a := assert.New(t)

	cm := NewConfigManager(TestConfigPath)
	_ = cm.CleanConfig()

	name := "test_config"
	value := "test_config"
	err := cm.SetString(name, value)
	a.Nil(err)

	hasConfig := cm.HasConfig(name)
	a.Nil(err)
	a.True(hasConfig)

	value2, err := cm.GetString(name)
	a.Nil(err)
	a.Equal(value, value2)

	fmt.Printf("result ==> name: %s, value: %s\n", name, value2)
}

func TestSaveAnGetObject(t *testing.T) {
	a := assert.New(t)

	cm := NewConfigManager(TestConfigPath)
	_ = cm.CleanConfig()

	name := "test_config"
	value := []TestDTO{
		{Name: "zhangsan", Age: 18},
		{Name: "lisi", Age: 20},
	}
	err := cm.SetObject(name, value)
	a.Nil(err)

	hasConfig := cm.HasConfig(name)
	a.Nil(err)
	a.True(hasConfig)

	value2, err := cm.GetObject(name)
	a.Nil(err)
	a.Equal(len(value), len(value2.([]TestDTO)))

	fmt.Printf("result ==> name: %s, value: %v\n", name, value2)
}

func TestCheckedConfig(t *testing.T) {
	a := assert.New(t)

	cm := NewConfigManager(TestConfigPath)
	_ = cm.CleanConfig()
	hasConfig := cm.HasConfig("xxx")
	a.False(hasConfig)
}
