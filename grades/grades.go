package grades

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
	data "github.com/tokuchi765/npb-analysis/entity/player"
	csvReader "github.com/tokuchi765/npb-analysis/infrastructure/csv"
	"github.com/tokuchi765/npb-analysis/interfaces/repository"
	"github.com/tokuchi765/npb-analysis/util"
)

// GradesInteractor 成績情報処理のInteractor
type GradesInteractor struct {
	repository.GradesRepository
	repository.TeamRepository
	csvReader.GradesReader
	util.TeamUtil
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
	teamID := Interactor.TeamUtil.GetTeamID(initial)
	teamName := Interactor.TeamRepository.GetTeamName(teamID)
	Interactor.GradesRepository.InsertTeamPlayers(teamID, teamName, players)
}

// GetPlayers 引数で受け取った x_players.csv ファイルを読み取って、配列にして返す
func (Interactor *GradesInteractor) GetPlayers(csvPath string, initial string) (players [][]string) {
	return Interactor.GradesReader.GetPlayers(csvPath, initial)
}

// ReadCareers 引数で受け取った選手リストをもとに、経歴をまとめたデータクラスのリストを返す
func (Interactor *GradesInteractor) ReadCareers(csvPath string, initial string, players [][]string) (careerList []data.CAREER) {
	for _, player := range players {
		id := extractionPlayerID(player[0])
		career, exists := Interactor.GradesReader.ReadCareer(csvPath, initial, id, player[1])
		if exists {
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
func (Interactor *GradesInteractor) ReadGradesMap(csvPath string, initial string, players [][]string) (picherMap map[string][]data.PICHERGRADES, batterMap map[string][]data.BATTERGRADES) {
	picherMap = make(map[string][]data.PICHERGRADES)
	batterMap = make(map[string][]data.BATTERGRADES)
	for _, player := range players {
		id := strings.Replace(strings.Replace(player[0], "/bis/players/", "", 1), ".html", "", 1)

		picherGrades, batterGrades, exist := Interactor.GradesReader.ReadGrades(csvPath, initial, id, player[1])

		if exist {
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
