Feature: Steps
  In order to develop behavior driven
  As a go developer
  I want to define steps

  Scenario: Undefined step
    Given a step with no implementation
    When I run cucumber
    Then cucumber should indicate that the step is undefined

