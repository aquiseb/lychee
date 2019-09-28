import React from "react";
import Form from "./Form";
import { Route, Router, Switch } from "./Router/index";

export const Routes = () => {
  return (
    <Router>
      <Switch>
        <Route exact path="/" component={Form} />
      </Switch>
    </Router>
  );
};