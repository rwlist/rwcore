import React, { Component } from 'react';

class FastArea extends Component {
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
                <textarea value={this.state.text} onChange={this.onChange}></textarea>
                <button onClick={this.handleClick}>Send</button>
            </div>
        )
    }
}

export default FastArea;