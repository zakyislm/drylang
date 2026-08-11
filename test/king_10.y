// king test 10
// web scraper mockup

try {
  cns html_9 = req("GET", "https://example.com")
  cns length_9 = len(html_9)
  pt(length_9 > 0)
} err(e) {
  pt("Failed request")
}
