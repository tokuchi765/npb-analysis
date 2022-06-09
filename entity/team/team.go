package team

import (
	"database/sql"
	"math"
)

// TEAMDATA チーム情報
type TEAMDATA struct {
	TeamID     string // チームID
	TeamNameID string // チーム名ID
	League     string // リーグ
}

// TEAMNAME チーム名
type TEAMNAME struct {
	TeamNameID string // チーム名ID
	TeamName   string // チーム名
}

// TeamBatting チーム打撃成績
type TeamBatting struct {
	TeamID                 string          // チームID
	Year                   string          // 年度
	BattingAverage         float64         // 打率
	Games                  int             // 試合
	PlateAppearance        int             // 打席
	AtBat                  int             // 打数
	Score                  int             // 得点
	Hit                    int             // 安打
	Double                 int             // 二塁打
	Triple                 int             // 三塁打
	HomeRun                int             // 本塁打
	BaseHit                int             // 塁打
	RunsBattedIn           int             // 打点
	StolenBase             int             // 盗塁
	CaughtStealing         int             // 盗塁刺
	SacrificeHits          int             // 犠打
	SacrificeFlies         int             // 犠飛
	BaseOnBalls            int             // 四球
	IntentionalWalk        int             // 故意四球
	HitByPitches           int             // 死球
	StrikeOut              int             // 三振
	StrikeOutRate          sql.NullFloat64 // 三振率
	GroundedIntoDoublePlay int             // 併殺打
	SluggingPercentage     float64         // 長打率
	OnBasePercentage       float64         // 出塁率
	BABIP                  float64         // BABIP
}

// SetBABIP BABIPを算出して設定する
func (teamBatting *TeamBatting) SetBABIP() {
	teamBatting.BABIP = (float64(teamBatting.Hit) - float64(teamBatting.HomeRun)) / (float64(teamBatting.AtBat) - float64(teamBatting.StrikeOut) - float64(teamBatting.HomeRun) + float64(teamBatting.SacrificeFlies))
	if math.IsNaN(teamBatting.BABIP) {
		teamBatting.BABIP = 0.0
	}
}

// SetStrikeOutRate 三振率を算出して設定する
func (teamBatting *TeamBatting) SetStrikeOutRate() {
	strikeOutRate := (float64(teamBatting.StrikeOut)) / float64(teamBatting.PlateAppearance)

	if math.IsNaN(strikeOutRate) {
		strikeOutRate = 0.0
	}

	teamBatting.StrikeOutRate = sql.NullFloat64{
		Float64: strikeOutRate,
	}
}

// TeamPitching チーム投手成績
type TeamPitching struct {
	TeamID           string  // チームID
	Year             string  // 年度
	EarnedRunAverage float64 // 防御率
	Games            int     // 試合
	Win              int     // 勝利
	Lose             int     // 敗北
	Save             int     // セーブ
	Hold             int     // ホールド
	HoldPoint        int     // ホールドポイント
	CompleteGame     int     // 完投
	Shutout          int     // 完封
	NoWalks          int     // 無四球
	WinningRate      float64 // 勝率
	Batter           int     // 打者
	InningsPitched   int     // 投球回数
	Hit              int     // 安打
	HomeRun          int     // ホームラン
	BaseOnBalls      int     // 四球
	IntentionalWalk  int     // 故意四球
	HitByPitches     int     // 死球
	StrikeOut        int     // 三振
	WildPitches      int     // 暴投
	Balk             int     // ボーク
	RunsAllowed      int     // 失点
	EarnedRun        int     // 自責点
	BABIP            float64 // 被BABIP
	StrikeOutRate    float64 // 奪三振率
}

// SetBABIP 被BABIPを算出して設定する
func (teamPitching *TeamPitching) SetBABIP() {
	teamPitching.BABIP = (float64(teamPitching.Hit) - float64(teamPitching.HomeRun)) / (float64(teamPitching.Batter) - (float64(teamPitching.BaseOnBalls) + float64(teamPitching.HitByPitches)) - float64(teamPitching.StrikeOut) - float64(teamPitching.HomeRun))
	if math.IsNaN(teamPitching.BABIP) {
		teamPitching.BABIP = 0.0
	}
}

// SetStrikeOutRate 奪三振率を算出して設定する
func (teamPitching *TeamPitching) SetStrikeOutRate() {
	teamPitching.StrikeOutRate = (float64(teamPitching.StrikeOut) * 9) / float64(teamPitching.InningsPitched)
	if math.IsNaN(teamPitching.StrikeOutRate) {
		teamPitching.StrikeOutRate = 0.0
	}
}

// TeamLeagueStats チームシーズン成績
type TeamLeagueStats struct {
	TeamID                 string  // チームID
	Year                   string  // 年度
	Manager                string  // 監督
	Games                  int     // 試合
	Win                    int     // 勝利
	Lose                   int     // 敗北
	Draw                   int     // 引き分け
	WinningRate            float64 // 勝率
	ExchangeWin            int     // 交流戦勝利
	ExchangeLose           int     // 交流戦敗北
	ExchangeDraw           int     // 交流戦引き分け
	HomeWin                int     // ホーム勝利
	HomeLose               int     // ホーム敗北
	HomeDraw               int     // ホーム引き分け
	LoadWin                int     // ロード勝利
	LoadLose               int     // ロード敗北
	LoadDraw               int     // ロード引き分け
	PythagoreanExpectation float64 // ピタゴラス勝率
}

// TeamMatchResults チーム対戦成績
type TeamMatchResults struct {
	TeamID            string // チームID
	Year              string // 年度
	CompetitiveTeamID string // 対戦チームID
	VsType            string // 対戦タイプ
	Win               int    // 勝利
	Lose              int    // 敗北
	Draw              int    // 引き分け
}
