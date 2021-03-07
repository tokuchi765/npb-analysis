import React, { useEffect, useState } from 'react';
import { RouteComponentProps } from 'react-router-dom';
import { useHistory } from 'react-router-dom';
import GenericTemplate from '../templates/GenericTemplate';
import TablePages, { HeadCell } from './TablePages';
import axios from 'axios';
import _ from 'lodash';
import * as H from 'history';

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

function createPlayerIds(
  careers: {
    PlayerID: string;
    Name: string;
  }[]
) {
  const playerIdMap: Map<string, string> = new Map<string, string>();
  careers.forEach((career) => {
    playerIdMap.set(career.Name, career.PlayerID);
  });
  return playerIdMap;
}

function getTeamId(teamName: string) {
  for (const index in teamNameList) {
    if (teamNameList[index] === teamName) {
      return _.padStart(String(Number(index) + 1), 2, '0');
    }
  }
}

interface PageProps extends RouteComponentProps<{ id: string }> {
  history: H.History;
  location: H.Location<any>;
}

const PlayersPage: React.FC<PageProps> = (props) => {
  const [playerDates, setPlayerDates] = useState<PlayerDate[]>([]);
  const [playerIdMap, setPlayerIds] = useState<Map<string, string>>(new Map<string, string>());
  const [initTeam, setInitTeam] = useState<string>('');
  const history = useHistory();

  const getPlayerList = async (teamName: string) => {
    const teamID = getTeamId(teamName);
    const result = await axios.get(`http://localhost:8081/team/careers/${teamID}/2020`);
    setPlayerIds(createPlayerIds(result.data.careers));
    setPlayerDates(createPlayerDates(result.data.careers));
    setInitTeam(teamName);

    history.push({
      state: { teamName: teamName },
    });
  };

  const getTeamName = (location: any) => {
    let teamName = location.state && location.state.teamName;
    if (teamName === undefined) {
      teamName = 'Hawks';
    }
    return teamName;
  };

  useEffect(() => {
    (async () => {
      const teamName = getTeamName(props.location);
      getPlayerList(teamName);
    })();
  }, []);

  return (
    <GenericTemplate title="選手一覧ページ">
      <TablePages
        title={'選手一覧'}
        setSelect={setInitTeam}
        getDataList={getPlayerList}
        datas={playerDates}
        selects={teamNameList}
        headCells={headCells}
        initSorted={'main'}
        initSelect={initTeam}
        selectLabel={'チーム'}
        mainLink={true}
        linkValues={playerIdMap}
        path={'/player/'}
      />
    </GenericTemplate>
  );
};

export default PlayersPage;
