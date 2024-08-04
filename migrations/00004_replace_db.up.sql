-- システム設定登録
INSERT INTO system_setting(setting,value) VALUES('created_team_players','false');
INSERT INTO system_setting(setting,value) VALUES('created_player_careers','false');
INSERT INTO system_setting(setting,value) VALUES('created_player_battings','false');
INSERT INTO system_setting(setting,value) VALUES('created_player_pitchings','false');

-- 登録済みCSVテーブル追加
create table registered_csv (
  csv_type character varying not null
  , file_path character varying not null
  , constraint registered_csv_PKC primary key (csv_type,file_path)
) ;

comment on table registered_csv is '登録済みCSV';
comment on column registered_csv.csv_type is 'CSV種別';
comment on column registered_csv.file_path is 'ファイルパス';
