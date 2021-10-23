package grades

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	data "github.com/tokuchi765/npb-analysis/entity/player"
	"github.com/tokuchi765/npb-analysis/infrastructure"
	"github.com/tokuchi765/npb-analysis/team"
)

// GradesInteractor 成績情報処理のInteractor
type GradesInteractor struct {
	infrastructure.GradesRepository
	infrastructure.TeamRepository
}

// GetPitching 個人投手成績一覧を取得する
func (Interactor *GradesInteractor) GetPitching(playerID string) (pitchings []data.PICHERGRADES) {
	return Interactor.GradesRepository.GetPitchings(playerID)
}

// GetBatting 個人打撃成績一覧を取得する
func (Interactor *GradesInteractor) GetBatting(playerID string) (battings []data.BATTERGRADES) {
	return Interactor.GradesRepository.GetBattings(playerID)
}

// GetCareer 選手情報を取得する
func (Interactor *GradesInteractor) GetCareer(playerID string) (career data.CAREER) {
	return Interactor.GradesRepository.GetCareer(playerID)
}

// GetPlayersByTeamIDAndYear チームIDと年から選手一覧を取得する
func (Interactor *GradesInteractor) GetPlayersByTeamIDAndYear(teamID string, year string) (players []data.PLAYER) {
	return Interactor.GradesRepository.GetPlayersByTeamIDAndYear(teamID, year)
}

// InsertTeamPlayers 年度ごとの選手一覧をDBに登録する
func (Interactor *GradesInteractor) InsertTeamPlayers(initial string, players [][]string) {
	teamID := team.GetTeamID(initial)
	teamName := Interactor.TeamRepository.GetTeamName(teamID)
	Interactor.GradesRepository.InsertTeamPlayers(teamID, teamName, players)
}

// GetPlayers 引数で受け取った x_players.csv ファイルを読み取って、配列にして返す
func GetPlayers(url string) (players [][]string) {
	// バイト列を読み込む
	file, err := os.Open(url)
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

// ReadCareers 引数で受け取った選手リストをもとに、経歴をまとめたデータクラスのリストを返す
func ReadCareers(playersPath string, initial string, players [][]string) (careerList []data.CAREER) {
	for _, player := range players {
		id := extractionPlayerID(player[0])
		url := playersPath + initial + "/careers/" + id + "_" + player[1] + "_career.csv"
		if exists(url) {
			career := readCareer(url)
			careerList = append(careerList, career)
		}
	}
	return careerList
}

func extractionPlayerID(url string) string {
	return strings.Replace(strings.Replace(url, "/bis/players/", "", 1), ".html", "", 1)
}

// ExtractionCareers 引数で受け取ったCAREERリストから重複選手を除外する
func (Interactor *GradesInteractor) ExtractionCareers(careers *[]data.CAREER) {
	Interactor.GradesRepository.ExtractionCareers(careers)
}

// InsertCareers 引数で受け取った CAREER をDBへ登録する
func (Interactor *GradesInteractor) InsertCareers(careers []data.CAREER) {
	Interactor.GradesRepository.InsertCareers(careers)
}

// ReadGradesMap 引数のplayersに設定されている選手成績を読み込み、Mapにして返す
func ReadGradesMap(playersPath string, initial string, players [][]string) (picherMap map[string][]data.PICHERGRADES, batterMap map[string][]data.BATTERGRADES) {
	picherMap = make(map[string][]data.PICHERGRADES)
	batterMap = make(map[string][]data.BATTERGRADES)
	for _, player := range players {
		id := strings.Replace(strings.Replace(player[0], "/bis/players/", "", 1), ".html", "", 1)
		path := playersPath + initial + "/grades/" + id + "_" + player[1] + "_grades.csv"

		if exists(path) {
			picherGrades, batterGrades := readGrades(path)
			if picherGrades != nil {
				picherMap[id] = picherGrades
			} else {
				batterMap[id] = batterGrades
			}
		}
	}
	return picherMap, batterMap
}

// ExtractionPicherGrades 引数で受け取ったPICHERGRADESリストから重複選手を除外する
func (Interactor *GradesInteractor) ExtractionPicherGrades(picherMap *map[string][]data.PICHERGRADES, teamID string) {
	Interactor.GradesRepository.ExtractionPicherGrades(picherMap, teamID)
}

// InsertPicherGrades 引数で受け取ったPICHERGRADESリストから重複選手を除外する
func (Interactor *GradesInteractor) InsertPicherGrades(picherMap map[string][]data.PICHERGRADES) {
	Interactor.GradesRepository.InsertPicherGrades(picherMap)
}

// ExtractionBatterGrades 引数で受け取ったBATTERGRADESリストから重複選手を除外する
func (Interactor *GradesInteractor) ExtractionBatterGrades(batterMap *map[string][]data.BATTERGRADES, teamID string) {
	Interactor.GradesRepository.ExtractionBatterGrades(batterMap, teamID)
}

// InsertBatterGrades 引数で受け取ったBATTERGRADESをDBに登録する
func (Interactor *GradesInteractor) InsertBatterGrades(batterMap map[string][]data.BATTERGRADES, current string) {
	// 加重出塁率の計算に必要なconfigファイルを読み込む
	config, _ := loadConfig(current)

	for key, value := range batterMap {
		for _, batter := range value {
			setSingle(&batter)
			setWoba(&batter, config)
			Interactor.GradesRepository.InsertBatterGrades(key, batter)
		}
	}
}

func setWoba(batterGrades *data.BATTERGRADES, config *config) {
	molecule := config.BaseOnBallsAndHitByPitches*(float64(batterGrades.BaseOnBalls)+float64(batterGrades.HitByPitches)) +
		config.Single*float64(batterGrades.Single) +
		config.Double*float64(batterGrades.Double) +
		config.Triple*float64(batterGrades.Triple) +
		config.HomeRun*float64(batterGrades.HomeRun)
	denominator := (float64(batterGrades.AtBat) + float64(batterGrades.BaseOnBalls) + float64(batterGrades.HitByPitches) + float64(batterGrades.SacrificeFlies))
	batterGrades.Woba = molecule / denominator
}

type config struct {
	Single                     float64 `json:"single"`
	BaseOnBallsAndHitByPitches float64 `json:"baseOnBallsAndHitByPitches"`
	Double                     float64 `json:"double"`
	Triple                     float64 `json:"triple"`
	HomeRun                    float64 `json:"homeRun"`
}

func loadConfig(current string) (*config, error) {
	f, err := os.Open(current + "/grades/property/config.json")
	if err != nil {
		log.Fatal("loadConfig os.Open err:", err)
		return nil, err
	}
	defer f.Close()

	var cfg config
	err = json.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}

func setSingle(batterGrades *data.BATTERGRADES) {
	batterGrades.Single = batterGrades.Hit - batterGrades.Double - batterGrades.Triple - batterGrades.HomeRun
}

func readGrades(path string) (picherGradesList []data.PICHERGRADES, batterGradesList []data.BATTERGRADES) {
	// バイト列を読み込む
	file, err := os.Open(path)
	if err != nil {
		log.Print(err)
	}
	// 	終わったらファイルを閉じる
	defer file.Close()

	reader := csv.NewReader(file)

	// ヘッダーを取得
	header, err := reader.Read()
	if err != nil {
		log.Print(err)
	}

	for {
		line, err := reader.Read()

		if err != nil {
			break
		}

		if header[3] == "登板" {
			picherGradesList = append(picherGradesList, setPicherGrades(line))
		} else {
			batterGradesList = append(batterGradesList, setBatterGrades(line))
		}
	}
	return picherGradesList, batterGradesList
}

func setPicherGrades(line []string) (grades data.PICHERGRADES) {
	grades.Year = line[1]
	grades.TeamID = team.GetTeamID(line[2])
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

func setBatterGrades(line []string) (grades data.BATTERGRADES) {
	grades.Year = line[1]
	grades.TeamID = team.GetTeamID(line[2])
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

func readCareer(path string) (career data.CAREER) {
	// バイト列を読み込む
	file, err := os.Open(path)
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

	return setCareer(lines)
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
