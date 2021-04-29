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

//　Enzymeの設定
Enzyme.configure({
  adapter: new Adapter(),
});

describe('ルーティング', () => {
  it('シーズン成績ページ', () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={['/season']}>
        <App />
      </MemoryRouter>
    );
    expect(wrapper.find(SeasonPage)).toHaveLength(1);
  });

  it('トップページ', () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={['/']}>
        <App />
      </MemoryRouter>
    );
    expect(wrapper.find(HomePage)).toHaveLength(1);
  });

  it('打撃成績ページ', () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={['/batting']}>
        <App />
      </MemoryRouter>
    );
    expect(wrapper.find(BattingPage)).toHaveLength(1);
  });

  it('投手成績ページ', () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={['/pitching']}>
        <App />
      </MemoryRouter>
    );
    expect(wrapper.find(PitchingPage)).toHaveLength(1);
  });

  it('選手一覧ページ', () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={['/players']}>
        <App />
      </MemoryRouter>
    );
    expect(wrapper.find(PlayersPage)).toHaveLength(1);
  });

  it('選手個人ページ', () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={['/player/00000001']}>
        <App />
      </MemoryRouter>
    );
    expect(wrapper.find(PlayerPage)).toHaveLength(1);
  });

  it('選手個人ページ', () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={['/manager']}>
        <App />
      </MemoryRouter>
    );
    expect(wrapper.find(ManagerPage)).toHaveLength(1);
  });
});
