export default class Tools {
    constructor(props) {
        this.files = props.files;
        this.multiselected = props.multiselected;
        console.log('constructor', this.files, this.multiselected);
    }
}