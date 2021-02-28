import React, { useEffect, useState } from 'react';
import { RouteComponentProps } from 'react-router-dom';
import GenericTemplate from '../templates/GenericTemplate';
import TablePages, { HeadCell } from './TablePages';
import axios from 'axios';
import _ from 'lodash';

type PageProps = RouteComponentProps<{ id: string }>;

interface BattingDate {
  main: string;
  team: string;
  atBat: number;
  onBasePercentage: number;
  battingAverage: number;
  sluggingPercentage: number;
  homeRun: number;
  strikeOut: number;
  groundedIntoDoublePlay: number;
}

const batterHeadCells: HeadCell[] = [
  { id: 'main', numeric: false, disablePadding: true, label: '年' },
  { id: 'team', numeric: false, disablePadding: true, label: 'チーム' },
  { id: 'atBat', numeric: true, disablePadding: true, label: '打数' },
  { id: 'onBasePercentage', numeric: true, disablePadding: true, label: '出塁率' },
  { id: 'battingAverage', numeric: true, disablePadding: true, label: '打率' },
  { id: 'sluggingPercentage', numeric: true, disablePadding: true, label: '長打率' },
  { id: 'homeRun', numeric: true, disablePadding: true, label: '本塁打' },
  { id: 'strikeOut', numeric: true, disablePadding: true, label: '三振' },
  { id: 'groundedIntoDoublePlay', numeric: true, disablePadding: true, label: '併殺打' },
];

function createBattingDatas(
  battingList: {
    Year: string;
    Team: string;
    AtBat: string;
    OnBasePercentage: string;
    BattingAverage: string;
    SluggingPercentage: string;
    HomeRun: string;
    StrikeOut: string;
    GroundedIntoDoublePlay: string;
  }[]
) {
  const battings: BattingDate[] = [];
  battingList.forEach((batting) => {
    battings.push({
      main: batting.Year === 'nan' ? '通算' : batting.Year,
      team: batting.Team,
      atBat: Number(batting.AtBat),
      onBasePercentage: Number(batting.OnBasePercentage),
      battingAverage: Number(batting.BattingAverage),
      sluggingPercentage: Number(batting.SluggingPercentage),
      homeRun: Number(batting.HomeRun),
      strikeOut: Number(batting.StrikeOut),
      groundedIntoDoublePlay: Number(batting.GroundedIntoDoublePlay),
    });
  });
  return battings;
}

interface PitchingDate {
  main: string;
  team: string;
  piched: number;
  inningsPitched: number;
  earnedRunAverage: number;
  batter: number;
  strikeOut: number;
  homeRun: number;
  baseOnBalls: number;
  hitByPitches: number;
}

const pitcherHeadCells: HeadCell[] = [
  { id: 'main', numeric: false, disablePadding: true, label: '年' },
  { id: 'team', numeric: false, disablePadding: true, label: 'チーム' },
  { id: 'piched', numeric: true, disablePadding: true, label: '登板' },
  { id: 'inningsPitched', numeric: true, disablePadding: true, label: '投球回' },
  { id: 'earnedRunAverage', numeric: true, disablePadding: true, label: '防御率' },
  { id: 'batter', numeric: true, disablePadding: true, label: '打者' },
  { id: 'strikeOut', numeric: true, disablePadding: true, label: '三振' },
  { id: 'homeRun', numeric: true, disablePadding: true, label: '被本塁打' },
  { id: 'baseOnBalls', numeric: true, disablePadding: true, label: '四球' },
  { id: 'hitByPitches', numeric: true, disablePadding: true, label: '死球' },
];

function createPitchingDatas(
  pitchingList: {
    Year: string;
    Team: string;
    Piched: string;
    InningsPitched: string;
    EarnedRunAverage: string;
    Batter: string;
    StrikeOut: string;
    HomeRun: string;
    BaseOnBalls: string;
    HitByPitches: string;
  }[]
) {
  const pitchings: PitchingDate[] = [];
  pitchingList.forEach((pitching) => {
    pitchings.push({
      main: pitching.Year === 'nan' ? '通算' : pitching.Year,
      team: pitching.Team,
      piched: Number(pitching.Piched),
      inningsPitched: Number(pitching.InningsPitched),
      earnedRunAverage: Number(pitching.EarnedRunAverage),
      batter: Number(pitching.Batter),
      strikeOut: Number(pitching.StrikeOut),
      homeRun: Number(pitching.HomeRun),
      baseOnBalls: Number(pitching.BaseOnBalls),
      hitByPitches: Number(pitching.HitByPitches),
    });
  });
  return pitchings;
}

const PlayerPage: React.FC<PageProps> = (props) => {
  const [playerName, setPlayerName] = useState<string>('');
  const [battingDates, setBattingDates] = useState<BattingDate[]>([]);
  const [pitchingDates, setPitchingDates] = useState<PitchingDate[]>([]);

  const getPlayerDatas = async () => {
    const playerID = props.match.params.id;
    const result = await axios.get(`http://localhost:8081/player/${playerID}`);
    const { career, batting, pitching } = result.data;
    setPlayerName(career.Name);
    setBattingDates(_.isEmpty(batting) ? [] : createBattingDatas(batting));
    setPitchingDates(_.isEmpty(pitching) ? [] : createPitchingDatas(pitching));
  };

  useEffect(() => {
    (async () => {
      getPlayerDatas();
    })();
  }, []);

  return (
    <GenericTemplate title={playerName}>
      <TablePages
        title={'打撃成績'}
        getDataList={getPlayerDatas}
        datas={battingDates}
        selects={[]}
        headCells={batterHeadCells}
        initSorted={'main'}
        initSelect={''}
        selectLabel={''}
        mainLink={false}
        linkValues={new Map<string, string>()}
        path={''}
      />
      <TablePages
        title={'投手成績'}
        getDataList={getPlayerDatas}
        datas={pitchingDates}
        selects={[]}
        headCells={pitcherHeadCells}
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
