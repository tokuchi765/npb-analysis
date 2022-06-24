import React from 'react';
import Enzyme from 'enzyme';
import renderer from 'react-test-renderer';
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import StrengthPage from '../../../components/pages/StrengthPage';
import { MemoryRouter } from 'react-router-dom';

Enzyme.configure({ adapter: new Adapter() });

describe('戦力ページテスト', () => {
  it('スナップショット作成', () => {
    const tree = renderer
      .create(
        <MemoryRouter>
          <StrengthPage
            years={['2020']}
            initYear={'2020'}
            maxTeamPitching={{ maxStrikeOutRate: 0, maxRunsAllowed: 0 }}
            minTeamPitching={{ minStrikeOutRate: 0, minRunsAllowed: 0 }}
            maxTeamBatting={{ maxHomeRun: 0, maxSluggingPercentage: 0, maxOnBasePercentage: 0 }}
            minTeamBatting={{ minHomeRun: 0, minSluggingPercentage: 0, minOnBasePercentage: 0 }}
          />
        </MemoryRouter>
      )
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
