import React from 'react';
import Enzyme from 'enzyme';
import renderer from 'react-test-renderer';
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import Chart from '../../../components/common/Chart';

Enzyme.configure({ adapter: new Adapter() });

describe('Chartテスト', () => {
  it('スナップショット作成', () => {
    const testData = [
      { x: 'test1', y: 1 },
      { x: 'test2', y: 2 },
    ];
    const testLineData = [
      ['test1', '#BAD3FF'],
      ['test2', '#FFD700'],
    ];
    const tree = renderer.create(
      <Chart title={'タイトル'} data={testData} label={'ラベル'} lineData={testLineData} />
    );
    expect(tree).toMatchSnapshot();
  });
});
