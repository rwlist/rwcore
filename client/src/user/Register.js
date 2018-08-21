import React, { Component } from 'react';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import { withStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import FormGroup from '@material-ui/core/FormGroup';

const styles = theme => ({
    root: {
        flexGrow: 1,
        padding: theme.spacing.unit * 2,
    },
    button: {
        marginTop: theme.spacing.unit,
    },
});

class Register extends Component {
    render() {
        const { classes, handlers, values } = this.props;

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
                    <TextField
                        autoFocus
                        margin="dense"
                        id="verifyPassword"
                        label="Verify Password"
                        type="password"
                        value={values.verifyPassword}
                        onChange={handlers.handleChange('verifyPassword')}
                    />
                    <Button
                        variant="contained"
                        color="primary"
                        className={classes.button}
                        onClick={handlers.register}
                    >
                        Register
                    </Button>
                    <Button
                        variant="contained"
                        color="primary"
                        className={classes.button}
                        onClick={handlers.wantLogin}
                    >
                        Log into existing account
                    </Button>
                </FormGroup>
            </div>
        );
    }
}

export default withStyles(styles)(Register);