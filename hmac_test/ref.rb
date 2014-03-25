begin
  # use `bundle install --standalone' to get this...
  require_relative './bundle/bundler/setup'
rescue LoadError
  # fall back to regular bundler if the developer hasn't bundled standalone
  require 'bundler'
  Bundler.setup
end
require "hmac/signer"
ts = ARGV[2]
date = Time.at(ts.to_i).strftime('%a, %e %b %Y %T GMT')
print HMAC::Signer.new().sign_url(ARGV[0], ARGV[1],{:date => date})
