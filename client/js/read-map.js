async function getMap() {
    const options = {
        method: "GET",
        headers: { 'Content-Type': 'application/json' }
    };

    return new Promise((resolve, reject) => {
        fetch('/generate', options)
            .then(response => resolve(response))
            .catch(error => reject(error));
    });
}


getMap()
    .then(response => response.json())
    .catch(error => console.log(error.message))
        .then(map => console.log(map))


