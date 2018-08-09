import React, { Component } from 'react';
import Grid from '@material-ui/core/Grid';
import AddIcon from '@material-ui/icons/Add';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { withStyles } from '@material-ui/core/styles';

const styles = theme => ({
    root: {
        flexGrow: 1
    },
    paper: {
        padding: theme.spacing.unit * 2,
        paddingTop: theme.spacing.unit,
        paddingBottom: theme.spacing.unit,
    },
    button: {
        margin: theme.spacing.unit
    }
});

class ExplorerControls extends Component {
    render() {
        const { classes } = this.props;
        // TODO: material design components
        return (
            <Paper className={classes.paper}>
                <Button variant="contained" color="primary" className={classes.button}>
                    <AddIcon/>
                    New file
                </Button>
                <Button variant="contained" color="primary" className={classes.button}>
                    <AddIcon/>
                    New directory
                </Button>
            </Paper>
        )
    }
}

export default withStyles(styles)(ExplorerControls);