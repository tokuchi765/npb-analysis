import React, { useEffect, useState } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import SeasonPage from './components/pages/SeasonPage';
import HomePage from './components/pages/HomePage';
import BattingPage from './components/pages/BattingPage';
import PitchingPage from './components/pages/PitchingPage';
import PlayersPage from './components/pages/PlayersPage';
import PlayerPage from './components/pages/PlayerPage';
import ManagerPage from './components/pages/ManagerPage';
import StrengthPage from './components/pages/StrengthPage';
import { getMaxTeamPitching, getMinTeamPitching } from './data/api/teamPitching';
import { getMaxTeamBatting, getMinTeamBatting } from './data/api/teamBatting';
import {
  MaxTeamPitchingResponse,
  MinTeamPitchingResponse,
  MaxTeamBattingResponse,
  MinTeamBattingResponse,
} from './data/type/index';

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

function App() {
  const [maxTeamPitching, setMaxTeamPitching] = useState<MaxTeamPitchingResponse>({
    maxStrikeOutRate: 0,
    maxRunsAllowed: 0,
  });
  const [minTeamPitching, setMinTeamPitching] = useState<MinTeamPitchingResponse>({
    minStrikeOutRate: 0,
    minRunsAllowed: 0,
  });
  const [maxTeamBatting, setMaxTeamBatting] = useState<MaxTeamBattingResponse>({
    maxHomeRun: 0,
    maxSluggingPercentage: 0,
    maxOnBasePercentage: 0,
  });
  const [minTeamBatting, setMinTeamBatting] = useState<MinTeamBattingResponse>({
    minHomeRun: 0,
    minSluggingPercentage: 0,
    minOnBasePercentage: 0,
  });
  useEffect(() => {
    (async () => {
      const responses = await Promise.all([
        getMaxTeamPitching(),
        getMinTeamPitching(),
        getMaxTeamBatting(),
        getMinTeamBatting(),
      ]);

      setMaxTeamPitching(responses[0]);
      setMinTeamPitching(responses[1]);
      setMaxTeamBatting(responses[2]);
      setMinTeamBatting(responses[3]);
    })();
  }, []);
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
          render={() => (
            <StrengthPage
              years={years}
              initYear={initYear}
              maxTeamPitching={maxTeamPitching}
              minTeamPitching={minTeamPitching}
              maxTeamBatting={maxTeamBatting}
              minTeamBatting={minTeamBatting}
            />
          )}
          exact
        />
        <Route path="/players" component={PlayersPage} exact />
        <Route path="/player/:id" component={PlayerPage} exact />
        <Route path="/manager" component={ManagerPage} exact />
      </Switch>
    </Router>
  );
}

export default App;
