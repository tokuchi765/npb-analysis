package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
	testUtil "github.com/tokuchi765/npb-analysis/test"
)

func TestSyastemRepository_GetSystemSetting(t *testing.T) {
	type args struct {
		setting string
	}
	tests := []struct {
		name      string
		args      args
		wantValue string
	}{
		{
			"選手成績作成済みフラグ確認",
			args{"created_player_grades"},
			"false",
		},
		{
			"チーム成績作成済みフラグ確認",
			args{"created_team_stats"},
			"false",
		},
		{
			"追加データ作成済みフラグ確認",
			args{"created_add_value"},
			"false",
		},
	}
	resource, pool := testUtil.CreateContainer()
	defer testUtil.CloseContainer(resource, pool)
	db := testUtil.ConnectDB(resource, pool)
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = db
	Repository := SyastemRepository{SQLHandler: *sqlHandler}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantValue, Repository.GetSystemSetting(tt.args.setting))
		})
	}
}

func TestSyastemRepository_SetSystemSetting(t *testing.T) {
	type args struct {
		setting string
		value   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"選手成績作成済みフラグ確認",
			args{"created_player_grades", "true"},
		},
	}
	resource, pool := testUtil.CreateContainer()
	defer testUtil.CloseContainer(resource, pool)
	db := testUtil.ConnectDB(resource, pool)
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = db
	Repository := SyastemRepository{SQLHandler: *sqlHandler}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Repository.SetSystemSetting(tt.args.setting, tt.args.value)
			assert.Equal(t, tt.args.value, Repository.GetSystemSetting(tt.args.setting))
		})
	}
}
