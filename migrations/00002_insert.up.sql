
-- システム設定登録
INSERT INTO system_setting(setting,value) VALUES('created_player_grades','false');
INSERT INTO system_setting(setting,value) VALUES('created_team_stats','false');
INSERT INTO system_setting(setting,value) VALUES('created_add_value','false');

-- チーム名登録
INSERT INTO team_name(team_name_id,team_name) VALUES('01','Giants');
INSERT INTO team_name(team_name_id,team_name) VALUES('02','Baystars');
INSERT INTO team_name(team_name_id,team_name) VALUES('03','Tigers');
INSERT INTO team_name(team_name_id,team_name) VALUES('04','Carp');
INSERT INTO team_name(team_name_id,team_name) VALUES('05','Dragons');
INSERT INTO team_name(team_name_id,team_name) VALUES('06','Swallows');
INSERT INTO team_name(team_name_id,team_name) VALUES('07','Lions');
INSERT INTO team_name(team_name_id,team_name) VALUES('08','Hawks');
INSERT INTO team_name(team_name_id,team_name) VALUES('09','Eagles');
INSERT INTO team_name(team_name_id,team_name) VALUES('10','Marines');
INSERT INTO team_name(team_name_id,team_name) VALUES('11','Fighters');
INSERT INTO team_name(team_name_id,team_name) VALUES('12','Buffaloes');

-- チーム情報登録
INSERT INTO team_stats(team_id,team_name_id,league) VALUES('01','01','central');
INSERT INTO team_stats(team_id,team_name_id,league) VALUES('02','02','central');
INSERT INTO team_stats(team_id,team_name_id,league) VALUES('03','03','central');
INSERT INTO team_stats(team_id,team_name_id,league) VALUES('04','04','central');
INSERT INTO team_stats(team_id,team_name_id,league) VALUES('05','05','central');
INSERT INTO team_stats(team_id,team_name_id,league) VALUES('06','06','central');
INSERT INTO team_stats(team_id,team_name_id,league) VALUES('07','07','pacific');
INSERT INTO team_stats(team_id,team_name_id,league) VALUES('08','08','pacific');
INSERT INTO team_stats(team_id,team_name_id,league) VALUES('09','09','pacific');
INSERT INTO team_stats(team_id,team_name_id,league) VALUES('10','10','pacific');
INSERT INTO team_stats(team_id,team_name_id,league) VALUES('11','11','pacific');
INSERT INTO team_stats(team_id,team_name_id,league) VALUES('12','12','pacific');
