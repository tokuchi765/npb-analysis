import React, { useEffect, useState } from 'react';
import Chart, { ChartData } from '../common/Chart';
import GenericTemplate from '../templates/GenericTemplate';
import _ from 'lodash';
import { Grid } from '@mui/material';
import { getTeamBattingByYear } from '../../data/api/teamBatting';
import { BasePaper } from '../common/papers';
import { getTeamPitchingByYear } from '../../data/api/teamPitching';

function createCentralBattingAverages(
  teams: {
    Giants: any;
    Baystars: any;
    Tigers: any;
    Carp: any;
    Dragons: any;
    Swallows: any;
  }[],
  years: string[]
) {
  const battingAverage: {
    year: string;
    Giants: number;
    Baystars: number;
    Tigers: number;
    Carp: number;
    Dragons: number;
    Swallows: number;
  }[] = [];
  let i = 0;
  years.forEach((year) => {
    battingAverage.push({
      year: year,
      Giants: teams[i].Giants.BattingAverage,
      Baystars: teams[i].Baystars.BattingAverage,
      Tigers: teams[i].Tigers.BattingAverage,
      Carp: teams[i].Carp.BattingAverage,
      Dragons: teams[i].Dragons.BattingAverage,
      Swallows: teams[i].Swallows.BattingAverage,
    });
    i = i + 1;
  });
  return battingAverage;
}

function createPacificBattingAverages(
  teams: {
    Lions: any;
    Hawks: any;
    Eagles: any;
    Marines: any;
    Fighters: any;
    Buffaloes: any;
  }[],
  years: string[]
) {
  const battingAverage: {
    year: string;
    Lions: number;
    Hawks: number;
    Eagles: number;
    Marines: number;
    Fighters: number;
    Buffaloes: number;
  }[] = [];
  let i = 0;
  years.forEach((year) => {
    battingAverage.push({
      year: year,
      Lions: teams[i].Lions.BattingAverage,
      Hawks: teams[i].Hawks.BattingAverage,
      Eagles: teams[i].Eagles.BattingAverage,
      Marines: teams[i].Marines.BattingAverage,
      Fighters: teams[i].Fighters.BattingAverage,
      Buffaloes: teams[i].Buffaloes.BattingAverage,
    });
    i = i + 1;
  });
  return battingAverage;
}

function createCentralEarnedRunAverages(
  teams: {
    Giants: any;
    Baystars: any;
    Tigers: any;
    Carp: any;
    Dragons: any;
    Swallows: any;
  }[],
  years: string[]
) {
  const earnedRunAverages: {
    year: string;
    Giants: number;
    Baystars: number;
    Tigers: number;
    Carp: number;
    Dragons: number;
    Swallows: number;
  }[] = [];
  let i = 0;
  years.forEach((year) => {
    earnedRunAverages.push({
      year: year,
      Giants: teams[i].Giants.EarnedRunAverage,
      Baystars: teams[i].Baystars.EarnedRunAverage,
      Tigers: teams[i].Tigers.EarnedRunAverage,
      Carp: teams[i].Carp.EarnedRunAverage,
      Dragons: teams[i].Dragons.EarnedRunAverage,
      Swallows: teams[i].Swallows.EarnedRunAverage,
    });
    i = i + 1;
  });
  return earnedRunAverages;
}

function createPacificEarnedRunAverages(
  teams: {
    Lions: any;
    Hawks: any;
    Eagles: any;
    Marines: any;
    Fighters: any;
    Buffaloes: any;
  }[],
  years: string[]
) {
  const earnedRunAverages: {
    year: string;
    Lions: number;
    Hawks: number;
    Eagles: number;
    Marines: number;
    Fighters: number;
    Buffaloes: number;
  }[] = [];
  let i = 0;
  years.forEach((year) => {
    earnedRunAverages.push({
      year: year,
      Lions: teams[i].Lions.EarnedRunAverage,
      Hawks: teams[i].Hawks.EarnedRunAverage,
      Eagles: teams[i].Eagles.EarnedRunAverage,
      Marines: teams[i].Marines.EarnedRunAverage,
      Fighters: teams[i].Fighters.EarnedRunAverage,
      Buffaloes: teams[i].Buffaloes.EarnedRunAverage,
    });
    i = i + 1;
  });
  return earnedRunAverages;
}

function HomePage(props: { years: string[] }) {
  const [centralData, setCentralData] = useState<Array<{ year: string; Giants: number }>>(Array);
  const [pacificData, setPacificData] = useState<Array<{ year: string; Lions: number }>>(Array);
  const [centralPitchingData, setCentralPitchingData] =
    useState<Array<{ year: string; Giants: number }>>(Array);
  const [pacificPitchingData, setPacificPitchingData] =
    useState<Array<{ year: string; Lions: number }>>(Array);
  const width = 400;
  const height = 300;

  const centralChartDatas: ChartData[] = [
    { key: 'Giants', name: 'Giants', stroke: '#FF4F02' },
    { key: 'Baystars', name: 'Baystars', stroke: '#00FFFF' },
    { key: 'Tigers', name: 'Tigers', stroke: '#ffbf00' },
    { key: 'Carp', name: 'Carp', stroke: '#FF0000' },
    { key: 'Dragons', name: 'Dragons', stroke: '#005FFF' },
    { key: 'Swallows', name: 'Swallows', stroke: '#000055' },
  ];

  const pacificChartDatas: ChartData[] = [
    { key: 'Lions', name: 'Lions', stroke: '#BAD3FF' },
    { key: 'Hawks', name: 'Hawks', stroke: '#FFD700' },
    { key: 'Eagles', name: 'Eagles', stroke: '#FF0461' },
    { key: 'Marines', name: 'Marines', stroke: '#555555' },
    { key: 'Fighters', name: 'Fighters', stroke: '#000011' },
    { key: 'Buffaloes', name: 'Buffaloes', stroke: '#4B0082' },
  ];

  useEffect(() => {
    (async () => {
      const battingResult = await getTeamBattingByYear('2005', '2023');
      const centralTeams = _.map(battingResult.data.teamBatting, (teamBatting) => {
        const teamBattings = {
          Giants: _.filter(teamBatting, { TeamID: '01' })[0],
          Baystars: _.filter(teamBatting, { TeamID: '02' })[0],
          Tigers: _.filter(teamBatting, { TeamID: '03' })[0],
          Carp: _.filter(teamBatting, { TeamID: '04' })[0],
          Dragons: _.filter(teamBatting, { TeamID: '05' })[0],
          Swallows: _.filter(teamBatting, { TeamID: '06' })[0],
        };
        return teamBattings;
      });

      setCentralData(createCentralBattingAverages(centralTeams, props.years));

      const pacificTeams = _.map(battingResult.data.teamBatting, (teamBatting) => {
        const teamBattings = {
          Lions: _.filter(teamBatting, { TeamID: '07' })[0],
          Hawks: _.filter(teamBatting, { TeamID: '08' })[0],
          Eagles: _.filter(teamBatting, { TeamID: '09' })[0],
          Marines: _.filter(teamBatting, { TeamID: '10' })[0],
          Fighters: _.filter(teamBatting, { TeamID: '11' })[0],
          Buffaloes: _.filter(teamBatting, { TeamID: '12' })[0],
        };
        return teamBattings;
      });

      setPacificData(createPacificBattingAverages(pacificTeams, props.years));

      const pitchingResult = await getTeamPitchingByYear('2005', '2023');
      const centralPitchings = _.map(pitchingResult.data.teamPitching, (teamPitching) => {
        const teamPitchings = {
          Giants: _.filter(teamPitching, { TeamID: '01' })[0],
          Baystars: _.filter(teamPitching, { TeamID: '02' })[0],
          Tigers: _.filter(teamPitching, { TeamID: '03' })[0],
          Carp: _.filter(teamPitching, { TeamID: '04' })[0],
          Dragons: _.filter(teamPitching, { TeamID: '05' })[0],
          Swallows: _.filter(teamPitching, { TeamID: '06' })[0],
        };
        return teamPitchings;
      });

      setCentralPitchingData(createCentralEarnedRunAverages(centralPitchings, props.years));

      const pacificPitchings = _.map(pitchingResult.data.teamPitching, (teamPitching) => {
        const teamPitchings = {
          Lions: _.filter(teamPitching, { TeamID: '07' })[0],
          Hawks: _.filter(teamPitching, { TeamID: '08' })[0],
          Eagles: _.filter(teamPitching, { TeamID: '09' })[0],
          Marines: _.filter(teamPitching, { TeamID: '10' })[0],
          Fighters: _.filter(teamPitching, { TeamID: '11' })[0],
          Buffaloes: _.filter(teamPitching, { TeamID: '12' })[0],
        };
        return teamPitchings;
      });

      setPacificPitchingData(createPacificEarnedRunAverages(pacificPitchings, props.years));
    })();
  }, []);

  return (
    <GenericTemplate title="トップページ">
      <Grid container alignItems="center" justifyContent="center">
        <BasePaper>
          <Chart
            data={centralData}
            title={'（セ）チーム打率推移'}
            label={'打率'}
            chartDatas={centralChartDatas}
            width={width}
            height={height}
          />
          <Chart
            data={pacificData}
            title={'（パ）チーム打率推移'}
            label={'打率'}
            chartDatas={pacificChartDatas}
            width={width}
            height={height}
          />
        </BasePaper>
        <BasePaper>
          <Chart
            data={centralPitchingData}
            title={'（セ）チーム防御率推移'}
            label={'防御率'}
            chartDatas={centralChartDatas}
            width={width}
            height={height}
          />
          <Chart
            data={pacificPitchingData}
            title={'（パ）チーム打率推移'}
            label={'防御率'}
            chartDatas={pacificChartDatas}
            width={width}
            height={height}
          />
        </BasePaper>
      </Grid>
    </GenericTemplate>
  );
}

export default HomePage;
