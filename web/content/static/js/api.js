class Request {
    constructor(host) {
        this.host = host;
        this.body = null;
    }

    send(method, callback) {
        let client = new XMLHttpRequest();
        client.open(method, this.host);

        if (this.body) {
            client.setRequestHeader("Content-Type", "application/json");
            client.send(JSON.stringify(this.body));
        } else {
            client.send();
        }

        client.onreadystatechange = function() {
            if (client.readyState === 4) {
                callback(client.response, client.status);
            }
        }
    }

    get(callback) {
        this.send("GET", callback);
    }

    post(callback) {
        this.send("POST", callback);
    }

    patch(callback) {
        this.send("PATCH", callback);
    }

    put(callback) {
        this.send("PUT", callback);
    }

    delete(callback) {
        this.send("DELETE", callback);
    }
}

class Model {
    constructor(path) {
        this.host = "http://localhost:8080" + path;
    }

    handleResponse(callback) {
        return function(response, status) {
            if (status === 200) {
                callback(JSON.parse(response), status);
            } else {
                callback(null, status);
            }
        };
    }

    get(callback, data = null) {
        let request = new Request(this.host);

        if (data !== null) {
            request.host += data.id;
        }

        request.get(this.handleResponse(callback));
    }

    create(callback, data) {
        let request = new Request(this.host);

        request.body = data;

        request.post(this.handleResponse(callback));
    }

    update(callback, data) {
        let request = new Request(this.host + data.id);

        request.body = data;

        request.patch(this.handleResponse(callback));
    }

    delete(callback, data) {
        let request = new Request((this.host + data.id));

        request.delete(this.handleResponse(callback));
    }
}

shop = {
    categories: new Model("/categories/"),
    items: new Model("/items/"),
    roles: new Model("/roles/"),
    tags: new Model("/tags/"),
    users: new Model("/users/")
};