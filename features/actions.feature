Feature: Step actions and behavior
  In order to test my application
  As a go developer
  I want to define the actions and behavior of steps

  Scenario: Invoke code of a step
    Given a go wire server
    And the following feature:
      """
      Feature: feature
        Scenario: scenario
          Given a step with code
      """
    When I run cucumber
    Then the code was executed

