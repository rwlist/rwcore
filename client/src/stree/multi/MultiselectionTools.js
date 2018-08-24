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
import Filter from './Filter';
import Tools from './Tools';
import BulkUpdate from './BulkUpdate';

const styles = theme => ({
    root: {
        padding: theme.spacing.unit * 3,
    },
    deleteButton: {
        margin: theme.spacing.unit,
        color: theme.palette.getContrastText(red[500]),
        backgroundColor: red[500],
        '&:hover': {
            backgroundColor: red[700],
        },
    },
    divider: {
        marginTop: theme.spacing.unit * 2,
        marginBottom: theme.spacing.unit * 2,
    },
});

class MultiselectionTools extends Component {
    constructor(props) {
        super(props);
    }

    tools = () => new Tools(this.props);

    handleChange = name => value => {
        this.setState({
            [name]: value,
        });
    }
    
    createButton(content, action) {
        return (
            <Grid item>
                <Button
                    variant="contained"
                    color="primary"
                    onClick={action}
                >
                    {content}
                </Button>
            </Grid>
        )
    }

    render() {
        const { classes } = this.props;
        const btn = this.createButton;

        return (
            <Paper className={classes.root}>
                <Typography variant="display2" gutterBottom className={classes.head}>
                    Multiselection tools
                </Typography>
                <Grid container spacing={16}>
                    {btn("Clear selection", this.tools().clearSelection)}
                    {btn("Invert selection", this.tools().invertSelecion)}
                </Grid>
                <Divider className={classes.divider}/>
                <Filter tools={this.tools()}/>
                <Divider className={classes.divider}/>
                <BulkUpdate tools={this.tools()}/>
            </Paper>
        )
    }
}

export default withStyles(styles)(MultiselectionTools);