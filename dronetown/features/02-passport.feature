Feature: We need to know all the places where Drone has been.
  The passport is a write-ahead log, handled by the Information kiosk.

  Scenario Outline: If Drone arrives to the Junkyard it's because it can't 
    complete their errand, as stated in her passport
    Given a shopless DroneTown
    When I send Drone on an errand to get "<item>"
    Then a Drone arrives in "Junkyard" with their passport having
      | <item> |   |

    Examples:
      | item     |
      | tomato   |
      | oregano  |
      | a beamer |


  Scenario Outline: If Drone arrives Dock it's because it can
    complete their errand, as stated in her passport
    Given a shopless DroneTown
    And 1 shop opens providing free "<item>"
    When I send Drone on an errand to get "<item>"
    Then a Drone arrives "Dock" with their passport having
      | <item> | <item>0 |

    Examples:
      | item     |
      | tomato   |
      | oregano  |
      | a beamer |


  Scenario: Drone's passport when it arrives at dock includes
    all errands
    Given a shopless DroneTown
    And 1 shop opens providing "pasta" requiring "salsa"
    And 1 shop opens providing "salsa" requiring "tomato"
    And 1 shop opens providing free "tomato"
    When I send Drone on an errand to get "pasta"
    Then a Drone arrives "Dock" with their passport having
      | pasta  | pasta0  |
      | salsa  | salsa0  |
      | tomato | tomato0 |


  Scenario: Drone's passport when it arrives in Junkyard includes
    some errands. The last one was not achieved.
    Given a shopless DroneTown
    And 1 shop opens providing "pasta" requiring "salsa"
    And 1 shop opens providing "salsa" requiring "tomato"
    When I send Drone on an errand to get "pasta"
    Then a Drone arrives in "Junkyard" with their passport having
      | pasta  |   |
      | salsa  |   |
      | tomato |   |


  Scenario: Drone's passport when she arrives in Junkyard includes
    some errands. The last one was not achieved.
    Given a shopless DroneTown
    And 1 shop opens providing "pasta" requiring "salsa"
    And 1 shop opens providing free "tomato"
    When I send Drone on an errand to get "pasta"
    Then a Drone arrives in "Junkyard" with their passport having
      | pasta  |   |
      | salsa  |   |


  Scenario: No two passports can be the same
    Given a shopless DroneTown
    And 2 shops open providing free "tomato"
    When I send Drone on an errand to get "tomato"
    Then a Drone arrives "Dock" with their passport having
      | tomato  | tomato0 |
    And a Drone arrives "Dock" with their passport having
      | tomato  | tomato1 |
