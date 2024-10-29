import React from 'react';
import Enzyme from 'enzyme';
import renderer from 'react-test-renderer';
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import { match, MemoryRouter, RouteComponentProps } from 'react-router-dom';
import SearchGradesPage, {
  PeriodType,
  SearchGrades,
  SearchGradesCondition,
  Type,
} from '../../../components/pages/SearchGradesPage';
import * as H from 'history';
import { ThresholdType } from '../../../data/type';

Enzyme.configure({ adapter: new Adapter() });

describe('戦力ページテスト', () => {
  it('スナップショット作成', () => {
    const location: H.Location<SearchGrades> = {
      pathname: '',
      search: '',
      state: {
        baseCondition: {
          Total: false,
          TotalYear: undefined,
          FromYear: undefined,
          ToYear: undefined,
          TeamID: '',
        },
        batterCondition: {
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
        },
        pitcherCondition: {
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
        },
        type: Type.Batter,
        periodType: PeriodType.Period,
      },
      hash: '',
    };
    const history: H.History<SearchGrades> = {
      length: 0,
      action: 'PUSH',
      location: location,
      push: function (
        location: H.LocationDescriptor<SearchGrades>,
        state?: SearchGrades | undefined
      ): void {
        throw new Error('Function not implemented.');
      },
      replace: function (
        location: H.LocationDescriptor<SearchGrades>,
        state?: SearchGrades | undefined
      ): void {
        throw new Error('Function not implemented.');
      },
      go: function (n: number): void {
        throw new Error('Function not implemented.');
      },
      goBack: function (): void {
        throw new Error('Function not implemented.');
      },
      goForward: function (): void {
        throw new Error('Function not implemented.');
      },
      block: function (
        prompt?: string | boolean | H.TransitionPromptHook<SearchGrades> | undefined
      ): H.UnregisterCallback {
        throw new Error('Function not implemented.');
      },
      listen: function (listener: H.LocationListener<SearchGrades>): H.UnregisterCallback {
        throw new Error('Function not implemented.');
      },
      createHref: function (location: H.LocationDescriptorObject<SearchGrades>): H.Href {
        throw new Error('Function not implemented.');
      },
    };
    const condition: match<{}> = {
      params: {},
      isExact: false,
      path: '',
      url: '',
    };
    const tree = renderer
      .create(
        <MemoryRouter>
          <SearchGradesPage history={history} location={location} match={condition} />
        </MemoryRouter>
      )
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
