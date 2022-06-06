import React, { useEffect, useState } from 'react';
import Chart, { ChartData } from '../common/Chart';
import GenericTemplate from '../templates/GenericTemplate';
import _ from 'lodash';
import { Grid } from '@mui/material';
import { getTeamBattingByYear } from '../../data/api/teamBatting';
import { BasePaper } from '../common/papers';

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
    battingAverage.push(
      createCentralBattingAverage(
        year,
        teams[i].Giants.BattingAverage,
        teams[i].Baystars.BattingAverage,
        teams[i].Tigers.BattingAverage,
        teams[i].Carp.BattingAverage,
        teams[i].Dragons.BattingAverage,
        teams[i].Swallows.BattingAverage
      )
    );
    i = i + 1;
  });
  return battingAverage;
}

function createCentralBattingAverage(
  year: string,
  Giants: number,
  Baystars: number,
  Tigers: number,
  Carp: number,
  Dragons: number,
  Swallows: number
) {
  return { year, Giants, Baystars, Tigers, Carp, Dragons, Swallows };
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
    battingAverage.push(
      createPacificBattingAverage(
        year,
        teams[i].Lions.BattingAverage,
        teams[i].Hawks.BattingAverage,
        teams[i].Eagles.BattingAverage,
        teams[i].Marines.BattingAverage,
        teams[i].Fighters.BattingAverage,
        teams[i].Buffaloes.BattingAverage
      )
    );
    i = i + 1;
  });
  return battingAverage;
}

function createPacificBattingAverage(
  year: string,
  Lions: number,
  Hawks: number,
  Eagles: number,
  Marines: number,
  Fighters: number,
  Buffaloes: number
) {
  return { year, Lions, Hawks, Eagles, Marines, Fighters, Buffaloes };
}

function HomePage(props: { years: string[] }) {
  const [centralData, setCentralData] = useState<Array<{ year: string; Giants: number }>>(Array);
  const [pacificData, setPacificData] = useState<Array<{ year: string; Lions: number }>>(Array);
  const width = 400;
  const height = 300;

  const centralChartDatas: ChartData[] = [
    { key: 'Giants', name: 'Giants', stroke: '#FF4F02' },
    { key: 'Baystars', name: 'Baystars', stroke: '#00FFFF' },
    { key: 'Tigers', name: 'Tigers', stroke: '#FFFF00' },
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
      const result = await getTeamBattingByYear('2005', '2021');
      const centralTeams = _.map(result.data.teamBatting, (teamBatting) => {
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

      const pacificTeams = _.map(result.data.teamBatting, (teamBatting) => {
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
      </Grid>
    </GenericTemplate>
  );
}

export default HomePage;
