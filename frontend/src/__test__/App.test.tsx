import Enzyme, { mount } from 'enzyme';
import App from '../App';
import { MemoryRouter } from 'react-router';
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import SeasonPage from '../components/pages/SeasonPage';
import HomePage from '../components/pages/HomePage';
import BattingPage from '../components/pages/BattingPage';
import PitchingPage from '../components/pages/PitchingPage';
import PlayersPage from '../components/pages/PlayersPage';
import PlayerPage from '../components/pages/PlayerPage';
import ManagerPage from '../components/pages/ManagerPage';
import * as teamPitchingModule from '../data/api/teamPitching';
import * as teamBattingModule from '../data/api/teamBatting';
import * as teamStatsModule from '../data/api/teamStats';
import * as teamCareersModule from '../data/api/teamCareers';
import * as playerModule from '../data/api/player';
import {
  MaxTeamBattingResponse,
  MaxTeamPitchingResponse,
  MinTeamBattingResponse,
  MinTeamPitchingResponse,
} from '../data/type';

//　Enzymeの設定
Enzyme.configure({
  adapter: new Adapter(),
});

describe('ルーティング', () => {
  let getMaxTeamPitching: jest.SpyInstance<Promise<any>>;
  let getMinTeamPitching: jest.SpyInstance<Promise<any>>;
  let getMaxTeamBatting: jest.SpyInstance<Promise<any>>;
  let getMinTeamBatting: jest.SpyInstance<Promise<any>>;
  let getTeamStatsByYear: jest.SpyInstance<Promise<any>>;
  let getTeamBattingByYear: jest.SpyInstance<Promise<any>>;
  let getTeamPitchingByYear: jest.SpyInstance<Promise<any>>;
  let getCareers: jest.SpyInstance<Promise<any>>;
  let getPlayer: jest.SpyInstance<Promise<any>>;
  beforeEach(() => {
    getMaxTeamPitching = jest.spyOn(teamPitchingModule, 'getMaxTeamPitching').mockReturnValueOnce(
      new Promise<MaxTeamPitchingResponse>(() => {
        return { maxStrikeOutRate: 0, maxRunsAllowed: 0 };
      })
    );
    getMinTeamPitching = jest.spyOn(teamPitchingModule, 'getMinTeamPitching').mockReturnValueOnce(
      new Promise<MinTeamPitchingResponse>(() => {
        return { minStrikeOutRate: 0, minRunsAllowed: 0 };
      })
    );
    getMaxTeamBatting = jest.spyOn(teamBattingModule, 'getMaxTeamBatting').mockReturnValueOnce(
      new Promise<MaxTeamBattingResponse>(() => {
        return { maxHomeRun: 0, maxSluggingPercentage: 0, maxOnBasePercentage: 0 };
      })
    );
    getMinTeamBatting = jest.spyOn(teamBattingModule, 'getMinTeamBatting').mockReturnValueOnce(
      new Promise<MinTeamBattingResponse>(() => {
        return { minHomeRun: 0, minSluggingPercentage: 0, minOnBasePercentage: 0 };
      })
    );
    getTeamStatsByYear = jest.spyOn(teamStatsModule, 'getTeamStatsByYear').mockReturnValueOnce(
      new Promise<any>(() => {
        return {};
      })
    );
    getTeamBattingByYear = jest
      .spyOn(teamBattingModule, 'getTeamBattingByYear')
      .mockReturnValueOnce(
        new Promise<any>(() => {
          return {};
        })
      );
    getTeamPitchingByYear = jest
      .spyOn(teamPitchingModule, 'getTeamPitchingByYear')
      .mockReturnValueOnce(
        new Promise<any>(() => {
          return {};
        })
      );
    getCareers = jest.spyOn(teamCareersModule, 'getCareers').mockReturnValueOnce(
      new Promise<any>(() => {
        return {};
      })
    );
    getPlayer = jest.spyOn(playerModule, 'getPlayer').mockReturnValueOnce(
      new Promise<any>(() => {
        return {};
      })
    );
  });

  it('シーズン成績ページ', () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={['/season']}>
        <App />
      </MemoryRouter>
    );
    expect(wrapper.find(SeasonPage)).toHaveLength(1);
    expect(getTeamStatsByYear).toHaveBeenCalled();
    expect(getMaxTeamPitching).toHaveBeenCalled();
    expect(getMinTeamPitching).toHaveBeenCalled();
    expect(getMaxTeamBatting).toHaveBeenCalled();
    expect(getMinTeamBatting).toHaveBeenCalled();
  });

  it('トップページ', () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={['/']}>
        <App />
      </MemoryRouter>
    );
    expect(wrapper.find(HomePage)).toHaveLength(1);

    expect(getMaxTeamPitching).toHaveBeenCalled();
    expect(getMinTeamPitching).toHaveBeenCalled();
    expect(getMaxTeamBatting).toHaveBeenCalled();
    expect(getMinTeamBatting).toHaveBeenCalled();
  });

  it('打撃成績ページ', () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={['/batting']}>
        <App />
      </MemoryRouter>
    );
    expect(wrapper.find(BattingPage)).toHaveLength(1);
    expect(getTeamBattingByYear).toHaveBeenCalled();
    expect(getMaxTeamPitching).toHaveBeenCalled();
    expect(getMinTeamPitching).toHaveBeenCalled();
    expect(getMaxTeamBatting).toHaveBeenCalled();
    expect(getMinTeamBatting).toHaveBeenCalled();
  });

  it('投手成績ページ', () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={['/pitching']}>
        <App />
      </MemoryRouter>
    );
    expect(wrapper.find(PitchingPage)).toHaveLength(1);
    expect(getTeamPitchingByYear).toHaveBeenCalled();
    expect(getMaxTeamPitching).toHaveBeenCalled();
    expect(getMinTeamPitching).toHaveBeenCalled();
    expect(getMaxTeamBatting).toHaveBeenCalled();
    expect(getMinTeamBatting).toHaveBeenCalled();
  });

  it('選手一覧ページ', () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={['/players']}>
        <App />
      </MemoryRouter>
    );
    expect(wrapper.find(PlayersPage)).toHaveLength(1);
    expect(getCareers).toHaveBeenCalled();
    expect(getMaxTeamPitching).toHaveBeenCalled();
    expect(getMinTeamPitching).toHaveBeenCalled();
    expect(getMaxTeamBatting).toHaveBeenCalled();
    expect(getMinTeamBatting).toHaveBeenCalled();
  });

  it('選手個人ページ', () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={['/player/00000001']}>
        <App />
      </MemoryRouter>
    );
    expect(wrapper.find(PlayerPage)).toHaveLength(1);
    expect(getPlayer).toHaveBeenCalled();
    expect(getMaxTeamPitching).toHaveBeenCalled();
    expect(getMinTeamPitching).toHaveBeenCalled();
    expect(getMaxTeamBatting).toHaveBeenCalled();
    expect(getMinTeamBatting).toHaveBeenCalled();
  });

  it('監督成績ページ', () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={['/manager']}>
        <App />
      </MemoryRouter>
    );
    expect(wrapper.find(ManagerPage)).toHaveLength(1);
    expect(getTeamStatsByYear).toHaveBeenCalled();
    expect(getMaxTeamPitching).toHaveBeenCalled();
    expect(getMinTeamPitching).toHaveBeenCalled();
    expect(getMaxTeamBatting).toHaveBeenCalled();
    expect(getMinTeamBatting).toHaveBeenCalled();
  });
});
