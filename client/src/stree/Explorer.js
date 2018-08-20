import React, { Component } from 'react';
import Grid from '@material-ui/core/Grid';
import ExplorerControls from './ExplorerControls';
import Paper from '@material-ui/core/Paper';
import { withStyles } from '@material-ui/core/styles';
import Files from './Files';
import FileInfo from './FileInfo';
import NewDirectoryDialog from './NewDirectoryDialog';
import NewFileDialog from './NewFileDialog';
import RenameDialog from './RenameDialog';

const styles = theme => ({
    root: {
        flexGrow: 1
    }
});

class Explorer extends Component {
    constructor(props) {
        super(props);
        this.fetcher = props.fetcher;
        this.state = {
            path: [props.root],
            selected: null,
            status: 'loading',
            files: null,
            dialog: null,
        };
    }

    componentDidMount() {
        this.refresh();
    }

    refresh = () => {
        this.fetchFiles(this.state.path);
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
        this.setState({
            selected: null,
            status: 'loading',
            files: null,
            path,
        });
        const dir = path[path.length - 1];
        this.fetcher.get('/stree/ListDirectory/' + dir.ID)
            .then(it => this.onFilesLoaded(dir, it));
    }

    getDir() {
        return this.state.path[this.state.path.length - 1];
    }

    createDirectory = (name) => {
        this.setState({ dialog: null });
        const dir = this.getDir();
        this.fetcher.postJSON('/stree/CreateDir/' + dir.ID, {Name: name}, true)
            .then(() => this.refresh());
    }

    createFile = (name, content) => {
        this.setState({ dialog: null });
        const dir = this.getDir();
        this.fetcher.postJSON(
            '/stree/CreateFile/' + dir.ID + '?name=' + encodeURIComponent(name),
            content
        )
            .then(() => this.refresh());
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

    onSelect = (selected) => {
        this.setState({ selected })
    }

    delete = it => {
        this.fetcher.postJSON('/stree/Delete/' + it.ID)
            .then(() => this.refresh());
    }

    rename = newName => {
        this.setState({ dialog: null });
        this.fetcher.postJSON(
            '/stree/Rename/' + this.state.selected.ID + '?newName=' + encodeURIComponent(newName)
        )
        fetch('/stree/Rename/' + this.state.selected.ID + '?newName=' + encodeURIComponent(newName), {
            method: 'POST'
        })
            .then(() => this.refresh());
    }

    render() {
        const { classes } = this.props;

        return (
            <Grid container spacing={16}>
                <Grid item xs={12}>
                    <ExplorerControls 
                        onDialog={dialog => this.setState({ dialog })}
                        onRefresh={this.refresh}
                        selected={this.state.selected}
                    />
                </Grid>

                <Grid item xs={12} md={6}>
                    <Files 
                        path={this.state.path}
                        files={this.state.files}
                        selected={this.state.selected}
                        status={this.state.status}
                        onOpen={this.onOpen}
                        onSelect={this.onSelect}
                    />
                </Grid>

                <Grid item xs={12} md={6}>
                    <FileInfo
                        file={this.state.selected}
                        onDelete={() => this.delete(this.state.selected)}
                    />
                </Grid>

                <NewDirectoryDialog
                    open={this.state.dialog === 'newDirectory'}
                    handleClose={() => this.setState({ dialog: null })}
                    handleAction={this.createDirectory}
                />

                <NewFileDialog
                    open={this.state.dialog === 'newFile'}
                    handleClose={() => this.setState({ dialog: null })}
                    handleAction={this.createFile}
                />

                <RenameDialog
                    open={this.state.dialog === 'rename'}
                    handleClose={() => this.setState({ dialog: null })}
                    handleAction={this.rename}
                    selected={this.state.selected}
                />

            </Grid>
        )
    }
}

export default withStyles(styles)(Explorer);