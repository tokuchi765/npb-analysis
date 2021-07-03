import React from 'react';
import { Link } from 'react-router-dom';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
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
import { ClassNameMap } from '@material-ui/core/styles/withStyles';

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

function TableComponentTitleBar(props: {
  classes: ClassNameMap<'formControl' | 'title' | 'table' | 'grid' | 'visuallyHidden' | 'paper'>;
  initSelect: string;
  title: string;
  selectLabel: string;
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
  return (
    <Typography className={props.classes.title} variant="h6" id="tableTitle" component="div">
      <Grid container className={props.classes.grid}>
        <Grid key={1} item>
          <Paper className={props.classes.paper}>
            {props.initSelect}
            {props.title}
          </Paper>
        </Grid>
        <Grid key={2} item>
          <Selectable
            formControl={props.classes.formControl}
            selectLabel={props.selectLabel}
            initSelect={props.initSelect}
            selects={props.selects}
            handleChange={props.handleChange}
          />
        </Grid>
      </Grid>
    </Typography>
  );
}

function TableComponentHader(props: {
  classes: ClassNameMap<'formControl' | 'title' | 'table' | 'grid' | 'visuallyHidden' | 'paper'>;
  headCells: HeadCell[];
  orderBy: string;
  order: Order;
  createSortHandler: (property: string) => (event: React.MouseEvent<unknown>) => void;
}) {
  return (
    <TableHead>
      <TableRow>
        {props.headCells.map((headCell) => (
          <TableCell
            key={headCell.id}
            align={headCell.numeric ? 'right' : 'left'}
            padding={headCell.disablePadding ? 'none' : 'default'}
            sortDirection={props.orderBy === headCell.id ? props.order : false}
          >
            <TableSortLabel
              active={props.orderBy === headCell.id}
              direction={props.orderBy === headCell.id ? props.order : 'asc'}
              onClick={props.createSortHandler(headCell.id)}
            >
              {headCell.label}
              {props.orderBy === headCell.id ? (
                <span className={props.classes.visuallyHidden}>
                  {props.order === 'desc' ? 'sorted descending' : 'sorted ascending'}
                </span>
              ) : null}
            </TableSortLabel>
          </TableCell>
        ))}
      </TableRow>
    </TableHead>
  );
}

function TableComponentBody(props: { datas: { main: string }[]; order: Order; orderBy: string }) {
  return (
    <TableBody>
      {stableSort(props.datas, getComparator(props.order, props.orderBy)).map((teamData, index) => {
        const labelId = `enhanced-table-checkbox-${index}`;
        return (
          <TableRow hover tabIndex={-1} key={teamData.main}>
            <TableCell component="th" id={labelId} scope="row" padding="none" key={teamData.main}>
              <div>{teamData.main}</div>
            </TableCell>
            {_.map(teamData, (val, key) => {
              if (key === 'main') {
                return;
              }
              return (
                <TableCell align="right" key={key}>
                  {val}
                </TableCell>
              );
            })}
          </TableRow>
        );
      })}
    </TableBody>
  );
}

function TableLinkComponentBody(props: {
  datas: { main: string }[];
  order: Order;
  orderBy: string;
  mainLink: boolean;
  linkValues: Map<string, string>;
  path: string;
}) {
  return (
    <TableBody>
      {stableSort(props.datas, getComparator(props.order, props.orderBy)).map((teamData, index) => {
        const labelId = `enhanced-table-checkbox-${index}`;
        return (
          <TableRow hover tabIndex={-1} key={teamData.main}>
            <TableCell component="th" id={labelId} scope="row" padding="none" key={teamData.main}>
              <Link
                to={{
                  pathname: `${props.path}${props.linkValues.get(teamData.main)}`,
                }}
              >
                {teamData.main}
              </Link>
            </TableCell>
            {_.map(teamData, (val, key) => {
              if (key === 'main') {
                return;
              }
              return (
                <TableCell align="right" key={key}>
                  {val}
                </TableCell>
              );
            })}
          </TableRow>
        );
      })}
    </TableBody>
  );
}

export function TableComponent(props: {
  title: string;
  setSelect: (select: string) => void;
  getDataList: (year: string) => void;
  datas: { main: string }[];
  selects: string[];
  headCells: HeadCell[];
  initSorted: string;
  initSelect: string;
  selectLabel: string;
}) {
  const classes = useStyles();
  const [order, setOrder] = React.useState<Order>('desc');
  const [orderBy, setOrderBy] = React.useState<string>(props.initSorted);
  const handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    props.setSelect(event.target.value as string);
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
        <TableComponentTitleBar
          classes={classes}
          initSelect={props.initSelect}
          title={props.title}
          selectLabel={props.selectLabel}
          selects={props.selects}
          handleChange={handleChange}
        />
        <Table className={classes.table} aria-label="simple table">
          <TableComponentHader
            classes={classes}
            headCells={props.headCells}
            orderBy={orderBy}
            order={order}
            createSortHandler={createSortHandler}
          />
          <TableComponentBody datas={props.datas} order={order} orderBy={orderBy} />
        </Table>
      </TableContainer>
    </React.Fragment>
  );
}

export function TableLinkComponent(props: {
  title: string;
  setSelect: (select: string) => void;
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
  const [order, setOrder] = React.useState<Order>('desc');
  const [orderBy, setOrderBy] = React.useState<string>(props.initSorted);
  const handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    props.setSelect(event.target.value as string);
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
        <TableComponentTitleBar
          classes={classes}
          initSelect={props.initSelect}
          title={props.title}
          selectLabel={props.selectLabel}
          selects={props.selects}
          handleChange={handleChange}
        />
        <Table className={classes.table} aria-label="simple table">
          <TableComponentHader
            classes={classes}
            headCells={props.headCells}
            orderBy={orderBy}
            order={order}
            createSortHandler={createSortHandler}
          />
          <TableLinkComponentBody
            datas={props.datas}
            order={order}
            orderBy={orderBy}
            mainLink={props.mainLink}
            linkValues={props.linkValues}
            path={props.path}
          />
        </Table>
      </TableContainer>
    </React.Fragment>
  );
}
