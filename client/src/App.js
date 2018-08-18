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

const styles = {
    root: {
      flexGrow: 1,
    },
    flex: {
      flexGrow: 1,
    },
  };
  

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            tab: 0,
        }
    }

    handleTab = (event, tab) => {
        this.setState({ tab });
    }

    render() {
        const { classes } = this.props;

        let content;
        if (this.state.tab === 0) {
            content = <UserPage/>
        }
        if (this.state.tab === 1) {
            content = <Lists/>
        }
        if (this.state.tab === 2) {
            content = <STree/>
        }
        return (
            <MuiThemeProvider theme={theme}>
                <CssBaseline/>
                <AppBar position="static" color="primary">
                    <Toolbar>
                        <Typography variant="title" color="inherit" className={classes.flex}>
                            rwlist.io
                        </Typography>
                        <Tabs value={this.state.tab} onChange={this.handleTab}>
                            <Tab label="User" />
                            <Tab label="Lists" />
                            <Tab label="My STree" />
                        </Tabs>
                    </Toolbar>
                </AppBar>
                {content}
            </MuiThemeProvider>
        )
    }
}

export default withStyles(styles)(App);
