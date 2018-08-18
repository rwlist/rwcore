import React, { Component } from 'react';
import Grid from '@material-ui/core/Grid';
import ExplorerControls from './ExplorerControls';
import Paper from '@material-ui/core/Paper';
import { withStyles } from '@material-ui/core/styles';
import Files from './Files';
import FileInfo from './FileInfo';
import NewDirectoryDialog from './NewDirectoryDialog';

const styles = theme => ({
    root: {
        flexGrow: 1
    }
});

class Explorer extends Component {
    constructor(props) {
        super(props);
        this.state = {
            path: [props.root],
            selected: null,
            status: 'loading',
            files: null,
            dialog: null,
        };
    }

    componentDidMount() {
        this.fetchFiles();
    }

    onFilesLoaded(dir, files) {
        if (dir.ID !== this.getDir().ID) {
            console.error('mismatch', dir, this.getDir());
            return;
        }
        this.setState({
            selected: null,
            status: 'ready',
            files,
        });
    }

    fetchFiles = (path) => {
        const _path = (path ? path : this.state.path);
        this.setState({
            selected: null,
            status: 'loading',
            files: null,
            path: _path,
        });
        const dir = _path[_path.length - 1];
        fetch('/stree/ListDirectory/' + dir.ID, { method: 'GET' })
            .then(it => it.json())
            .then(it => {
                if (it.Error) {
                    throw it;
                }
                this.onFilesLoaded(dir, it);
            })
            .catch(it => {
                console.error('error while listing directory');
                throw it;
            })
    }

    getDir() {
        return this.state.path[this.state.path.length - 1];
    }

    createDirectory = (name) => {
        this.setState({ dialog: null });
        const dir = this.getDir();
        fetch('/stree/CreateDir/' + dir.ID, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json; charset=utf-8'
            },
            body: JSON.stringify({
                Name: name
            })
        })
            .then(it => it.json())
            .then(it => {
                if (it.Error) {
                    throw it;
                }
                console.log('created directory', it);
                this.fetchFiles();
            })
            .catch(it => {
                console.error('error while creating directory');
                throw it;
            });
    }

    onOpen = (file) => {
        if (!file) {
            console.error('open incorrect file', file);
            return;
        }
        const dir = this.getDir();
        if (file.ID === dir.ParentID) {
            this.fetchFiles(this.state.path.slice(0, -1));
        } else {
            this.fetchFiles(this.state.path.concat([file]));
        }
    }

    render() {
        const { classes } = this.props;
        // TODO: material design components
        return (
            <Grid container spacing={16}>
                <Grid item xs={12}>
                    <ExplorerControls 
                        onDialog={dialog => this.setState({ dialog })}
                        onRefresh={this.fetchFiles}
                    />
                </Grid>

                <Grid item xs={12} md={6}>
                    <Files 
                        path={this.state.path}
                        files={this.state.files}
                        selected={this.state.selected}
                        status={this.state.status}
                        onOpen={this.onOpen}
                    />
                </Grid>

                <Grid item xs={12} md={6}>
                    <FileInfo/>
                </Grid>

                <NewDirectoryDialog
                    open={this.state.dialog === 'directory'}
                    handleClose={() => this.setState({ dialog: null })}
                    handleAction={this.createDirectory}
                />
            </Grid>
        )
    }
}

export default withStyles(styles)(Explorer);