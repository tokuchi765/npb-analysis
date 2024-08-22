package grades

import (
	"encoding/json"
	"log"
	"math"
	"os"

	_ "github.com/lib/pq"
	"github.com/tokuchi765/npb-analysis/entity/player"
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
func (Interactor *GradesInteractor) GetPitching(playerID string) (pitchings []player.PitcherGrades) {
	return Interactor.GradesRepository.GetPitchings(playerID)
}

// GetBatting 個人打撃成績一覧を取得する
func (Interactor *GradesInteractor) GetBatting(playerID string) (battings []player.BatterGrades) {
	return Interactor.GradesRepository.GetBattings(playerID)
}

// GetCareer 選手情報を取得する
func (Interactor *GradesInteractor) GetCareer(playerID string) (career player.Players) {
	return Interactor.GradesRepository.GetPlayers(playerID)
}

// SearchCareerByName 選手名から選手データを検索する
func (Interactor *GradesInteractor) SearchCareerByName(name string) (career []player.Players) {
	return Interactor.GradesRepository.SearchCareerByName(name)
}

// InsertCareers 引数で受け取った CAREER をDBへ登録する
func (Interactor *GradesInteractor) InsertCareers(csvPath string) {
	careers := Interactor.GradesReader.ReadCareers(csvPath)
	var players []player.Players
	for _, career := range careers {
		career.SetSearchName()
		players = append(players, Interactor.mappingCareers(career))
	}
	Interactor.GradesRepository.InsertPlayers(players)
}

func (Interactor *GradesInteractor) mappingCareers(career player.CAREER) player.Players {
	return player.Players{
		PlayerID:           career.PlayerID,
		Name:               career.Name,
		Position:           career.Position,
		PitchingAndBatting: career.PitchingAndBatting,
		Height:             career.Height,
		Weight:             career.Weight,
		Birthday:           career.Birthday,
		Career:             career.Career,
		Draft:              career.Draft,
		SearchName:         career.SearchName,
	}
}

// InsertPicherGrades 選手投手成績をDBに登録します
func (Interactor *GradesInteractor) InsertPicherGrades(csvPath string) {
	pitcherGrades := Interactor.GradesReader.ReadPitcherGrades(csvPath)

	for key, pichers := range pitcherGrades {
		for _, picher := range pichers {
			picher.SetInningsPitched()
			picher.SetBABIP()
			picher.SetStrikeOutRate()
			Interactor.GradesRepository.InsertPicherGrades(Interactor.mappingPicherGrades(key, picher))
		}
	}
}

func (Interactor *GradesInteractor) mappingPicherGrades(key string, picher player.PICHERGRADES) player.PitcherGrades {
	return player.PitcherGrades{
		PlayerID:         key,
		Year:             picher.Year,
		TeamID:           picher.TeamID,
		Team:             picher.Team,
		Pitched:          picher.Pitched,
		Win:              picher.Win,
		Lose:             picher.Lose,
		Save:             picher.Save,
		Hold:             picher.Hold,
		HoldPoint:        picher.HoldPoint,
		CompleteGame:     picher.CompleteGame,
		Shutout:          picher.Shutout,
		NoWalks:          picher.NoWalks,
		WinningRate:      picher.WinningRate,
		Batter:           picher.Batter,
		InningsPitched:   picher.InningsPitched,
		Hit:              picher.Hit,
		HomeRun:          picher.HomeRun,
		BaseOnBalls:      picher.BaseOnBalls,
		HitByPitches:     picher.HitByPitches,
		StrikeOut:        picher.StrikeOut,
		WildPitches:      picher.WildPitches,
		Balk:             picher.Balk,
		RunsAllowed:      picher.RunsAllowed,
		EarnedRun:        picher.EarnedRun,
		EarnedRunAverage: picher.EarnedRunAverage,
		Babip:            picher.BABIP,
		StrikeOutRate:    picher.StrikeOutRate,
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
			Interactor.GradesRepository.InsertBatterGrades(Interactor.mappingBatterGrades(key, batter))
		}
	}
}

func (Interactor *GradesInteractor) mappingBatterGrades(key string, batter player.BATTERGRADES) player.BatterGrades {
	return player.BatterGrades{
		PlayerID:               key,
		Year:                   batter.Year,
		TeamID:                 batter.TeamID,
		Team:                   batter.Team,
		Games:                  batter.Games,
		PlateAppearance:        batter.PlateAppearance,
		AtBat:                  batter.AtBat,
		Score:                  batter.Score,
		Hit:                    batter.Hit,
		Single:                 batter.Single,
		Double:                 batter.Double,
		Triple:                 batter.Triple,
		HomeRun:                batter.HomeRun,
		BaseHit:                batter.BaseHit,
		RunsBattedIn:           batter.RunsBattedIn,
		StolenBase:             batter.StolenBase,
		CaughtStealing:         batter.CaughtStealing,
		SacrificeHits:          batter.SacrificeHits,
		SacrificeFlies:         batter.SacrificeFlies,
		BaseOnBalls:            batter.BaseOnBalls,
		HitByPitches:           batter.HitByPitches,
		StrikeOut:              batter.StrikeOut,
		StrikeOutRate:          batter.StrikeOutRate.Float64,
		GroundedIntoDoublePlay: batter.GroundedIntoDoublePlay,
		BattingAverage:         batter.BattingAverage,
		SluggingPercentage:     batter.SluggingPercentage,
		OnBasePercentage:       batter.OnBasePercentage,
		WOba:                   batter.Woba,
		RC:                     batter.RC,
		Babip:                  batter.BABIP,
	}
}

func setWoba(batterGrades *player.BATTERGRADES, config *config) {
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

func setSingle(batterGrades *player.BATTERGRADES) {
	batterGrades.Single = batterGrades.Hit - batterGrades.Double - batterGrades.Triple - batterGrades.HomeRun
}
