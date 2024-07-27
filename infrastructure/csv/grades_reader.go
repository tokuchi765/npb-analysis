package csv

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tokuchi765/npb-analysis/entity/player"
	"github.com/tokuchi765/npb-analysis/util"
)

// GradesReader チーム成績CSVの読み込みを管理する
type GradesReader struct {
	util.TeamUtil
}

// GetPlayers 引数で受け取った x_players.csv ファイルを読み取って、配列にして返す
func (GradesReader *GradesReader) GetPlayers(csvPath string, initial string, year string) (players [][]string) {
	// バイト列を読み込む
	file, err := os.Open(csvPath + "/teams/" + year + "/" + initial + "_players.csv")
	if err != nil {
		log.Print(err)
	}
	// 	終わったらファイルを閉じる
	defer file.Close()

	reader := csv.NewReader(file)

	// ヘッダーを読み飛ばす
	_, _ = reader.Read()

	for {
		line, err := reader.Read()
		if err != nil {
			break
		}

		players = append(players, []string{line[1], line[2]})
	}

	return players
}

// ReadCareers 選手キャリア情報を取得します
func (GradesReader *GradesReader) ReadCareers(csvPath string) (careers []player.CAREER) {
	pathes := getAllFilePathes(csvPath + "/players/careers")

	for _, path := range pathes {
		file, err := os.Open(path)

		if err != nil {
			log.Print(err)
		}

		reader := csv.NewReader(file)

		_, _ = reader.Read()
		var lines []string
		for {
			line, err := reader.Read()
			if err != nil {
				break
			}

			lines = append(lines, line[2])
		}
		careers = append(careers, setCareer(lines))
	}

	return careers
}

func setCareer(line []string) (career player.CAREER) {
	career.PlayerID = line[7]
	career.Name = line[8]
	career.Position = line[0]
	career.PitchingAndBatting = line[1]
	career.Height = line[2]
	career.Weight = line[3]
	career.Birthday = line[4]
	career.Career = line[5]
	career.Draft = line[6]

	return career
}

// ReadBatterGrades 選手打撃情報CSVを読み込みます
func (GradesReader *GradesReader) ReadBatterGrades(csvPath string) (batterGrades map[string][]player.BATTERGRADES) {
	batterGrades = make(map[string][]player.BATTERGRADES)
	pathes := getAllFilePathes(csvPath + "/players/batting_grades")

	for _, path := range pathes {
		file, err := os.Open(path)

		if err != nil {
			log.Print(err)
		}

		reader := csv.NewReader(file)

		var batterGradesList []player.BATTERGRADES

		_, _ = reader.Read()
		for {
			line, err := reader.Read()
			if err != nil {
				break
			}

			batterGradesList = append(batterGradesList, GradesReader.setBatterGrades(line))
		}
		id := extractionPlayerID(file.Name())
		batterGrades[id] = batterGradesList
	}

	return batterGrades
}

func (GradesReader *GradesReader) setBatterGrades(line []string) (grades player.BATTERGRADES) {
	grades.Year = strings.Replace(line[1], ".0", "", -1)
	grades.TeamID = GradesReader.TeamUtil.GetTeamID(line[2])
	grades.Team = line[2]
	grades.Games, _ = strconv.Atoi(line[3])
	grades.PlateAppearance, _ = strconv.Atoi(line[4])
	grades.AtBat, _ = strconv.Atoi(line[5])
	grades.Score, _ = strconv.Atoi(line[6])
	grades.Hit, _ = strconv.Atoi(line[7])
	grades.Double, _ = strconv.Atoi(line[8])
	grades.Triple, _ = strconv.Atoi(line[9])
	grades.HomeRun, _ = strconv.Atoi(line[10])
	grades.BaseHit, _ = strconv.Atoi(line[11])
	grades.RunsBattedIn, _ = strconv.Atoi(line[12])
	grades.StolenBase, _ = strconv.Atoi(line[13])
	grades.CaughtStealing, _ = strconv.Atoi(line[14])
	grades.SacrificeHits, _ = strconv.Atoi(line[15])
	grades.SacrificeFlies, _ = strconv.Atoi(line[16])
	grades.BaseOnBalls, _ = strconv.Atoi(line[17])
	grades.HitByPitches, _ = strconv.Atoi(line[18])
	grades.StrikeOut, _ = strconv.Atoi(line[19])
	grades.GroundedIntoDoublePlay, _ = strconv.Atoi(line[20])
	grades.BattingAverage, _ = strconv.ParseFloat(line[21], 64)
	grades.SluggingPercentage, _ = strconv.ParseFloat(line[22], 64)
	grades.OnBasePercentage, _ = strconv.ParseFloat(line[23], 64)

	return grades
}

// ReadPitcherGrades 選手投手情報CSVを読み込みます
func (GradesReader *GradesReader) ReadPitcherGrades(csvPath string) (pitcherGrades map[string][]player.PICHERGRADES) {
	pitcherGrades = make(map[string][]player.PICHERGRADES)
	pathes := getAllFilePathes(csvPath + "/players/pitching_grades")

	for _, path := range pathes {
		file, err := os.Open(path)

		if err != nil {
			log.Print(err)
		}

		reader := csv.NewReader(file)

		var pitcherGradesList []player.PICHERGRADES

		_, _ = reader.Read()
		for {
			line, err := reader.Read()
			if err != nil {
				break
			}

			pitcherGradesList = append(pitcherGradesList, GradesReader.setPicherGrades(line))
		}
		id := extractionPlayerID(file.Name())
		pitcherGrades[id] = pitcherGradesList
	}

	return pitcherGrades
}

func (GradesReader *GradesReader) setPicherGrades(line []string) (grades player.PICHERGRADES) {
	grades.Year = strings.Replace(line[1], ".0", "", -1)
	grades.TeamID = GradesReader.TeamUtil.GetTeamID(line[2])
	grades.Team = line[2]
	grades.Piched, _ = strconv.ParseFloat(line[3], 64)
	grades.Win, _ = strconv.ParseFloat(line[4], 64)
	grades.Lose, _ = strconv.ParseFloat(line[5], 64)
	grades.Save, _ = strconv.ParseFloat(line[6], 64)
	grades.Hold, _ = strconv.ParseFloat(line[7], 64)
	grades.HoldPoint, _ = strconv.ParseFloat(line[8], 64)
	grades.CompleteGame, _ = strconv.ParseFloat(line[9], 64)
	grades.Shutout, _ = strconv.ParseFloat(line[10], 64)
	grades.NoWalks, _ = strconv.ParseFloat(line[11], 64)
	grades.WinningRate, _ = strconv.ParseFloat(line[12], 64)
	grades.Batter, _ = strconv.ParseFloat(line[13], 64)
	grades.InningsPitched, _ = strconv.ParseFloat(line[14], 64)
	grades.Hit, _ = strconv.ParseFloat(line[15], 64)
	grades.HomeRun, _ = strconv.ParseFloat(line[16], 64)
	grades.BaseOnBalls, _ = strconv.ParseFloat(line[17], 64)
	grades.HitByPitches, _ = strconv.ParseFloat(line[18], 64)
	grades.StrikeOut, _ = strconv.ParseFloat(line[19], 64)
	grades.WildPitches, _ = strconv.ParseFloat(line[20], 64)
	grades.Balk, _ = strconv.ParseFloat(line[21], 64)
	grades.RunsAllowed, _ = strconv.ParseFloat(line[22], 64)
	grades.EarnedRun, _ = strconv.ParseFloat(line[23], 64)
	grades.EarnedRunAverage, _ = strconv.ParseFloat(line[24], 64)

	return grades
}
