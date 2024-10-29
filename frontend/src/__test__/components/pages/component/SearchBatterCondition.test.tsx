import React from 'react';
import Enzyme from 'enzyme';
import renderer from 'react-test-renderer';
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import { MemoryRouter } from 'react-router-dom';
import { SearchBatterCondition } from '../../../../components/pages/component/SearchBatterCondition';
import { SearchBatterGradesCondition, ThresholdType } from '../../../../data/type';

Enzyme.configure({ adapter: new Adapter() });

describe('戦力ページテスト', () => {
  it('スナップショット作成', () => {
    const tree = renderer
      .create(
        <MemoryRouter>
          <SearchBatterCondition
            searchBatterGradesCondition={{
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
            }}
            setSearchBatterGradesCondition={function (
              value: React.SetStateAction<SearchBatterGradesCondition>
            ): void {
              throw new Error('Function not implemented.');
            }}
          />
        </MemoryRouter>
      )
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
