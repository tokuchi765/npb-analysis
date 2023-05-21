import React, { useEffect, useState } from 'react';
import { IconButton, InputBase, Paper, TableContainer } from '@mui/material';
import _ from 'lodash';
import { TableSearchComponent } from '../common/TableComponent';
import GenericTemplate from '../templates/GenericTemplate';
import { searchPlayer } from '../../data/api/player';
import SearchIcon from '@mui/icons-material/Search';
import Title from '../common/Title';
import { RouteComponentProps, useHistory } from 'react-router-dom';
import * as H from 'history';
import {
  CareerHeadCells,
  createPlayerDatas,
  createPlayerIds,
  PlayerData,
} from '../util/PlayerUtil';

interface Search {
  name: string;
}

export interface SearchCondition extends RouteComponentProps<{ id: string }> {
  history: H.History<Search>;
  location: H.Location<Search>;
}

function SearchPage(props: SearchCondition) {
  const [name, setName] = useState('');
  const [playerDatas, setPlayerDatas] = useState<PlayerData[]>([]);
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
        setPlayerDatas([]);
      } else {
        setNoSearchResults(false);
        setPlayerIds(createPlayerIds(result.careers));
        setPlayerDatas(createPlayerDatas(result.careers));
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
        setPlayerDatas(createPlayerDatas(result.careers));
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
          datas={playerDatas}
          headCells={CareerHeadCells}
          initSorted={'main'}
          linkValues={playerIdMap}
        />
      </TableContainer>
    </GenericTemplate>
  );
}

export default SearchPage;
