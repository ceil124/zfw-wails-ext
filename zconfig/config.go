package zconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ceil124/zfw-wails-ext/zlog"
	"log"
	"os"
)

const ConfigPath = "config.json"

type ConfigManager struct {
	path   string
	params map[string]any
}

func NewConfigManager() *ConfigManager {
	cm := ConfigManager{
		path: ConfigPath,
	}
	cm.createFileIfNoExists()
	return &cm
}

func NewConfigManagerWithFile(path string) *ConfigManager {
	cm := ConfigManager{
		path: path,
	}
	cm.createFileIfNoExists()
	return &cm
}

// 检查配置文件是否存在，如不存在则完成初始化
func (cm *ConfigManager) createFileIfNoExists() {
	params := map[string]any{}

	// 检查文件是否存在
	if _, err := os.Stat(cm.path); os.IsNotExist(err) {
		// 生成空json
		empty, _ := json.Marshal(make(map[string]any))
		// 写入新文件
		err = os.WriteFile(cm.path, empty, 0644)
		if err != nil {
			fmt.Println("写入配置文件失败:", err)
			return
		}
		zlog.Infof("config file: %s init done...", cm.path)
	} else {
		// 读取配置文件，由于前面检查过，这里不会抛异常
		data, _ := os.ReadFile(cm.path)
		err := json.Unmarshal(data, &params)
		if err != nil {
			log.Fatal(err) // 输出异常并中断程序
		}
		zlog.Infof("config file: %s load done...", cm.path)
	}

	// 缓存params
	cm.params = params
}

// HasConfig 检查配置是否存在
func (cm *ConfigManager) HasConfig(name string) bool {
	for k, _ := range cm.params {
		if k == name {
			return true
		}
	}
	return false
}

// GetString 查询配置值，返回字符串
func (cm *ConfigManager) GetString(name string) (string, error) {
	obj, err := cm.GetObject(name)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", obj), nil
}

// GetObject 查询配置值 对象会返回map[string]interface{}，需要自行json序列化转换
func (cm *ConfigManager) GetObject(name string) (any, error) {
	if name == "" {
		return "", errors.New("键不能为空格")
	}
	return cm.params[name], nil
}

// SetString 设置配置
func (cm *ConfigManager) SetString(name string, value string) error {
	err := cm.SetObject(name, value)
	if err != nil {
		return err
	}
	return nil
}

// SetObject 设置配置
func (cm *ConfigManager) SetObject(name string, value any) error {
	// 将配置写入本地缓存
	cm.params[name] = value

	// 序列化配置信息
	marshal, err := json.Marshal(cm.params)
	if err != nil {
		zlog.Error(err)
		return err
	}

	// 配置写入本地文件
	err = os.WriteFile(cm.path, marshal, 0644)
	if err != nil {
		return err
	}
	return nil
}

// CleanConfig 清空配置文件内容（慎用）
func (cm *ConfigManager) CleanConfig() error {
	// 生成空json
	emptyJsonBytes, _ := json.Marshal(make(map[string]any))
	// 写入文件
	err := os.WriteFile(cm.path, emptyJsonBytes, 0644)
	if err != nil {
		zlog.Error(err)
		return err
	}
	return nil
}
