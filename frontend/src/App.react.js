import * as React from "react";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";

import HomePage from "./HomePage.react";
import EmployeeForm from "./EmployeeForm";

import "tabler-react/dist/Tabler.css";

type Props = {||};

function App(props: Props): React.Node {
  return (
    <React.Fragment>
      <Router>
        <Switch>
          <Route exact path="/" component={HomePage} />
          <Route exact path="/employee-add" component={EmployeeForm} />
          <Route exact path="/employee-list" component={HomePage} />
        </Switch>
      </Router>
    </React.Fragment>
  );
}

export default App;
