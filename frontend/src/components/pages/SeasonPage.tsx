import React, { useEffect, useState } from 'react';
import GenericTemplate from '../templates/GenericTemplate';
import axios from 'axios';
import _ from 'lodash';
import { TableComponent, HeadCell } from '../common/TableComponent';

const createTeamData = (
  main: string,
  winningRate: number,
  pythagoreanExpectation: number,
  winningRateDifference: number,
  win: number,
  lose: number,
  draw: number
) => {
  return { main, winningRate, pythagoreanExpectation, winningRateDifference, win, lose, draw };
};

interface CentralTeams {
  Giants: { WinningRate: number; Win: number; Lose: number; Draw: number };
  Baystars: { WinningRate: number; Win: number; Lose: number; Draw: number };
  Tigers: { WinningRate: number; Win: number; Lose: number; Draw: number };
  Carp: { WinningRate: number; Win: number; Lose: number; Draw: number };
  Dragons: { WinningRate: number; Win: number; Lose: number; Draw: number };
  Swallows: { WinningRate: number; Win: number; Lose: number; Draw: number };
}

interface PacificTeams {
  Lions: { WinningRate: number; Win: number; Lose: number; Draw: number };
  Hawks: { WinningRate: number; Win: number; Lose: number; Draw: number };
  Eagles: { WinningRate: number; Win: number; Lose: number; Draw: number };
  Marines: { WinningRate: number; Win: number; Lose: number; Draw: number };
  Fighters: { WinningRate: number; Win: number; Lose: number; Draw: number };
  Buffaloes: { WinningRate: number; Win: number; Lose: number; Draw: number };
}

function createTeamDataList(teams: CentralTeams[] | PacificTeams[]) {
  const teamDataList: {
    main: string;
    winningRate: number;
    pythagoreanExpectation: number;
    winningRateDifference: number;
    win: number;
    lose: number;
    draw: number;
  }[] = [];

  teams.forEach((team: any) => {
    _.forEach(team, (val, key) => {
      teamDataList.push(
        createTeamData(
          key,
          val.WinningRate,
          val.PythagoreanExpectation,
          val.WinningRate - val.PythagoreanExpectation,
          val.Win,
          val.Lose,
          val.Draw
        )
      );
    });
  });

  return teamDataList;
}

const headCells: HeadCell[] = [
  { id: 'main', numeric: false, disablePadding: true, label: 'チーム名' },
  { id: 'winningRate', numeric: true, disablePadding: false, label: '勝率' },
  { id: 'pythagoreanExpectation', numeric: true, disablePadding: false, label: 'ピタゴラス勝率' },
  { id: 'winningRateDifference', numeric: true, disablePadding: false, label: '勝率との差分' },
  { id: 'win', numeric: true, disablePadding: false, label: '勝数' },
  { id: 'lose', numeric: true, disablePadding: false, label: '負数' },
  { id: 'draw', numeric: true, disablePadding: false, label: '引き分け' },
];

interface TeamStatsResponse {
  teanStats: any;
}

function SeasonPage(props: { years: string[] }) {
  const [initCentralYear, setCentralYear] = useState<string>('');
  const [centralTeamDatas, setCentralTeamlData] = useState<
    { main: string; winningRate: number; win: number; lose: number; draw: number }[]
  >([]);

  const getTeamCentralDataList = async (year: string) => {
    const result = await axios.get<TeamStatsResponse>(
      `http://localhost:8081/team/stats?from_year=${year}&to_year=${year}`
    );

    const teams: CentralTeams[] = _.map(result.data.teanStats, (teanStats) => {
      const teanStatses = {
        Giants: _.filter(teanStats, { TeamID: '01' })[0],
        Baystars: _.filter(teanStats, { TeamID: '02' })[0],
        Tigers: _.filter(teanStats, { TeamID: '03' })[0],
        Carp: _.filter(teanStats, { TeamID: '04' })[0],
        Dragons: _.filter(teanStats, { TeamID: '05' })[0],
        Swallows: _.filter(teanStats, { TeamID: '06' })[0],
      };
      return teanStatses;
    });

    setCentralYear(year);
    setCentralTeamlData(createTeamDataList(teams));
  };

  const [initPacificYear, setPacificYear] = useState<string>('');
  const [pacificTeamDatas, setPacificTeamlData] = useState<
    { main: string; winningRate: number; win: number; lose: number; draw: number }[]
  >([]);

  const getTeamPacificDataList = async (year: string) => {
    const result = await axios.get<TeamStatsResponse>(
      `http://localhost:8081/team/stats?from_year=${year}&to_year=${year}`
    );

    const teams: PacificTeams[] = _.map(result.data.teanStats, (teanStats) => {
      const teanStatses = {
        Lions: _.filter(teanStats, { TeamID: '07' })[0],
        Hawks: _.filter(teanStats, { TeamID: '08' })[0],
        Eagles: _.filter(teanStats, { TeamID: '09' })[0],
        Marines: _.filter(teanStats, { TeamID: '10' })[0],
        Fighters: _.filter(teanStats, { TeamID: '11' })[0],
        Buffaloes: _.filter(teanStats, { TeamID: '12' })[0],
      };
      return teanStatses;
    });

    setPacificYear(year);
    setPacificTeamlData(createTeamDataList(teams));
  };

  useEffect(() => {
    (async () => {
      getTeamCentralDataList('2020');
      getTeamPacificDataList('2020');
    })();
  }, []);

  return (
    <GenericTemplate title="チーム成績ページ">
      <TableComponent
        title={'シーズン成績(セ)'}
        setSelect={setCentralYear}
        getDataList={getTeamCentralDataList}
        datas={centralTeamDatas}
        selects={props.years}
        headCells={headCells}
        initSorted={'winningRate'}
        initSelect={initCentralYear}
        selectLabel={'年'}
      />
      <TableComponent
        title={'シーズン成績(パ)'}
        setSelect={setPacificYear}
        getDataList={getTeamPacificDataList}
        datas={pacificTeamDatas}
        selects={props.years}
        headCells={headCells}
        initSorted={'winningRate'}
        initSelect={initPacificYear}
        selectLabel={'年'}
      />
    </GenericTemplate>
  );
}

export default SeasonPage;
