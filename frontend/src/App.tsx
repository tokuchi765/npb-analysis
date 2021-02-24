import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import SeasonPage from './components/pages/SeasonPage';
import HomePage from './components/pages/HomePage';
import BattingPage from './components/pages/BattingPage';
import PitchingPage from './components/pages/PitchingPage';
import PlayersPage from './components/pages/PlayersPage';
import PlayerPage from './components/pages/PlayerPage';

const App: React.FC = () => {
  return (
    <Router>
      <Switch>
        <Route path="/season" component={SeasonPage} exact />
        <Route path="/" component={HomePage} exact />
        <Route path="/batting" component={BattingPage} exact />
        <Route path="/pitching" component={PitchingPage} exact />
        <Route path="/players" component={PlayersPage} exact />
        <Route path="/player/:id" component={PlayerPage} exact />
      </Switch>
    </Router>
  );
};

export default App;
