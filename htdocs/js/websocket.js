var serversocket = new WebSocket('wss://' + window.location.host + '/echo');

serversocket.onopen = function() {
    serversocket.send("Connection init");
}

// Write message on receive
serversocket.onmessage = function(e) {
    document.getElementById('output').innerHTML += "Received: " + e.data + "<br>";
};

function senddata() {
    var data = document.getElementById('sendtext').value;
    serversocket.send(data);
    document.getElementById('output').innerHTML += "Sent: " + data + "<br>";
}

