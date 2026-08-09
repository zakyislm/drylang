// template: html-render
// desc: HTTP Server that returns dynamic HTML pages

fn render_page(title, content) {
    rev `
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>${title}</title>
        <style>
            body { font-family: system-ui, sans-serif; max-width: 800px; margin: 2rem auto; padding: 0 1rem; }
            h1 { color: #333; }
            .card { padding: 1rem; border: 1px solid #ddd; border-radius: 8px; margin-top: 1rem; }
        </style>
    </head>
    <body>
        <header>
            <h1>${title}</h1>
            <nav><a href="/">Home</a> | <a href="/about">About</a></nav>
        </header>
        <main>
            ${content}
        </main>
        <footer>
            <hr>
            <p><small>Rendered with dryLang Server</small></p>
        </footer>
    </body>
    </html>
    `
}

fn handler(req) {
    if req["path"] = "/" {
        content = `
            <div class="card">
                <h2>Welcome</h2>
                <p>This page was generated dynamically on the server!</p>
                <p>Server Time: <strong>${date()["format"]}</strong></p>
            </div>
        `
        html = render_page("Home - dryLang", content)
        
        rev { "status": 200, "headers": {"Content-Type": "text/html"}, "body": html }
        
    } elif req["path"] = "/about" {
        content = `
            <div class="card">
                <h2>About Us</h2>
                <p>We build incredibly fast apps with minimal syntax.</p>
            </div>
        `
        html = render_page("About - dryLang", content)
        
        rev { "status": 200, "headers": {"Content-Type": "text/html"}, "body": html }
        
    } el {
        html = render_page("404 Not Found", "<h2>Oops!</h2><p>The page you requested doesn't exist.</p>")
        rev { "status": 404, "headers": {"Content-Type": "text/html"}, "body": html }
    }
}

pt "HTML server running on http://localhost:8080"
op(8080, handler, "mul", 100)
