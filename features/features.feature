Feature: Calculating the complexity
  Scenario: simple file
    When analyze path "internal/simple.go"
    Then it returns no error
    And cyclomatic complexity of function NoComplexity equals 1
    And cyclomatic complexity of function OneIf equals 2
  Scenario: subdirectory
    When analyze path "internal/sub/"
    Then it returns no error
    And cyclomatic complexity of function FuncInASubDirectory equals 3
  Scenario: function with pointer receiver
    When analyze path "internal/pointer.go"
    Then it returns no error
    And cyclomatic complexity of function WithPointerReceiver equals 2