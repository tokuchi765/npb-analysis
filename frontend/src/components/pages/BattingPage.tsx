import React, { useEffect, useState } from 'react';
import GenericTemplate from '../templates/GenericTemplate';
import TablePages, { HeadCell } from '../common/TablePages';
import axios from 'axios';
import _ from 'lodash';

const years = [
  '2005',
  '2006',
  '2007',
  '2008',
  '2009',
  '2010',
  '2011',
  '2012',
  '2013',
  '2014',
  '2015',
  '2016',
  '2017',
  '2018',
  '2019',
  '2020',
];

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

const BattingPage: React.FC = () => {
  const [initCentralYear, setCentralYear] = useState<string>('');
  const [centralBattingDatas, setCentralBattingData] = useState<BattingData[]>([]);

  const getBattingCentralDataList = async (year: string) => {
    const result = await axios.get(
      `http://localhost:8081/team/batting?from_year=${year}&to_year=${year}`
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

    setCentralYear(year);
    setCentralBattingData(createBattingDataList(teams));
  };

  const [initPacificYear, setPacificYear] = useState<string>('');
  const [pacificBattingDatas, setPacificBattingData] = useState<BattingData[]>([]);

  const getBattingPacificDataList = async (year: string) => {
    const result = await axios.get(
      `http://localhost:8081/team/batting?from_year=${year}&to_year=${year}`
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

    setPacificYear(year);
    setPacificBattingData(createBattingDataList(pacificTeams));
  };

  useEffect(() => {
    (async () => {
      getBattingCentralDataList('2020');
      getBattingPacificDataList('2020');
    })();
  }, []);

  return (
    <GenericTemplate title="チーム打撃成績ページ">
      <TablePages
        title={'シーズン打撃成績(セ)'}
        setSelect={setCentralYear}
        getDataList={getBattingCentralDataList}
        datas={centralBattingDatas}
        selects={years}
        headCells={headCells}
        initSorted={'battingAverage'}
        initSelect={initCentralYear}
        selectLabel={'年'}
        mainLink={false}
        linkValues={new Map<string, string>()}
        path={''}
      />
      <TablePages
        title={'シーズン打撃成績(パ)'}
        setSelect={setPacificYear}
        getDataList={getBattingPacificDataList}
        datas={pacificBattingDatas}
        selects={years}
        headCells={headCells}
        initSorted={'battingAverage'}
        initSelect={initPacificYear}
        selectLabel={'年'}
        mainLink={false}
        linkValues={new Map<string, string>()}
        path={''}
      />
    </GenericTemplate>
  );
};

export default BattingPage;
