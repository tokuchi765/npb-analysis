import React from 'react';
import Enzyme from 'enzyme';
import renderer from 'react-test-renderer';
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import RadarChartComponent from '../../../components/common/RadarChartComponent';
import { MemoryRouter } from 'react-router-dom';

Enzyme.configure({ adapter: new Adapter() });

const data = [
  {
    item: '本塁打',
    A: 50,
    B: 80,
    fullMark: 100,
  },
];

describe('レーダーチャートテスト', () => {
  it('スナップショット作成', () => {
    const tree = renderer
      .create(
        <MemoryRouter>
          <RadarChartComponent
            title="戦力チャート"
            data={data}
            nameA={'teamA'}
            keyA={'A'}
            nameB={'teamB'}
            keyB={'B'}
            help={'注釈'}
          />
        </MemoryRouter>
      )
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
