import React from 'react';
import { Link } from 'react-router-dom';
import { createStyles, lighten, makeStyles, Theme } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import Typography from '@material-ui/core/Typography';
import TableContainer from '@material-ui/core/TableContainer';
import TableSortLabel from '@material-ui/core/TableSortLabel';
import FormControl from '@material-ui/core/FormControl';
import InputLabel from '@material-ui/core/InputLabel';
import Select from '@material-ui/core/Select';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import MenuItem from '@material-ui/core/MenuItem';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import _ from 'lodash';

type Order = 'asc' | 'desc';

function descendingComparator<T>(a: T, b: T, orderBy: keyof T) {
  if (b[orderBy] < a[orderBy]) {
    return -1;
  }
  if (b[orderBy] > a[orderBy]) {
    return 1;
  }
  return 0;
}

function getComparator<Key extends keyof any>(
  order: Order,
  orderBy: Key
): (a: { [key in Key]: number | string }, b: { [key in Key]: number | string }) => number {
  return order === 'desc'
    ? (a, b) => descendingComparator(a, b, orderBy)
    : (a, b) => -descendingComparator(a, b, orderBy);
}

function stableSort<T>(array: T[], comparator: (a: T, b: T) => number) {
  const stabilizedThis = array.map((el, index) => [el, index] as [T, number]);
  stabilizedThis.sort((a, b) => {
    const order = comparator(a[0], b[0]);
    if (order !== 0) return order;
    return a[1] - b[1];
  });
  return stabilizedThis.map((el) => el[0]);
}

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    table: {
      minWidth: 650,
    },
    visuallyHidden: {
      border: 0,
      clip: 'rect(0 0 0 0)',
      height: 1,
      margin: -1,
      overflow: 'hidden',
      padding: 0,
      position: 'absolute',
      top: 20,
      width: 1,
    },
    title: {
      flex: 'auto',
      paddingLeft: '10px',
    },
    formControl: {
      margin: theme.spacing(1),
      minWidth: 120,
    },
    paper: {
      textAlign: 'center',
      height: 50,
      width: 300,
      padding: theme.spacing(1, 2),
    },
    grid: {
      paddingTop: 10,
    },
  })
);

export interface HeadCell {
  disablePadding: boolean;
  id: string;
  label: string;
  numeric: boolean;
}

function Main(props: { mainLink: boolean; main: string; value: string | undefined; path: string }) {
  if (props.mainLink) {
    return (
      <Link
        to={{
          pathname: `${props.path}${props.value}`,
        }}
      >
        {props.main}
      </Link>
    );
  } else {
    return <div>{props.main}</div>;
  }
}

function Selectable(props: {
  formControl: string;
  selectLabel: string;
  initSelect: string;
  selects: string[];
  handleChange:
    | ((
        event: React.ChangeEvent<{
          name?: string | undefined;
          value: unknown;
        }>,
        child: React.ReactNode
      ) => void)
    | undefined;
}) {
  if (!_.isEmpty(props.selects)) {
    return (
      <FormControl className={props.formControl}>
        <InputLabel id="demo-simple-select-label">{props.selectLabel}</InputLabel>
        <Select
          labelId="demo-simple-select-label"
          id="demo-simple-select"
          value={props.initSelect}
          onChange={props.handleChange}
        >
          {props.selects.map((select) => {
            return (
              <MenuItem key={select} value={select}>
                {select}
              </MenuItem>
            );
          })}
        </Select>
      </FormControl>
    );
  } else {
    return <div></div>;
  }
}

export default function TablePages(props: {
  title: string;
  getDataList: (year: string) => void;
  datas: { main: string }[];
  selects: string[];
  headCells: HeadCell[];
  initSorted: string;
  initSelect: string;
  selectLabel: string;
  mainLink: boolean;
  linkValues: Map<string, string>;
  path: string;
}) {
  const classes = useStyles();
  const [initSelect, setYear] = React.useState(props.initSelect);
  const [order, setOrder] = React.useState<Order>('desc');
  const [orderBy, setOrderBy] = React.useState<string>(props.initSorted);
  const handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setYear(event.target.value as string);
    props.getDataList(String(event.target.value));
  };
  const handleRequestSort = (event: React.MouseEvent<unknown>, property: string) => {
    const isAsc = orderBy === property && order === 'asc';
    setOrder(isAsc ? 'desc' : 'asc');
    setOrderBy(property);
  };
  const createSortHandler = (property: string) => (event: React.MouseEvent<unknown>) => {
    handleRequestSort(event, property);
  };

  return (
    <React.Fragment>
      <TableContainer component={Paper}>
        <Typography className={classes.title} variant="h6" id="tableTitle" component="div">
          <Grid container className={classes.grid}>
            <Grid key={1} item>
              <Paper className={classes.paper}>
                {initSelect}
                {props.title}
              </Paper>
            </Grid>
            <Grid key={2} item>
              <Selectable
                formControl={classes.formControl}
                selectLabel={props.selectLabel}
                initSelect={initSelect}
                selects={props.selects}
                handleChange={handleChange}
              />
            </Grid>
          </Grid>
        </Typography>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
              {props.headCells.map((headCell) => (
                <TableCell
                  key={headCell.id}
                  align={headCell.numeric ? 'right' : 'left'}
                  padding={headCell.disablePadding ? 'none' : 'default'}
                  sortDirection={orderBy === headCell.id ? order : false}
                >
                  <TableSortLabel
                    active={orderBy === headCell.id}
                    direction={orderBy === headCell.id ? order : 'asc'}
                    onClick={createSortHandler(headCell.id)}
                  >
                    {headCell.label}
                    {orderBy === headCell.id ? (
                      <span className={classes.visuallyHidden}>
                        {order === 'desc' ? 'sorted descending' : 'sorted ascending'}
                      </span>
                    ) : null}
                  </TableSortLabel>
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {stableSort(props.datas, getComparator(order, orderBy)).map((teamData, index) => {
              const labelId = `enhanced-table-checkbox-${index}`;
              return (
                <TableRow hover tabIndex={-1} key={teamData.main}>
                  <TableCell component="th" id={labelId} scope="row" padding="none">
                    <Main
                      mainLink={props.mainLink}
                      main={teamData.main}
                      value={props.linkValues.get(teamData.main)}
                      path={props.path}
                    />
                  </TableCell>
                  {_.map(teamData, (val, key) => {
                    if (key === 'main') {
                      return;
                    }
                    return <TableCell align="right">{val}</TableCell>;
                  })}
                </TableRow>
              );
            })}
          </TableBody>
        </Table>
      </TableContainer>
    </React.Fragment>
  );
}
