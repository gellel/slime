(function() {
    var src = "http://localhost:5000";
    var i = document.createElement("iframe");
    window.addEventListener("message", function(e) {
        if (e.origin === src) {
            console.log("!", e);
            i.parentElement.removeChild(i);
            try {
                document.cookie = ("uuid"+ "=" + e.data) + ";" + ("expires" + "=" + new Date(new Date().getTime()+60*60*1000*24).toGMTString()) + ";" + ("path" + "=" + "/");
            } catch(err) {
                console.warn("document.cookie", err);
            }
            try {
                localStorage.setItem("uuid", e.data);
            } catch(err) {
                console.warn("localStorage", err);
            }
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
            console.warn("iframe.contentWindow.postMessage", err);
        }
    });
    document.body.appendChild(i);
}());