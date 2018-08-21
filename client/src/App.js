import React, { Component } from 'react';
import Lists from './lists/Lists';
import STree from './stree/STree';
import CssBaseline from '@material-ui/core/CssBaseline';
import theme from './theme';
import { MuiThemeProvider } from '@material-ui/core';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import { withStyles } from '@material-ui/core/styles';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import UserPage from './user/UserPage';
import Fetcher from './util/Fetcher';

const styles = {
    flex: {
      flexGrow: 1,
    },
  };
  

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            tab: 0,
        };
        this.fetcher = new Fetcher();
    }

    handleTab = (event, tab) => {
        this.setState({ tab });
    }

    render() {
        const { classes } = this.props;

        const tabs = [
            [<UserPage fetcher={this.fetcher}/>, "User"],
            [<Lists fetcher={this.fetcher}/>, "Lists"],
            [<STree fetcher={this.fetcher}/>, "My STree"],
        ]

        const content = tabs[this.state.tab][0];
        return (
            <MuiThemeProvider theme={theme}>
                <CssBaseline/>
                <AppBar position="static" color="primary">
                    <Toolbar>
                        <Typography variant="title" color="inherit" className={classes.flex}>
                            rwlist.io
                        </Typography>
                        <Tabs value={this.state.tab} onChange={this.handleTab}>
                            {tabs.map((it, index) => <Tab label={it[1]} key={index}/>)}
                        </Tabs>
                    </Toolbar>
                </AppBar>
                {content}
            </MuiThemeProvider>
        )
    }
}

export default withStyles(styles)(App);
