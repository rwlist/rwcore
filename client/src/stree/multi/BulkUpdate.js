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

class BulkUpdate extends Component {
    constructor(props) {
        super(props);
        this.state = {
            code: '(it, api) => false',
            error: null,
        }
    }

    onCodeChange = code => {
        this.setState({ code })
    }

    onCompile = () => {
        this.setState({ result: null });
        try {
            const action = eval(this.state.code);
            this.setState({
                action,
                status: 'ok',
            });
        } catch (error) {
            this.setState({
                status: 'fail',
                error,
            });
        }
    }

    apply = () => {
        this.setState({
            result: this.props.tools.executeAction(this.state.action)
        })
    }

    render() {
        const { classes } = this.props;

        let result = null;
        if (this.state.result) {
            result = (
                <React.Fragment>
                    <Typography gutterBottom>
                        Processed {this.state.result.processed} items.
                    </Typography>
                    {!!this.state.result.error && (
                        <AceEditor
                            mode="javascript"
                            theme="github"
                            name="displayErrorQuick"
                            value={this.state.result.error.message}
                        />
                    )}
                </React.Fragment>
            )
        }
        return (
            <React.Fragment>
                <Typography variant="headline" gutterBottom className={classes.head}>
                    Bulk action
                </Typography>
                <Grid container spacing={16}>
                    <Grid item xs={12} sm={6}>
                        <AceEditor
                            mode="javascript"
                            theme="github"
                            onChange={this.onCodeChange}
                            name="actionCode"
                            editorProps={{$blockScrolling: true}}
                            value={this.state.code}
                            width="100%"
                        />
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        {this.state.status === 'ok' && (
                            <React.Fragment>
                                <Typography gutterBottom>
                                    Action compiled ok. Ready to apply action to {this.props.tools.selectedCount()} items.
                                </Typography>
                            </React.Fragment>
                        )}
                        {this.state.status === 'fail' && (
                            <React.Fragment>
                                <Typography gutterBottom>
                                    Action fails with error.
                                </Typography>
                                <AceEditor
                                    mode="javascript"
                                    theme="github"
                                    name="displayErrorQuick"
                                    value={this.state.error.message}
                                />
                            </React.Fragment>
                        )}
                        {result}
                        <br/>
                        <Button
                            variant="contained"
                            color="primary"
                            onClick={this.onCompile}
                        >
                            Compile
                        </Button>
                        <Button
                            variant="contained"
                            color="primary"
                            onClick={this.apply}
                            disabled={this.state.status !== 'ok'}
                        >
                            Run action
                        </Button>
                    </Grid>
                </Grid>
            </React.Fragment>
        )
    }
}

export default withStyles(styles)(BulkUpdate);