import React, { Component } from 'react';
import Explorer from './Explorer';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import { withStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import CircularProgress from '@material-ui/core/CircularProgress';

const styles = theme => ({
    root: {
        flexGrow: 1,
        padding: theme.spacing.unit * 2
    },
    grid: {
    
    },
    progress: {
        margin: theme.spacing.unit * 2,
    },
});

class STree extends Component {
    constructor(props) {
        super(props);
        this.fetcher = props.fetcher;
        this.state = {
            status: 'loading'
        }
    }
    
    componentDidMount() {
        this.load();
    }

    load() {
        this.setState({ status: 'loading' });
        this.fetcher.get('/stree/GetRoot')
            .then(it => {
                this.setState({
                    status: 'explorer',
                    root: it,
                });
            })
    }

    render() {
        const { classes } = this.props;

        let content;
        if (this.state.status === 'loading') {
            content = <CircularProgress className={classes.progress} />;
        } else {
            content = <Explorer root={this.state.root} fetcher={this.fetcher}/>;
        }
        return (
            <div className={classes.root}>
                <Grid container className={classes.grid}>
                    <Grid item xs={12}>
                        <Typography variant="display3" gutterBottom>
                            Explorer
                        </Typography>
                        {content}
                    </Grid>
                </Grid>
            </div>
        )
    }
}

export default withStyles(styles)(STree);