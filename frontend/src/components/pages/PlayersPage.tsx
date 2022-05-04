import React, { useEffect, useState, useCallback } from 'react';
import { RouteComponentProps } from 'react-router-dom';
import { useHistory } from 'react-router-dom';
import GenericTemplate from '../templates/GenericTemplate';
import { TableLinkComponent, SelectItem, HeadCell } from '../common/TableComponent';
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

const years = ['2020', '2021'];

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

export function getTeamId(teamName: string) {
  for (const index in teamNameList) {
    if (teamNameList[index] === teamName) {
      return _.padStart(String(Number(index) + 1), 2, '0');
    }
  }
}

export interface PageProps extends RouteComponentProps<{ id: string }> {
  history: H.History;
  location: H.Location<any>;
}

interface TeamCareersResponse {
  careers: any;
}

function PlayersPage(props: PageProps) {
  const [playerDates, setPlayerDates] = useState<PlayerDate[]>([]);
  const [playerIdMap, setPlayerIds] = useState<Map<string, string>>(new Map<string, string>());
  const [team, setTeam] = useState<string>('');
  const [year, setYear] = useState<string>('');
  const history = useHistory();

  const getPlayerList = async () => {
    const teamID = getTeamId(team);
    const result = await axios.get<TeamCareersResponse>(
      `http://localhost:8081/team/careers/${teamID}/${year}`
    );
    setPlayerIds(createPlayerIds(result.data.careers));
    setPlayerDates(createPlayerDates(result.data.careers));

    history.push({
      state: { teamName: team, year: year },
    });
  };

  const getTeamName = (location: any) => {
    let teamName = location.state && location.state.teamName;
    if (teamName === undefined) {
      teamName = 'Hawks';
    }
    return teamName;
  };

  const getYear = (location: any) => {
    let year = location.state && location.state.year;
    if (year === undefined) {
      year = '2021';
    }
    return year;
  };

  useEffect(() => {
    (async () => {
      if (_.isEmpty(team)) {
        setTeam(getTeamName(props.location));
      } else if (_.isEmpty(year)) {
        setYear(getYear(props.location));
      } else {
        getPlayerList();
      }
    })();
  }, [team, year]);

  return (
    <GenericTemplate title="選手一覧ページ">
      <TableLinkComponent
        title={'選手一覧'}
        datas={playerDates}
        headCells={headCells}
        initSorted={'main'}
        selectItems={[
          new SelectItem(team, 'チーム', teamNameList, setTeam),
          new SelectItem(year, '年', years, setYear),
        ]}
        linkValues={playerIdMap}
        path={'/player/'}
      />
    </GenericTemplate>
  );
}

export default PlayersPage;
