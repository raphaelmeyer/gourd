Feature: Snippets
  In order to quickly add new steps
  I want code snippets for undefined steps

  Scenario: Suggest go code snippets for undefined steps
    Given a go wire server
    And the following feature:
      """
      Feature: feature
        Scenario: scenario
          Given an undefined step
          When another undefined step
          Then yet another undefined step
      """
    When I run cucumber
    Then the output should contain:
      """
      cucumber.Given("an undefined step").Pending
      """
    And the output should contain:
      """
      cucumber.When("another undefined step").Pending
      """
    And the output should contain:
      """
      cucumber.Then("yet another undefined step").Pending
      """

