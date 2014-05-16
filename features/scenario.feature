Feature: Scenario
  As a go developer
  I want to run scenarios

  Scenario: A scenario has a context
    Given a wire server running on port 2345
    When cucumber connects to port 2345
    And a new scenario starts
    And the scenario ends
    And cucumber closes the connection
    Then a new context has been created

