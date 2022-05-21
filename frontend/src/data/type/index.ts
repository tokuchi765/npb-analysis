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

export type {
  MaxTeamPitchingResponse,
  MinTeamPitchingResponse,
  MaxTeamBattingResponse,
  MinTeamBattingResponse,
};
