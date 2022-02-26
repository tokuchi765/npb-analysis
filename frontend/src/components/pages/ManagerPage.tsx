import React, { useEffect, useState } from 'react';
import GenericTemplate from '../templates/GenericTemplate';
import axios from 'axios';
import _ from 'lodash';
import { TableComponent, HeadCell, SelectItem } from '../common/TableComponent';

const THREE = '(3年以上)';
const ALL = '(全て)';

const years = [
  '2005',
  '2006',
  '2007',
  '2008',
  '2009',
  '2010',
  '2011',
  '2012',
  '2013',
  '2014',
  '2015',
  '2016',
  '2017',
  '2018',
  '2019',
  '2020',
];

const selects = [ALL, THREE];

const headCells: HeadCell[] = [
  { id: 'main', numeric: false, disablePadding: true, label: '監督' },
  { id: 'years', numeric: true, disablePadding: true, label: '年数' },
  { id: 'winningRateAverage', numeric: true, disablePadding: true, label: '平均勝率' },
  {
    id: 'winningRateDifferenceAverage',
    numeric: true,
    disablePadding: true,
    label: 'ピタゴラス勝率との平均差分',
  },
];

interface ManagerData {
  winningRate: number;
  pythagoreanExpectation: number;
  winningRateDifference: number;
}

interface Manager {
  main: string;
  years: number;
  winningRateAverage: number;
  winningRateDifferenceAverage: number;
}

function createManagerAverage(select: string, managerMap: Map<string, Array<ManagerData>>) {
  const managerList: Manager[] = [];
  if (select === THREE) {
    managerMap.forEach((value: ManagerData[], key: string) => {
      if (value.length >= 3) {
        const winningRateDifferenceAverage = _.meanBy(value, 'winningRateDifference');
        const winningRateAverage = _.meanBy(value, 'winningRate');
        managerList.push({
          main: key,
          years: value.length,
          winningRateAverage: winningRateAverage,
          winningRateDifferenceAverage: winningRateDifferenceAverage,
        });
      }
    });
  } else {
    managerMap.forEach((value: ManagerData[], key: string) => {
      const average = _.meanBy(value, 'winningRateDifference');
      const winningRateAverage = _.meanBy(value, 'winningRate');
      managerList.push({
        main: key,
        years: value.length,
        winningRateAverage: winningRateAverage,
        winningRateDifferenceAverage: average,
      });
    });
  }

  return managerList;
}

function ManagerPage() {
  const [initSelect, setSelect] = useState<string>('');
  const [centralManager, setManager] = useState<Manager[]>([]);

  const getTeamDataList = async (select: string) => {
    const managerMap = new Map<string, Array<ManagerData>>();
    for (const year of years) {
      const result = await axios.get(
        `http://localhost:8081/team/stats?from_year=${year}&to_year=${year}`
      );

      const { teanStats } = result.data;
      const stats: Array<ManagerData> = teanStats[year];

      _.forEach(stats, (element: any) => {
        if (managerMap.has(element.Manager)) {
          managerMap.get(element.Manager)?.push({
            winningRate: element.WinningRate,
            pythagoreanExpectation: element.PythagoreanExpectation,
            winningRateDifference: element.WinningRate - element.PythagoreanExpectation,
          });
        } else {
          const array: Array<ManagerData> = [
            {
              winningRate: element.WinningRate,
              pythagoreanExpectation: element.PythagoreanExpectation,
              winningRateDifference: element.WinningRate - element.PythagoreanExpectation,
            },
          ];
          managerMap.set(element.Manager, array);
        }
      });
    }
    setSelect(select);
    setManager(createManagerAverage(select, managerMap));
  };

  useEffect(() => {
    (async () => {
      getTeamDataList(THREE);
    })();
  }, []);

  return (
    <GenericTemplate title="監督ページ">
      <TableComponent
        title={'ピタゴラス勝率'}
        setSelect={setSelect}
        getDataList={getTeamDataList}
        datas={centralManager}
        headCells={headCells}
        initSorted={'winningRateDifferenceAverage'}
        selectItems={[new SelectItem(initSelect, '選択', selects)]}
      />
    </GenericTemplate>
  );
}

export default ManagerPage;
