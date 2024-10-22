import React, { useEffect, useState } from 'react';
import GenericTemplate from '../templates/GenericTemplate';
import {
  TextField,
  Select,
  MenuItem,
  FormControl,
  Paper,
  RadioGroup,
  FormControlLabel,
  Radio,
  Grid,
  Typography,
  SelectChangeEvent,
  Button,
  Accordion,
  AccordionSummary,
  AccordionDetails,
  Box,
} from '@mui/material';
import { searchBatterGrades, searchPitcherGrades } from '../../data/api/player';
import { HeadCell, TableSearchComponent } from '../common/TableComponent';
import {
  ThresholdType,
  BatterGrades,
  PitcherGrades,
  SearchBatterGradesCondition,
  SearchBaseGradesCondition,
  SearchPitcherGradesCondition,
} from '../../data/type';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import { SearchBatterCondition } from './component/SearchBatterCondition';
import { SearchPitchrCondition } from './component/SearchPitcherCondition';
import Loading from '../common/Loading';

enum Type {
  Batter = 'batter',
  Pitcher = 'pitcher',
}

enum PeriodType {
  Period = 'period',
  Total = 'total',
}

interface BatterGradesData {
  main: string;
  year: number;
  team: string;
  battingAverage: number;
  plateAppearance: number;
  onBasePercentage: number;
  baseOnBalls: number;
  hit: number;
  single: number;
  double: number;
  triple: number;
  homeRun: number;
  stolenBase: number;
  strikeOut: number;
  strikeOutRate: number;
  wOba: number;
  rc: number;
  babip: number;
  groundedIntoDoublePlay: number;
  sluggingPercentage: number;
}

const BatterGradesHeadCells: HeadCell[] = [
  { id: 'main', numeric: false, disablePadding: true, label: '選手名' },
  { id: 'year', numeric: true, disablePadding: true, label: '年度' },
  { id: 'team', numeric: true, disablePadding: false, label: 'チーム' },
  { id: 'plateAppearance', numeric: true, disablePadding: true, label: '打席' },
  { id: 'onBasePercentage', numeric: true, disablePadding: true, label: '出塁率' },
  { id: 'battingAverage', numeric: true, disablePadding: true, label: '打率' },
  { id: 'baseOnBalls', numeric: true, disablePadding: true, label: '四球' },
  { id: 'hit', numeric: true, disablePadding: true, label: '安打' },
  { id: 'single', numeric: true, disablePadding: true, label: '単打' },
  { id: 'double', numeric: true, disablePadding: true, label: '二塁打' },
  { id: 'triple', numeric: true, disablePadding: true, label: '三塁打' },
  { id: 'homeRun', numeric: true, disablePadding: true, label: '本塁打' },
  { id: 'stolenBase', numeric: true, disablePadding: true, label: '盗塁' },
  { id: 'strikeOut', numeric: true, disablePadding: true, label: '三振' },
  { id: 'strikeOutRate', numeric: true, disablePadding: true, label: '三振率' },
  { id: 'wOba', numeric: true, disablePadding: true, label: '加重出塁率' },
  { id: 'rc', numeric: true, disablePadding: true, label: '創出得点' },
  { id: 'babip', numeric: true, disablePadding: true, label: 'BABIP' },
  { id: 'groundedIntoDoublePlay', numeric: true, disablePadding: true, label: '併殺打' },
  { id: 'sluggingPercentage', numeric: true, disablePadding: true, label: '長打率' },
];

function buildBatterGradesDatas(responses: BatterGrades[]) {
  const gradesDate: BatterGradesData[] = [];
  responses.forEach((response) => {
    gradesDate.push({
      main: response.Name,
      year: response.Year,
      team: response.Team,
      plateAppearance: response.PlateAppearance,
      onBasePercentage: response.OnBasePercentage,
      battingAverage: response.BattingAverage,
      baseOnBalls: response.BaseOnBalls,
      hit: response.Hit,
      single: response.Single,
      double: response.Double,
      triple: response.Triple,
      homeRun: response.HomeRun,
      stolenBase: response.StolenBase,
      strikeOut: response.StrikeOut,
      strikeOutRate: response.StrikeOutRate,
      wOba: response.WOba,
      rc: response.RC,
      babip: response.Babip,
      groundedIntoDoublePlay: response.GroundedIntoDoublePlay,
      sluggingPercentage: response.SluggingPercentage,
    });
  });
  return gradesDate;
}

interface PitcherGradesData {
  main: string;
  year: number;
  team: string;
  pitched: number;
  inningsPitched: number;
  earnedRunAverage: number;
  babip: number;
  strikeOutRate: number;
  strikeOut: number;
  hit: number;
  baseOnBalls: number;
  homeRun: number;
  win: number;
  lose: number;
  save: number;
  hold: number;
  holdPoint: number;
  completeGame: number;
  shutout: number;
  winningRate: number;
}

const PitcherGradesHeadCells: HeadCell[] = [
  { id: 'main', numeric: false, disablePadding: true, label: '選手名' },
  { id: 'year', numeric: true, disablePadding: true, label: '年度' },
  { id: 'team', numeric: true, disablePadding: false, label: 'チーム' },
  { id: 'pitched', numeric: true, disablePadding: true, label: '登板' },
  { id: 'inningsPitched', numeric: true, disablePadding: true, label: '投球回数' },
  { id: 'earnedRunAverage', numeric: true, disablePadding: true, label: '防御率' },
  { id: 'babip', numeric: true, disablePadding: true, label: '被BABIP' },
  { id: 'strikeOutRate', numeric: true, disablePadding: true, label: '奪三振率' },
  { id: 'strikeOut', numeric: true, disablePadding: true, label: '三振' },
  { id: 'hit', numeric: true, disablePadding: true, label: '被安打' },
  { id: 'baseOnBalls', numeric: true, disablePadding: true, label: '四球' },
  { id: 'homeRun', numeric: true, disablePadding: true, label: '被本塁打' },
  { id: 'win', numeric: true, disablePadding: true, label: '勝利' },
  { id: 'lose', numeric: true, disablePadding: true, label: '敗北' },
  { id: 'save', numeric: true, disablePadding: true, label: 'セーブ' },
  { id: 'hold', numeric: true, disablePadding: true, label: 'ホールド' },
  { id: 'holdPoint', numeric: true, disablePadding: true, label: 'HP' },
  { id: 'completeGame', numeric: true, disablePadding: true, label: '完投' },
  { id: 'shutout', numeric: true, disablePadding: true, label: '完封' },
  { id: 'winningRate', numeric: true, disablePadding: true, label: '勝率' },
];

function buildPitcherGradesDatas(responses: PitcherGrades[]) {
  const gradesDate: PitcherGradesData[] = [];
  responses.forEach((response) => {
    gradesDate.push({
      main: response.Name,
      year: response.Year,
      team: response.Team,
      pitched: response.Pitched,
      inningsPitched: response.InningsPitched,
      earnedRunAverage: response.EarnedRunAverage,
      babip: response.Babip,
      strikeOutRate: response.StrikeOutRate,
      strikeOut: response.StrikeOut,
      hit: response.Hit,
      baseOnBalls: response.BaseOnBalls,
      homeRun: response.HomeRun,
      win: response.Win,
      lose: response.Lose,
      save: response.Save,
      hold: response.Hold,
      holdPoint: response.HoldPoint,
      completeGame: response.CompleteGame,
      shutout: response.Shutout,
      winningRate: response.WinningRate,
    });
  });
  return gradesDate;
}

function createPlayerIds(responses: BatterGrades[] | PitcherGrades[]) {
  const playerIdMap: Map<string, string> = new Map<string, string>();
  responses.forEach((response) => {
    playerIdMap.set(response.Name, response.PlayerID);
  });
  return playerIdMap;
}

function SearchRadioGroup(props: {
  name: string;
  labels: Map<string, string>;
  initValue: string;
  handleChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
}) {
  return (
    <React.Fragment>
      <FormControl component="fieldset">
        <RadioGroup
          aria-labelledby={`${props.name}-label`}
          name={props.name}
          value={props.initValue}
          onChange={props.handleChange}
          row
        >
          {Array.from(props.labels.entries()).map(([key, value]) => {
            return (
              <FormControlLabel
                key={key}
                value={key}
                control={<Radio sx={{ transform: 'scale(0.6)' }} />}
                label={value}
                sx={{
                  '& .MuiFormControlLabel-label': { fontSize: '12px' },
                }}
              />
            );
          })}
        </RadioGroup>
      </FormControl>
    </React.Fragment>
  );
}

export function ThresholdRadioGroup(props: {
  name: string;
  initValue: ThresholdType;
  setThresholdType: (value: ThresholdType) => void;
}) {
  const handleChangeThresholdType = (event: React.ChangeEvent<HTMLInputElement>) => {
    props.setThresholdType(
      Number((event.target as HTMLInputElement).value) === ThresholdType.GreaterOrEqual
        ? ThresholdType.GreaterOrEqual
        : ThresholdType.LessOrEqual
    );
  };

  return (
    <React.Fragment>
      <FormControl component="fieldset">
        <RadioGroup
          aria-labelledby={`${props.name}-label`}
          name={props.name}
          value={props.initValue}
          onChange={handleChangeThresholdType}
          row
        >
          {Array.from(
            new Map<number, string>([
              [ThresholdType.GreaterOrEqual, '以上'],
              [ThresholdType.LessOrEqual, '以下'],
            ])
          ).map(([key, value]) => {
            return (
              <FormControlLabel
                key={key}
                value={key}
                control={<Radio sx={{ transform: 'scale(0.6)' }} />}
                label={value}
                sx={{
                  '& .MuiFormControlLabel-label': { fontSize: '12px' },
                  marginLeft: 0,
                  marginRight: 0,
                }}
              />
            );
          })}
        </RadioGroup>
      </FormControl>
    </React.Fragment>
  );
}

export function InputFloatText(props: {
  name: string;
  minValue: number;
  maxValue: number;
  maxLength: number;
  placeholder: string;
  height: number;
  width: number;
  value: number | undefined;
  setValue: (i: number | undefined) => void;
}) {
  const [value, setValue] = useState('');
  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const inputValue = event.target.value;

    if (inputValue === '') {
      setValue('');
      props.setValue(undefined);
      return;
    }

    // 数値に変換できるかチェック
    if (!isNaN(Number(inputValue)) && inputValue.length <= props.maxLength) {
      setValue(inputValue);
      const numericValue = Number(inputValue);
      props.setValue(numericValue);
    }
  };

  const handleBlur = () => {
    if (value === '') {
      setValue('');
      props.setValue(undefined);
      return;
    }

    if (!isNaN(Number(props.value))) {
      let numericValue = Number(props.value);

      // 上限内に数値を収める
      if (numericValue < props.minValue) {
        numericValue = props.minValue;
      } else if (numericValue > props.maxValue) {
        numericValue = props.maxValue;
      }

      setValue(String(numericValue));
      props.setValue(numericValue);
    }
  };

  useEffect(() => {
    if (props.value === undefined) {
      setValue('');
    } else {
      setValue(String(props.value));
    }
  }, [props.value]);

  return (
    <React.Fragment>
      <TextField
        name={props.name}
        value={value}
        onChange={handleChange}
        onBlur={handleBlur}
        placeholder={props.placeholder}
        inputProps={{
          sx: {
            '&::placeholder': {
              fontSize: '12px',
            },
          },
          inputMode: 'numeric',
          maxLength: props.maxLength,
          style: { height: props.height, width: props.width },
        }}
      />
    </React.Fragment>
  );
}

export function InputNumberText(props: {
  name: string;
  minValue: number;
  maxValue: number;
  maxLength: number;
  placeholder: string;
  height: number;
  width: number;
  value: number | undefined;
  error?: boolean;
  setValue: (i: number | undefined) => void;
}) {
  const [value, setValue] = useState('');
  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const inputValue = event.target.value;

    if (inputValue === '') {
      setValue('');
      props.setValue(undefined);
      return;
    }

    // 数値に変換できるかチェック
    if (
      !isNaN(Number(inputValue)) &&
      inputValue.length <= props.maxLength &&
      /^\d*$/.test(inputValue)
    ) {
      setValue(inputValue);
      const numericValue = Number(inputValue);
      props.setValue(numericValue);
    }
  };

  const handleBlur = () => {
    if (value === '') {
      setValue('');
      props.setValue(undefined);
      return;
    }

    if (!isNaN(Number(props.value))) {
      let numericValue = Number(props.value);

      // 上限内に数値を収める
      if (numericValue < props.minValue) {
        numericValue = props.minValue;
      } else if (numericValue > props.maxValue) {
        numericValue = props.maxValue;
      }

      setValue(String(numericValue));
      props.setValue(numericValue);
    }
  };

  useEffect(() => {
    if (props.value === undefined) {
      setValue('');
    } else {
      setValue(String(props.value));
    }
  }, [props.value]);

  return (
    <React.Fragment>
      <TextField
        name={props.name}
        value={value}
        onChange={handleChange}
        onBlur={handleBlur}
        placeholder={props.placeholder}
        inputProps={{
          sx: {
            '&::placeholder': {
              fontSize: '12px',
            },
          },
          inputMode: 'numeric',
          maxLength: props.maxLength,
          style: {
            height: props.height,
            width: props.width,
            backgroundColor: props.error ? '#FFCDD2' : 'inherit',
          },
        }}
      />
    </React.Fragment>
  );
}

export function SearchLabel(props: { label: string }) {
  return (
    <React.Fragment>
      <Typography paddingLeft={'10px'} fontSize={'15px'} fontWeight="bold">
        {props.label}
      </Typography>
    </React.Fragment>
  );
}

function ErrorLabel(props: { label: string }) {
  return (
    <React.Fragment>
      <Typography paddingLeft={'10px'} fontSize={'10px'} fontWeight="bold" color="#FF0000">
        {props.label}
      </Typography>
    </React.Fragment>
  );
}

function SearchDropdown(props: {
  name: string;
  selectItems: Map<string, string>;
  selectedValue: string;
  setValue: (team: string) => void;
}) {
  const handleChange = (event: SelectChangeEvent<string>) => {
    props.setValue(event.target.value);
  };

  return (
    <FormControl>
      <Select
        sx={{ width: 110, height: 33, fontSize: 13 }}
        labelId={props.name}
        value={props.selectedValue}
        onChange={handleChange}
      >
        {Array.from(props.selectItems.entries()).map(([key, value]) => {
          return (
            <MenuItem key={key} value={key}>
              {value}
            </MenuItem>
          );
        })}
      </Select>
    </FormControl>
  );
}

function SearchGradesPage() {
  // 共通
  const [errors, setErrors] = useState<string[]>([]);
  const [playerIdMap, setPlayerIds] = useState<Map<string, string>>(new Map<string, string>());
  const [type, setType] = useState<Type>(Type.Batter);
  const [periodType, setPeriodType] = useState<PeriodType>(PeriodType.Period);
  const [fromPeriodError, setFromPeriodError] = useState<boolean>(false);
  const [toPeriodError, setToPeriodError] = useState<boolean>(false);
  const [totalYearError, setTotalYearError] = useState<boolean>(false);
  const [searchBaseGradesCondition, setSearchBaseGradesCondition] =
    useState<SearchBaseGradesCondition>({
      Total: false,
      TotalYear: undefined,
      FromYear: undefined,
      ToYear: undefined,
      TeamID: '',
    });

  const handleChangeType = (event: React.ChangeEvent<HTMLInputElement>) => {
    const typeVal =
      (event.target as HTMLInputElement).value === Type.Batter ? Type.Batter : Type.Pitcher;
    setType(typeVal);
  };

  const onClickClearSearchCondition = () => {
    setSearchBaseGradesCondition({
      Total: false,
      TotalYear: undefined,
      FromYear: undefined,
      ToYear: undefined,
      TeamID: '',
    });
    setSearchBatterGradesCondition({
      PlateAppearanceThresholdType: ThresholdType.GreaterOrEqual,
      PlateAppearance: undefined,
      OnBasePercentageThresholdType: ThresholdType.GreaterOrEqual,
      OnBasePercentage: undefined,
      BattingAverageThresholdType: ThresholdType.GreaterOrEqual,
      BattingAverage: undefined,
      BaseOnBallsThresholdType: ThresholdType.GreaterOrEqual,
      BaseOnBalls: undefined,
      HitThresholdType: ThresholdType.GreaterOrEqual,
      Hit: undefined,
      SingleThresholdType: ThresholdType.GreaterOrEqual,
      Single: undefined,
      DoubleThresholdType: ThresholdType.GreaterOrEqual,
      Double: undefined,
      TripleThresholdType: ThresholdType.GreaterOrEqual,
      Triple: undefined,
      SluggingPercentageThresholdType: ThresholdType.GreaterOrEqual,
      SluggingPercentage: undefined,
      HomeRunThresholdType: ThresholdType.GreaterOrEqual,
      HomeRun: undefined,
      StrikeOutThresholdType: ThresholdType.GreaterOrEqual,
      StrikeOut: undefined,
      StrikeOutRateThresholdType: ThresholdType.GreaterOrEqual,
      StrikeOutRate: undefined,
      StolenBaseThresholdType: ThresholdType.GreaterOrEqual,
      StolenBase: undefined,
      GroundedIntoDoublePlayThresholdType: ThresholdType.GreaterOrEqual,
      GroundedIntoDoublePlay: undefined,
      WObaThresholdType: ThresholdType.GreaterOrEqual,
      WOba: undefined,
      RCThresholdType: ThresholdType.GreaterOrEqual,
      RC: undefined,
      BabipThresholdType: ThresholdType.GreaterOrEqual,
      Babip: undefined,
    });
    setSearchPitcherGradesCondition({
      PitchedThresholdType: ThresholdType.GreaterOrEqual,
      Pitched: undefined,
      InningsPitchedThresholdType: ThresholdType.GreaterOrEqual,
      InningsPitched: undefined,
      EarnedRunAverageThresholdType: ThresholdType.GreaterOrEqual,
      EarnedRunAverage: undefined,
      BabipThresholdType: ThresholdType.GreaterOrEqual,
      Babip: undefined,
      StrikeOutRateThresholdType: ThresholdType.GreaterOrEqual,
      StrikeOutRate: undefined,
      StrikeOutThresholdType: ThresholdType.GreaterOrEqual,
      StrikeOut: undefined,
      HitThresholdType: ThresholdType.GreaterOrEqual,
      Hit: undefined,
      BaseOnBallsThresholdType: ThresholdType.GreaterOrEqual,
      BaseOnBalls: undefined,
      HomeRunThresholdType: ThresholdType.GreaterOrEqual,
      HomeRun: undefined,
      WinThresholdType: ThresholdType.GreaterOrEqual,
      Win: undefined,
      LoseThresholdType: ThresholdType.GreaterOrEqual,
      Lose: undefined,
      SaveThresholdType: ThresholdType.GreaterOrEqual,
      Save: undefined,
      HoldThresholdType: ThresholdType.GreaterOrEqual,
      Hold: undefined,
      HoldPointThresholdType: ThresholdType.GreaterOrEqual,
      HoldPoint: undefined,
      CompleteGameThresholdType: ThresholdType.GreaterOrEqual,
      CompleteGame: undefined,
      ShutoutThresholdType: ThresholdType.GreaterOrEqual,
      Shutout: undefined,
      WinningRateThresholdType: ThresholdType.GreaterOrEqual,
      WinningRate: undefined,
    });
  };

  const handleChangePeriodType = (event: React.ChangeEvent<HTMLInputElement>) => {
    const periodType =
      (event.target as HTMLInputElement).value === PeriodType.Period
        ? PeriodType.Period
        : PeriodType.Total;

    setPeriodType(periodType);
    if (periodType === PeriodType.Period) {
      setSearchBaseGradesCondition({
        ...searchBaseGradesCondition,
        Total: false,
        TotalYear: undefined,
      });
    } else {
      setSearchBaseGradesCondition({
        ...searchBaseGradesCondition,
        FromYear: undefined,
        ToYear: undefined,
        TeamID: '',
        Total: true,
      });
    }
  };

  const clearErrorState = () => {
    setErrors([]);
    setFromPeriodError(false);
    setToPeriodError(false);
    setTotalYearError(false);
    setBatterErrors([]);
    setPitcherErrors([]);
  };

  // 野手検索条件
  const [batterErrors, setBatterErrors] = useState<string[]>([]);
  const [batterGradesDatas, setBatterGradesDatas] = useState<BatterGradesData[]>([]);
  const [searchBatterGradesCondition, setSearchBatterGradesCondition] =
    useState<SearchBatterGradesCondition>({
      PlateAppearanceThresholdType: ThresholdType.GreaterOrEqual,
      PlateAppearance: undefined,
      OnBasePercentageThresholdType: ThresholdType.GreaterOrEqual,
      OnBasePercentage: undefined,
      BattingAverageThresholdType: ThresholdType.GreaterOrEqual,
      BattingAverage: undefined,
      BaseOnBallsThresholdType: ThresholdType.GreaterOrEqual,
      BaseOnBalls: undefined,
      HitThresholdType: ThresholdType.GreaterOrEqual,
      Hit: undefined,
      SingleThresholdType: ThresholdType.GreaterOrEqual,
      Single: undefined,
      DoubleThresholdType: ThresholdType.GreaterOrEqual,
      Double: undefined,
      TripleThresholdType: ThresholdType.GreaterOrEqual,
      Triple: undefined,
      SluggingPercentageThresholdType: ThresholdType.GreaterOrEqual,
      SluggingPercentage: undefined,
      HomeRunThresholdType: ThresholdType.GreaterOrEqual,
      HomeRun: undefined,
      StrikeOutThresholdType: ThresholdType.GreaterOrEqual,
      StrikeOut: undefined,
      StrikeOutRateThresholdType: ThresholdType.GreaterOrEqual,
      StrikeOutRate: undefined,
      StolenBaseThresholdType: ThresholdType.GreaterOrEqual,
      StolenBase: undefined,
      GroundedIntoDoublePlayThresholdType: ThresholdType.GreaterOrEqual,
      GroundedIntoDoublePlay: undefined,
      WObaThresholdType: ThresholdType.GreaterOrEqual,
      WOba: undefined,
      RCThresholdType: ThresholdType.GreaterOrEqual,
      RC: undefined,
      BabipThresholdType: ThresholdType.GreaterOrEqual,
      Babip: undefined,
    });

  const [searchPitcherGradesCondition, setSearchPitcherGradesCondition] =
    useState<SearchPitcherGradesCondition>({
      PitchedThresholdType: ThresholdType.GreaterOrEqual,
      Pitched: undefined,
      InningsPitchedThresholdType: ThresholdType.GreaterOrEqual,
      InningsPitched: undefined,
      EarnedRunAverageThresholdType: ThresholdType.GreaterOrEqual,
      EarnedRunAverage: undefined,
      BabipThresholdType: ThresholdType.GreaterOrEqual,
      Babip: undefined,
      StrikeOutRateThresholdType: ThresholdType.GreaterOrEqual,
      StrikeOutRate: undefined,
      StrikeOutThresholdType: ThresholdType.GreaterOrEqual,
      StrikeOut: undefined,
      HitThresholdType: ThresholdType.GreaterOrEqual,
      Hit: undefined,
      BaseOnBallsThresholdType: ThresholdType.GreaterOrEqual,
      BaseOnBalls: undefined,
      HomeRunThresholdType: ThresholdType.GreaterOrEqual,
      HomeRun: undefined,
      WinThresholdType: ThresholdType.GreaterOrEqual,
      Win: undefined,
      LoseThresholdType: ThresholdType.GreaterOrEqual,
      Lose: undefined,
      SaveThresholdType: ThresholdType.GreaterOrEqual,
      Save: undefined,
      HoldThresholdType: ThresholdType.GreaterOrEqual,
      Hold: undefined,
      HoldPointThresholdType: ThresholdType.GreaterOrEqual,
      HoldPoint: undefined,
      CompleteGameThresholdType: ThresholdType.GreaterOrEqual,
      CompleteGame: undefined,
      ShutoutThresholdType: ThresholdType.GreaterOrEqual,
      Shutout: undefined,
      WinningRateThresholdType: ThresholdType.GreaterOrEqual,
      WinningRate: undefined,
    });

  const isAllBatterConditionUndefind = () => {
    return (
      !searchBatterGradesCondition.PlateAppearance &&
      !searchBatterGradesCondition.OnBasePercentage &&
      !searchBatterGradesCondition.BattingAverage &&
      !searchBatterGradesCondition.SluggingPercentage &&
      !searchBatterGradesCondition.HomeRun &&
      !searchBatterGradesCondition.StrikeOut &&
      !searchBatterGradesCondition.StrikeOutRate &&
      !searchBatterGradesCondition.GroundedIntoDoublePlay &&
      !searchBatterGradesCondition.WOba &&
      !searchBatterGradesCondition.RC &&
      !searchBatterGradesCondition.Babip
    );
  };

  // 投手検索条件
  const [pitcherErrors, setPitcherErrors] = useState<string[]>([]);
  const [pitcherGradesDatas, setPitcherGradesDatas] = useState<PitcherGradesData[]>([]);

  const isAllPitcherConditionUndefind = () => {
    return (
      !searchPitcherGradesCondition.Pitched &&
      !searchPitcherGradesCondition.InningsPitched &&
      !searchPitcherGradesCondition.EarnedRunAverage &&
      !searchPitcherGradesCondition.StrikeOutRate &&
      !searchPitcherGradesCondition.StrikeOut &&
      !searchPitcherGradesCondition.Hit &&
      !searchPitcherGradesCondition.BaseOnBalls &&
      !searchPitcherGradesCondition.HomeRun &&
      !searchPitcherGradesCondition.Win &&
      !searchPitcherGradesCondition.Lose &&
      !searchPitcherGradesCondition.Save &&
      !searchPitcherGradesCondition.Hold &&
      !searchPitcherGradesCondition.HoldPoint &&
      !searchPitcherGradesCondition.CompleteGame &&
      !searchPitcherGradesCondition.Shutout &&
      !searchPitcherGradesCondition.WinningRate
    );
  };

  const handleSearch = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    clearErrorState();

    const errorMessages: string[] = [];
    // エラーチェック
    if (periodType === PeriodType.Period) {
      const from = searchBaseGradesCondition.FromYear;
      const to = searchBaseGradesCondition.ToYear;
      if (!from || !to) {
        errorMessages.push('期間を選択した場合、FROM年、TO年を入力してください');
        setFromPeriodError(true);
        setToPeriodError(true);
      }
      if (from && to && from > to) {
        errorMessages.push('FROM年はTO年以下の値を入力してください');
        setFromPeriodError(true);
        setToPeriodError(true);
      }
    } else if (periodType === PeriodType.Total) {
      if (!searchBaseGradesCondition.TotalYear || searchBaseGradesCondition.TotalYear === 0) {
        errorMessages.push('通算を選択した場合、年数を入力してください');
        setTotalYearError(true);
      }
    }

    const batterErrorMessages: string[] = [];
    if (type === Type.Batter) {
      if (isAllBatterConditionUndefind()) {
        batterErrorMessages.push('野手検索条件を最低でも1つ入力してください');
      }
    }

    const pitcherErrorMessages: string[] = [];
    if (type === Type.Pitcher) {
      if (isAllPitcherConditionUndefind()) {
        pitcherErrorMessages.push('投手検索条件を最低でも1つ入力してください');
      }
    }

    if (errorMessages.length !== 0) {
      setErrors(errorMessages);
      return;
    }

    if (batterErrorMessages.length !== 0) {
      setBatterErrors(batterErrorMessages);
      return;
    }

    if (pitcherErrorMessages.length !== 0) {
      setPitcherErrors(pitcherErrorMessages);
      return;
    }

    setLoading(true);

    if (type === Type.Batter) {
      const response = await searchBatterGrades(
        searchBaseGradesCondition,
        searchBatterGradesCondition
      );
      setBatterGradesDatas(buildBatterGradesDatas(response.results));
      setPlayerIds(createPlayerIds(response.results));
    } else {
      const response = await searchPitcherGrades(
        searchBaseGradesCondition,
        searchPitcherGradesCondition
      );
      setPitcherGradesDatas(buildPitcherGradesDatas(response.results));
      setPlayerIds(createPlayerIds(response.results));
    }

    clearErrorState();

    setLoading(false);
  };

  const [loading, setLoading] = useState(false);

  return (
    <GenericTemplate title="選手成績検索ページ">
      <Loading loading={loading} />
      <Paper
        component="form"
        onSubmit={handleSearch}
        sx={{ p: '2px 4px', display: 'flex', alignItems: 'center' }}
      >
        <Grid container alignItems="center">
          <Grid item xs={12}>
            <Typography fontWeight="bold">{'検索条件'}</Typography>
          </Grid>
          <Grid item xs={12} display={'flex'}>
            <Grid item display={'flex'} alignItems="center">
              <SearchLabel label={'種別：'} />
              <SearchRadioGroup
                name={'type'}
                labels={
                  new Map<Type, string>([
                    [Type.Batter, '野手'],
                    [Type.Pitcher, '投手'],
                  ])
                }
                initValue={type}
                handleChange={handleChangeType}
              />
            </Grid>
            <Grid item display={'flex'} alignItems="center">
              <SearchLabel label={'期間種別：'} />
              <SearchRadioGroup
                name={'periodType'}
                labels={
                  new Map<PeriodType, string>([
                    [PeriodType.Period, '期間'],
                    [PeriodType.Total, '通算'],
                  ])
                }
                initValue={periodType}
                handleChange={handleChangePeriodType}
              />
            </Grid>
            {periodType === PeriodType.Period && (
              <React.Fragment>
                <Grid item display={'flex'} alignItems="center">
                  <InputNumberText
                    name={'fromYear'}
                    minValue={1986}
                    maxValue={2030}
                    maxLength={4}
                    height={0}
                    width={50}
                    placeholder={'FROM年'}
                    value={searchBaseGradesCondition.FromYear}
                    setValue={(i) =>
                      setSearchBaseGradesCondition({
                        ...searchBaseGradesCondition,
                        FromYear: i,
                      })
                    }
                    error={fromPeriodError}
                  />
                  {'～'}
                  <InputNumberText
                    name={'toYear'}
                    minValue={1986}
                    maxValue={2030}
                    maxLength={4}
                    height={0}
                    width={50}
                    placeholder={'TO年'}
                    value={searchBaseGradesCondition.ToYear}
                    setValue={(i) =>
                      setSearchBaseGradesCondition({
                        ...searchBaseGradesCondition,
                        ToYear: i,
                      })
                    }
                    error={toPeriodError}
                  />
                </Grid>
                <Grid item display={'flex'} alignItems="center">
                  <SearchLabel label={'チーム：'} />
                  <SearchDropdown
                    name={'team'}
                    selectItems={
                      new Map<string, string>([
                        ['', '---'],
                        ['01', 'Giants'],
                        ['02', 'Baystars'],
                        ['03', 'Tigers'],
                        ['04', 'Carp'],
                        ['05', 'Dragons'],
                        ['06', 'Swallows'],
                        ['07', 'Lions'],
                        ['08', 'Hawks'],
                        ['09', 'Eagles'],
                        ['10', 'Marines'],
                        ['11', 'Fighters'],
                        ['12', 'Buffaloes'],
                      ])
                    }
                    selectedValue={searchBaseGradesCondition.TeamID}
                    setValue={(team) =>
                      setSearchBaseGradesCondition({
                        ...searchBaseGradesCondition,
                        TeamID: team,
                      })
                    }
                  />
                </Grid>
              </React.Fragment>
            )}
            {periodType === PeriodType.Total && (
              <Grid item display={'flex'} alignItems="center">
                <InputNumberText
                  name={'totalYear'}
                  minValue={0}
                  maxValue={30}
                  maxLength={2}
                  height={0}
                  width={35}
                  placeholder={'年数'}
                  value={searchBaseGradesCondition.TotalYear}
                  setValue={(i) =>
                    setSearchBaseGradesCondition({
                      ...searchBaseGradesCondition,
                      TotalYear: i,
                    })
                  }
                  error={totalYearError}
                />
              </Grid>
            )}
            <Grid item paddingLeft={'10px'} display={'flex'} alignItems="center">
              <Button sx={{ height: 33 }} type="submit" variant="contained" color="primary">
                {'検索'}
              </Button>
            </Grid>
            <Grid item paddingLeft={'10px'} display={'flex'} alignItems="center">
              <Button
                sx={{ height: 33 }}
                type="button"
                variant="contained"
                color="secondary"
                onClick={onClickClearSearchCondition}
              >
                {'検索条件クリア'}
              </Button>
            </Grid>
          </Grid>
          {errors.length !== 0 && (
            <Grid item xs={12} display={'flex'}>
              {errors.map((error, index) => {
                return <ErrorLabel label={error} key={index} />;
              })}
            </Grid>
          )}

          <Accordion defaultExpanded={true} disableGutters style={{ width: '100%' }}>
            <AccordionSummary
              expandIcon={<ExpandMoreIcon />}
              aria-controls="panel1a-content"
              id="panel1a-header"
            >
              {type === Type.Batter && <Typography fontWeight="bold">{'野手検索条件'}</Typography>}
              {type === Type.Pitcher && <Typography fontWeight="bold">{'投手検索条件'}</Typography>}
            </AccordionSummary>
            <AccordionDetails>
              {type === Type.Batter ? (
                <SearchBatterCondition
                  searchBatterGradesCondition={searchBatterGradesCondition}
                  setSearchBatterGradesCondition={setSearchBatterGradesCondition}
                />
              ) : (
                <SearchPitchrCondition
                  searchPitcherGradesCondition={searchPitcherGradesCondition}
                  setSearchPitcherGradesCondition={setSearchPitcherGradesCondition}
                />
              )}
            </AccordionDetails>
          </Accordion>
          <Grid item xs={12} display={'flex'}>
            {batterErrors.length !== 0 && (
              <Grid item xs={12} display={'flex'}>
                {batterErrors.map((error, index) => {
                  return <ErrorLabel label={error} key={index} />;
                })}
              </Grid>
            )}
            {pitcherErrors.length !== 0 && (
              <Grid item xs={12} display={'flex'}>
                {pitcherErrors.map((error, index) => {
                  return <ErrorLabel label={error} key={index} />;
                })}
              </Grid>
            )}
          </Grid>
        </Grid>
      </Paper>
      {type === Type.Batter && (
        <Box
          sx={{
            width: '100%',
            overflow: 'auto',
            border: '1px solid #ccc',
          }}
        >
          <TableSearchComponent
            datas={batterGradesDatas}
            headCells={BatterGradesHeadCells}
            initSorted={'main'}
            linkValues={playerIdMap}
            width={1600}
          />
        </Box>
      )}
      {type === Type.Pitcher && (
        <Box
          sx={{
            width: '100%',
            overflow: 'auto',
            border: '1px solid #ccc',
          }}
        >
          <TableSearchComponent
            datas={pitcherGradesDatas}
            headCells={PitcherGradesHeadCells}
            initSorted={'main'}
            linkValues={playerIdMap}
            width={1500}
          />
        </Box>
      )}
    </GenericTemplate>
  );
}

export default SearchGradesPage;
