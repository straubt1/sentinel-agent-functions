import "module" "agent-functions" {
  source = "./functions/agent-functions.sentinel"
}

policy "demo-agent-functions" {
  source            = "./policies/demo-agent-functions.sentinel"
  enforcement_level = "advisory"
}