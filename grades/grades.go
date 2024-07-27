package grades

import (
	"encoding/json"
	"log"
	"math"
	"os"

	_ "github.com/lib/pq"
	data "github.com/tokuchi765/npb-analysis/entity/player"
	"github.com/tokuchi765/npb-analysis/interfaces/reader"
	"github.com/tokuchi765/npb-analysis/interfaces/repository"
	"github.com/tokuchi765/npb-analysis/util"
)

// GradesInteractor 成績情報処理のInteractor
type GradesInteractor struct {
	repository.GradesRepository
	repository.TeamRepository
	reader.GradesReader
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

// SearchCareerByName 選手名から選手データを検索する
func (Interactor *GradesInteractor) SearchCareerByName(name string) (career []data.CAREER) {
	return Interactor.GradesRepository.SearchCareerByName(name)
}

// GetPlayers 引数で受け取った x_players.csv ファイルを読み取って、配列にして返す
func (Interactor *GradesInteractor) GetPlayers(csvPath string, initial string, year string) (players [][]string) {
	return Interactor.GradesReader.GetPlayers(csvPath, initial, year)
}

// InsertCareers 引数で受け取った CAREER をDBへ登録する
func (Interactor *GradesInteractor) InsertCareers(csvPath string) {
	careers := Interactor.GradesReader.ReadCareers(csvPath)
	Interactor.GradesRepository.InsertCareers(careers)
}

// InsertPicherGrades 選手投手成績をDBに登録します
func (Interactor *GradesInteractor) InsertPicherGrades(csvPath string) {
	pitcherGrades := Interactor.GradesReader.ReadPitcherGrades(csvPath)

	for key, pichers := range pitcherGrades {
		for _, picher := range pichers {
			picher.SetBABIP()
			picher.SetStrikeOutRate()
			Interactor.GradesRepository.InsertPicherGrades(key, picher)
		}
	}
}

// InsertBatterGrades 選手打撃成績をDBに登録します
func (Interactor *GradesInteractor) InsertBatterGrades(current string) {
	batterGrades := Interactor.GradesReader.ReadBatterGrades(current + "/csv")

	// 加重出塁率の計算に必要なconfigファイルを読み込む
	config, _ := loadConfig(current)

	for key, value := range batterGrades {
		for _, batter := range value {
			setSingle(&batter)
			setWoba(&batter, config)
			batter.SetRC()
			batter.SetBABIP()
			batter.SetStrikeOutRate()
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
	if math.IsNaN(batterGrades.Woba) {
		batterGrades.Woba = 0.0
	}
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
