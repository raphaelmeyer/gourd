require 'tempfile'
require 'rspec/expectations'

Given (/^the following feature:$/) do |feature|
  @feature_file = Tempfile.new("feature")
  @feature_file.write(feature)
  @feature_file.close
end

When (/^I run cucumber$/) do
  @result = `cucumber test/ #{@feature_file.path}`
end

Then(/^the output should contain:$/) do |output|
  expect(@result).to match(/#{Regexp.quote(output)}/)
end

