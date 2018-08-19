import React, { Component } from 'react';

class FastInput extends Component {
    constructor(props) {
        super(props);
        this.state = {
            text: ""
        }

        this.handleClick = this.handleClick.bind(this);
        this.onChange = this.onChange.bind(this);
    }

    handleClick() {
        this.props.handle(this.state.text);
    }

    onChange(e) {
        this.setState({ text: e.target.value });
    }

    render() {
        return (
            <div>
                <input type="text" value={this.state.text} onChange={this.onChange}></input>
                <button onClick={this.handleClick}>Send</button>
            </div>
        )
    }
}

export default FastInput;