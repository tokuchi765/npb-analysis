import React, { useEffect, useState } from 'react';
import Chart from '../common/Chart';
import GenericTemplate from '../templates/GenericTemplate';
import Paper from '@material-ui/core/Paper';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';
import _ from 'lodash';
import Grid from '@material-ui/core/Grid';
import { getTeamBattingByYear } from '../../data/api/teamBatting';

const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
  },
  toolbar: {
    paddingRight: 24, // keep right padding when drawer closed
  },
  toolbarIcon: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'flex-end',
    padding: '0 8px',
    ...theme.mixins.toolbar,
  },
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  },
  appBarShift: {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  menuButton: {
    marginRight: 36,
  },
  menuButtonHidden: {
    display: 'none',
  },
  title: {
    flexGrow: 1,
  },
  drawerPaper: {
    position: 'relative',
    whiteSpace: 'nowrap',
    width: drawerWidth,
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  drawerPaperClose: {
    overflowX: 'hidden',
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    width: theme.spacing(7),
    [theme.breakpoints.up('sm')]: {
      width: theme.spacing(9),
    },
  },
  appBarSpacer: theme.mixins.toolbar,
  content: {
    flexGrow: 1,
    height: '100vh',
    overflow: 'auto',
  },
  container: {
    paddingTop: theme.spacing(4),
    paddingBottom: theme.spacing(4),
  },
  paper: {
    position: 'relative',
    padding: theme.spacing(2),
    display: 'flex',
    overflow: 'auto',
    flexDirection: 'row',
  },
  fixedHeight: {
    height: 400,
  },
}));

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

const centralLineData = [
  ['Giants', '#FF4F02'],
  ['Baystars', '#00FFFF'],
  ['Tigers', '#FFFF00'],
  ['Carp', '#FF0000'],
  ['Dragons', '#005FFF'],
  ['Swallows', '#000055'],
];

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

const pacificLineData = [
  ['Lions', '#BAD3FF'],
  ['Hawks', '#FFD700'],
  ['Eagles', '#FF0461'],
  ['Marines', '#555555'],
  ['Fighters', '#000011'],
  ['Buffaloes', '#4B0082'],
];

function HomePage(props: { years: string[] }) {
  const [centralData, setCentralData] = useState<Array<{ year: string; Giants: number }>>(Array);
  const [pacificData, setPacificData] = useState<Array<{ year: string; Lions: number }>>(Array);
  const classes = useStyles();

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

  const fixedHeightPaper = clsx(classes.paper, classes.fixedHeight);
  return (
    <GenericTemplate title="トップページ">
      <Grid container alignItems="center" justifyContent="center">
        <Paper className={fixedHeightPaper}>
          <Chart
            data={centralData}
            title={'（セ）チーム打率推移'}
            label={'打率'}
            lineData={centralLineData}
          />
          <Chart
            data={pacificData}
            title={'（パ）チーム打率推移'}
            label={'打率'}
            lineData={pacificLineData}
          />
        </Paper>
      </Grid>
    </GenericTemplate>
  );
}

export default HomePage;
