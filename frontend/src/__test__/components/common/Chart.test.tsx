import React from 'react';
import Enzyme from 'enzyme';
import renderer from 'react-test-renderer';
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import Chart, { ChartData } from '../../../components/common/Chart';

Enzyme.configure({ adapter: new Adapter() });

describe('Chartテスト', () => {
  it('スナップショット作成', () => {
    const testData = [
      { x: 'test1', y: 1 },
      { x: 'test2', y: 2 },
    ];
    const pacificChartDatas: ChartData[] = [
      { key: 'test1', name: 'test1', stroke: '#BAD3FF' },
      { key: 'test2', name: 'test2', stroke: '#FFD700' },
    ];
    const width = 400;
    const height = 300;
    const tree = renderer.create(
      <Chart
        title={'タイトル'}
        data={testData}
        label={'ラベル'}
        chartDatas={pacificChartDatas}
        width={width}
        height={height}
      />
    );
    expect(tree).toMatchSnapshot();
  });
});
