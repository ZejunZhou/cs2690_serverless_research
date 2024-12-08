require_relative 'bundle/bundler/setup'
require 'json'

payload = STDIN.read
if payload != ""
  payload = JSON.parse(payload)
  
  # payload contains checks
  if payload["input"] 
    puts payload["input"]
  end
end

raise "Something went terribly wrong!"
