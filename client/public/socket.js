const socket = new WebSocket("ws://localhost:8080/ws")

const messageSubscribers = []

socket.addEventListener("open", function (event) {
	console.log("Connected to WebSocket server at ws://localhost:8080")
})

socket.addEventListener("message", function (event) {
	console.log("Message from server ", event.data)

	messageSubscribers.forEach((callback) => callback(event.data))
})

export function sendDataToServer(data) {
	if (socket.readyState === WebSocket.OPEN) {
		socket.send(data)
		console.log("Data sent to server:", data)
	} else {
		console.log("WebSocket is not open. ReadyState:", socket.readyState)
	}
}

export function onMessage(callback) {
	messageSubscribers.push(callback)
}
