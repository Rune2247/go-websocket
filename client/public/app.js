// public/app.js

import { onMessage, sendDataToServer } from "./socket.js"

// Send data to server
document.querySelector("button#send").addEventListener("click", () => {
	const message = document.querySelector("input#message").value
	sendDataToServer(message)
})

// Handle incoming messages from the server
onMessage((data) => {
	console.log("Received data in app.js:", data)
	document.querySelector("#output").textContent = `Server says: ${data}`
})
