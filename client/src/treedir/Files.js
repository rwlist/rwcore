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
import File from './File';

const styles = theme => ({
    path: {
        padding: theme.spacing.unit * 2,
        paddingBottom: theme.spacing.unit * 1
    }
});

class Files extends Component {
    render() {
        const { classes } = this.props;
        // TODO: material design components
        return (
            <Paper>
                <Typography variant="body2" gutterBottom className={classes.path}>
                    /home/kek/mda/lol/
                </Typography>
                <Divider/>
                <List>
                    <File name=".." icon="folder"/>
                    <File name="Dir1" icon="folder"/>
                    <File name="Dir2" icon="folder"/>
                    <File name="File1" icon="file"/>
                    <File name="File2" icon="file"/>
                    <File name="File3" icon="file"/>
                    <File name="FileKek" icon="file"/>
                </List>
            </Paper>
        )
    }
}

export default withStyles(styles)(Files);