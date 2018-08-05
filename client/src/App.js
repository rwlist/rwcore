import React, { Component } from 'react';
import Lists from './lists/Lists';
import './App.css';

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            name: "default",
            show: false,
        }
        
        this.onChange = this.onChange.bind(this);
        this.onClick = this.onClick.bind(this);
    }

    onChange(e) {
        this.setState({ name: e.target.value });
    }

    onClick() {
        this.setState({ show: true });
    }

    render() {
        if (this.state.show) {
            return (
                <div className="App">
                    <Lists name={this.state.name}/>
                </div>
            );
        }

        return (
            <div className="App">
                <input type="text" onChange={this.onChange}/>
                <button onClick={this.onClick}>Go</button>
            </div>
        )
    }
}

export default App;
