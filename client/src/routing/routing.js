const fs = require("fs")
const path = require("path")

// A function to serve static files
const serveStaticFile = (res, filename, contentType) => {
	const filePath = path.join(__dirname, "../../public", filename)
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

const routes = {
	GET: {
		"/": (req, res) => serveStaticFile(res, "index.html", "text/html"),
		/* "/about": (req, res) => serveStaticFile(res, "about.html", "text/html"),
		"/contact": (req, res) => serveStaticFile(res, "contact.html", "text/html"), */
		"/api/data": (req, res) => {
			const data = { message: "GET request received" }
			res.writeHead(200, { "Content-Type": "application/json" })
			res.end(JSON.stringify(data))
		}
	},
	POST: {
		/* "/api/data": (req, res) => {
			let body = ""
			req.on("data", (chunk) => {
				body += chunk.toString()
			})
			req.on("end", () => {
				const parsedData = JSON.parse(body)
				res.writeHead(200, { "Content-Type": "application/json" })
				res.end(
					JSON.stringify({ message: "POST request received", data: parsedData })
				)
			})
		} */
	}
}

module.exports = { routes }
