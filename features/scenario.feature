Feature: Scenario
  As a go developer
  I want to run scenarios

  Scenario: An scenario with no steps
    Given a wire server running on port 2345
    And no given, when or then step
    When cucumber runs the scenario
    Then cucumber connects to port 2345
    And a new scenario starts
    And a new context is created
    And the wire server returns "success"
    And the scenario ends
    And the wire server returns "success"
    And cucumber closes the connection

