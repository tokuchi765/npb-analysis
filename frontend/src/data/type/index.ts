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
};
