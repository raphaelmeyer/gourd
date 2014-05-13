Feature: Wire server
  As a go developer
  I want a go wire server
  In order to use cucumber for behavior driven development

  Scenario: Connection on the default port
    Given a wire server running on port 1847
    Then cucumber can connect to port 1847

  @wip
  Scenario: Connection on a specified port
    Given a wire server running on port 2345
    Then cucumber can connect to port 2345
