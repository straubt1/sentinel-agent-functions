import "module" "agent-functions" {
  source = "./functions/agent-functions.sentinel"
}

import "plugin" "plugin_example" {
  source = "./policies/sentinel-plugin-tfresources"
}

policy "demo-agent-functions" {
  source            = "./policies/demo-agent-functions.sentinel"
  enforcement_level = "advisory"
}