# Results
test {
  rules = {
    main = true
  }
}

mock "agent-functions" {
  module {
    source = "../../../functions/agent-functions.sentinel"
  }
}