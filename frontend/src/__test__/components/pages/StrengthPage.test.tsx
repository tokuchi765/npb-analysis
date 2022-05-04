import React from 'react';
import Enzyme from 'enzyme';
import renderer from 'react-test-renderer';
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import StrengthPage from '../../../components/pages/StrengthPage';
import { MemoryRouter } from 'react-router-dom';
import {
  MaxTeamBattingResponse,
  MaxTeamPitchingResponse,
  MinTeamBattingResponse,
  MinTeamPitchingResponse,
} from '../../../App';

Enzyme.configure({ adapter: new Adapter() });

describe('戦力ページテスト', () => {
  it('スナップショット作成', () => {
    const tree = renderer
      .create(
        <MemoryRouter>
          <StrengthPage
            years={['2020']}
            initYear={'2020'}
            maxTeamPitching={new MaxTeamPitchingResponse(0, 0)}
            minTeamPitching={new MinTeamPitchingResponse(0, 0)}
            maxTeamBatting={new MaxTeamBattingResponse(0, 0, 0)}
            minTeamBatting={new MinTeamBattingResponse(0, 0, 0)}
          />
        </MemoryRouter>
      )
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
