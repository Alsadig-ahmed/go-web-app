function makeAjaxCall() {
    fetch('/api/request')
    .then(function(response) {
        return response.json();
    })
    .then(function(msg) {
      let message = msg.Body;
     let output = document.getElementById("output");
     output.innerHTML = message
     output.style.boxShadow = "0 0 1rem .2rem";
     output.style.padding ="1rem 0"
    });
}