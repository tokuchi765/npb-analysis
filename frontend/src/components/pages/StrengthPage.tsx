import React, { useEffect, useState } from 'react';
import RadarChartComponent from '../common/RadarChartComponent';
import GenericTemplate from '../templates/GenericTemplate';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import clsx from 'clsx';
import Title from '../common/Title';
import { Box } from '@mui/material';
import { Selectable } from '../common/TableComponent';
import { getTeamId } from '../pages/PlayersPage';
import axios from 'axios';

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

const useStyles = makeStyles((theme) => ({
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
  formControl: {
    margin: theme.spacing(1),
    minWidth: 120,
  },
}));

const fullMark = 100;

interface TeamPitchingResponse {
  teamPitching: {
    StrikeOutRate: number;
    RunsAllowed: number;
  };
}

interface TeamBattingResponse {
  teamBatting: {
    HomeRun: number;
    SluggingPercentage: number;
    OnBasePercentage: number;
  };
}

interface MaxTeamBattingResponse {
  maxHomeRun: number;
  maxSluggingPercentage: number;
  maxOnBasePercentage: number;
}

interface MaxTeamPitchingResponse {
  maxStrikeOutRate: number;
  maxRunsAllowed: number;
}

interface MinTeamBattingResponse {
  minHomeRun: number;
  minSluggingPercentage: number;
  minOnBasePercentage: number;
}

interface MinTeamPitchingResponse {
  minStrikeOutRate: number;
  minRunsAllowed: number;
}

export function StrengthPage(props: { years: string[]; initYear: string }) {
  const classes = useStyles();
  const [yearA, setYearA] = useState<string>(props.initYear);
  const [teamA, setTeamA] = useState<string>('Hawks');
  const [yearB, setYearB] = useState<string>(props.initYear);
  const [teamB, setTeamB] = useState<string>('Buffaloes');
  const [data, setData] = useState<any>();
  const fixedHeightPaper = clsx(classes.paper, classes.fixedHeight);

  const setStrength = async () => {
    const {
      data: { maxStrikeOutRate, maxRunsAllowed },
    } = await axios.get<MaxTeamPitchingResponse>(`http://localhost:8081/team/pitching/max`);
    const {
      data: { minStrikeOutRate, minRunsAllowed },
    } = await axios.get<MinTeamPitchingResponse>(`http://localhost:8081/team/pitching/min`);
    const {
      data: { maxHomeRun, maxSluggingPercentage, maxOnBasePercentage },
    } = await axios.get<MaxTeamBattingResponse>(`http://localhost:8081/team/batting/max`);
    const {
      data: { minHomeRun, minSluggingPercentage, minOnBasePercentage },
    } = await axios.get<MinTeamBattingResponse>(`http://localhost:8081/team/batting/min`);

    const teamIdA = getTeamId(teamA);
    const {
      data: {
        teamPitching: { StrikeOutRate: strikeOutRateA, RunsAllowed: runsAllowedA },
      },
    } = await axios.get<TeamPitchingResponse>(
      `http://localhost:8081/team/pitching/${teamIdA}/${yearA}`
    );
    const {
      data: {
        teamBatting: {
          HomeRun: homeRunA,
          SluggingPercentage: sluggingPercentageA,
          OnBasePercentage: onBasePercentageA,
        },
      },
    } = await axios.get<TeamBattingResponse>(
      `http://localhost:8081/team/batting/${teamIdA}/${yearA}`
    );
    const teamIdB = getTeamId(teamB);
    const {
      data: {
        teamPitching: { StrikeOutRate: strikeOutRateB, RunsAllowed: runsAllowedB },
      },
    } = await axios.get<TeamPitchingResponse>(
      `http://localhost:8081/team/pitching/${teamIdB}/${yearB}`
    );
    const {
      data: {
        teamBatting: {
          HomeRun: homeRunB,
          SluggingPercentage: sluggingPercentageB,
          OnBasePercentage: onBasePercentageB,
        },
      },
    } = await axios.get<TeamBattingResponse>(
      `http://localhost:8081/team/batting/${teamIdB}/${yearB}`
    );

    const createData = [
      {
        item: '本塁打',
        A: ((homeRunA - minHomeRun) / (maxHomeRun - minHomeRun)) * 100,
        B: ((homeRunB - minHomeRun) / (maxHomeRun - minHomeRun)) * 100,
        fullMark: fullMark,
      },
      {
        item: '長打率',
        A:
          ((sluggingPercentageA - minSluggingPercentage) /
            (maxSluggingPercentage - minSluggingPercentage)) *
          100,
        B:
          ((sluggingPercentageB - minSluggingPercentage) /
            (maxSluggingPercentage - minSluggingPercentage)) *
          100,
        fullMark: fullMark,
      },
      {
        item: '出塁率',
        A:
          ((onBasePercentageA - minOnBasePercentage) /
            (maxOnBasePercentage - minOnBasePercentage)) *
          100,
        B:
          ((onBasePercentageB - minOnBasePercentage) /
            (maxOnBasePercentage - minOnBasePercentage)) *
          100,
        fullMark: fullMark,
      },
      {
        item: '奪三振率',
        A: ((strikeOutRateA - minStrikeOutRate) / (maxStrikeOutRate - minStrikeOutRate)) * 100,
        B: ((strikeOutRateB - minStrikeOutRate) / (maxStrikeOutRate - minStrikeOutRate)) * 100,
        fullMark: fullMark,
      },
      {
        item: '失点数',
        A: (1 - (runsAllowedA - minRunsAllowed) / (maxRunsAllowed - minRunsAllowed)) * 100,
        B: (1 - (runsAllowedB - minRunsAllowed) / (maxRunsAllowed - minRunsAllowed)) * 100,
        fullMark: fullMark,
      },
    ];
    setData(createData);
  };

  useEffect(() => {
    (async () => {
      setStrength();
    })();
  }, [yearA, teamA, yearB, teamB]);

  return (
    <GenericTemplate title="チーム戦力チャート">
      <Grid container alignItems="center" justifyContent="center">
        <Paper className={fixedHeightPaper}>
          <React.Fragment>
            <Box display="flex" flexDirection="column" p={1} width={400} height={300}>
              <Title>{'チーム選択'}</Title>
              <Selectable
                key={'yearA'}
                formControl={classes.formControl}
                selectLabel={'チームA年度'}
                initSelect={yearA}
                selects={props.years}
                setSelect={setYearA}
              />
              <Selectable
                key={'teamA'}
                formControl={classes.formControl}
                selectLabel={'チームA'}
                initSelect={teamA}
                selects={teamNameList}
                setSelect={setTeamA}
              />
              <Selectable
                key={'yearB'}
                formControl={classes.formControl}
                selectLabel={'チームB年度'}
                initSelect={yearB}
                selects={props.years}
                setSelect={setYearB}
              />
              <Selectable
                key={'teamB'}
                formControl={classes.formControl}
                selectLabel={'チームB'}
                initSelect={teamB}
                selects={teamNameList}
                setSelect={setTeamB}
              />
            </Box>
          </React.Fragment>
        </Paper>
        <Paper className={fixedHeightPaper}>
          <RadarChartComponent
            title="戦力チャート"
            data={data}
            nameA={teamA}
            keyA={'A'}
            nameB={teamB}
            keyB={'B'}
            help={'過去のデータの最大値を100、最小値を0として、点数換算した値を戦力値として表示'}
          />
        </Paper>
      </Grid>
    </GenericTemplate>
  );
}
export default StrengthPage;
