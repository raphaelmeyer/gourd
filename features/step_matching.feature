Feature: Patterns to match steps
  In order to reuse steps with different parameters
  I want to use regex when defining a step pattern

  @wip
  Scenario: Match a number
    Given step with pattern "^a number (\d+)$"
    And a go wire server
    And the following feature:
      """
      Feature: feature
        Scenario: scenario
          Given a number 1234
      """
    Then the given step matches number 1234

