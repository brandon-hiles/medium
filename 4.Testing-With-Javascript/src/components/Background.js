import React from 'react';

export default class Background extends React.Component {
    
    constructor(props) {
        super(props)
        this.state = {
            color: "white"
        };

        this.handleChange = this.handleChange.bind(this);
    }



    handleChange(evt) {
        // runs on every keystroke event
        this.setState({
          [evt.target.name]: evt.target.value
        })
      }
    

    render() {

        const StyleObj = {
            background: this.state.color,
            height: "100vh"
        }

        const formObj = {
            position: "absolute",
            left: "50%",
            marginLeft: "-50px"
        }
        return (
            <div style={StyleObj} className="Background-Test">
                <form style={formObj}>
                    <input type="text"
                        name="color"
                        placeholder="Enter your color"
                        onChange={this.handleChange}
                    />
                </form>
            </div>
        )
    }
}