import React, { Component } from 'react';
import Grid from '@material-ui/core/Grid';
import ExplorerControls from './ExplorerControls';
import Paper from '@material-ui/core/Paper';
import { withStyles } from '@material-ui/core/styles';
import Files from './Files';
import FileInfo from './FileInfo';
import NewDirectoryDialog from './dialog/NewDirectory';
import NewFileDialog from './dialog/NewFile';
import RenameDialog from './dialog/Rename';
import Multiselection from './multi/Multiselection';
import MultiselectionTools from './multi/MultiselectionTools';
import ExplorerAPI from './ExplorerAPI';
import Tools from './multi/Tools';

class Explorer extends Component {
    constructor(props) {
        super(props);
        this.fetcher = props.fetcher;
        this.api = new ExplorerAPI(this.fetcher);

        this.state = {
            path: [props.root], // TODO: derived state
            selected: null,
            files: null,

            status: 'loading',
            dialog: null,
        };
    }

    componentDidMount() {
        this.refresh();
    }

    refresh = () => {
        this.fetchFiles(this.state.path);
    }

    fetchFiles = (path) => {
        this.setState({ // TODO: async setState
            status: 'loading',
            prevStatus: this.state.status,
            files: null,
            path,
        });
        this.api.ListDirectory(ExplorerAPI.getDir(path))
            .then(files => {
                if (this.state.prevStatus === 'multiselect') {
                    this.setState({
                        selected: new Tools({files}).containsFilter(this.state.selected),
                        status: 'multiselect',
                        files,
                        path,
                    });
                } else {
                    let newSelected = null;
                    if (new Tools({files}).contains(this.state.selected)) {
                        newSelected = this.state.selected;
                    }
                    this.setState({
                        selected: newSelected,
                        status: 'ready',
                        files,
                        path,
                    });
                }
            });
    }

    createDirectory = (name) => {
        this.setState({ dialog: null });
        this.api.CreateDir(
            ExplorerAPI.getDir(this.state.path),
            name
        )
            .then(this.refresh);
    }

    createFile = (name, file) => {
        this.setState({ dialog: null });
        this.api.CreateFile(
            ExplorerAPI.getDir(this.state.path),
            name,
            file,
        )
            .then(() => this.refresh());
    }

    onOpen = (file) => {
        if (!file) {
            console.error('open incorrect file', file);
            return;
        }
        this.fetchFiles(
            ExplorerAPI.go(this.state.path, file)
        );
    }

    onSelect = (node) => {
        console.log('onSelect', node);
        if (!node) return;

        this.setState(state => {
            const { status } = state;
            if (status === 'ready') {
                return { ...state, selected: node };
            }
            if (status === 'multiselect') {
                return {
                    ...state,
                    selected: ExplorerAPI.select(
                        state.selected,
                        node,
                    )
                };
            }
            return state;
        });
    }

    delete = it => {
        this.api.Delete(it)
            .then(this.refresh);
    }

    rename = newName => {
        this.setState({ dialog: null });
        this.api.Rename(this.state.selected, newName)
            .then(this.refresh);
    }

    onMultiselect = () => {
        this.setState(state => {
            if (state.status === 'multiselect') {
                return { ...state, status: 'ready', selected: null };
            }
            if (state.status === 'ready') {
                return { ...state, selected: {}, status: 'multiselect' };
            }
            return state;
        })
    }

    render() {
        let info;
        if (this.state.status === 'multiselect') {
            info = (
                <Multiselection
                    files={this.state.selected}
                />
            )
        } else {
            info = (
                <FileInfo
                    file={this.state.selected}
                    dir={ExplorerAPI.getDir(this.state.path)}
                    onDelete={() => this.delete(this.state.selected)}
                />
            )
        }

        return (
            <Grid container spacing={16}>
                <Grid item xs={12}>
                    <ExplorerControls 
                        onDialog={dialog => this.setState({ dialog })}
                        onRefresh={this.refresh}
                        selected={this.state.selected}
                        onMultiselect={this.onMultiselect}
                        status={this.state.status}
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
                    {info}
                </Grid>

                {this.state.status === 'multiselect' && 
                    <Grid item xs={12}>
                        <MultiselectionTools
                            api={this.api}
                            files={this.state.files}
                            selected={this.state.selected}
                            onChangeSelection={it => this.setState({ selected: it })}
                            onSelect={this.onSelect}
                            refresh={this.refresh}
                        />
                    </Grid>
                }

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

export default Explorer;