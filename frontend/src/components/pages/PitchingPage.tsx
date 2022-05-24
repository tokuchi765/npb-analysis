import React, { useEffect, useState } from 'react';
import GenericTemplate from '../templates/GenericTemplate';
import { TableComponent, HeadCell, SelectItem } from '../common/TableComponent';
import _ from 'lodash';
import { getTeamPitchingByYear } from '../../data/api/teamPitching';

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
  strikeOutRate: number;
  babip: number;
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
  { id: 'strikeOutRate', numeric: true, disablePadding: false, label: '奪三振率' },
  { id: 'babip', numeric: true, disablePadding: false, label: '被BABIP' },
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
  strikeOut: number,
  strikeOutRate: number,
  babip: number
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
    strikeOutRate,
    babip,
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
          val.StrikeOut,
          val.StrikeOutRate,
          val.BABIP
        )
      );
    });
  });
  return datas;
}

function PitchingPage(props: { years: string[]; initYear: string }) {
  const [centralYear, setCentralYear] = useState<string>('');
  const [centralDatas, setCentralData] = useState<PitchingData[]>([]);

  const getCentralPitchingDataList = async () => {
    const result = await getTeamPitchingByYear(centralYear, centralYear);

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

  useEffect(() => {
    (async () => {
      if (_.isEmpty(centralYear)) {
        setCentralYear(props.initYear);
      } else {
        getCentralPitchingDataList();
      }
    })();
  }, [centralYear]);

  const [pacificYear, setPacificYear] = useState<string>('');
  const [pacificDatas, setPacificData] = useState<PitchingData[]>([]);

  const getPacificPitchingDataList = async () => {
    const result = await getTeamPitchingByYear(pacificYear, pacificYear);

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
      if (_.isEmpty(pacificYear)) {
        setPacificYear(props.initYear);
      } else {
        getPacificPitchingDataList();
      }
    })();
  }, [pacificYear]);

  return (
    <GenericTemplate title="チーム投手成績ページ">
      <TableComponent
        title={'シーズン投手成績(セ)'}
        datas={centralDatas}
        headCells={headCells}
        initSorted={'earnedRunAverage'}
        selectItems={[new SelectItem(centralYear, '年', props.years, setCentralYear)]}
      />
      <TableComponent
        title={'シーズン投手成績(パ)'}
        datas={pacificDatas}
        headCells={headCells}
        initSorted={'earnedRunAverage'}
        selectItems={[new SelectItem(pacificYear, '年', props.years, setPacificYear)]}
      />
    </GenericTemplate>
  );
}

export default PitchingPage;
