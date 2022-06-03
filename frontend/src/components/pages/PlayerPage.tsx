import React, { useEffect, useState } from 'react';
import { RouteComponentProps } from 'react-router-dom';
import GenericTemplate from '../templates/GenericTemplate';
import { Button } from '@mui/material';
import { TableComponent, HeadCell } from '../common/TableComponent';
import _ from 'lodash';
import { getPlayer } from '../../data/api/player';

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
  woba: number;
  rc: number;
  babip: number;
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
  { id: 'woba', numeric: true, disablePadding: true, label: '加重出塁率' },
  { id: 'rc', numeric: true, disablePadding: true, label: '創出得点' },
  { id: 'babip', numeric: true, disablePadding: true, label: 'BABIP' },
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
    Woba: string;
    RC: string;
    BABIP: string;
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
      woba: Number(batting.Woba),
      rc: Number(batting.RC),
      babip: Number(batting.BABIP),
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
  strikeOutRate: number;
  homeRun: number;
  baseOnBalls: number;
  hitByPitches: number;
  babip: number;
}

const pitcherHeadCells: HeadCell[] = [
  { id: 'main', numeric: false, disablePadding: true, label: '年' },
  { id: 'team', numeric: false, disablePadding: true, label: 'チーム' },
  { id: 'piched', numeric: true, disablePadding: true, label: '登板' },
  { id: 'inningsPitched', numeric: true, disablePadding: true, label: '投球回' },
  { id: 'earnedRunAverage', numeric: true, disablePadding: true, label: '防御率' },
  { id: 'batter', numeric: true, disablePadding: true, label: '打者' },
  { id: 'strikeOut', numeric: true, disablePadding: true, label: '三振' },
  { id: 'strikeOutRate', numeric: true, disablePadding: false, label: '奪三振率' },
  { id: 'homeRun', numeric: true, disablePadding: true, label: '被本塁打' },
  { id: 'baseOnBalls', numeric: true, disablePadding: true, label: '四球' },
  { id: 'hitByPitches', numeric: true, disablePadding: true, label: '死球' },
  { id: 'babip', numeric: true, disablePadding: true, label: '被BABIP' },
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
    StrikeOutRate: string;
    HomeRun: string;
    BaseOnBalls: string;
    HitByPitches: string;
    BABIP: string;
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
      strikeOutRate: Number(pitching.StrikeOutRate),
      homeRun: Number(pitching.HomeRun),
      baseOnBalls: Number(pitching.BaseOnBalls),
      hitByPitches: Number(pitching.HitByPitches),
      babip: Number(pitching.BABIP),
    });
  });
  return pitchings;
}

function PlayerPage(props: PageProps) {
  const [playerName, setPlayerName] = useState<string>('');
  const [battingDates, setBattingDates] = useState<BattingDate[]>([]);
  const [pitchingDates, setPitchingDates] = useState<PitchingDate[]>([]);

  const getPlayerDatas = async () => {
    const playerID = props.match.params.id;
    const { career, batting, pitching } = await getPlayer(playerID);
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
      <Button onClick={() => props.history.goBack()} variant="contained" color="primary">
        戻る
      </Button>
      <TableComponent
        title={'打撃成績'}
        datas={battingDates}
        headCells={batterHeadCells}
        initSorted={'main'}
        selectItems={[]}
      />
      <TableComponent
        title={'投手成績'}
        datas={pitchingDates}
        headCells={pitcherHeadCells}
        initSorted={'main'}
        selectItems={[]}
      />
    </GenericTemplate>
  );
}

export default PlayerPage;
