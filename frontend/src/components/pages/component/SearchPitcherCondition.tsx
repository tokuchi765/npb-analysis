import React from 'react';
import { SearchPitcherGradesCondition } from '../../../data/type';
import { Grid } from '@mui/material';
import {
  InputFloatText,
  InputNumberText,
  SearchLabel,
  ThresholdRadioGroup,
} from '../SearchGradesPage';

export interface SearchPitcherConditionProps {
  searchPitcherGradesCondition: SearchPitcherGradesCondition;
  setSearchPitcherGradesCondition: React.Dispatch<
    React.SetStateAction<SearchPitcherGradesCondition>
  >;
}

export function SearchPitcherCondition(props: SearchPitcherConditionProps) {
  return (
    <React.Fragment>
      <Grid item xs={12} display={'flex'}>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'登板数：'} />
          <InputNumberText
            name="pitched"
            minValue={0}
            maxValue={9999}
            maxLength={4}
            height={0}
            width={45}
            placeholder={'登板'}
            value={props.searchPitcherGradesCondition.Pitched}
            setValue={(i) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                Pitched: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'pitchedThresholdType'}
            initValue={props.searchPitcherGradesCondition.PitchedThresholdType}
            setThresholdType={(value) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                PitchedThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'投球回数：'} />
          <InputNumberText
            name={'inningsPitched'}
            minValue={0}
            maxValue={9999}
            maxLength={4}
            height={0}
            width={50}
            placeholder={'投球回'}
            value={props.searchPitcherGradesCondition.InningsPitched}
            setValue={(i) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                InningsPitched: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'inningsPitchedThresholdType'}
            initValue={props.searchPitcherGradesCondition.InningsPitchedThresholdType}
            setThresholdType={(value) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                InningsPitchedThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'防御率：'} />
          <InputFloatText
            name={'earnedRunAverage'}
            minValue={0}
            maxValue={999.99}
            maxLength={6}
            height={0}
            width={50}
            placeholder={'防御率'}
            value={props.searchPitcherGradesCondition.EarnedRunAverage}
            setValue={(i) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                EarnedRunAverage: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'earnedRunAverageThresholdType'}
            initValue={props.searchPitcherGradesCondition.EarnedRunAverageThresholdType}
            setThresholdType={(value) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                EarnedRunAverageThresholdType: value,
              })
            }
          />
        </Grid>
      </Grid>
      <Grid item xs={12} display={'flex'}>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'被BABIP：'} />
          <InputFloatText
            name={'babip'}
            minValue={0}
            maxValue={1.0}
            maxLength={4}
            height={0}
            width={50}
            placeholder={'被BABIP'}
            value={props.searchPitcherGradesCondition.Babip}
            setValue={(i) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                Babip: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'babipThresholdType'}
            initValue={props.searchPitcherGradesCondition.BabipThresholdType}
            setThresholdType={(value) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                BabipThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'奪三振率：'} />
          <InputFloatText
            name={'strikeOutRate'}
            minValue={0}
            maxValue={30.0}
            maxLength={5}
            height={0}
            width={50}
            placeholder={'奪三振率'}
            value={props.searchPitcherGradesCondition.StrikeOutRate}
            setValue={(i) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                StrikeOutRate: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'strikeOutRateThresholdType'}
            initValue={props.searchPitcherGradesCondition.StrikeOutRateThresholdType}
            setThresholdType={(value) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                StrikeOutRateThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'奪三振数：'} />
          <InputNumberText
            name={'strikeOut'}
            minValue={0}
            maxValue={9999}
            maxLength={4}
            height={0}
            width={50}
            placeholder={'奪三振数'}
            value={props.searchPitcherGradesCondition.StrikeOut}
            setValue={(i) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                StrikeOut: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'strikeOutThresholdType'}
            initValue={props.searchPitcherGradesCondition.StrikeOutThresholdType}
            setThresholdType={(value) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                StrikeOutThresholdType: value,
              })
            }
          />
        </Grid>
      </Grid>
      <Grid item xs={12} display={'flex'}>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'被安打数：'} />
          <InputNumberText
            name={'hit'}
            minValue={0}
            maxValue={9999}
            maxLength={4}
            height={0}
            width={50}
            placeholder={'被安打数'}
            value={props.searchPitcherGradesCondition.Hit}
            setValue={(i) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                Hit: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'hitThresholdType'}
            initValue={props.searchPitcherGradesCondition.HitThresholdType}
            setThresholdType={(value) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                HitThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'四球数：'} />
          <InputNumberText
            name={'baseOnBalls'}
            minValue={0}
            maxValue={999}
            maxLength={3}
            height={0}
            width={50}
            placeholder={'四球数'}
            value={props.searchPitcherGradesCondition.BaseOnBalls}
            setValue={(i) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                BaseOnBalls: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'baseOnBallsThresholdType'}
            initValue={props.searchPitcherGradesCondition.BaseOnBallsThresholdType}
            setThresholdType={(value) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                BaseOnBallsThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'被本塁打数：'} />
          <InputNumberText
            name={'homeRun'}
            minValue={0}
            maxValue={999}
            maxLength={3}
            height={0}
            width={50}
            placeholder={'被本塁打数'}
            value={props.searchPitcherGradesCondition.HomeRun}
            setValue={(i) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                HomeRun: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'homeRunThresholdType'}
            initValue={props.searchPitcherGradesCondition.HomeRunThresholdType}
            setThresholdType={(value) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                HomeRunThresholdType: value,
              })
            }
          />
        </Grid>
      </Grid>
      <Grid item xs={12} display={'flex'}>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'完投数：'} />
          <InputNumberText
            name={'completeGame'}
            minValue={0}
            maxValue={99}
            maxLength={2}
            height={0}
            width={50}
            placeholder={'完投数'}
            value={props.searchPitcherGradesCondition.CompleteGame}
            setValue={(i) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                CompleteGame: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'completeGameThresholdType'}
            initValue={props.searchPitcherGradesCondition.CompleteGameThresholdType}
            setThresholdType={(value) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                CompleteGameThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'完封数：'} />
          <InputNumberText
            name={'shutout'}
            minValue={0}
            maxValue={99}
            maxLength={2}
            height={0}
            width={50}
            placeholder={'完封数'}
            value={props.searchPitcherGradesCondition.Shutout}
            setValue={(i) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                Shutout: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'shutoutThresholdType'}
            initValue={props.searchPitcherGradesCondition.ShutoutThresholdType}
            setThresholdType={(value) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                ShutoutThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'勝率数：'} />
          <InputFloatText
            name={'winningRate'}
            minValue={0}
            maxValue={1.0}
            maxLength={4}
            height={0}
            width={50}
            placeholder={'勝率数'}
            value={props.searchPitcherGradesCondition.WinningRate}
            setValue={(i) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                WinningRate: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'winningRateThresholdType'}
            initValue={props.searchPitcherGradesCondition.WinningRateThresholdType}
            setThresholdType={(value) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                WinningRateThresholdType: value,
              })
            }
          />
        </Grid>
      </Grid>
      <Grid item xs={12} display={'flex'}>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'勝利数：'} />
          <InputNumberText
            name={'win'}
            minValue={0}
            maxValue={999}
            maxLength={3}
            height={0}
            width={50}
            placeholder={'勝利数'}
            value={props.searchPitcherGradesCondition.Win}
            setValue={(i) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                Win: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'winThresholdType'}
            initValue={props.searchPitcherGradesCondition.WinThresholdType}
            setThresholdType={(value) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                WinThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'敗北数：'} />
          <InputNumberText
            name={'lose'}
            minValue={0}
            maxValue={999}
            maxLength={3}
            height={0}
            width={50}
            placeholder={'敗北数'}
            value={props.searchPitcherGradesCondition.Lose}
            setValue={(i) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                Lose: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'loseThresholdType'}
            initValue={props.searchPitcherGradesCondition.LoseThresholdType}
            setThresholdType={(value) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                LoseThresholdType: value,
              })
            }
          />
        </Grid>
      </Grid>
      <Grid item xs={12} display={'flex'}>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'セーブ数：'} />
          <InputNumberText
            name={'save'}
            minValue={0}
            maxValue={999}
            maxLength={3}
            height={0}
            width={50}
            placeholder={'セーブ数'}
            value={props.searchPitcherGradesCondition.Save}
            setValue={(i) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                Save: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'saveThresholdType'}
            initValue={props.searchPitcherGradesCondition.SaveThresholdType}
            setThresholdType={(value) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                SaveThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'ホールド数：'} />
          <InputNumberText
            name={'hold'}
            minValue={0}
            maxValue={999}
            maxLength={3}
            height={0}
            width={50}
            placeholder={'ホールド数'}
            value={props.searchPitcherGradesCondition.Hold}
            setValue={(i) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                Hold: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'holdThresholdType'}
            initValue={props.searchPitcherGradesCondition.HoldThresholdType}
            setThresholdType={(value) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                HoldThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'HP数：'} />
          <InputNumberText
            name={'holdPoint'}
            minValue={0}
            maxValue={999}
            maxLength={3}
            height={0}
            width={50}
            placeholder={'HP数'}
            value={props.searchPitcherGradesCondition.HoldPoint}
            setValue={(i) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                HoldPoint: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'holdPointThresholdType'}
            initValue={props.searchPitcherGradesCondition.HoldPointThresholdType}
            setThresholdType={(value) =>
              props.setSearchPitcherGradesCondition({
                ...props.searchPitcherGradesCondition,
                HoldPointThresholdType: value,
              })
            }
          />
        </Grid>
      </Grid>
    </React.Fragment>
  );
}
