(function() {
    var src = "http://localhost:5000";
    var i = document.createElement("iframe");
    window.addEventListener("message", function(e) {
        if (e.origin === src) {
            console.log("!", e);
            i.parentElement.removeChild(i);
            console.log("done!");
        }
    });
    i.frameBorder = 0;
    i.height = i.width = 1;
    i.id = "slime";
    i.src = src + "/" + "f";
    i.addEventListener("load", function() {
        try {
            i.contentWindow.postMessage(i.id, i.src);
        } catch(err) {
            console.warn(err);
        }
    });
    document.body.appendChild(i);
}());