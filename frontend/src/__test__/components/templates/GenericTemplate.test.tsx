import React from 'react';
import Enzyme from 'enzyme';
import renderer from 'react-test-renderer';
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import GenericTemplate from '../../../components/templates/GenericTemplate';
import { MemoryRouter } from 'react-router-dom';

Enzyme.configure({ adapter: new Adapter() });

describe('GenericTemplateテスト', () => {
  it('スナップショット作成', () => {
    const tree = renderer
      .create(
        <MemoryRouter>
          <GenericTemplate title="タイトル">
            <div>テスト</div>
          </GenericTemplate>
        </MemoryRouter>
      )
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
