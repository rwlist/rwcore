export default class ExplorerAPI {
    constructor(fetcher) {
        this.fetcher = fetcher;
    }

    static getDir(path) {
        return path[path.length - 1];
    }

    static go(path, node) {
        const dir = ExplorerAPI.getDir(path);
        if (node.ID === dir.ParentID) {
            return path.slice(0, -1);
        }
        return path.concat([node]);
    }

    static select(_selection, node) {
        const selection = Object.assign({}, _selection);
        const { ID } = node;
        if (selection[ID] !== undefined) {
            delete selection[ID];
        } else {
            selection[ID] = node;
        }
        return selection;
    }

    static isSelected(node, selection) {
        if (!node) {
            return false;
        }
        return selection[node.ID] !== undefined;
        
    }

    ListDirectory(dir) {
        return this.fetcher.get('/stree/ListDirectory/' + dir.ID);
    }

    CreateDir(dir, name) {
        return this.fetcher.postJSON('/stree/CreateDir/' + dir.ID, {Name: name}, true);
    }

    CreateFile(dir, name, file) {
        return this.fetcher.postJSON(
            '/stree/CreateFile/' + dir.ID + '?name=' + encodeURIComponent(name),
            file,
        );
    }

    Delete(node) {
        return this.fetcher.postJSON('/stree/Delete/' + node.ID);
    }

    Rename(node, newName) {
        return this.fetcher.postJSON(
            '/stree/Rename/' + node.ID + '?newName=' + encodeURIComponent(newName)
        );
    }
}