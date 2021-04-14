import React, { useEffect, useState } from 'react';
import GenericTemplate from '../templates/GenericTemplate';
import { TableComponent, HeadCell } from '../common/TableComponent';
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

interface PitchingData {
  main: string;
  earnedRunAverage: number;
  games: number;
  runsAllowed: number;
  save: number;
  hold: number;
  homeRun: number;
  baseOnBalls: number;
  strikeOut: number;
}

const headCells: HeadCell[] = [
  { id: 'main', numeric: false, disablePadding: true, label: 'チーム名' },
  { id: 'earnedRunAverage', numeric: true, disablePadding: false, label: '防御率' },
  { id: 'games', numeric: true, disablePadding: false, label: '試合' },
  { id: 'runsAllowed', numeric: true, disablePadding: false, label: '失点' },
  { id: 'save', numeric: true, disablePadding: false, label: 'セーブ' },
  { id: 'hold', numeric: true, disablePadding: false, label: 'ホールド' },
  { id: 'homeRun', numeric: true, disablePadding: false, label: '被本塁打' },
  { id: 'baseOnBalls', numeric: true, disablePadding: false, label: '与四球' },
  { id: 'strikeOut', numeric: true, disablePadding: false, label: '三振' },
];

function createPitchingData(
  main: string,
  earnedRunAverage: number,
  games: number,
  runsAllowed: number,
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
    runsAllowed,
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
          val.RunsAllowed,
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
  const [initCentralYear, setCentralYear] = useState<string>('');
  const [centralDatas, setCentralData] = useState<PitchingData[]>([]);

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

    setCentralYear(year);

    setCentralData(createPitchingDataList(centralPitchings));
  };

  const [initPacificYear, setPacificYear] = useState<string>('');
  const [pacificDatas, setPacificData] = useState<PitchingData[]>([]);

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

    setPacificYear(year);

    setPacificData(createPitchingDataList(pacificPitchings));
  };

  useEffect(() => {
    (async () => {
      getCentralPitchingDataList('2020');
      getPacificPitchingDataList('2020');
    })();
  }, []);

  return (
    <GenericTemplate title="チーム投手成績ページ">
      <TableComponent
        title={'シーズン投手成績(セ)'}
        setSelect={setCentralYear}
        getDataList={getCentralPitchingDataList}
        datas={centralDatas}
        selects={years}
        headCells={headCells}
        initSorted={'earnedRunAverage'}
        initSelect={initCentralYear}
        selectLabel={'年'}
      />
      <TableComponent
        title={'シーズン投手成績(パ)'}
        setSelect={setPacificYear}
        getDataList={getPacificPitchingDataList}
        datas={pacificDatas}
        selects={years}
        headCells={headCells}
        initSorted={'earnedRunAverage'}
        initSelect={initPacificYear}
        selectLabel={'年'}
      />
    </GenericTemplate>
  );
};

export default PitchingPage;
