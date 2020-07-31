import React from 'react';
import { Link } from 'react-router-dom'

import '../sass/app.scss'

export default class Header extends React.Component {

    render() {
        return (
            <header id="public-header">
                <ul>
                    <li> <Link to="/"> Home </Link> </li>
                    <span className="right-links">
                        <li> <Link to="/property-change"> Test 1 </Link></li>
                        <li> <Link to="/form-validation"> Test 2 </Link></li>
                    </span>
                </ul>
            </header>
        )
    }
}