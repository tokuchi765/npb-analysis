type MaxTeamPitchingResponse = {
  maxStrikeOutRate: number;
  maxRunsAllowed: number;
};

type MinTeamPitchingResponse = {
  minStrikeOutRate: number;
  minRunsAllowed: number;
};

type MaxTeamBattingResponse = {
  maxHomeRun: number;
  maxSluggingPercentage: number;
  maxOnBasePercentage: number;
};

type MinTeamBattingResponse = {
  minHomeRun: number;
  minSluggingPercentage: number;
  minOnBasePercentage: number;
};

type TeamPitchingResponse = {
  teamPitching: {
    StrikeOutRate: number;
    RunsAllowed: number;
  };
};

type TeamBattingResponse = {
  teamBatting: {
    HomeRun: number;
    SluggingPercentage: number;
    OnBasePercentage: number;
  };
};

type PlayerResponse = {
  career: any;
  batting: any;
  pitching: any;
};

type PlayersResponse = {
  careers: CareerResponse[];
};

type CareerResponse = {
  PlayerID: string;
  Name: string;
  Position: string;
  PitchingAndBatting: string;
  Height: string;
  Weight: string;
  Birthday: string;
  Career: string;
  Draft: string;
};

export enum ThresholdType {
  GreaterOrEqual = 0,
  LessOrEqual = 1,
}

type SearchBaseGradesCondition = {
  Total: boolean;
  TotalYear: number | undefined;
  FromYear: number | undefined;
  ToYear: number | undefined;
  TeamID: string;
};

type SearchBatterGradesCondition = {
  PlateAppearanceThresholdType: ThresholdType;
  PlateAppearance: number | undefined;
  OnBasePercentageThresholdType: ThresholdType;
  OnBasePercentage: number | undefined;
  BattingAverageThresholdType: ThresholdType;
  BattingAverage: number | undefined;
  BaseOnBallsThresholdType: ThresholdType;
  BaseOnBalls: number | undefined;
  HitThresholdType: ThresholdType;
  Hit: number | undefined;
  SingleThresholdType: ThresholdType;
  Single: number | undefined;
  DoubleThresholdType: ThresholdType;
  Double: number | undefined;
  TripleThresholdType: ThresholdType;
  Triple: number | undefined;
  SluggingPercentageThresholdType: ThresholdType;
  SluggingPercentage: number | undefined;
  HomeRunThresholdType: ThresholdType;
  HomeRun: number | undefined;
  StrikeOutThresholdType: ThresholdType;
  StrikeOut: number | undefined;
  StrikeOutRateThresholdType: ThresholdType;
  StrikeOutRate: number | undefined;
  StolenBaseThresholdType: ThresholdType;
  StolenBase: number | undefined;
  GroundedIntoDoublePlayThresholdType: ThresholdType;
  GroundedIntoDoublePlay: number | undefined;
  WObaThresholdType: ThresholdType;
  WOba: number | undefined;
  RCThresholdType: ThresholdType;
  RC: number | undefined;
  BabipThresholdType: ThresholdType;
  Babip: number | undefined;
};

type BatterGradesResponse = {
  results: BatterGrades[];
};

type BatterGrades = {
  PlayerID: string;
  Name: string;
  Year: number;
  Team: string;
  Games: number;
  PlateAppearance: number;
  AtBat: number;
  OnBasePercentage: number;
  BattingAverage: number;
  SluggingPercentage: number;
  WOba: number;
  RC: number;
  Babip: number;
  BaseOnBalls: number;
  HitByPitches: number;
  StrikeOut: number;
  StrikeOutRate: number;
  Score: number;
  Hit: number;
  HomeRun: number;
  RunsBattedIn: number;
  Single: number;
  Double: number;
  Triple: number;
  BaseHit: number;
  StolenBase: number;
  CaughtStealing: number;
  SacrificeHits: number;
  SacrificeFlies: number;
  GroundedIntoDoublePlay: number;
};

type SearchPitcherGradesCondition = {
  PitchedThresholdType: ThresholdType;
  Pitched: number | undefined;
  InningsPitchedThresholdType: ThresholdType;
  InningsPitched: number | undefined;
  EarnedRunAverageThresholdType: ThresholdType;
  EarnedRunAverage: number | undefined;
  BabipThresholdType: ThresholdType;
  Babip: number | undefined;
  StrikeOutRateThresholdType: ThresholdType;
  StrikeOutRate: number | undefined;
  StrikeOutThresholdType: ThresholdType;
  StrikeOut: number | undefined;
  HitThresholdType: ThresholdType;
  Hit: number | undefined;
  BaseOnBallsThresholdType: ThresholdType;
  BaseOnBalls: number | undefined;
  HomeRunThresholdType: ThresholdType;
  HomeRun: number | undefined;
  WinThresholdType: ThresholdType;
  Win: number | undefined;
  LoseThresholdType: ThresholdType;
  Lose: number | undefined;
  SaveThresholdType: ThresholdType;
  Save: number | undefined;
  HoldThresholdType: ThresholdType;
  Hold: number | undefined;
  HoldPointThresholdType: ThresholdType;
  HoldPoint: number | undefined;
  CompleteGameThresholdType: ThresholdType;
  CompleteGame: number | undefined;
  ShutoutThresholdType: ThresholdType;
  Shutout: number | undefined;
  WinningRateThresholdType: ThresholdType;
  WinningRate: number | undefined;
};

type PitcherGradesResponse = {
  results: PitcherGrades[];
};

type PitcherGrades = {
  PlayerID: string;
  Name: string;
  Year: number;
  Team: string;
  Pitched: number;
  InningsPitched: number;
  EarnedRunAverage: number;
  Babip: number;
  StrikeOutRate: number;
  StrikeOut: number;
  Hit: number;
  BaseOnBalls: number;
  HomeRun: number;
  Win: number;
  Lose: number;
  Save: number;
  Hold: number;
  HoldPoint: number;
  CompleteGame: number;
  Shutout: number;
  NoWalks: number;
  WinningRate: number;
  Batter: number;
  HitByPitches: number;
  WildPitches: number;
  Balk: number;
  RunsAllowed: number;
  EarnedRun: number;
};

export type {
  MaxTeamPitchingResponse,
  MinTeamPitchingResponse,
  MaxTeamBattingResponse,
  MinTeamBattingResponse,
  TeamPitchingResponse,
  TeamBattingResponse,
  PlayerResponse,
  PlayersResponse,
  CareerResponse,
  SearchBaseGradesCondition,
  BatterGradesResponse,
  BatterGrades,
  SearchBatterGradesCondition,
  SearchPitcherGradesCondition,
  PitcherGradesResponse,
  PitcherGrades,
};
