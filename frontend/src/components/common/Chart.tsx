import React from 'react';
import { Box, useTheme } from '@mui/material';
import { LineChart, Line, XAxis, YAxis, Label, CartesianGrid, Tooltip, Legend } from 'recharts';
import Title from './Title';

export interface ChartData {
  key: string;
  name: string;
  stroke: string;
}

export default function Chart(props: {
  title: string;
  data: any;
  label: string;
  chartDatas: ChartData[];
  width: number;
  height: number;
}) {
  const theme = useTheme();

  return (
    <React.Fragment>
      <Box display="flex" flexDirection="column" p={1}>
        <Title>{props.title}</Title>
        <LineChart
          width={props.width}
          height={props.height}
          data={props.data}
          margin={{
            top: 16,
            right: 16,
            bottom: 0,
            left: 24,
          }}
        >
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="year" stroke={theme.palette.text.secondary} />
          <YAxis stroke={theme.palette.text.secondary}>
            <Label
              position="left"
              style={{ textAnchor: 'middle', fill: theme.palette.text.primary }}
            >
              {props.label}
            </Label>
          </YAxis>
          <Tooltip />
          <Legend />
          {props.chartDatas.map((value) => {
            return (
              <Line
                type="monotone"
                dataKey={value.key}
                name={value.name}
                stroke={value.stroke}
                key={value.key}
              />
            );
          })}
        </LineChart>
      </Box>
    </React.Fragment>
  );
}
