import React from 'react';
import { BrowserRouter as Router, Route } from "react-router-dom";

import Header from './Header';
import Background from './Background';
import Form from './Form';

export default class App extends React.Component {

    render() {
        return (
            <Router>
                <Header />
                <Route path="/property-change" component={() => <Background />} />
                <Route path="/form-validation" component={() => <Form />} />
            </Router>
        )
    }
}