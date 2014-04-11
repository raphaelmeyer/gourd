#Feature: Steps
#  In order to run a feature
#  As a cucumber
#  I want to invoke steps

Feature: Steps
  In order to develop behavior driven
  As a go developer
  I want to define steps

  Scenario: Undefined step
    Given a scenario with step "arbitrary step"
    When I run cucumber
    Then it should return step "arbitrary step" is undefined

