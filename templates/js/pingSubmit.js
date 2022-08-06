const pingform = document.querySelector('.searchform');
const display_ping = document.querySelector('.ping-response');

var socket = new WebSocket('ws://localhost:3000/ws')
socket.onmessage = function(e){
    var message = e.data;
    message = JSON.parse(message);
    let temp = '';
    message.map(mes =>{
        temp = temp + mes + '<br>';
    })
    display_ping.innerHTML = temp;
}

function isInViewport(element) {
    const rect = element.getBoundingClientRect();
    return (
        rect.top >= 0 &&
        rect.left >= 0 &&
        rect.bottom <= (window.innerHeight || document.documentElement.clientHeight) &&
        rect.right <= (window.innerWidth || document.documentElement.clientWidth)
    );
}

async function submitPingURL(url) {
    const temp = {"URL": url};
    let response = await fetch('/ping', {
        headers: {
            'Content-Type': 'application/json',
          },
        method: 'POST',
        body: JSON.stringify(temp),
        }
    )

    socket.send("response")
    
    return
}

pingform.addEventListener('submit', (e) =>{  // on form submission, prevent default
    e.preventDefault();
    const formData = new FormData(pingform);
    let url = formData.get('searchurl').toLowerCase();
    submitPingURL(url);
    setInterval(function(){
        console.log("OKAY")
        if (isInViewport){
            socket.send("GET LOG")
        }
    }, 5000)
});