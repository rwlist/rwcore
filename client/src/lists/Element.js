import React, { Component } from 'react';

class Element extends Component {
    render() {
        return (
            <div>
                <pre><code>{JSON.stringify(this.props.e, null, 4)}</code></pre>
            </div>
        )
    }
}

export default Element;