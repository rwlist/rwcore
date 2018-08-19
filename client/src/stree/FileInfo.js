import React, { Component } from 'react';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import { withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import RefreshIcon from '@material-ui/icons/Refresh';
import Button from '@material-ui/core/Button';
import Divider from '@material-ui/core/Divider';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import ListItemText from '@material-ui/core/ListItemText';
import DeleteIcon from '@material-ui/icons/Delete';
import red from '@material-ui/core/colors/red';

const styles = theme => ({
    text: {
        flexGrow: 0,
    },
    root: {
        padding: theme.spacing.unit
    },
    deleteButton: {
        margin: theme.spacing.unit,
        color: theme.palette.getContrastText(red[500]),
        backgroundColor: red[500],
        '&:hover': {
            backgroundColor: red[700],
        },
    },
});

class FileInfo extends Component {
    render() {
        const { classes } = this.props;

        let content;
        if (this.props.file) {
            content = (
                <React.Fragment>
                    <pre>
                        <code>
                            {JSON.stringify(this.props.file, null, 4)}
                        </code>
                    </pre>
                    <Button
                        variant="contained"
                        className={classes.deleteButton}
                        onClick={this.props.onDelete}
                    >
                        <DeleteIcon/>
                        Delete
                    </Button>
                    {/* <Grid item>
                        <Typography variant="body2" align="left" className={classes.text}>
                            "mda" file
                        </Typography>
                    </Grid>
                    <Grid item>
                        <Typography variant="body2" gutterBottom>
                            Created on August 10, 2018
                        </Typography>
                    </Grid> */}
                </React.Fragment>
            );
        } else {
            content = (
                <Typography variant="caption" gutterBottom align="center">
                    Select an item to view details.
                </Typography>
            );
        }
        return (
            <Paper className={classes.root}>
                <Grid container spacing={8} direction="column">
                    {content}
                </Grid>
            </Paper>
        )
    }
}

export default withStyles(styles)(FileInfo);