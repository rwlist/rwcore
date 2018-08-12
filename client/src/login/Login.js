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
});

class Login extends Component {
    constructor(props) {
        super(props);
        this.state = {
            login: '',
            password: '',
            remember: false,
        }
    }

    handleChange = name => event => {
        this.setState({
            [name]: event.target.value,
        });
    }
    
    login() {
        var formData = new FormData();
        formData.append('username', this.state.login);
        formData.append('password', this.state.password);
        formData.append('remember', this.state.remember);
        fetch('/login', {
            method: "POST",
            body: formData,
        });
    }

    render() {
        const { classes } = this.props;

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
                        onChange={this.handleChange('name')}
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
                        onChange={this.handleChange('remember')}
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
                </FormGroup>
            </div>
        )
    }
}

export default withStyles(styles)(Login);