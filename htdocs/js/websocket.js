var serversocket = new WebSocket('wss://' + window.location.host + '/echo');


serversocket.onopen = function() {
    var msg = {
        topic:  "null",
        sender: "null",
        message: "Connection init",
        date: Date.now()
    };
     
    serversocket.send(JSON.stringify(msg));
}

// Write message on receive
serversocket.onmessage = function(e) {
    document.getElementById('output').innerHTML += "Received: " + e.data + "<br>";
};

function senddata() {
    // Construct a msg object containing the data the server needs to process the message from the chat client.
    var msg = {
        topic:  document.getElementById("topic").value,
        sender: document.getElementById("sender").value,
        message: document.getElementById("message").value,
        like: document.getElementById("like").value,
        date: Date.now()
    };
     
    serversocket.send(JSON.stringify(msg));
    document.getElementById('output').innerHTML += "Sent: " + JSON.stringify(msg) + "<br>";
}

$.plot("#placeholder", [ 
    {data: [[260,0],[260,100]], color: 'black', lines: {lineWidth:4}} 
],{
    yaxis:{
        show: false
    },
    xaxis:{
        min: 0,
        max: 400
    },
    grid: {
        show: true,
        borderWidth: 0,
        margin: {bottom: 90},
        labelMargin: -90,
        color: 'white'
    }
});
