package csv

import (
	"encoding/csv"
	"log"
	"os"

	data "github.com/tokuchi765/npb-analysis/entity/player"
)

// GradesReader チーム成績CSVの読み込みを管理する
type GradesReader struct{}

// GetPlayers 引数で受け取った x_players.csv ファイルを読み取って、配列にして返す
func (GradesReader *GradesReader) GetPlayers(csvPath string, initial string) (players [][]string) {
	// バイト列を読み込む
	file, err := os.Open(csvPath + "/teams/" + initial + "_players.csv")
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

// ReadCareer 引数で受け取ったチームイニシャル、プレイヤーID、プレイヤー名を元にCSVを読み込む
func (GradesReader *GradesReader) ReadCareer(csvPath string, initial string, playerID string, playerName string) (career data.CAREER, exsist bool) {
	url := csvPath + "/players/" + initial + "/careers/" + playerID + "_" + playerName + "_career.csv"

	exsist = exists(url)
	if !exsist {
		return career, exsist
	}

	// バイト列を読み込む
	file, err := os.Open(url)
	if err != nil {
		log.Print(err)
	}
	// 	終わったらファイルを閉じる
	defer file.Close()

	reader := csv.NewReader(file)
	var lines []string

	// ヘッダーを取得
	_, err = reader.Read()
	if err != nil {
		log.Print(err)
	}

	for {
		line, err := reader.Read()
		if err != nil {
			break
		}

		lines = append(lines, line[2])
	}

	return setCareer(lines), exsist
}

func setCareer(line []string) (career data.CAREER) {
	career.PlayerID = line[7]
	career.Name = line[8]
	career.Position = line[0]
	career.PitchingAndBatting = line[1]
	career.Height = line[2]
	career.Weight = line[6]
	career.Birthday = line[3]
	career.Career = line[4]
	career.Draft = line[5]

	return career
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
