Feature: Steps
  In order to develop behavior driven
  As a go developer
  I want to define steps

  @wip
  Scenario: Undefined step
    Given no step implementation
    When I run cucumber
    And a new scenario begins
    And the scenario has a step
    Then the scenario ends
    And reports an undefined step

