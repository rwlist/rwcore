import React, { Component } from 'react';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import { withStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import CircularProgress from '@material-ui/core/CircularProgress';
import Login from './Login';
import Register from './Register';

const styles = theme => ({
    root: {
        flexGrow: 1,
        padding: theme.spacing.unit * 2,
    },
    status: {
        padding: theme.spacing.unit * 2,
        margin: theme.spacing.unit * 2,
        maxWidth: 400,
    },
    progress: {
        margin: theme.spacing.unit * 2,
    },
});

class UserPage extends Component {
    constructor(props) {
        super(props);
        this.fetcher = props.fetcher; // TODO
        this.state = {
            login: '',
            password: '',
            verifyPassword: '',
            remember: false,

            status: 'loading',
            page: 'login',
        }
    }

    componentDidMount() {
        this.updateUser();
    }
    
    onUserUpdate = (user) => {
        if (user.Status === "User") {
            this.setState({
                status: 'logged-in',
                user,
            });
        } else {
            this.setState({
                status: 'forms',
                user,
            });
        }
    }

    updateUser = () => {
        this.setState({ status: 'loading' })
        this.fetcher.get('user/current')
            .then(this.onUserUpdate);
    }

    login = () => {
        var formData = new FormData();
        formData.append('username', this.state.login);
        formData.append('password', this.state.password);
        formData.append('remember', this.state.remember);
        this.fetcher.post('/login', formData)
            .then(this.updateUser);
    }

    register = () => {
        var formData = new FormData();
        formData.append('username', this.state.login);
        formData.append('password', this.state.password);
        formData.append('verifyPassword', this.state.verifyPassword);
        this.fetcher.post('/register', formData)
            .then(this.updateUser);
    }

    logout = () => {
        this.fetcher.post('/logout')
            .then(this.updateUser);
    }

    handleChange = name => event => {
        this.setState({
            [name]: event.target.value,
        });
    }

    handleCheckbox = name => event => {
        this.setState({
            [name]: event.target.checked,
        });
    }

    render() {
        const { classes } = this.props;
        const { handleChange, handleCheckbox } = this;
        const { login, password, verifyPassword, remember } = this.state;

        if (this.state.status === 'loading') {
            return (
                <CircularProgress className={classes.progress} />
            );
        }
        if (this.state.status === 'logged-in') {
            return (
                <div className={classes.root}>
                    <Typography variant="display2" gutterBottom>
                        Alright, you are in
                    </Typography>
                    <Paper className={classes.status}>
                        <pre>
                            <code>
                                {JSON.stringify(this.state.user, null, 4)}
                            </code>
                        </pre>
                    </Paper>
                    <Button variant="contained" color="primary" onClick={this.logout}>
                        Logout
                    </Button>
                </div>
            )
        }

        if (this.state.page === 'login') {
            return (
                <Login
                    values={{
                        login,
                        password,
                        remember,
                    }}
                    handlers={{
                        handleChange,
                        handleCheckbox,
                        login: this.login,
                        wantRegister: () => this.setState({ page: 'register' }),
                    }}
                />
            )
        }
        if (this.state.page === 'register') {
            return (
                <Register
                    values={{
                        login,
                        password,
                        verifyPassword,
                    }}
                    handlers={{
                        handleChange,
                        register: this.register,
                        wantLogin: () => this.setState({ page: 'login' }),
                    }}
                />
            )
        }
    }
}

export default withStyles(styles)(UserPage);