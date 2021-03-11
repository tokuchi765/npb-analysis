package player

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
}

// BATTERGRADES 成績
type BATTERGRADES struct {
	Year                   string  // 年度
	TeamID                 string  // チームID
	Team                   string  // 所属球団
	Games                  int     // 試合
	PlateAppearance        int     // 打席
	AtBat                  int     // 打数
	Score                  int     // 得点
	Hit                    int     // 安打
	Single                 int     //単打
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
	GroundedIntoDoublePlay int     // 併殺打
	BattingAverage         float64 // 打率
	SluggingPercentage     float64 // 長打率
	OnBasePercentage       float64 // 出塁率
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
