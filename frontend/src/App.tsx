import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import SeasonPage from './components/pages/SeasonPage';
import HomePage from './components/pages/HomePage';
import BattingPage from './components/pages/BattingPage';
import PitchingPage from './components/pages/PitchingPage';
import PlayersPage from './components/pages/PlayersPage';
import PlayerPage from './components/pages/PlayerPage';
import ManagerPage from './components/pages/ManagerPage';
import StrengthPage from './components/pages/StrengthPage';

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
  '2021',
];

const initYear = '2021';

const App: React.FC = () => {
  return (
    <Router>
      <Switch>
        <Route
          path="/season"
          render={() => <SeasonPage years={years} initYear={initYear} />}
          exact
        />
        <Route path="/" render={() => <HomePage years={years} />} exact />
        <Route
          path="/batting"
          render={() => <BattingPage years={years} initYear={initYear} />}
          exact
        />
        <Route
          path="/pitching"
          render={() => <PitchingPage years={years} initYear={initYear} />}
          exact
        />
        <Route
          path="/strength"
          render={() => <StrengthPage years={years} initYear={initYear} />}
          exact
        />
        <Route path="/players" component={PlayersPage} exact />
        <Route path="/player/:id" component={PlayerPage} exact />
        <Route path="/manager" component={ManagerPage} exact />
      </Switch>
    </Router>
  );
};

export default App;
