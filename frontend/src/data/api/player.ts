import { rest } from '../rest';
import {
  BatterGradesResponse,
  PitcherGradesResponse,
  PlayerResponse,
  PlayersResponse,
  SearchBaseGradesCondition,
  SearchBatterGradesCondition,
  SearchPitcherGradesCondition,
} from '../type';

const baseUri = '/player';

const getPlayer = async (playerID: string): Promise<PlayerResponse> => {
  try {
    const { data } = await rest.get<PlayerResponse>(`${baseUri}/${playerID}`);
    return data;
  } catch (error: any) {
    throw new Error(error);
  }
};

const searchPlayer = async (name: string): Promise<PlayersResponse> => {
  try {
    const { data } = await rest.getParams(baseUri + '/search', {
      params: {
        Name: name,
      },
    });
    return data;
  } catch (error: any) {
    throw new Error(error);
  }
};

const searchBatterGrades = async (
  baseCondition: SearchBaseGradesCondition,
  condition: SearchBatterGradesCondition
): Promise<BatterGradesResponse> => {
  try {
    const { data } = await rest.getParams(baseUri + '/search/batter-grades', {
      params: {
        total: baseCondition.Total,
        totalYear: baseCondition.TotalYear,
        fromYear: baseCondition.FromYear,
        toYear: baseCondition.ToYear,
        teamID: baseCondition.TeamID,
        plateAppearanceThresholdType: condition.PlateAppearanceThresholdType,
        plateAppearance: condition.PlateAppearance,
        onBasePercentageThresholdType: condition.OnBasePercentageThresholdType,
        onBasePercentage: condition.OnBasePercentage,
        battingAverageThresholdType: condition.BattingAverageThresholdType,
        battingAverage: condition.BattingAverage,
        sluggingPercentageThresholdType: condition.SluggingPercentageThresholdType,
        sluggingPercentage: condition.SluggingPercentage,
        homeRunThresholdType: condition.HomeRunThresholdType,
        homeRun: condition.HomeRun,
        strikeOutThresholdType: condition.StrikeOutThresholdType,
        strikeOut: condition.StrikeOut,
        strikeOutRateThresholdType: condition.StrikeOutRateThresholdType,
        strikeOutRate: condition.StrikeOutRate,
        groundedIntoDoublePlayThresholdType: condition.GroundedIntoDoublePlayThresholdType,
        groundedIntoDoublePlay: condition.GroundedIntoDoublePlay,
        wObaThresholdType: condition.WObaThresholdType,
        wOba: condition.WOba,
        rCThresholdType: condition.RCThresholdType,
        rC: condition.RC,
        babipThresholdType: condition.BabipThresholdType,
        babip: condition.Babip,
      },
    });
    return data;
  } catch (error: any) {
    throw new Error(error);
  }
};

const searchPitcherGrades = async (
  baseCondition: SearchBaseGradesCondition,
  condition: SearchPitcherGradesCondition
): Promise<PitcherGradesResponse> => {
  try {
    const { data } = await rest.getParams(baseUri + '/search/pitcher-grades', {
      params: {
        total: baseCondition.Total,
        totalYear: baseCondition.TotalYear,
        fromYear: baseCondition.FromYear,
        toYear: baseCondition.ToYear,
        teamID: baseCondition.TeamID,
        pitchedThresholdType: condition.PitchedThresholdType,
        pitched: condition.Pitched,
        inningsPitchedThresholdType: condition.InningsPitchedThresholdType,
        inningsPitched: condition.InningsPitched,
        earnedRunAverageThresholdType: condition.EarnedRunAverageThresholdType,
        earnedRunAverage: condition.EarnedRunAverage,
        babipThresholdType: condition.BabipThresholdType,
        babip: condition.Babip,
        strikeOutRateThresholdType: condition.StrikeOutRateThresholdType,
        strikeOutRate: condition.StrikeOutRate,
        strikeOutThresholdType: condition.StrikeOutThresholdType,
        strikeOut: condition.StrikeOut,
        hitThresholdType: condition.HitThresholdType,
        hit: condition.Hit,
        baseOnBallsThresholdType: condition.BaseOnBallsThresholdType,
        baseOnBalls: condition.BaseOnBalls,
        homeRunThresholdType: condition.HomeRunThresholdType,
        homeRun: condition.HomeRun,
        winThresholdType: condition.WinThresholdType,
        win: condition.Win,
        loseThresholdType: condition.LoseThresholdType,
        lose: condition.Lose,
        saveThresholdType: condition.SaveThresholdType,
        save: condition.Save,
        holdThresholdType: condition.HoldThresholdType,
        hold: condition.Hold,
        holdPointThresholdType: condition.HoldPointThresholdType,
        holdPoint: condition.HoldPoint,
        completeGameThresholdType: condition.CompleteGameThresholdType,
        completeGame: condition.CompleteGame,
        shutoutThresholdType: condition.ShutoutThresholdType,
        shutout: condition.Shutout,
        winningRateThresholdType: condition.WinningRateThresholdType,
        winningRate: condition.WinningRate,
      },
    });
    return data;
  } catch (error: any) {
    throw new Error(error);
  }
};

export { getPlayer, searchPlayer, searchBatterGrades, searchPitcherGrades };
