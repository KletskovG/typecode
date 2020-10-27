import React, { useEffect } from 'react';
import { Switch, Route, Redirect } from 'wouter';
import HomePage from './components/pages/Home/HomePage';
import Frontend from './components/pages/Frontend/FrontendPage';
import DevOpsPage from './components/pages/DevOps/DevOpsPage';
import RavesPage from './components/pages/Raves/RavesPage';

function App() {
  
  return (
    <>
      <Switch>
        <Route path="/" component={() => <HomePage />} >  </Route>
        <Route path="/frontend" component={() => <Frontend />}></Route>
        <Route path="/devops" component={() => <DevOpsPage/> }></Route>
        <Route path="/raves" component={() => <RavesPage />}></Route>
        <Route path="/:rest*">
          <Redirect to={'/'} />
        </Route>
      </Switch>
    </>
  );
}

export default App;