import React, { Component } from 'react';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import FolderIcon from '@material-ui/icons/Folder';
import FileIcon from '@material-ui/icons/Description';
import HelpOutlineIcon from '@material-ui/icons/HelpOutline';
import { withStyles } from '@material-ui/core/styles';
import Checkbox from '@material-ui/core/Checkbox';

const styles = theme => ({
    selected: {
        backgroundColor: "#e8f0fe",
        color: "#1967d2",
    }
});


class File extends Component {
    render() {
        const { classes } = this.props;

        let icon = <HelpOutlineIcon />;
        if (this.props.type === 'directory') {
            icon = <FolderIcon />;
        }
        if (this.props.type === 'file') {
            icon = <FileIcon />;
        }
        if (icon) {
            icon = (
                <ListItemIcon>
                    {icon}
                </ListItemIcon>
            );
        }
        if (this.props.multiselect) {
            return (
                <ListItem
                    button
                    disableRipple
                    onClick={this.props.onSelect}
                >
                    <Checkbox
                        checked={this.props.selected}
                        tabIndex={-1}
                        disableRipple
                    />
                    {icon}
                    <ListItemText
                        primary={this.props.name}
                    />
                </ListItem>
            )
        } else {
            return (
                <ListItem
                    button
                    disableRipple
                    onDoubleClick={this.props.onOpen}
                    onClick={this.props.onSelect}
                    className={this.props.selected ? classes.selected : ''}
                >
                    {icon}
                    <ListItemText
                        primary={this.props.name}
                    />
                </ListItem>
            )
        }
    }
}

export default withStyles(styles)(File);