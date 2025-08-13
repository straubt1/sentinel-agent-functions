# Results
test {
  rules = {
    main = true
  }
}

import "plugin" "plugin_example" {
    source = "../../../plugin/sentinel-plugin-example"
}

mock "agent-functions" {
  module {
    source = "../../../functions/agent-functions.sentinel"
  }
}