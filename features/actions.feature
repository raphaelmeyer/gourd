Feature: Step actions and behavior
  In order to validate implementations against acceptance criterias
  As a go developer
  I want to define actions and behavior of cucumber steps

  Scenario: Invoke user defined code of a step
    Given a go wire server
    And the following feature:
      """
      Feature: feature
        Scenario: scenario
          Given a step with code
      """
    When I run cucumber
    Then the code was executed

  Scenario: Report the message of a failing step
    Given a step with pattern "failure step" that fails with message "failure message"
    And a go wire server
    And the following feature:
      """
      Feature: feature
        Scenario: scenario
          Given failure step
      """
    When I run cucumber
    Then the output should contain:
      """
      failure message
      """

