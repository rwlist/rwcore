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

import brace from 'brace';
import AceEditor from 'react-ace';

import 'brace/mode/javascript';
import 'brace/theme/github';

const styles = theme => ({
    head: {
        marginTop: theme.spacing.unit,
    },
});

class Filter extends Component {
    constructor(props) {
        super(props);
        this.state = {
            filterCode: '',
        }
    }

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
            <React.Fragment>
                <Typography variant="headline" gutterBottom className={classes.head}>
                    Select by filter
                </Typography>
                <Grid container>
                    <Grid item xs={12} sm={6}>
                        <AceEditor
                            mode="javascript"
                            theme="github"
                            onChange={this.handleChange('filterCode')}
                            name="UNIQUE_ID_OF_DIV"
                            editorProps={{$blockScrolling: true}}
                            value={this.state.filterCode}
                            width="100%"
                        />
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <AceEditor
                            mode="javascript"
                            theme="github"
                            onChange={this.handleChange('filterCode')}
                            name="UNIQUE_ID_OF_DIV"
                            editorProps={{$blockScrolling: true}}
                            value={this.state.filterCode}
                            width="100%"
                        />
                    </Grid>
                </Grid>
            </React.Fragment>
        )
    }
}

export default withStyles(styles)(Filter);