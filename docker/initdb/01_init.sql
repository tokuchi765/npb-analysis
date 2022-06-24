-- Project Name : npm-scraping
-- Date/Time    : 2022/06/08 23:04:29
-- Author       : hiroki
-- RDBMS Type   : PostgreSQL
-- Application  : A5:SQL Mk-2

/*
  << 注意！！ >>
  BackupToTempTable, RestoreFromTempTable疑似命令が付加されています。
  これにより、drop table, create table 後もデータが残ります。
  この機能は一時的に $$TableName のような一時テーブルを作成します。
  この機能は A5:SQL Mk-2でのみ有効であることに注意してください。
*/

-- 選手一覧テーブル
--* RestoreFromTempTable
create table team_players (
  year character varying not null
  , team_id character varying not null
  , team_name character varying
  , player_id character varying not null
  , player_name character varying
  , constraint team_players_PKC primary key (year,team_id,player_id)
) ;

-- システム設定
--* RestoreFromTempTable
create table system_setting (
  setting character varying
  , value character varying
  , constraint system_setting_PKC primary key (setting)
) ;

-- 対戦成績
--* RestoreFromTempTable
create table TEAM_MATCH_RESULTS (
  team_id character varying
  , year character varying
  , competitive_team_id character varying
  , vs_type character varying
  , win integer
  , lose integer
  , draw integer
  , constraint TEAM_MATCH_RESULTS_PKC primary key (team_id,year,competitive_team_id)
) ;

-- チームシーズン成績
--* RestoreFromTempTable
create table TEAM_SEASON_STATS (
  team_id character varying not null
  , year character varying not null
  , manager character varying
  , games integer
  , win integer
  , lose integer
  , draw integer
  , winning_rate real
  , exchange_win integer
  , exchange_lose integer
  , exchange_draw integer
  , home_win integer
  , home_lose integer
  , home_draw integer
  , load_win integer
  , load_lose integer
  , load_draw integer
  , pythagorean_expectation real
  , constraint TEAM_SEASON_STATS_PKC primary key (team_id,year)
) ;

-- チーム投手成績
--* RestoreFromTempTable
create table TEAM_PITCHING (
  team_id character varying not null
  , year character varying not null
  , earned_run_average real
  , games integer
  , win integer
  , lose integer
  , save integer
  , hold integer
  , hold_point integer
  , complete_game integer
  , shutout integer
  , no_walks integer
  , winning_rate real
  , batter integer
  , innings_pitched integer
  , hit integer
  , home_run integer
  , base_on_balls integer
  , intentional_walk integer
  , hit_by_ptches integer
  , strike_out integer
  , wild_pitches integer
  , balk integer
  , runs_allowed integer
  , earned_run integer
  , babip real
  , strike_out_rate real
  , constraint TEAM_PITCHING_PKC primary key (team_id,year)
) ;

-- チーム打撃成績
--* RestoreFromTempTable
create table TEAM_BATTING (
  team_id character varying not null
  , year character varying not null
  , batting_average real
  , games integer
  , plate_appearance integer
  , at_bat integer
  , score integer
  , hit integer
  , double integer
  , triple integer
  , home_run integer
  , base_hit integer
  , runs_batted_in integer
  , stolen_base integer
  , caught_stealing integer
  , sacrifice_hits integer
  , sacrifice_flies integer
  , base_on_balls integer
  , intentional_walk integer
  , hit_by_pitches integer
  , strike_out integer
  , strike_out_rate real
  , grounded_into_double_play integer
  , slugging_percentage real
  , on_base_percentage real
  , babip real
  , constraint TEAM_BATTING_PKC primary key (team_id,year)
) ;

-- チーム名
--* RestoreFromTempTable
create table TEAM_NAME (
  team_name_id character varying not null
  , team_name character varying
  , constraint TEAM_NAME_PKC primary key (team_name_id)
) ;

-- チーム情報
--* RestoreFromTempTable
create table TEAM_STATS (
  team_id character varying not null
  , team_name_id character varying
  , league character varying
  , constraint TEAM_STATS_PKC primary key (team_id)
) ;

-- 野手成績
--* RestoreFromTempTable
create table BATTER_GRADES (
  player_id character varying not null
  , year character varying
  , team_id character varying
  , team character varying
  , games integer
  , plate_appearance integer
  , at_bat integer
  , score integer
  , hit integer
  , single integer
  , double integer
  , triple integer
  , home_run integer
  , base_hit integer
  , runs_batted_in integer
  , stolen_base integer
  , caught_stealing integer
  , sacrifice_hits integer
  , sacrifice_flies integer
  , base_on_balls integer
  , hit_by_pitches integer
  , strike_out integer
  , strike_out_rate real
  , grounded_into_double_play integer
  , batting_average real
  , slugging_percentage real
  , on_base_percentage real
  , w_oba real
  , rc real
  , babip real
  , constraint BATTER_GRADES_PKC primary key (player_id,year,team_id)
) ;

-- 選手テーブル
--* RestoreFromTempTable
create table Players (
  player_id character varying not null
  , name character varying
  , position character varying
  , pitching_and_batting character varying
  , height character varying
  , weight character varying
  , birthday character varying
  , draft character varying
  , career character varying
  , constraint Players_PKC primary key (player_id)
) ;

-- 投手成績
--* RestoreFromTempTable
create table PICHER_GRADES (
  player_id character varying not null
  , year character varying
  , team_id character varying
  , team character varying
  , piched real
  , win real
  , lose real
  , save real
  , hold real
  , hold_point real
  , complete_game real
  , shutout real
  , no_walks real
  , winning_rate real
  , batter real
  , innings_pitched real
  , hit real
  , home_run real
  , base_on_balls real
  , hit_by_ptches real
  , strike_out real
  , wild_pitches real
  , balk real
  , runs_allowed real
  , earned_run real
  , earned_run_average real
  , babip real
  , strike_out_rate real
  , constraint PICHER_GRADES_PKC primary key (player_id,year,team_id)
) ;

comment on table team_players is '選手一覧テーブル';
comment on column team_players.year is '年';
comment on column team_players.team_id is 'チームID';
comment on column team_players.team_name is 'チーム名';
comment on column team_players.player_id is '選手ID';
comment on column team_players.player_name is '選手名';

comment on table system_setting is 'システム設定';
comment on column system_setting.setting is '設定';
comment on column system_setting.value is '値';

comment on table TEAM_MATCH_RESULTS is '対戦成績';
comment on column TEAM_MATCH_RESULTS.team_id is 'チームID';
comment on column TEAM_MATCH_RESULTS.year is '年';
comment on column TEAM_MATCH_RESULTS.competitive_team_id is '対戦チームID';
comment on column TEAM_MATCH_RESULTS.vs_type is '対戦タイプ';
comment on column TEAM_MATCH_RESULTS.win is '勝利';
comment on column TEAM_MATCH_RESULTS.lose is '敗北';
comment on column TEAM_MATCH_RESULTS.draw is '引き分け';

comment on table TEAM_SEASON_STATS is 'チームシーズン成績';
comment on column TEAM_SEASON_STATS.team_id is 'チームID';
comment on column TEAM_SEASON_STATS.year is '年';
comment on column TEAM_SEASON_STATS.manager is '監督';
comment on column TEAM_SEASON_STATS.games is '試合';
comment on column TEAM_SEASON_STATS.win is '勝利';
comment on column TEAM_SEASON_STATS.lose is '敗北';
comment on column TEAM_SEASON_STATS.draw is '引き分け';
comment on column TEAM_SEASON_STATS.winning_rate is '勝率';
comment on column TEAM_SEASON_STATS.exchange_win is '交流戦勝利';
comment on column TEAM_SEASON_STATS.exchange_lose is '交流戦敗北';
comment on column TEAM_SEASON_STATS.exchange_draw is '交流戦引き分け';
comment on column TEAM_SEASON_STATS.home_win is 'ホーム勝利';
comment on column TEAM_SEASON_STATS.home_lose is 'ホーム敗北';
comment on column TEAM_SEASON_STATS.home_draw is 'ホーム引き分け';
comment on column TEAM_SEASON_STATS.load_win is 'ロード勝利';
comment on column TEAM_SEASON_STATS.load_lose is 'ロード敗北';
comment on column TEAM_SEASON_STATS.load_draw is 'ロード引き分け';
comment on column TEAM_SEASON_STATS.pythagorean_expectation is 'ピタゴラス勝率';

comment on table TEAM_PITCHING is 'チーム投手成績';
comment on column TEAM_PITCHING.team_id is 'チームID';
comment on column TEAM_PITCHING.year is '年';
comment on column TEAM_PITCHING.earned_run_average is '防御率';
comment on column TEAM_PITCHING.games is '試合';
comment on column TEAM_PITCHING.win is '勝利';
comment on column TEAM_PITCHING.lose is '敗北';
comment on column TEAM_PITCHING.save is 'セーブ';
comment on column TEAM_PITCHING.hold is 'ホールド';
comment on column TEAM_PITCHING.hold_point is 'ホールドポイント';
comment on column TEAM_PITCHING.complete_game is '完投';
comment on column TEAM_PITCHING.shutout is '完封';
comment on column TEAM_PITCHING.no_walks is '無四球';
comment on column TEAM_PITCHING.winning_rate is '勝率';
comment on column TEAM_PITCHING.batter is '打者';
comment on column TEAM_PITCHING.innings_pitched is '投球回数';
comment on column TEAM_PITCHING.hit is '安打';
comment on column TEAM_PITCHING.home_run is '本塁打';
comment on column TEAM_PITCHING.base_on_balls is '四球';
comment on column TEAM_PITCHING.intentional_walk is '故意四球';
comment on column TEAM_PITCHING.hit_by_ptches is '死球';
comment on column TEAM_PITCHING.strike_out is '三振';
comment on column TEAM_PITCHING.wild_pitches is '暴投';
comment on column TEAM_PITCHING.balk is 'ボーク';
comment on column TEAM_PITCHING.runs_allowed is '失点';
comment on column TEAM_PITCHING.earned_run is '自責点';
comment on column TEAM_PITCHING.babip is '被BABIP';
comment on column TEAM_PITCHING.strike_out_rate is '奪三振率';

comment on table TEAM_BATTING is 'チーム打撃成績';
comment on column TEAM_BATTING.team_id is 'チームID';
comment on column TEAM_BATTING.year is '年';
comment on column TEAM_BATTING.batting_average is '打率';
comment on column TEAM_BATTING.games is '試合';
comment on column TEAM_BATTING.plate_appearance is '打席';
comment on column TEAM_BATTING.at_bat is '打数';
comment on column TEAM_BATTING.score is '得点';
comment on column TEAM_BATTING.hit is '安打';
comment on column TEAM_BATTING.double is '二塁打';
comment on column TEAM_BATTING.triple is '三塁打';
comment on column TEAM_BATTING.home_run is '本塁打';
comment on column TEAM_BATTING.base_hit is '塁打';
comment on column TEAM_BATTING.runs_batted_in is '打点';
comment on column TEAM_BATTING.stolen_base is '盗塁';
comment on column TEAM_BATTING.caught_stealing is '盗塁刺';
comment on column TEAM_BATTING.sacrifice_hits is '犠打';
comment on column TEAM_BATTING.sacrifice_flies is '犠飛';
comment on column TEAM_BATTING.base_on_balls is '四球';
comment on column TEAM_BATTING.intentional_walk is '故意四';
comment on column TEAM_BATTING.hit_by_pitches is '死球';
comment on column TEAM_BATTING.strike_out is '三振';
comment on column TEAM_BATTING.strike_out_rate is '三振率';
comment on column TEAM_BATTING.grounded_into_double_play is '併殺打';
comment on column TEAM_BATTING.slugging_percentage is '長打率';
comment on column TEAM_BATTING.on_base_percentage is '出塁率';
comment on column TEAM_BATTING.babip is 'BABIP';

comment on table TEAM_NAME is 'チーム名';
comment on column TEAM_NAME.team_name_id is 'チーム名ID';
comment on column TEAM_NAME.team_name is 'チーム名';

comment on table TEAM_STATS is 'チーム情報';
comment on column TEAM_STATS.team_id is 'チームID';
comment on column TEAM_STATS.team_name_id is 'チーム名ID';
comment on column TEAM_STATS.league is 'リーグ';

comment on table BATTER_GRADES is '野手成績';
comment on column BATTER_GRADES.player_id is '選手ID';
comment on column BATTER_GRADES.year is '年';
comment on column BATTER_GRADES.team_id is 'チームID';
comment on column BATTER_GRADES.team is '所属球団';
comment on column BATTER_GRADES.games is '試合';
comment on column BATTER_GRADES.plate_appearance is '打席';
comment on column BATTER_GRADES.at_bat is '打数';
comment on column BATTER_GRADES.score is '得点';
comment on column BATTER_GRADES.hit is '安打';
comment on column BATTER_GRADES.single is '単打';
comment on column BATTER_GRADES.double is '二塁打';
comment on column BATTER_GRADES.triple is '三塁打';
comment on column BATTER_GRADES.home_run is '本塁打';
comment on column BATTER_GRADES.base_hit is '塁打';
comment on column BATTER_GRADES.runs_batted_in is '打点';
comment on column BATTER_GRADES.stolen_base is '盗塁';
comment on column BATTER_GRADES.caught_stealing is '盗塁刺';
comment on column BATTER_GRADES.sacrifice_hits is '犠打';
comment on column BATTER_GRADES.sacrifice_flies is '犠飛';
comment on column BATTER_GRADES.base_on_balls is '四球';
comment on column BATTER_GRADES.hit_by_pitches is '死球';
comment on column BATTER_GRADES.strike_out is '三振';
comment on column BATTER_GRADES.strike_out_rate is '三振率';
comment on column BATTER_GRADES.grounded_into_double_play is '併殺打';
comment on column BATTER_GRADES.batting_average is '打率';
comment on column BATTER_GRADES.slugging_percentage is '長打率';
comment on column BATTER_GRADES.on_base_percentage is '出塁率';
comment on column BATTER_GRADES.w_oba is '加重出塁率';
comment on column BATTER_GRADES.rc is '創出得点';
comment on column BATTER_GRADES.babip is 'BABIP';

comment on table Players is '選手テーブル';
comment on column Players.player_id is '選手ID';
comment on column Players.name is '選手名';
comment on column Players.position is 'ポジション';
comment on column Players.pitching_and_batting is '投打';
comment on column Players.height is '身長';
comment on column Players.weight is '体重';
comment on column Players.birthday is '生年月日';
comment on column Players.draft is 'ドラフト';
comment on column Players.career is '経歴';

comment on table PICHER_GRADES is '投手成績';
comment on column PICHER_GRADES.player_id is '選手ID';
comment on column PICHER_GRADES.year is '年';
comment on column PICHER_GRADES.team_id is 'チームID';
comment on column PICHER_GRADES.team is '所属球団';
comment on column PICHER_GRADES.piched is '登板';
comment on column PICHER_GRADES.win is '勝利';
comment on column PICHER_GRADES.lose is '敗北';
comment on column PICHER_GRADES.save is 'セーブ';
comment on column PICHER_GRADES.hold is 'ホールド';
comment on column PICHER_GRADES.hold_point is 'ホールドポイント';
comment on column PICHER_GRADES.complete_game is '完投';
comment on column PICHER_GRADES.shutout is '完封';
comment on column PICHER_GRADES.no_walks is '無四球';
comment on column PICHER_GRADES.winning_rate is '勝率';
comment on column PICHER_GRADES.batter is '打者';
comment on column PICHER_GRADES.innings_pitched is '投球回数';
comment on column PICHER_GRADES.hit is '安打';
comment on column PICHER_GRADES.home_run is 'ホームラン';
comment on column PICHER_GRADES.base_on_balls is '四球';
comment on column PICHER_GRADES.hit_by_ptches is '死球';
comment on column PICHER_GRADES.strike_out is '三振';
comment on column PICHER_GRADES.wild_pitches is '暴投';
comment on column PICHER_GRADES.balk is 'ボーク';
comment on column PICHER_GRADES.runs_allowed is '失点';
comment on column PICHER_GRADES.earned_run is '自責点';
comment on column PICHER_GRADES.earned_run_average is '防御率';
comment on column PICHER_GRADES.babip is '被BABIP';
comment on column PICHER_GRADES.strike_out_rate is '奪三振率';
