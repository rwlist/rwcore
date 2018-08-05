import React, { Component } from 'react';

class Spoiler extends Component {
    constructor(props) {
        super(props);
        this.state = {
            hidden: true
        }

        this.handleClick = this.handleClick.bind(this);
    }

    handleClick() {
        this.setState({
            hidden: !this.state.hidden
        });
    }

    render() {
        return (
            <div>
                <a href="#" onClick={this.handleClick}>{this.props.text}</a>
                {!this.state.hidden && this.props.children}
            </div>
        )
    }
}

export default Spoiler;