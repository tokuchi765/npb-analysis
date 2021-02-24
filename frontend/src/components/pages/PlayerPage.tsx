import React, { useEffect, useState } from 'react';
import { RouteComponentProps } from 'react-router-dom';
import GenericTemplate from '../templates/GenericTemplate';
import TablePages, { HeadCell } from './TablePages';
import axios from 'axios';

type PageProps = RouteComponentProps<{ id: string }>;

interface PlayerData {
  main: string;
}

const headCells: HeadCell[] = [
  { id: 'main', numeric: false, disablePadding: true, label: '選手名' },
];

const PlayerPage: React.FC<PageProps> = (props) => {
  const [playerDates, setPlayerDates] = useState<PlayerData[]>([]);

  const getPlayerDatas = async () => {
    const playerID = props.match.params.id;
    const result = await axios.get(`http://localhost:8081/team/player/${playerID}`);
  };

  useEffect(() => {
    (async () => {
      getPlayerDatas();
    })();
  }, []);

  return (
    <GenericTemplate title="打撃成績">
      <TablePages
        title={'打撃成績'}
        getDataList={getPlayerDatas}
        datas={playerDates}
        selects={[]}
        headCells={headCells}
        initSorted={'main'}
        initSelect={''}
        selectLabel={''}
        mainLink={false}
        linkValues={new Map<string, string>()}
        path={''}
      />
    </GenericTemplate>
  );
};

export default PlayerPage;
