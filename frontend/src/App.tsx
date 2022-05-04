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
import axios from 'axios';

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

export class MaxTeamBattingResponse {
  maxHomeRun: number;
  maxSluggingPercentage: number;
  maxOnBasePercentage: number;
  constructor(maxHomeRun: number, maxSluggingPercentage: number, maxOnBasePercentage: number) {
    this.maxHomeRun = maxHomeRun;
    this.maxSluggingPercentage = maxSluggingPercentage;
    this.maxOnBasePercentage = maxOnBasePercentage;
  }
}

export class MaxTeamPitchingResponse {
  maxStrikeOutRate: number;
  maxRunsAllowed: number;
  constructor(maxStrikeOutRate: number, maxRunsAllowed: number) {
    this.maxStrikeOutRate = maxStrikeOutRate;
    this.maxRunsAllowed = maxRunsAllowed;
  }
}

export class MinTeamBattingResponse {
  minHomeRun: number;
  minSluggingPercentage: number;
  minOnBasePercentage: number;
  constructor(minHomeRun: number, minSluggingPercentage: number, minOnBasePercentage: number) {
    this.minHomeRun = minHomeRun;
    this.minSluggingPercentage = minSluggingPercentage;
    this.minOnBasePercentage = minOnBasePercentage;
  }
}

export class MinTeamPitchingResponse {
  minStrikeOutRate: number;
  minRunsAllowed: number;
  constructor(minStrikeOutRate: number, minRunsAllowed: number) {
    this.minStrikeOutRate = minStrikeOutRate;
    this.minRunsAllowed = minRunsAllowed;
  }
}

function App() {
  const [maxTeamPitching, setMaxTeamPitching] = useState<MaxTeamPitchingResponse>(
    new MaxTeamPitchingResponse(0, 0)
  );
  const [minTeamPitching, setMinTeamPitching] = useState<MinTeamPitchingResponse>(
    new MinTeamPitchingResponse(0, 0)
  );
  const [maxTeamBatting, setMaxTeamBatting] = useState<MaxTeamBattingResponse>(
    new MaxTeamBattingResponse(0, 0, 0)
  );
  const [minTeamBatting, setMinTeamBatting] = useState<MinTeamBattingResponse>(
    new MinTeamBattingResponse(0, 0, 0)
  );
  useEffect(() => {
    (async () => {
      axios
        .get<MaxTeamPitchingResponse>(`http://localhost:8081/team/pitching/max`)
        .then((response) => {
          setMaxTeamPitching(response.data);
        });
      axios
        .get<MinTeamPitchingResponse>(`http://localhost:8081/team/pitching/min`)
        .then((response) => {
          setMinTeamPitching(response.data);
        });
      axios
        .get<MaxTeamBattingResponse>(`http://localhost:8081/team/batting/max`)
        .then((response) => {
          setMaxTeamBatting(response.data);
        });
      axios
        .get<MinTeamBattingResponse>(`http://localhost:8081/team/batting/min`)
        .then((response) => {
          setMinTeamBatting(response.data);
        });
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
