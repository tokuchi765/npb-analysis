import React from 'react';
import Enzyme from 'enzyme';
import renderer from 'react-test-renderer';
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import PlayerPage from '../../../components/pages/PlayerPage';
import { MemoryRouter } from 'react-router-dom';

Enzyme.configure({ adapter: new Adapter() });

describe('選手詳細ページテスト', () => {
  it('スナップショット作成', () => {
    const tree = renderer
      .create(
        <MemoryRouter>
          <PlayerPage
            history={{} as any}
            location={{} as any}
            match={{ params: { id: 'test' }, isExact: true, path: 'path', url: 'url' }}
          />
        </MemoryRouter>
      )
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
