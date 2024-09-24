const { routes } = require("./routing/routing.js")
const http = require("http")
const fs = require("fs")
const path = require("path")

const hostname = "127.0.0.1"
const port = 3000

const mimeTypes = {
	".html": "text/html",
	".js": "text/javascript",
	".css": "text/css",
	".json": "application/json",
	".png": "image/png",
	".jpg": "image/jpg",
}

const server = http.createServer((req, res) => {
	const method = req.method
	const url = req.url

	if (routes[method] && routes[method][url]) {
		routes[method][url](req, res)
	} else {
		// Handle static files or return 404
		let extname = String(path.extname(req.url)).toLowerCase()
		let contentType = mimeTypes[extname] || "application/octet-stream"
		let filePath = path.join(__dirname, "../public", req.url)

		fs.readFile(filePath, (err, content) => {
			if (err) {
				res.writeHead(404)
				res.end("404: File Not Found")
			} else {
				res.writeHead(200, { "Content-Type": contentType })
				res.end(content, "utf-8")
			}
		})
	}
})

server.listen(port, hostname, () => {
	console.log(`Server running at http://${hostname}:${port}/`)
})
