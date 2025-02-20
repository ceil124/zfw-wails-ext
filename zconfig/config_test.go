package zconfig

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveAnGetConfig(t *testing.T) {
	a := assert.New(t)

	cm := NewConfigManager()
	_ = cm.CleanConfig()

	name := "test_config"
	value := "test_config"
	err := cm.SetConfig(name, value)
	a.Nil(err)

	hasConfig := cm.CheckConfig(name)
	a.Nil(err)
	a.True(hasConfig)

	value2, err := cm.GetConfig(name)
	a.Nil(err)
	a.Equal(value, value2)

	fmt.Printf("result ==> name: %s, value: %s\n", name, value2)
}

func TestCheckedConfig(t *testing.T) {
	a := assert.New(t)

	cm := NewConfigManager()
	_ = cm.CleanConfig()
	hasConfig := cm.CheckConfig("xxx")
	a.False(hasConfig)
}
