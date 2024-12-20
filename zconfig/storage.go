package zconfig

import (
	"database/sql"
	"e.coding.net/fangletianhua/mys/zfw-wails-ext/zlog"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

type ConfigManager struct {
	DriverName string
	FileName   string
}

func NewConfigManager(fileName string) *ConfigManager {
	cs := ConfigManager{
		FileName:   fileName,
		DriverName: DatabaseDriverName,
	}

	cs.CreateTableIfNotExists()
	return &cs
}

func (cm *ConfigManager) CheckConfig(name string) (bool, error) {
	db, err := cm.openConnectionWithConfigStorage()
	if err != nil {
		return false, err
	}
	defer db.Close()

	// 查询数据
	stmt, err := db.Prepare("SELECT count(1) FROM settings where name = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(name)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	hasNext := rows.Next()
	if !hasNext {
		return false, errors.New("查询异常")
	}

	var count int
	err = rows.Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (cm *ConfigManager) GetConfig(name string) (ConfigPO, error) {
	db, err := cm.openConnectionWithConfigStorage()
	if err != nil {
		return ConfigPO{}, err
	}
	defer db.Close()

	// 查询数据
	stmt, err := db.Prepare("SELECT name, value FROM settings where name = ?")
	if err != nil {
		return ConfigPO{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(name)
	if err != nil {
		return ConfigPO{}, err
	}
	defer rows.Close()

	hasNext := rows.Next()
	if !hasNext {
		return ConfigPO{}, errors.New("查无对象")
	}

	var configName, configValue string
	err = rows.Scan(&configName, &configValue)

	configPO := ConfigPO{
		Name:  configName,
		Value: configValue,
	}
	return configPO, nil
}

func (cm *ConfigManager) SetConfig(name string, value string) error {
	db, err := cm.openConnectionWithConfigStorage()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM settings where name = ?", name)
	if err != nil {
		return err
	}

	// 初始化数据
	stmt, err := db.Prepare("INSERT INTO settings(name, value) VALUES(?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, value)
	if err != nil {
		return err
	}

	return nil
}

func (cm *ConfigManager) CreateTableIfNotExists() {
	db, err := cm.openConnectionWithConfigStorage()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS settings (name TEXT PRIMARY KEY, value TEXT);`)
	if err != nil {
		panic(err)
	}
	zlog.Info("table settings init done...")
}

func (cm *ConfigManager) openConnectionWithConfigStorage() (*sql.DB, error) {
	return openConnection(cm.DriverName, cm.FileName)
}

// 创建或打开 SQLite 数据库
func openConnection(driverName, fileName string) (*sql.DB, error) {
	// 缓存存在的情况下直接读缓存的链接对象
	db, err := sql.Open(driverName, fileName)
	if err != nil {
		zlog.Error(err)
		panic(err)
	}
	return db, nil
}
