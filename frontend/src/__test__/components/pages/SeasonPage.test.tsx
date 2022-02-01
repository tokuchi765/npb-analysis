import React from 'react';
import Enzyme from 'enzyme';
import renderer from 'react-test-renderer';
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import SeasonPage from '../../../components/pages/SeasonPage';
import { MemoryRouter } from 'react-router-dom';

Enzyme.configure({ adapter: new Adapter() });

describe('チーム成績ページテスト', () => {
  it('スナップショット作成', () => {
    const tree = renderer
      .create(
        <MemoryRouter>
          <SeasonPage years={['2020']} />
        </MemoryRouter>
      )
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
