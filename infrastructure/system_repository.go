package infrastructure

import (
	"github.com/tokuchi765/npb-analysis/entity/system"
)

// SyastemRepository システムデータアクセスを管理するリポジトリ
type SyastemRepository struct {
	SQLHandler
}

// GetSystemSetting システム設定を取得する
func (Repository *SyastemRepository) GetSystemSetting(setting string) (value string) {
	var user system.SystemSetting
	Repository.GormDB.Where("setting = ?", setting).First(&user)

	return user.Value
}

// SetSystemSetting システム設定を登録する
func (Repository *SyastemRepository) SetSystemSetting(setting string, value string) {
	Repository.GormDB.Model(&system.SystemSetting{}).Where("setting = ?", setting).Update("value", value)
}
