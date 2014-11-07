Feature: Steps
  In order to develop behavior driven
  As a go developer
  I want to define steps

  Scenario: Undefined step
    Given a go wire server
    And the following feature:
      """
      Feature: feature
        Scenario: scenario
          Given an undefined step
      """
    When I run cucumber
    Then the output should contain:
      """
      1 scenario (1 undefined)
      1 step (1 undefined)
      """

  Scenario: Pending step
    Given a go wire server
    And the following feature:
      """
      Feature: feature
        Scenario: scenario
          Given a step which is pending
      """
    When I run cucumber
    Then the output should contain:
      """
      1 scenario (1 pending)
      1 step (1 pending)
      """

  Scenario: Passing step
    Given a go wire server
    And the following feature:
      """
      Feature: feature
        Scenario: scenario
          Given a step which passes
      """
    When I run cucumber
    Then the output should contain:
      """
      1 scenario (1 passed)
      1 step (1 passed)
      """

