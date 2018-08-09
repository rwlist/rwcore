import React, { Component } from 'react';
import Grid from '@material-ui/core/Grid';
import ExplorerControls from './ExplorerControls';
import Paper from '@material-ui/core/Paper';
import { withStyles } from '@material-ui/core/styles';
import Files from './Files';
import FileInfo from './FileInfo';

const styles = theme => ({
    root: {
        flexGrow: 1
    }
});

class Explorer extends Component {
    render() {
        const { classes } = this.props;
        // TODO: material design components
        return (
            <Grid container spacing={16}>
                <Grid item xs={12}>
                    <ExplorerControls/>
                </Grid>

                <Grid item xs={12} md={6}>
                    <Files/>
                </Grid>

                <Grid item xs={12} md={6}>
                    <FileInfo/>
                </Grid>
            </Grid>
        )
    }
}

export default withStyles(styles)(Explorer);