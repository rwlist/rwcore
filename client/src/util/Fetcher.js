export default class Fetcher {
    fetch(input, init) {
        console.log('fetch: ', input, init);
        return fetch(input, init)
            .then(it => it.json())
            .then(it => {
                if (it.Error) {
                    throw it;
                }
                console.log('fetch response: ', it);
                return it;
            })
            .catch(it => {
                console.error('fetch error: ', it);
                // TODO: check for unauthorized error
                throw it;
            });
    }

    postJSON(url, body, stringify) {
        const init = { method: 'POST' };
        if (body) {
            if (stringify) {
                init.body = JSON.stringify(body);
            } else {
                init.body = body;
            }
            init.headers = {
                'Content-Type': 'application/json; charset=utf-8'
            };
        }
        return this.fetch(url, init);
    }

    get(url) {
        return this.fetch(url, { method: 'GET' });
    }
}