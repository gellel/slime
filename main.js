(function() {
    window.addEventListener("message", function(e) {
        console.log("!!", e);
    });
    var i = document.createElement("iframe");
    i.id = "slime";
    i.addEventListener("load", function() {
        i.contentWindow.postMessage(i.id, i.src);
    });
    document.body.appendChild(i);
    i.src = "http://localhost:5000/i";
}());