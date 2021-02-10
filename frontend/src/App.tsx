import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import SeasonPage from './components/pages/SeasonPage';
import HomePage from './components/pages/HomePage';

const App: React.FC = () => {
  return (
    <Router>
      <Switch>
        <Route path="/season" component={SeasonPage} exact />
        <Route path="/" component={HomePage} exact />
      </Switch>
    </Router>
  );
};

export default App;
