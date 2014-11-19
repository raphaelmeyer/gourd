Feature: Step actions and behavior
  In order to test my application
  As a go developer
  I want to define the actions and behavior of steps

  @wip
  Scenario: Execute user defined step code
    Given a go wire server
    And the following feature:
      """
      Feature: feature
        Scenario: scenario
          When an action is triggered
          Then the action 
      """
    When I run cucumber
    Then the output should contain:
      """
      
      """
