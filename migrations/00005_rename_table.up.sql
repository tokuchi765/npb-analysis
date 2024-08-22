-- テーブル名誤記修正
ALTER TABLE picher_grades RENAME TO pitcher_grades;

-- カラム誤記修正
ALTER TABLE pitcher_grades RENAME COLUMN piched TO pitched;
ALTER TABLE pitcher_grades RENAME COLUMN hit_by_ptches TO hit_by_pitches;
ALTER TABLE team_pitching RENAME COLUMN hit_by_ptches TO hit_by_pitches;