import { Grid } from '@mui/material';
import React from 'react';
import { SearchBatterGradesCondition } from '../../../data/type';
import {
  InputFloatText,
  InputNumberText,
  SearchLabel,
  ThresholdRadioGroup,
} from '../SearchGradesPage';

export interface SearchBatterConditionProps {
  searchBatterGradesCondition: SearchBatterGradesCondition;
  setSearchBatterGradesCondition: React.Dispatch<React.SetStateAction<SearchBatterGradesCondition>>;
}

export function SearchBatterCondition(props: SearchBatterConditionProps) {
  return (
    <React.Fragment>
      <Grid item xs={12} display={'flex'}>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'打席数：'} />
          <InputNumberText
            name={'plateAppearance'}
            minValue={0}
            maxValue={99999}
            maxLength={5}
            height={0}
            width={50}
            placeholder={'打席'}
            value={props.searchBatterGradesCondition.PlateAppearance}
            setValue={(i) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                PlateAppearance: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'plateAppearanceThresholdType'}
            initValue={props.searchBatterGradesCondition.PlateAppearanceThresholdType}
            setThresholdType={(value) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                PlateAppearanceThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'出塁率：'} />
          <InputFloatText
            name={'onBasePercentage'}
            minValue={0.0}
            maxValue={0.99}
            maxLength={4}
            height={0}
            width={50}
            placeholder={'出塁率'}
            value={props.searchBatterGradesCondition.OnBasePercentage}
            setValue={(i) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                OnBasePercentage: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'onBasePercentageThresholdType'}
            initValue={props.searchBatterGradesCondition.OnBasePercentageThresholdType}
            setThresholdType={(value) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                OnBasePercentageThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'打率：'} />
          <InputFloatText
            name={'battingAverage'}
            minValue={0.0}
            maxValue={0.99}
            maxLength={4}
            height={0}
            width={50}
            placeholder={'打率'}
            value={props.searchBatterGradesCondition.BattingAverage}
            setValue={(i) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                BattingAverage: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'battingAverageThresholdType'}
            initValue={props.searchBatterGradesCondition.BattingAverageThresholdType}
            setThresholdType={(value) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                BattingAverageThresholdType: value,
              })
            }
          />
        </Grid>
      </Grid>
      <Grid item xs={12} display={'flex'}>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'本塁打：'} />
          <InputNumberText
            name={'homeRun'}
            minValue={0}
            maxValue={999}
            maxLength={3}
            height={0}
            width={50}
            placeholder={'本塁打'}
            value={props.searchBatterGradesCondition.HomeRun}
            setValue={(i) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                HomeRun: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'homeRunThresholdType'}
            initValue={props.searchBatterGradesCondition.HomeRunThresholdType}
            setThresholdType={(value) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                HomeRunThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'三振：'} />
          <InputNumberText
            name={'strikeOut'}
            minValue={0}
            maxValue={9999}
            maxLength={4}
            height={0}
            width={50}
            placeholder={'三振'}
            value={props.searchBatterGradesCondition.StrikeOut}
            setValue={(i) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                StrikeOut: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'strikeOutThresholdType'}
            initValue={props.searchBatterGradesCondition.StrikeOutThresholdType}
            setThresholdType={(value) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                StrikeOutThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'三振率：'} />
          <InputFloatText
            name={'strikeOutRate'}
            minValue={0}
            maxValue={0.99}
            maxLength={4}
            height={0}
            width={50}
            placeholder={'三振率'}
            value={props.searchBatterGradesCondition.StrikeOutRate}
            setValue={(i) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                StrikeOutRate: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'strikeOutRateThresholdType'}
            initValue={props.searchBatterGradesCondition.StrikeOutRateThresholdType}
            setThresholdType={(value) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                StrikeOutRateThresholdType: value,
              })
            }
          />
        </Grid>
      </Grid>
      <Grid item xs={12} display={'flex'}>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'加重出塁率：'} />
          <InputFloatText
            name={'wOba'}
            minValue={0}
            maxValue={0.99}
            maxLength={4}
            height={0}
            width={50}
            placeholder={'加重出塁率'}
            value={props.searchBatterGradesCondition.WOba}
            setValue={(i) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                WOba: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'wObaThresholdType'}
            initValue={props.searchBatterGradesCondition.WObaThresholdType}
            setThresholdType={(value) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                WObaThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'創出得点：'} />
          <InputNumberText
            name={'rC'}
            minValue={0}
            maxValue={9999}
            maxLength={4}
            height={0}
            width={50}
            placeholder={'創出得点'}
            value={props.searchBatterGradesCondition.RC}
            setValue={(i) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                RC: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'rCThresholdType'}
            initValue={props.searchBatterGradesCondition.RCThresholdType}
            setThresholdType={(value) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                RCThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'BABIP：'} />
          <InputFloatText
            name={'babip'}
            minValue={0}
            maxValue={0.99}
            maxLength={4}
            height={0}
            width={50}
            placeholder={'BABIP'}
            value={props.searchBatterGradesCondition.Babip}
            setValue={(i) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                Babip: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'babipThresholdType'}
            initValue={props.searchBatterGradesCondition.BabipThresholdType}
            setThresholdType={(value) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                BabipThresholdType: value,
              })
            }
          />
        </Grid>
      </Grid>
      <Grid item xs={12} display={'flex'}>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'併殺打：'} />
          <InputNumberText
            name={'groundedIntoDoublePlay'}
            minValue={0}
            maxValue={999}
            maxLength={3}
            height={0}
            width={50}
            placeholder={'併殺打'}
            value={props.searchBatterGradesCondition.GroundedIntoDoublePlay}
            setValue={(i) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                GroundedIntoDoublePlay: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'groundedIntoDoublePlayThresholdType'}
            initValue={props.searchBatterGradesCondition.GroundedIntoDoublePlayThresholdType}
            setThresholdType={(value) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                GroundedIntoDoublePlayThresholdType: value,
              })
            }
          />
        </Grid>
        <Grid item display={'flex'} alignItems="center">
          <SearchLabel label={'長打率：'} />
          <InputFloatText
            name={'sluggingPercentage'}
            minValue={0.0}
            maxValue={0.99}
            maxLength={4}
            height={0}
            width={50}
            placeholder={'長打率'}
            value={props.searchBatterGradesCondition.SluggingPercentage}
            setValue={(i) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                SluggingPercentage: i,
              })
            }
          />
          <ThresholdRadioGroup
            name={'sluggingPercentageThresholdType'}
            initValue={props.searchBatterGradesCondition.SluggingPercentageThresholdType}
            setThresholdType={(value) =>
              props.setSearchBatterGradesCondition({
                ...props.searchBatterGradesCondition,
                SluggingPercentageThresholdType: value,
              })
            }
          />
        </Grid>
      </Grid>
    </React.Fragment>
  );
}
