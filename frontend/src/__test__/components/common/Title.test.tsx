import React from 'react';
import Enzyme from 'enzyme';
import renderer from 'react-test-renderer';
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import Title from '../../../components/common/Title';

Enzyme.configure({ adapter: new Adapter() });

describe('Titleコンポーネントテスト', () => {
  it('スナップショット作成', () => {
    const tree = renderer.create(<Title children={'タイトル'} />).toJSON();
    expect(tree).toMatchSnapshot();
  });
});
