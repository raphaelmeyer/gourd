Feature: Wire server
  As a go developer
  I want a go wire server
  In order to use cucumber for behavior driven development

  Scenario: Connecting to the wire server
    Given a wire server running on port 2345
    When cucumber connects to port 2345
    And cucumber closes the connection
    Then the wire server exits

