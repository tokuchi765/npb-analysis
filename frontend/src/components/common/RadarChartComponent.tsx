import React from 'react';
import { Radar, RadarChart, PolarGrid, PolarAngleAxis, PolarRadiusAxis, Legend } from 'recharts';
import Title from './Title';
import { Box, Grid } from '@mui/material';
import Tooltip, { TooltipProps, tooltipClasses } from '@mui/material/Tooltip';
import { styled } from '@mui/material/styles';
import HelpOutlineIcon from '@mui/icons-material/HelpOutline';
import _ from 'lodash';

const LightTooltip = styled(({ className, ...props }: TooltipProps) => (
  <Tooltip {...props} classes={{ popper: className }} />
))(({ theme }) => ({
  [`& .${tooltipClasses.tooltip}`]: {
    backgroundColor: theme.palette.common.white,
    color: 'rgba(0, 0, 0, 0.87)',
    boxShadow: theme.shadows[1],
    fontSize: 11,
  },
}));

const style: React.CSSProperties = {
  marginTop: '5px',
};

export default function RadarChartComponent(props: {
  title: string;
  data: any;
  nameA: string;
  keyA: string;
  nameB: string;
  keyB: string;
  help: string;
}) {
  return (
    <React.Fragment>
      <Box display="auto" flexDirection="column" p={1}>
        <Grid container>
          <Title>{props.title}</Title>
          {props.help && (
            <LightTooltip title={props.help}>
              <HelpOutlineIcon style={style} />
            </LightTooltip>
          )}
        </Grid>
        <RadarChart width={400} height={300} cx="50%" cy="50%" outerRadius="80%" data={props.data}>
          <PolarGrid />
          <PolarAngleAxis dataKey="item" />
          <PolarRadiusAxis domain={[0, 100]} />
          <Radar
            name={props.nameA}
            dataKey={props.keyA}
            stroke="#8884d8"
            fill="#8884d8"
            fillOpacity={0.6}
          />
          <Radar
            name={props.nameB}
            dataKey={props.keyB}
            stroke="#82ca9d"
            fill="#82ca9d"
            fillOpacity={0.6}
          />
          <Legend />
        </RadarChart>
      </Box>
    </React.Fragment>
  );
}
