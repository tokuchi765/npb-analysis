import { HeadCell } from '../common/TableComponent';

export interface PlayerData {
  main: string;
  position: string;
  height: string;
  draft: string;
  career: string;
}

export const CareerHeadCells: HeadCell[] = [
  { id: 'main', numeric: false, disablePadding: true, label: '選手名' },
  { id: 'position', numeric: false, disablePadding: true, label: 'ポジション' },
  { id: 'height', numeric: false, disablePadding: true, label: '身長' },
  { id: 'draft', numeric: false, disablePadding: true, label: 'ドラフト' },
  { id: 'career', numeric: false, disablePadding: true, label: '経歴' },
];

function createPlayerData(
  main: string,
  position: string,
  height: string,
  draft: string,
  career: string
) {
  const result: PlayerData = { main, position, height, draft, career };
  return result;
}

export function createPlayerDatas(
  careers: {
    Name: string;
    Position: string;
    Height: string;
    Draft: string;
    Career: string;
  }[]
) {
  const playerDateList: PlayerData[] = [];
  careers.forEach((career) => {
    playerDateList.push(
      createPlayerData(career.Name, career.Position, career.Height, career.Draft, career.Career)
    );
  });
  return playerDateList;
}

export function createPlayerIds(
  careers: {
    PlayerID: string;
    Name: string;
  }[]
) {
  const playerIdMap: Map<string, string> = new Map<string, string>();
  careers.forEach((career) => {
    playerIdMap.set(career.Name, career.PlayerID);
  });
  return playerIdMap;
}
