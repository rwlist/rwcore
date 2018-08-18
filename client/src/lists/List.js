import React, { Component } from 'react';
import Element from './Element';
import FastArea from './FastArea';
import Spoiler from './Spoiler';

class List extends Component {
    constructor(props) {
        super(props);
        this.state = {
            status: "loading"
        };

        this.insertOne = this.insertOne.bind(this);
        this.insertMany = this.insertMany.bind(this);
        this.clear = this.clear.bind(this);
    }

    componentDidMount() {
        this.load();
    }

    onDataReceived(data) {
        if (this.state.status !== 'loading') {
            console.error('incorrect status');
            return;
        }
        this.setState({
            data,
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
        fetch('/lists/' + this.props.name + '/data', {
            method: 'GET'
        })
        .then(resp => resp.json())
        .then(it => {
            if (it.Error) throw it;
            return it;
        })
        .then(it => {
            console.log(it);
            this.onDataReceived(it);
        })
        .catch(handleErr)
    }

    link() {
        return '/lists/' + this.state.data.Name;
    }

    onError(err) {
        console.error(err);
    }

    onInfo(info) {
        console.log(info);
    }

    insertOne(text) {
        console.log(text);
        fetch(this.link() + '/insertOne', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json; charset=utf-8'
            },
            body: text
        })
        .then(resp => resp.json())
        .then(it => {
            if (it.Error) throw it;
            return it;
        })
        .then(it => {
            if (it.Error) {
                this.onError(it);
            } else {
                this.onInfo(it);
            }
            this.load();
        })
        .catch(err => this.onError(err))
    }

    insertMany(text) {
        console.log(text);
        fetch(this.link() + '/insertMany', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json; charset=utf-8'
            },
            body: text
        })
        .then(resp => resp.json())
        .then(it => {
            if (it.Error) {
                this.onError(it);
            } else {
                this.onInfo(it);
            }
            this.load();
        })
        .catch(err => this.onError(err))
    }

    clear() {
        console.log('CLEAR ACTION!')
        fetch(this.link() + '/clear', {
            method: 'POST'
        })
        .then(resp => resp.json())
        .then(it => {
            if (it.Error) {
                this.onError(it);
            } else {
                this.onInfo(it);
            }
            this.load();
        })
        .catch(err => this.onError(err))
    }

    render() {
        let elements;
        let controls;
        if (this.state.status === "loaded") {
            const backupLink = this.link() + '/backup';
            elements = this.state.data.Elements.map(it => (
                <Element e={it} key={it._id} />
            ));
            if (!elements) {
                elements = (
                    <div>No elements available!</div>
                );
            }
            controls = (
                <ul>
                    <li><Spoiler text="Insert one">
                        <br/>
                        <FastArea handle={this.insertOne} />
                    </Spoiler></li>
                    <li><Spoiler text="Insert many">
                        <br/>
                        <FastArea handle={this.insertMany} />
                    </Spoiler></li>
                    <li>
                        <a href={backupLink}>Backup all</a>
                    </li>
                    <li><Spoiler text="Clear">
                        <br/>
                        <button onClick={this.clear}>Clear!</button>
                    </Spoiler></li>
                </ul>
            );
        } else if (this.state.status === "loading") {
            elements = (
                <h3>Loading...</h3>
            );
        }
        return (
            <div>
                <h1>List '{this.props.name}'</h1>
                <hr/>
                {elements}
                <hr/>
                {controls}
            </div>
        )
    }
}

export default List;