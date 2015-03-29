Feature: Scenario validation
  In order to delevop behavior driven
  As a go developer
  I want to get feedback about success or failure of steps and scenarios

  Scenario: An undefined step
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

  Scenario: A pending step
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

  Scenario: A passing step
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

  Scenario: A failing step
    Given a go wire server
    And the following feature:
      """
      Feature: feature
        Scenario: scenario
          Given a step which fails
      """
    When I run cucumber
    Then the output should contain:
      """
      1 scenario (1 failed)
      1 step (1 failed)
      """

  Scenario: A combination of undefined, passing and failing steps
    Given a go wire server
    And the following feature:
      """
      Feature: feature
        Scenario: scenario
          Given a step which passes
          And a step which fails
          Then a step which is pending
          And a step which passes
      """
    When I run cucumber
    Then the output should contain:
      """
      1 scenario (1 failed)
      4 steps (1 failed, 2 skipped, 1 passed)
      """

