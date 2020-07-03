(function() {
    var src = "http://localhost:5000";
    var i = document.createElement("iframe");
    window.addEventListener("message", function(e) {
        console.log(e.origin, src, e.origin === src);
        if (e.origin === src) console.log("!", e);
    });
    i.frameBorder = 0;
    i.height = i.width = 1;
    i.id = "slime";
    i.src = src + "/" + "i";
    i.addEventListener("load", function() {
        try {
            i.contentWindow.postMessage(i.id, i.src);
        } catch(err) {
            console.warn(err);
        }
    });
    document.body.appendChild(i);
}());