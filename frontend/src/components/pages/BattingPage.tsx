import React, { useEffect, useState } from 'react';
import GenericTemplate from '../templates/GenericTemplate';
import { TableComponent, HeadCell, SelectItem } from '../common/TableComponent';
import axios from 'axios';
import _ from 'lodash';

interface BattingData {
  main: string;
  battingAverage: number;
  games: number;
  score: number;
  hit: number;
  homeRun: number;
  baseOnBalls: number;
  strikeOut: number;
  onBasePercentage: number;
}

const headCells: HeadCell[] = [
  { id: 'main', numeric: false, disablePadding: true, label: 'チーム名' },
  { id: 'battingAverage', numeric: true, disablePadding: false, label: '打率' },
  { id: 'games', numeric: true, disablePadding: false, label: '試合' },
  { id: 'score', numeric: true, disablePadding: false, label: '得点' },
  { id: 'hit', numeric: true, disablePadding: false, label: '安打' },
  { id: 'homeRun', numeric: true, disablePadding: false, label: '本塁打' },
  { id: 'baseOnBalls', numeric: true, disablePadding: false, label: '四球' },
  { id: 'strikeOut', numeric: true, disablePadding: false, label: '三振' },
  { id: 'onBasePercentage', numeric: true, disablePadding: false, label: '出塁率' },
];

function createBattingData(
  main: string,
  battingAverage: number,
  games: number,
  score: number,
  hit: number,
  homeRun: number,
  baseOnBalls: number,
  strikeOut: number,
  onBasePercentage: number
) {
  const result: BattingData = {
    main,
    battingAverage,
    games,
    score,
    hit,
    homeRun,
    baseOnBalls,
    strikeOut,
    onBasePercentage,
  };
  return result;
}

function createBattingDataList(
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
  const teamDataList: BattingData[] = [];
  teams.forEach((team: any) => {
    _.forEach(team, (val, key) => {
      teamDataList.push(
        createBattingData(
          key,
          val.BattingAverage,
          val.Games,
          val.Score,
          val.Hit,
          val.HomeRun,
          val.BaseOnBalls,
          val.StrikeOut,
          val.OnBasePercentage
        )
      );
    });
  });
  return teamDataList;
}

interface TeamBattingResponse {
  teamBatting: any;
}

function BattingPage(props: { years: string[]; initYear: string }) {
  const [initCentralYear, setCentralYear] = useState<string>('');
  const [centralBattingDatas, setCentralBattingData] = useState<BattingData[]>([]);

  const getBattingCentralDataList = async () => {
    const result = await axios.get<TeamBattingResponse>(
      `http://localhost:8081/team/batting?from_year=${initCentralYear}&to_year=${initCentralYear}`
    );

    const teams = _.map(result.data.teamBatting, (teamBatting) => {
      const teanStatses = {
        Giants: _.filter(teamBatting, { TeamID: '01' })[0],
        Baystars: _.filter(teamBatting, { TeamID: '02' })[0],
        Tigers: _.filter(teamBatting, { TeamID: '03' })[0],
        Carp: _.filter(teamBatting, { TeamID: '04' })[0],
        Dragons: _.filter(teamBatting, { TeamID: '05' })[0],
        Swallows: _.filter(teamBatting, { TeamID: '06' })[0],
      };
      return teanStatses;
    });

    setCentralBattingData(createBattingDataList(teams));
  };

  useEffect(() => {
    (async () => {
      if (_.isEmpty(initCentralYear)) {
        setCentralYear(props.initYear);
      } else {
        getBattingCentralDataList();
      }
    })();
  }, [initCentralYear]);

  const [initPacificYear, setPacificYear] = useState<string>('');
  const [pacificBattingDatas, setPacificBattingData] = useState<BattingData[]>([]);

  const getBattingPacificDataList = async () => {
    const result = await axios.get<TeamBattingResponse>(
      `http://localhost:8081/team/batting?from_year=${initPacificYear}&to_year=${initPacificYear}`
    );

    const pacificTeams = _.map(result.data.teamBatting, (teamBatting) => {
      const teamBattings = {
        Lions: _.filter(teamBatting, { TeamID: '07' })[0],
        Hawks: _.filter(teamBatting, { TeamID: '08' })[0],
        Eagles: _.filter(teamBatting, { TeamID: '09' })[0],
        Marines: _.filter(teamBatting, { TeamID: '10' })[0],
        Fighters: _.filter(teamBatting, { TeamID: '11' })[0],
        Buffaloes: _.filter(teamBatting, { TeamID: '12' })[0],
      };
      return teamBattings;
    });

    setPacificBattingData(createBattingDataList(pacificTeams));
  };

  useEffect(() => {
    (async () => {
      if (_.isEmpty(initPacificYear)) {
        setPacificYear(props.initYear);
      } else {
        getBattingPacificDataList();
      }
    })();
  }, [initPacificYear]);

  return (
    <GenericTemplate title="チーム打撃成績ページ">
      <TableComponent
        title={'シーズン打撃成績(セ)'}
        datas={centralBattingDatas}
        headCells={headCells}
        initSorted={'battingAverage'}
        selectItems={[new SelectItem(initCentralYear, '年', props.years, setCentralYear)]}
      />
      <TableComponent
        title={'シーズン打撃成績(パ)'}
        datas={pacificBattingDatas}
        headCells={headCells}
        initSorted={'battingAverage'}
        selectItems={[new SelectItem(initPacificYear, '年', props.years, setPacificYear)]}
      />
    </GenericTemplate>
  );
}

export default BattingPage;
