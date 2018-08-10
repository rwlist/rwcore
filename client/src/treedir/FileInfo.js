import React, { Component } from 'react';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import { withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography'
import Divider from '@material-ui/core/Divider';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import ListItemText from '@material-ui/core/ListItemText';
import FolderIcon from '@material-ui/icons/Folder';

const styles = theme => ({
    text: {
        flexGrow: 0,
    },
    root: {
        padding: theme.spacing.unit
    }
});

class FileInfo extends Component {
    render() {
        const { classes } = this.props;
        return (
            <Paper className={classes.root}>
                <Grid container spacing={8} direction="column">
                    <Grid item>
                        <Typography variant="body2" align="left" className={classes.text}>
                            "mda" file
                        </Typography>
                    </Grid>
                    <Grid item>
                        <Typography variant="body2" gutterBottom>
                            Created on August 10, 2018
                        </Typography>
                    </Grid>
                </Grid>
            </Paper>
        )
    }
}

export default withStyles(styles)(FileInfo);