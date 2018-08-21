import React, { Component } from 'react';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import { withStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import FormGroup from '@material-ui/core/FormGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';

const styles = theme => ({
    root: {
        flexGrow: 1,
        padding: theme.spacing.unit * 2,
    },
    button: {
        marginTop: theme.spacing.unit,
    },
});

class Login extends Component {
    render() {
        const { classes, handlers, values } = this.props;

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
                        value={values.login}
                        onChange={handlers.handleChange('login')}
                    />
                    <TextField
                        autoFocus
                        margin="dense"
                        id="password"
                        label="Password"
                        type="password"
                        value={values.password}
                        onChange={handlers.handleChange('password')}
                    />
                    <FormControlLabel
                        control={
                            <Checkbox
                                checked={values.remember}
                                onChange={handlers.handleCheckbox('remember')}
                                value="remember"
                            />
                        }
                        label="Remember"
                    />
                    <Button
                        variant="contained"
                        color="primary"
                        className={classes.button}
                        onClick={handlers.login}
                    >
                        Login
                    </Button>
                    <Button
                        variant="contained"
                        color="primary"
                        className={classes.button}
                        onClick={handlers.wantRegister}
                    >
                        Create new account
                    </Button>
                </FormGroup>
            </div>
        );
    }
}

export default withStyles(styles)(Login);