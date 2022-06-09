package player

import (
	"database/sql"
	"math"
)

// PICHERGRADES 成績
type PICHERGRADES struct {
	Year             string  // 年度
	TeamID           string  // チームID
	Team             string  // 所属球団
	Piched           float64 // 登板
	Win              float64 // 勝利
	Lose             float64 // 敗北
	Save             float64 // セーブ
	Hold             float64 // ホールド
	HoldPoint        float64 // ホールドポイント
	CompleteGame     float64 // 完投
	Shutout          float64 // 完封
	NoWalks          float64 // 無四球
	WinningRate      float64 // 勝率
	Batter           float64 // 打者
	InningsPitched   float64 // 投球回数
	Hit              float64 // 安打
	HomeRun          float64 // ホームラン
	BaseOnBalls      float64 // 四球
	HitByPitches     float64 // 死球
	StrikeOut        float64 // 三振
	WildPitches      float64 // 暴投
	Balk             float64 // ボーク
	RunsAllowed      float64 // 失点
	EarnedRun        float64 // 自責点
	EarnedRunAverage float64 // 防御率
	BABIP            float64 // 被BABIP
	StrikeOutRate    float64 // 奪三振率
}

// SetBABIP 被BABIPを算出して設定する
func (picherGrades *PICHERGRADES) SetBABIP() {
	picherGrades.BABIP = (float64(picherGrades.Hit) - float64(picherGrades.HomeRun)) / (float64(picherGrades.Batter) - (float64(picherGrades.BaseOnBalls) + float64(picherGrades.HitByPitches)) - float64(picherGrades.StrikeOut) - float64(picherGrades.HomeRun))
	if math.IsNaN(picherGrades.BABIP) {
		picherGrades.BABIP = 0.0
	}
}

// SetStrikeOutRate 奪三振率を算出して設定する
func (picherGrades *PICHERGRADES) SetStrikeOutRate() {
	picherGrades.StrikeOutRate = (picherGrades.StrikeOut * 9) / picherGrades.InningsPitched
	if math.IsNaN(picherGrades.StrikeOutRate) {
		picherGrades.StrikeOutRate = 0.0
	}
}

// BATTERGRADES 成績
type BATTERGRADES struct {
	Year                   string          // 年度
	TeamID                 string          // チームID
	Team                   string          // 所属球団
	Games                  int             // 試合
	PlateAppearance        int             // 打席
	AtBat                  int             // 打数
	Score                  int             // 得点
	Hit                    int             // 安打
	Single                 int             // 単打
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
	HitByPitches           int             // 死球
	StrikeOut              int             // 三振
	StrikeOutRate          sql.NullFloat64 // 三振率
	GroundedIntoDoublePlay int             // 併殺打
	BattingAverage         float64         // 打率
	SluggingPercentage     float64         // 長打率
	OnBasePercentage       float64         // 出塁率
	Woba                   float64         // 加重出塁率
	RC                     float64         // 創出得点
	BABIP                  float64         // BABIP
}

// SetStrikeOutRate 三振率を算出して設定する
func (batterGrades *BATTERGRADES) SetStrikeOutRate() {
	strikeOutRate := (float64(batterGrades.StrikeOut)) / float64(batterGrades.PlateAppearance)

	if math.IsNaN(strikeOutRate) {
		strikeOutRate = 0.0
	}

	batterGrades.StrikeOutRate = sql.NullFloat64{
		Float64: strikeOutRate,
	}
}

// SetRC RCを算出して設定する
func (batterGrades *BATTERGRADES) SetRC() {
	A := float64(batterGrades.Hit + batterGrades.BaseOnBalls + batterGrades.HitByPitches - batterGrades.CaughtStealing - batterGrades.GroundedIntoDoublePlay)
	B := float64(batterGrades.BaseHit) + (0.26 * float64(batterGrades.BaseOnBalls+batterGrades.HitByPitches)) + (0.53 * float64(batterGrades.SacrificeHits+batterGrades.SacrificeFlies)) + (0.64 * float64(batterGrades.StolenBase)) - (0.03 * float64(batterGrades.StrikeOut))
	C := float64(batterGrades.AtBat + batterGrades.BaseOnBalls + batterGrades.HitByPitches + batterGrades.SacrificeFlies + batterGrades.SacrificeHits)
	batterGrades.RC = ((A + (2.4 * C)) * (B + (3 * C)) / (9 * C)) - (0.9 * C)
	if math.IsNaN(batterGrades.RC) {
		batterGrades.RC = 0.0
	}
}

// SetBABIP BABIPを算出して設定する
func (batterGrades *BATTERGRADES) SetBABIP() {
	batterGrades.BABIP = (float64(batterGrades.Hit) - float64(batterGrades.HomeRun)) / (float64(batterGrades.AtBat) - float64(batterGrades.StrikeOut) - float64(batterGrades.HomeRun) + float64(batterGrades.SacrificeFlies))
	if math.IsNaN(batterGrades.BABIP) {
		batterGrades.BABIP = 0.0
	}
}

// CAREER 成績
type CAREER struct {
	PlayerID           string // 選手ID
	Name               string // 選手名
	Position           string // ポジション
	PitchingAndBatting string // 投打
	Height             string // 身長
	Weight             string // 体重
	Birthday           string // 生年月日
	Career             string // 経歴
	Draft              string // ドラフト
}

type PLAYER struct {
	Year     string // 年度
	TeamID   string // チームID
	PlayerID string // 選手ID
	Team     string // 所属球団
	Name     string // 選手名
}
