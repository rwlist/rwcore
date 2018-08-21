import React from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import { withStyles } from '@material-ui/core/styles';

const styles = theme => ({
    content: {
        fontFamily: "'Roboto Mono', monospace",
        fontSize: 15,
    }
});

class NewFileDialog extends React.Component {
    state = {
        content: "",
        name: "",
    };

    handleChange = name => event => {
        this.setState({
            [name]: event.target.value,
        });
    }

    render() {
        const { classes } = this.props;

        return (
            <div>
                <Dialog
                    open={this.props.open}
                    onClose={this.props.handleClose}
                    aria-labelledby="form-dialog-title"
                >
                    <DialogTitle id="form-dialog-title">Create new file</DialogTitle>
                    <DialogContent>
                        <TextField
                            autoFocus
                            margin="dense"
                            id="name"
                            label="File name"
                            type="text"
                            fullWidth
                            value={this.state.name}
                            onChange={this.handleChange('name')}
                        />
                        <TextField
                            margin="dense"
                            id="name"
                            label="Content"
                            type="text"
                            fullWidth
                            multiline
                            value={this.state.content}
                            onChange={this.handleChange('content')}
                            InputProps={{
                                className: classes.content,
                                spellCheck: "false",
                            }}
                        />
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={this.props.handleClose} color="primary">
                            Cancel
                        </Button>
                        <Button onClick={() => this.props.handleAction(this.state.name, this.state.content)} color="primary">
                            Create
                        </Button>
                    </DialogActions>
                </Dialog>
            </div>
        );
    }
}

export default withStyles(styles)(NewFileDialog);