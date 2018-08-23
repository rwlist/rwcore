export default class Tools {
    constructor(props) {
        Object.assign(this, props);
    }

    clearSelection = () => {
        this.onChangeSelection({});
    }

    invertSelecion = () => {
        this.files.forEach(it => {
            this.onSelect(it);
        });
    }

    applyFilter = filter => {
        this.files
            .filter(filter)
            .forEach(it => this.onSelect(it));
    }

    countMatches = filter => {
        let cnt = 0;
        this.files.forEach(it => {
            if (filter(it)) {
                cnt++;
            }
        });
        return cnt;
    }
}