// king test 5
// server handler mock

fn handle_4(req_obj) {
  if req_obj["method"] = "GET" {
    rev {"status": 200, "body": "OK"}
  }
  rev {"status": 404, "body": "NOT FOUND"}
}
pt(handle_4({"method": "GET"})["status"])
