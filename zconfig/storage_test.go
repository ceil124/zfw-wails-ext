package zconfig

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveAnGetConfig(t *testing.T) {
	a := assert.New(t)

	cs := NewConfigManager("storage.db")

	name := "test_config"
	value := "test_config"
	err := cs.SetConfig(name, value)
	a.Nil(err)

	hasConfig, err := cs.CheckConfig(name)
	a.Nil(err)
	a.True(hasConfig)
	config, err := cs.GetConfig(name)
	a.Nil(err)

	a.Equal(name, config.Name)
	a.Equal(value, config.Value)

	fmt.Printf("result ==> name: %s, value: %s\n", config.Value, config.Name)
}

func TestCheckedConfig(t *testing.T) {
	a := assert.New(t)

	cs := NewConfigManager("storage.db")
	hasConfig1, err := cs.CheckConfig("xxx")
	a.Nil(err)
	a.False(hasConfig1)
}
