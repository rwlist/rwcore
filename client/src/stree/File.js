import React, { Component } from 'react';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import FolderIcon from '@material-ui/icons/Folder';
import FileIcon from '@material-ui/icons/Description';
import HelpOutlineIcon from '@material-ui/icons/HelpOutline';

class File extends Component {
    render() {
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
        return (
            <ListItem button onDoubleClick={this.props.onOpen}>
                {icon}
                <ListItemText
                    primary={this.props.name}
                />
            </ListItem>
        )
    }
}

export default File;