import React, { Component } from 'react';
import List from './List';
import './Lists.css';

class Lists extends Component {
    constructor(props) {
        super(props);
        this.state = {
            status: "loading"
        };
        
        this.onChange = this.onChange.bind(this);
        this.showList = this.showList.bind(this);
        this.load = this.load.bind(this);
    }

    componentDidMount() {
        this.load();
    }

    onDataReceived(lists) {
        if (this.state.status !== 'loading') {
            console.error('incorrect status');
            return;
        }
        this.setState({
            lists,
            status: 'loaded'
        })
    }

    load() {
        const handleErr = (err) => {
            this.onError(err);
            console.log('error while loading, do it again')
            this.load();
        }
        this.setState({ status: 'loading' })
        fetch('/lists', { method: 'GET' })
        .then(resp => resp.json())
        .then(it => {
            if (it.Err) {
                handleErr(it);
            } else {
                console.log(it);
                this.onDataReceived(it);
            }
        })
        .catch(handleErr)
    }

    onChange(e) {
        this.setState({ name: e.target.value });
    }

    showList(name) {
        this.setState({
            name,
            status: 'list'
        })
    }

    render() {
        if (this.state.status === 'list') {
            return (
                <div>
                    <List name={this.state.name}/>
                    Back to <a onClick={this.load}>lists page</a>
                </div>
            )
        }
        let elements;
        if (this.state.status === 'loading') {
            elements = (
                <div>
                    <h2>Loading...</h2>
                </div>
            )
        }
        if (this.state.status === 'loaded') {
            elements = (
                <div>
                    {this.state.lists.map(it => (
                        <div className="Lists__card" key={it.Name}>
                            <a onClick={() => this.showList(it.Name)}>
                                <h3>{it.Name}</h3>
                            </a>
                            <b>Elements: </b>{it.Size}
                        </div>
                    ))}
                </div>
            )
        }
        return (
            <div>
                <h1>Lists</h1>
                {elements}
                <hr/>
                {this.state.status === 'loaded' &&
                    <div className="Lists__total">
                        <b>Total: </b> {this.state.lists.length}
                    </div>
                }
                <br/>
                <div>
                    Go to custom name list:
                    <input type="text" value={this.state.name} onChange={this.onChange}/>
                    <button onClick={() => this.showList(this.state.name)}>Go</button>
                </div>
            </div>
        )
    }
}

export default Lists;