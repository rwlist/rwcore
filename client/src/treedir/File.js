import React, { Component } from 'react';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import FolderIcon from '@material-ui/icons/Folder';
import FileIcon from '@material-ui/icons/Description';

class File extends Component {
    render() {
        let icon = null;
        if (this.props.icon === 'folder') {
            icon = <FolderIcon />;
        }
        if (this.props.icon === 'file') {
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
            <ListItem button>
                {icon}
                <ListItemText
                    primary={this.props.name}
                />
            </ListItem>
        )
    }
}

export default File;