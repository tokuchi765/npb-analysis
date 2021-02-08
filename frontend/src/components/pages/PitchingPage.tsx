import React, { useEffect, useState } from 'react';
import GenericTemplate from '../templates/GenericTemplate';
import axios from 'axios';
import _ from 'lodash';

interface PitchingData {
  main: string;
  earnedRunAverage: number;
  games: number;
  save: number;
  hold: number;
  homeRun: number;
  baseOnBalls: number;
  strikeOut: number;
}

function createPitchingData(
  main: string,
  earnedRunAverage: number,
  games: number,
  save: number,
  hold: number,
  homeRun: number,
  baseOnBalls: number,
  strikeOut: number
) {
  const result: PitchingData = {
    main,
    earnedRunAverage,
    games,
    save,
    hold,
    homeRun,
    baseOnBalls,
    strikeOut,
  };
  return result;
}

function createPitchingDataList(
  teams:
    | {
        Giants: any;
        Baystars: any;
        Tigers: any;
        Carp: any;
        Dragons: any;
        Swallows: any;
      }[]
    | {
        Lions: any;
        Hawks: any;
        Eagles: any;
        Marines: any;
        Fighters: any;
        Buffaloes: any;
      }[]
) {
  const datas: PitchingData[] = [];
  teams.forEach((team: any) => {
    _.forEach(team, (val, key) => {
      datas.push(
        createPitchingData(
          key,
          val.EarnedRunAverage,
          val.Games,
          val.Save,
          val.Hold,
          val.HomeRun,
          val.BaseOnBalls,
          val.StrikeOut
        )
      );
    });
  });
  return datas;
}

const PitchingPage: React.FC = () => {
  const [centralDatas, setCentralData] = useState<PitchingData[]>([]);
  const [pacificDatas, setPacificData] = useState<PitchingData[]>([]);

  const getCentralPitchingDataList = async (year: string) => {
    const result = await axios.get(
      `http://localhost:8081/team/pitching?from_year=${year}&to_year=${year}`
    );

    const centralPitchings = _.map(result.data.teamPitching, (teamPitching) => {
      return {
        Giants: _.filter(teamPitching, { TeamID: '01' })[0],
        Baystars: _.filter(teamPitching, { TeamID: '02' })[0],
        Tigers: _.filter(teamPitching, { TeamID: '03' })[0],
        Carp: _.filter(teamPitching, { TeamID: '04' })[0],
        Dragons: _.filter(teamPitching, { TeamID: '05' })[0],
        Swallows: _.filter(teamPitching, { TeamID: '06' })[0],
      };
    });

    setCentralData(createPitchingDataList(centralPitchings));
  };

  const getPacificPitchingDataList = async (year: string) => {
    const result = await axios.get(
      `http://localhost:8081/team/pitching?from_year=${year}&to_year=${year}`
    );

    const pacificPitchings = _.map(result.data.teamPitching, (teamPitching) => {
      return {
        Lions: _.filter(teamPitching, { TeamID: '07' })[0],
        Hawks: _.filter(teamPitching, { TeamID: '08' })[0],
        Eagles: _.filter(teamPitching, { TeamID: '09' })[0],
        Marines: _.filter(teamPitching, { TeamID: '10' })[0],
        Fighters: _.filter(teamPitching, { TeamID: '11' })[0],
        Buffaloes: _.filter(teamPitching, { TeamID: '12' })[0],
      };
    });

    setPacificData(createPitchingDataList(pacificPitchings));
  };

  useEffect(() => {
    (async () => {
      getCentralPitchingDataList('2020');
      getPacificPitchingDataList('2020');
    })();
  }, []);

  return <GenericTemplate title="チーム投手成績ページ">チーム投手成績ページ</GenericTemplate>;
};

export default PitchingPage;
