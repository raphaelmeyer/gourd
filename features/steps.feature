Feature: Steps
  In order to develop behavior driven
  As a go developer
  I want to define steps

  Scenario: Undefined step
    Given the following feature:
      """
      Feature: feature
        Scenario: scenario
          Given an undefined step
      """
    And a running wire server
    When I run cucumber
    Then the output should contain:
      """
      1 scenario (1 undefined)
      1 step (1 undefined)
      """

