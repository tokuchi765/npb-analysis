import React, { useEffect, useState } from 'react';
import GenericTemplate from '../templates/GenericTemplate';
import TablePages, { HeadCell } from './TablePages';
import axios from 'axios';
import _ from 'lodash';

const teamNameList = [
  'Giants',
  'Baystars',
  'Tigers',
  'Carp',
  'Dragons',
  'Swallows',
  'Lions',
  'Hawks',
  'Eagles',
  'Marines',
  'Fighters',
  'Buffaloes',
];

interface PlayerDate {
  main: string;
  position: string;
  height: string;
  draft: string;
  career: string;
}
const headCells: HeadCell[] = [
  { id: 'main', numeric: false, disablePadding: true, label: '選手名' },
  { id: 'position', numeric: false, disablePadding: true, label: 'ポジション' },
  { id: 'height', numeric: false, disablePadding: true, label: '身長' },
  { id: 'draft', numeric: false, disablePadding: true, label: 'ドラフト' },
  { id: 'career', numeric: false, disablePadding: true, label: '経歴' },
];

function createPlayerDate(
  main: string,
  position: string,
  height: string,
  draft: string,
  career: string
) {
  const result: PlayerDate = { main, position, height, draft, career };
  return result;
}

function createPlayerDates(
  careers: {
    Name: string;
    Position: string;
    Height: string;
    Draft: string;
    Career: string;
  }[]
) {
  const playerDateList: PlayerDate[] = [];
  careers.forEach((career) => {
    playerDateList.push(
      createPlayerDate(career.Name, career.Position, career.Height, career.Draft, career.Career)
    );
  });
  return playerDateList;
}

function getTeamId(teamName: string) {
  for (const index in teamNameList) {
    if (teamNameList[index] === teamName) {
      return _.padStart(String(Number(index) + 1), 2, '0');
    }
  }
}

const PlayersPage: React.FC = () => {
  const [playerDates, setPlayerDates] = useState<PlayerDate[]>([]);

  const getPlayerList = async (teamName: string) => {
    const teamID = getTeamId(teamName);
    const result = await axios.get(`http://localhost:8081/team/careers/${teamID}/2020`);
    setPlayerDates(createPlayerDates(result.data.careers));
  };

  useEffect(() => {
    (async () => {
      getPlayerList('Hawks');
    })();
  }, []);

  return (
    <GenericTemplate title="選手一覧ページ">
      <TablePages
        title={'選手一覧'}
        getTeamDataList={getPlayerList}
        teamDatas={playerDates}
        selects={teamNameList}
        headCells={headCells}
        initSorted={'main'}
        initSelect={'Hawks'}
        selectLabel={'チーム'}
      />
    </GenericTemplate>
  );
};

export default PlayersPage;
