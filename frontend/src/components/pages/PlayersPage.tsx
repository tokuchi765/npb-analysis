import React, { useEffect, useState } from 'react';
import { RouteComponentProps } from 'react-router-dom';
import { useHistory } from 'react-router-dom';
import GenericTemplate from '../templates/GenericTemplate';
import { TableLinkComponent, SelectItem } from '../common/TableComponent';
import _ from 'lodash';
import * as H from 'history';
import { getCareers } from '../../data/api/teamCareers';
import {
  CareerHeadCells,
  createPlayerDatas,
  createPlayerIds,
  PlayerData,
} from '../util/PlayerUtil';

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

const years = ['2020', '2021', '2022'];

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

function PlayersPage(props: PageProps) {
  const [playerDatas, setPlayerDatas] = useState<PlayerData[]>([]);
  const [playerIdMap, setPlayerIds] = useState<Map<string, string>>(new Map<string, string>());
  const [team, setTeam] = useState<string>('');
  const [year, setYear] = useState<string>('');
  const history = useHistory();

  const getPlayerList = async () => {
    const teamID = getTeamId(team);
    const result = await getCareers(teamID, year);
    setPlayerIds(createPlayerIds(result.data.careers));
    setPlayerDatas(createPlayerDatas(result.data.careers));

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
      year = '2022';
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
        datas={playerDatas}
        headCells={CareerHeadCells}
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
