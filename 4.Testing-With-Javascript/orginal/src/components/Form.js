import React from 'react';

export default class Form extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            firstName: "",
            lastName: "",
            email: "",
            password: ""
        }

        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
    }

    handleChange() {
    // runs on every keystroke event
    this.setState({
        [evt.target.name]: evt.target.value
      })    
    }

    handleSubmit() {
        var b = "B"
    }

    render() {
        return (
            <div>
                <form>
                    <input type="text"
                        name="firstName"
                        placeholder="First Name"
                        value={this.state.firstName}
                        onChange={this.handleChange}
                    />
                    <input type="text"
                        name="email"
                        placeholder="Email"
                        value={this.state.email}
                        onChange={this.handleChange}
                    />                
                    <input type="password"
                        name="password"
                        placeholder="Password"
                        value={this.state.password}
                        onChange={this.handleChange}
                    />                    
                </form>
            </div>
        )
    }
}