Feature: Calculating the complexity
  Scenario Outline: simple file
    When analyze path "internal/simple.go"
    Then it returns no error
    And cyclomatic complexity of function <func> equals <complexity>
    Examples:
      | func  | complexity |
      | NoComplexity| 1    |
      | OneIf| 2           |
      | And| 3             |
      | Or| 3              |
      | (S).AFunction| 1   |

  Scenario: subdirectory
    When analyze path "internal/sub/"
    Then it returns no error
    And cyclomatic complexity of function FuncInASubDirectory equals 3
  Scenario: function with pointer receiver
    When analyze path "internal/pointer.go"
    Then it returns no error
    And cyclomatic complexity of function WithPointerReceiver equals 2
  Scenario: return only top 5 most complex functions
    Given set top parameter to 5
    When analyze path "internal/"
    Then it returns no error
    And the size of the result should equal 5
  Scenario: return only top 3 most complex functions
    Given set top parameter to 3
    When analyze path "internal/"
    Then it returns no error
    And the size of the result should equal 3