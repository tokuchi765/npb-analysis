import Enzyme, { render, mount } from 'enzyme';
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

// このテストを実行するとエラーが起こる
// どうやらmaterial-uiのバグっぽい
// https://stackoverflow.com/questions/58070996/how-to-fix-the-warning-uselayouteffect-does-nothing-on-the-server
test('text', () => {
  const wrapper = render(
    <MemoryRouter initialEntries={['/']}>
      <App />
    </MemoryRouter>
  );
  expect(wrapper.text()).toBe(
    '管理画面トップページチーム成績ページ打撃成績ページ投手成績ページ選手一覧' +
      'ページ監督ページトップページ（セ）チーム打率推移0auto打率GiantsBaystars' +
      'TigersCarpDragonsSwallows（パ）チーム打率推移0auto打率LionsHawksEagles' +
      'MarinesFightersBuffaloesCopyright © 管理画面 2021.'
  );
});

// 以下のサイトを参考に実装
// https://medium.com/@antonybudianto/react-router-testing-with-jest-and-enzyme-17294fefd303
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
