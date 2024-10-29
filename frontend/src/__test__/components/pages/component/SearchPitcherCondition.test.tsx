import React from 'react';
import Enzyme from 'enzyme';
import renderer from 'react-test-renderer';
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import { MemoryRouter } from 'react-router-dom';
import { SearchPitcherCondition } from '../../../../components/pages/component/SearchPitcherCondition';
import { SearchPitcherGradesCondition, ThresholdType } from '../../../../data/type';

Enzyme.configure({ adapter: new Adapter() });

describe('戦力ページテスト', () => {
  it('スナップショット作成', () => {
    const tree = renderer
      .create(
        <MemoryRouter>
          <SearchPitcherCondition
            searchPitcherGradesCondition={{
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
            }}
            setSearchPitcherGradesCondition={function (
              value: React.SetStateAction<SearchPitcherGradesCondition>
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
