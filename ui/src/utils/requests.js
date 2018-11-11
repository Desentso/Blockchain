const baseUrl = "http://localhost:9090"

export const postRequest = (url, data) => {
    return fetch(baseUrl + url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json"
        },
        body: JSON.stringify(data)
    })
}

export const getRequest = (url) => {
    return fetch(baseUrl + url, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json"
        }
    })
    .then(resp => resp.json())
}
