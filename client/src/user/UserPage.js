import React, { Component } from 'react';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import { withStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import CircularProgress from '@material-ui/core/CircularProgress';
import TextField from '@material-ui/core/TextField';
import FormGroup from '@material-ui/core/FormGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';

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
    button: {
        marginTop: theme.spacing.unit,
    }
});

class UserPage extends Component {
    constructor(props) {
        super(props);
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
    
    onUserUpdate(user) {
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

    updateUser() {
        this.setState({ status: 'loading' })
        fetch('/user/current', { method: 'GET' })
        .then(it => it.json())
        .then(it => {
            if (it.Err) {
                throw it;
            }
            this.onUserUpdate(it);
        })
        .catch(it => {
            console.error(it);
        })
    }

    login = () => {
        var formData = new FormData();
        formData.append('username', this.state.login);
        formData.append('password', this.state.password);
        formData.append('remember', this.state.remember);
        fetch('/login', {
            method: "POST",
            body: formData,
        })
        .then(it => it.json())
        .then(it => {
            console.log(it);
            this.updateUser();
        })
        .catch(err => {
            console.error(err);
        })
    }

    register = () => {
        var formData = new FormData();
        formData.append('username', this.state.login);
        formData.append('password', this.state.password);
        formData.append('verifyPassword', this.state.verifyPassword);
        fetch('/register', {
            method: "POST",
            body: formData,
        })
        .then(it => it.json())
        .then(it => {
            console.log(it);
            this.updateUser();
        })
        .catch(err => {
            console.error(err);
        })
    }

    logout = () => {
        fetch('/logout', {
            method: "POST",
        })
        .then(it => it.json())
        .then(it => {
            console.log(it);
            this.updateUser();
        })
        .catch(err => {
            console.error(err);
        })
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
                    <Button
                        variant="contained"
                        color="primary"
                        className={classes.button}
                        onClick={this.logout}
                    >
                        Logout
                    </Button>
                </div>
            )
        }

        if (this.state.page === 'login') {
            return (
                <div className={classes.root}>
                    <FormGroup>
                        <Typography variant="display3" gutterBottom>
                            Login page
                        </Typography>
                        <TextField
                            autoFocus
                            margin="dense"
                            id="login"
                            label="Login"
                            type="text"
                            value={this.state.login}
                            onChange={this.handleChange('login')}
                        />
                        <TextField
                            autoFocus
                            margin="dense"
                            id="password"
                            label="Password"
                            type="password"
                            value={this.state.password}
                            onChange={this.handleChange('password')}
                        />
                        <FormControlLabel
                            control={
                                <Checkbox
                                    checked={this.state.remember}
                                    onChange={this.handleCheckbox('remember')}
                                    value="remember"
                                />
                            }
                            label="Remember"
                        />
                        <Button
                            variant="contained"
                            color="primary"
                            className={classes.button}
                            onClick={this.login}
                        >
                            Login
                        </Button>
                        <Button
                            variant="contained"
                            color="primary"
                            className={classes.button}
                            onClick={() => this.setState({ page: 'register' })}
                        >
                            I'd rather create new account
                        </Button>
                    </FormGroup>
                </div>
            )
        }
        if (this.state.page === 'register') {
            return (
                <div className={classes.root}>
                    <FormGroup>
                        <Typography variant="display3" gutterBottom>
                            Register page
                        </Typography>
                        <TextField
                            autoFocus
                            margin="dense"
                            id="login"
                            label="Login"
                            type="text"
                            value={this.state.login}
                            onChange={this.handleChange('login')}
                        />
                        <TextField
                            autoFocus
                            margin="dense"
                            id="password"
                            label="Password"
                            type="password"
                            value={this.state.password}
                            onChange={this.handleChange('password')}
                        />
                        <TextField
                            autoFocus
                            margin="dense"
                            id="verifyPassword"
                            label="Verify Password"
                            type="password"
                            value={this.state.verifyPassword}
                            onChange={this.handleChange('verifyPassword')}
                        />
                        <Button
                            variant="contained"
                            color="primary"
                            className={classes.button}
                            onClick={this.register}
                        >
                            Register
                        </Button>
                        <Button
                            variant="contained"
                            color="primary"
                            className={classes.button}
                            onClick={() => this.setState({ page: 'login' })}
                        >
                            I'd rather log into existing account
                        </Button>
                    </FormGroup>
                </div>
            )
        }
    }
}

export default withStyles(styles)(UserPage);