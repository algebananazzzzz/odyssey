function handler(event) {
    var request = event.request;
    var uri = request.uri;

    // If path has no extension
    if (!uri.includes('.')) {
        if (uri.endsWith('/')) {
            request.uri += 'index.html';
        } else {
            request.uri += '/index.html';
        }
    }
    return request;
}