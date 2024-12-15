import './style.css';
import './app.css';
import {EventsEmit, EventsOn} from "../wailsjs/runtime/runtime"

window.addEventListener("load", function() {
    let term = new Terminal();
    term.open(document.getElementById('terminal'));
    term.resize(100, 40);

    EventsEmit("client-ready", "");

    EventsOn("ptty-write", (data) => {
        term.write(data);
    })

    term.onData((data)=> {
        EventsEmit("craftty-write", data)
    }); 
});