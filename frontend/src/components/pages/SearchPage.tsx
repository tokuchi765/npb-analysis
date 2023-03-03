import React, { useEffect, useState } from 'react';
import { IconButton, InputBase, Paper, TableContainer } from '@mui/material';
import _ from 'lodash';
import { HeadCell, TableSearchComponent } from '../common/TableComponent';
import GenericTemplate from '../templates/GenericTemplate';
import { searchPlayer } from '../../data/api/player';
import SearchIcon from '@mui/icons-material/Search';
import Title from '../common/Title';
import { RouteComponentProps, useHistory } from 'react-router-dom';
import * as H from 'history';

const headCells: HeadCell[] = [
  { id: 'main', numeric: false, disablePadding: true, label: '選手名' },
  { id: 'position', numeric: false, disablePadding: true, label: 'ポジション' },
  { id: 'height', numeric: false, disablePadding: true, label: '身長' },
  { id: 'draft', numeric: false, disablePadding: true, label: 'ドラフト' },
  { id: 'career', numeric: false, disablePadding: true, label: '経歴' },
];

interface PlayerDate {
  main: string;
  position: string;
  height: string;
  draft: string;
  career: string;
}

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

interface Search {
  name: string;
}

export interface SearchCondition extends RouteComponentProps<{ id: string }> {
  history: H.History<Search>;
  location: H.Location<Search>;
}

function SearchPage(props: SearchCondition) {
  const [name, setName] = useState('');
  const [playerDates, setPlayerDates] = useState<PlayerDate[]>([]);
  const [playerIdMap, setPlayerIds] = useState<Map<string, string>>(new Map<string, string>());
  const [noSearchResults, setNoSearchResults] = useState(false);
  const history = useHistory<Search>();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const isNotEmptyName = name !== '';

    if (isNotEmptyName) {
      const result = await searchPlayer(name);
      if (_.isEmpty(result.careers)) {
        setNoSearchResults(true);
        setPlayerIds(new Map<string, string>());
        setPlayerDates([]);
      } else {
        setNoSearchResults(false);
        setPlayerIds(createPlayerIds(result.careers));
        setPlayerDates(createPlayerDates(result.careers));
        history.push({ state: { name: name } });
      }
    }
  };

  useEffect(() => {
    (async () => {
      if (props.location.state) {
        setName(props.location.state.name);
        const result = await searchPlayer(props.location.state.name);
        setNoSearchResults(false);
        setPlayerIds(createPlayerIds(result.careers));
        setPlayerDates(createPlayerDates(result.careers));
        history.push({ state: { name: name } });
      }
    })();
  }, []);

  return (
    <GenericTemplate title="選手検索ページ">
      <TableContainer component={Paper}>
        <Paper
          component="form"
          onSubmit={handleSubmit}
          sx={{ p: '2px 4px', display: 'flex', alignItems: 'center', width: 400 }}
        >
          <InputBase
            sx={{ ml: 1, flex: 1 }}
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="選手名を入力"
            inputProps={{ 'aria-label': 'search google maps' }}
          />
          <IconButton type="submit" sx={{ p: '10px' }} aria-label="search">
            <SearchIcon />
          </IconButton>
        </Paper>
        {noSearchResults && <Title>{'※検索結果がありません'}</Title>}
        <TableSearchComponent
          datas={playerDates}
          headCells={headCells}
          initSorted={'main'}
          linkValues={playerIdMap}
        />
      </TableContainer>
    </GenericTemplate>
  );
}

export default SearchPage;
