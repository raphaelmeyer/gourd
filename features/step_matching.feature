Feature: Patterns to match steps
  In order to reuse steps with different parameters
  I want to use regular expressions in step patterns

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
    When I run cucumber
    Then number 1234 is passed to the matching step

