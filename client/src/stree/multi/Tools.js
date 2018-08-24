export default class Tools {
    constructor(props) {
        Object.assign(this, props);
    }

    selectedCount = () => {
        return Object.keys(this.selected).length;
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

    contains = (node) => {
        if (!node) {
            return false;
        }
        return !!this.files.find(it => it.ID === node.ID);
    }

    containsFilter = (selected) => {
        const res = {};
        Object.values(selected)
            .filter(this.contains)
            .forEach(it => res[it.ID] = it);
        return res;
    }

    executeAction = action => {
        let processed = 0;
        let error = null;
        let promises = [];
        try {
            Object.values(this.selected)
                .forEach(it => {
                    promises.push(Promise.resolve(action(it, this.api)));
                    processed++;
                });
        } catch (e) {
            error = e;
        }
        Promise.all(promises)
            .then(values => {
                console.log("actions finished", values);
                this.refresh();
            });
        return {
            processed,
            error,
        }
    }
}