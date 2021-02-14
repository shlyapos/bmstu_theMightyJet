function getMap() {
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

getMap().then(response => console.log(response.json()))
        .catch(error => console.log(error.message));

/*

const request = new XMLHttpRequest();
const url = "/generate";
//request.responseType = "json";
request.open("GET", url, true);
// request.setRequestHeader('Destination', 'localhost:4444');

request.send();

console.log(request);
console.log(request.responseText);
*/