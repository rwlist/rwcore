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
        console.log('fetch', path);
        const dir = path[path.length - 1];
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

    refresh = () => {
        this.fetchFiles(this.state.path);
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
            this.refresh();
        })
        .catch(it => {
            console.error('error while creating directory', it);
        });
    }

    createFile = (name, content) => {
        this.setState({ dialog: null });
        const dir = this.getDir();
        fetch('/stree/CreateFile/' + dir.ID + '?name=' + encodeURIComponent(name), {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json; charset=utf-8'
            },
            body: content,
        })
        .then(it => it.json())
        .then(it => {
            if (it.Error) {
                throw it;
            }
            console.log('created file', it);
            this.refresh();
        })
        .catch(it => {
            console.error('error while creating file', it);
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

    onSelect = (selected) => {
        this.setState({ selected })
    }

    delete = it => {
        fetch('/stree/Delete/' + it.ID, {
            method: 'POST'
        })
        .then(it => it.json())
        .then(it => {
            if (it.Error) {
                throw it;
            }
            console.log('delete node', it);
            this.refresh();
        })
        .catch(it => {
            console.error('error while deleting directory', it);
        });
    }

    rename = newName => {
        this.setState({ dialog: null });
        fetch('/stree/Rename/' + this.state.selected.ID + '?newName=' + encodeURIComponent(newName), {
            method: 'POST'
        })
        .then(it => it.json())
        .then(it => {
            if (it.Error) {
                throw it;
            }
            console.log('rename node', it);
            this.refresh();
        })
        .catch(it => {
            console.error('error while renaming node', it);
        });
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