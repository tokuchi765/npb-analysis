package infrastructure

import "fmt"

// SyastemRepository システムデータアクセスを管理するリポジトリ
type SyastemRepository struct {
	SQLHandler
}

// GetSystemSetting システム設定を取得する
func (Repository *SyastemRepository) GetSystemSetting(setting string) (value string) {
	rows, err := Repository.Conn.Query("SELECT * FROM system_setting where setting = $1", setting)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var setting string
		rows.Scan(&setting, &value)
	}

	return value
}

// SetSystemSetting システム設定を登録する
func (Repository *SyastemRepository) SetSystemSetting(setting string, value string) {
	rows, err := Repository.Conn.Query("UPDATE system_setting SET value = $1 WHERE setting = $2", value, setting)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()
}
