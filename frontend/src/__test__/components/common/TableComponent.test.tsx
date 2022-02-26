import React from 'react';
import Enzyme from 'enzyme';
import renderer from 'react-test-renderer';
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import {
  TableComponent,
  TableLinkComponent,
  HeadCell,
  SelectItem,
} from '../../../components/common/TableComponent';
import { MemoryRouter } from 'react-router-dom';

Enzyme.configure({ adapter: new Adapter() });

describe('TableComponentテスト', () => {
  it('スナップショット作成', () => {
    const headCells: HeadCell[] = [
      { id: 'main', numeric: false, disablePadding: true, label: 'チーム名' },
      { id: 'battingAverage', numeric: true, disablePadding: false, label: '打率' },
    ];
    interface TestData {
      main: string;
      battingAverage: number;
    }

    const testData: TestData = { main: 'メイン', battingAverage: 0.3 };
    const tree = renderer
      .create(
        <TableComponent
          title={'タイトル'}
          setSelect={() => {}}
          getDataList={() => {}}
          datas={[testData]}
          headCells={headCells}
          initSorted={'初期ソート'}
          selectItems={[new SelectItem('選択1', 'ラベル名', ['選択1', '選択2'])]}
        />
      )
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});

describe('TableLinkComponentテスト', () => {
  it('スナップショット作成', () => {
    const headCells: HeadCell[] = [
      { id: 'main', numeric: false, disablePadding: true, label: 'チーム名' },
      { id: 'battingAverage', numeric: true, disablePadding: false, label: '打率' },
    ];
    interface TestData {
      main: string;
      battingAverage: number;
    }

    const testData: TestData = { main: 'メイン', battingAverage: 0.3 };
    const playerIdMap: Map<string, string> = new Map<string, string>();
    playerIdMap.set('key', 'value');
    const tree = renderer
      .create(
        <MemoryRouter>
          <TableLinkComponent
            title={'タイトル'}
            setSelect={() => {}}
            getDataList={() => {}}
            datas={[testData]}
            headCells={headCells}
            initSorted={'初期ソート'}
            selectItems={[new SelectItem('選択1', 'ラベル名', ['選択1', '選択2'])]}
            linkValues={playerIdMap}
            path={'/path/'}
          />
        </MemoryRouter>
      )
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
