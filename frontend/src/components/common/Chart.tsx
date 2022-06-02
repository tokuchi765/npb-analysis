import React from 'react';
import { Box, useTheme } from '@mui/material';
import { LineChart, Line, XAxis, YAxis, Label, CartesianGrid, Tooltip, Legend } from 'recharts';
import Title from './Title';

export default function Chart(props: {
  title: string;
  data: any;
  label: string;
  lineData: string[][];
}) {
  const theme = useTheme();

  return (
    <React.Fragment>
      <Box display="flex" flexDirection="column" p={1}>
        <Title>{props.title}</Title>
        <LineChart
          width={400}
          height={300}
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
          {props.lineData.map((val) => {
            return <Line type="monotone" key={val[0]} dataKey={val[0]} stroke={val[1]} />;
          })}
        </LineChart>
      </Box>
    </React.Fragment>
  );
}
