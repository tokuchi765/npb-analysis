package player

import (
	"database/sql"
	"math"
	"strings"

	"github.com/tokuchi765/npb-analysis/entity/sqlwrapper"
)

// PICHERGRADES 成績
type PICHERGRADES struct {
	Year             string  // 年度
	TeamID           string  // チームID
	Team             string  // 所属球団
	Pitched          float64 // 登板
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

// PitcherGrades 投手成績テーブルマッピング
type PitcherGrades struct {
	PlayerID         string  // 選手ID
	Year             string  // 年度
	TeamID           string  // チームID
	Team             string  // 所属球団
	Pitched          float64 // 登板
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
	Babip            float64 // 被BABIP
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

// SetInningsPitched 投球回数を正確な数値に変換します
func (picherGrades *PICHERGRADES) SetInningsPitched() {
	int, frac := math.Modf(picherGrades.InningsPitched)
	precision := 2
	multiplier := math.Pow(10, float64(precision))
	frac = math.Round(frac*multiplier) / multiplier
	picherGrades.InningsPitched = float64(int) + frac*3.0
}

// SearchPitcherGradesCondition 選手投手成績検索コンディション
type SearchPitcherGradesCondition struct {
	Total                         bool          // 通算
	TotalYear                     int           // 通算年数
	FromYear                      int           // 年度From
	ToYear                        int           // 年度To
	TeamID                        string        // チームID
	Pitched                       int           // 登板
	InningsPitched                int           // 投球回数
	EarnedRunAverageThresholdType ThresholdType // 防御率の以上、以下
	EarnedRunAverage              float64       // 防御率
	BabipThresholdType            ThresholdType // 被BABIPの以上、以下
	Babip                         float64       // 被BABIP
	StrikeOutRateThresholdType    ThresholdType // 奪三振率の以上、以下
	StrikeOutRate                 float64       // 奪三振率
	StrikeOutThresholdType        ThresholdType // 奪三振数の以上、以下
	StrikeOut                     int           // 奪三振数
	HitThresholdType              ThresholdType // 被安打数の以上、以下
	Hit                           int           // 被安打数
	BaseOnBallsThresholdType      ThresholdType // 四球数の以上、以下
	BaseOnBalls                   int           // 四球数
	HomeRunThresholdType          ThresholdType // 被ホームラン数の以上、以下
	HomeRun                       int           // 被ホームラン数
	WinThresholdType              ThresholdType // 勝利数の以上、以下
	Win                           int           // 勝利数
	LoseThresholdType             ThresholdType // 敗北数の以上、以下
	Lose                          int           // 敗北数
	SaveThresholdType             ThresholdType // セーブ数の以上、以下
	Save                          int           // セーブ数
	HoldThresholdType             ThresholdType // ホールド数の以上、以下
	Hold                          int           // ホールド数
	HoldPointThresholdType        ThresholdType // ホールドポイント数の以上、以下
	HoldPoint                     int           // ホールドポイント数
	CompleteGameThresholdType     ThresholdType // 完投数の以上、以下
	CompleteGame                  int           // 完投数
	ShutoutThresholdType          ThresholdType // 完封数の以上、以下
	Shutout                       int           // 完封数
	WinningRateThresholdType      ThresholdType // 勝率数の以上、以下
	WinningRate                   float64       // 勝率数
}

// SearchPitcherGradesResult 選手投手成績検索結果
type SearchPitcherGradesResult struct {
	PlayerID         string  // 選手ID
	Name             string  // 選手名
	Year             string  // 年度
	Team             string  // 所属球団
	Pitched          float64 // 登板
	InningsPitched   float64 // 投球回数
	EarnedRunAverage float64 // 防御率
	Babip            float64 // 被BABIP
	StrikeOutRate    float64 // 奪三振率
	StrikeOut        int     // 三振
	Hit              int     // 安打
	BaseOnBalls      int     // 四球
	HomeRun          int     // ホームラン
	Win              int     // 勝利
	Lose             int     // 敗北
	Save             int     // セーブ
	Hold             int     // ホールド
	HoldPoint        int     // ホールドポイント
	CompleteGame     int     // 完投
	Shutout          int     // 完封
	NoWalks          float64 // 無四球
	WinningRate      float64 // 勝率
	Batter           float64 // 打者
	HitByPitches     float64 // 死球
	WildPitches      float64 // 暴投
	Balk             float64 // ボーク
	RunsAllowed      float64 // 失点
	EarnedRun        float64 // 自責点
}

// BATTERGRADES 成績
type BATTERGRADES struct {
	Year                   string                 // 年度
	TeamID                 string                 // チームID
	Team                   string                 // 所属球団
	Games                  int                    // 試合
	PlateAppearance        int                    // 打席
	AtBat                  int                    // 打数
	Score                  int                    // 得点
	Hit                    int                    // 安打
	Single                 int                    // 単打
	Double                 int                    // 二塁打
	Triple                 int                    // 三塁打
	HomeRun                int                    // 本塁打
	BaseHit                int                    // 塁打
	RunsBattedIn           int                    // 打点
	StolenBase             int                    // 盗塁
	CaughtStealing         int                    // 盗塁刺
	SacrificeHits          int                    // 犠打
	SacrificeFlies         int                    // 犠飛
	BaseOnBalls            int                    // 四球
	HitByPitches           int                    // 死球
	StrikeOut              int                    // 三振
	StrikeOutRate          sqlwrapper.NullFloat64 // 三振率
	GroundedIntoDoublePlay int                    // 併殺打
	BattingAverage         float64                // 打率
	SluggingPercentage     float64                // 長打率
	OnBasePercentage       float64                // 出塁率
	Woba                   float64                // 加重出塁率
	RC                     float64                // 創出得点
	BABIP                  float64                // BABIP
}

// BatterGrades 野手成績テーブルマッピング
type BatterGrades struct {
	PlayerID               string  // 選手ID
	Year                   string  // 年度
	TeamID                 string  // チームID
	Team                   string  // 所属球団
	Games                  int     // 試合
	PlateAppearance        int     // 打席
	AtBat                  int     // 打数
	Score                  int     // 得点
	Hit                    int     // 安打
	Single                 int     // 単打
	Double                 int     // 二塁打
	Triple                 int     // 三塁打
	HomeRun                int     // 本塁打
	BaseHit                int     // 塁打
	RunsBattedIn           int     // 打点
	StolenBase             int     // 盗塁
	CaughtStealing         int     // 盗塁刺
	SacrificeHits          int     // 犠打
	SacrificeFlies         int     // 犠飛
	BaseOnBalls            int     // 四球
	HitByPitches           int     // 死球
	StrikeOut              int     // 三振
	StrikeOutRate          float64 // 三振率
	GroundedIntoDoublePlay int     // 併殺打
	BattingAverage         float64 // 打率
	SluggingPercentage     float64 // 長打率
	OnBasePercentage       float64 // 出塁率
	WOba                   float64 // 加重出塁率
	RC                     float64 // 創出得点
	Babip                  float64 // BABIP
}

// SetStrikeOutRate 三振率を算出して設定する
func (batterGrades *BATTERGRADES) SetStrikeOutRate() {
	strikeOutRate := (float64(batterGrades.StrikeOut)) / float64(batterGrades.PlateAppearance)

	if math.IsNaN(strikeOutRate) {
		strikeOutRate = 0.0
	}

	batterGrades.StrikeOutRate = sqlwrapper.NullFloat64{
		NullFloat64: sql.NullFloat64{
			Float64: strikeOutRate,
			Valid:   true,
		},
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
	SearchName         string // 検索用選手名
}

// Players 選手情報テーブルマッピング
type Players struct {
	PlayerID           string // 選手ID
	Name               string // 選手名
	Position           string // ポジション
	PitchingAndBatting string // 投打
	Height             string // 身長
	Weight             string // 体重
	Birthday           string // 生年月日
	Career             string // 経歴
	Draft              string // ドラフト
	SearchName         string // 検索用選手名
}

// SetSearchName 選手名検索に不要な文字列を除去した検索用選手名を設定する
func (career *CAREER) SetSearchName() {
	career.SearchName = strings.ReplaceAll(strings.ReplaceAll(career.Name, "　", ""), "・", "")
}

type PLAYER struct {
	Year     string // 年度
	TeamID   string // チームID
	PlayerID string // 選手ID
	Team     string // 所属球団
	Name     string // 選手名
}

// ThresholdType 以上、以下のタイプ
type ThresholdType int

const (
	GreaterOrEqual ThresholdType = iota // 以上
	LessOrEqual                         // 以下
)

func (d ThresholdType) String() string {
	return [...]string{">=", "<="}[d]
}

// SearchBatterGradesCondition 選手野手成績検索コンディション
type SearchBatterGradesCondition struct {
	Total                               bool          // 通算
	TotalYear                           int           // 通算年数
	FromYear                            int           // 年度From
	ToYear                              int           // 年度To
	TeamID                              string        // チームID
	PlateAppearance                     int           // 打席
	OnBasePercentageThresholdType       ThresholdType // 出塁率の以上、以下
	OnBasePercentage                    float64       // 出塁率
	SluggingPercentageThresholdType     ThresholdType // 長打率の以上、以下
	SluggingPercentage                  float64       // 長打率
	HomeRunThresholdType                ThresholdType // 本塁打の以上、以下
	HomeRun                             int           // 本塁打
	StrikeOutThresholdType              ThresholdType // 三振の以上、以下
	StrikeOut                           int           // 三振
	StrikeOutRateThresholdType          ThresholdType // 三振率の以上、以下
	StrikeOutRate                       float64       // 三振率
	GroundedIntoDoublePlayThresholdType ThresholdType // 併殺打の以上、以下
	GroundedIntoDoublePlay              int           // 併殺打
	WObaThresholdType                   ThresholdType // 加重出塁率の以上、以下
	WOba                                float64       // 加重出塁率
	RCThresholdType                     ThresholdType // 創出得点の以上、以下
	RC                                  float64       // 創出得点
	BabipThresholdType                  ThresholdType // BABIPの以上、以下
	Babip                               float64       // BABIP
}

// SearchBatterGradesResult 選手野手成績検索結果
type SearchBatterGradesResult struct {
	PlayerID               string                 // 選手ID
	Name                   string                 // 選手名
	Year                   string                 // 年度
	Team                   string                 // 所属球団
	Games                  int                    // 試合
	PlateAppearance        int                    // 打席
	AtBat                  int                    // 打数
	OnBasePercentage       float64                // 出塁率
	BattingAverage         float64                // 打率
	SluggingPercentage     float64                // 長打率
	WOba                   float64                // 加重出塁率
	RC                     float64                // 創出得点
	Babip                  float64                // BABIP
	BaseOnBalls            int                    // 四球
	HitByPitches           int                    // 死球
	StrikeOut              int                    // 三振
	StrikeOutRate          sqlwrapper.NullFloat64 // 三振率
	Score                  int                    // 得点
	Hit                    int                    // 安打
	HomeRun                int                    // 本塁打
	RunsBattedIn           int                    // 打点
	Single                 int                    // 単打
	Double                 int                    // 二塁打
	Triple                 int                    // 三塁打
	BaseHit                int                    // 塁打
	StolenBase             int                    // 盗塁
	CaughtStealing         int                    // 盗塁刺
	SacrificeHits          int                    // 犠打
	SacrificeFlies         int                    // 犠飛
	GroundedIntoDoublePlay int                    // 併殺打
}
