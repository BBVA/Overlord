Feature: When Drone is asked to get an errand with and &, it has to
  visit ANDrés' store in order to get it

  Scenario: Drone already has everthing required for their errand-with-ampersand
    Given a shopless DroneTown
    And Drone's backpack contains
      | tomato  |
      | oregano |
    When I send Drone on an errand to get "tomato" AND "oregano"
    Then a Drone arrives "Dock" with their backpack containing "tomato" AND "oregano"


  Scenario: Drone already is missing an item required for their
    errand-with-ampersand
    Given a shopless DroneTown
    And Drone's backpack contains
      | tomato  |
    When I send Drone on an errand to get "tomato" AND "oregano"
    Then a Drone arrives in "Junkyard" with their passport having
      | oregano |   |


  Scenario: Drone is sent to get an errand-with-ampersand that can be
    done
    Given a shopless DroneTown
    And 1 shop opens providing free "tomato"
    And 1 shop opens providing free "oregano"
    When I send Drone on an errand to get "tomato" AND "oregano"
    Then a Drone arrives "Dock" with their backpack containing "tomato" AND "oregano"
    And a Drone arrives in "Dock" with their passport having
      | oregano | oregano0 |
      | tomato  | tomato0  |


  Scenario: Drone is sent to get an errand-with-ampersand that can't be
    done
    Given a shopless DroneTown
    And 1 shop opens providing free "oregano"
    When I send Drone on an errand to get "tomato" AND "oregano"
    Then a Drone arrives in "Junkyard" with their passport having
      | oregano |          |
      | tomato  |          |


  Scenario: Drone is sent to get a simple errand that depends on a
    complex one
    Given a shopless DroneTown
    And 1 shop opens providing free "tomato"
    And 1 shop opens providing free "oregano"
    And 1 shop opens providing "salsa" requiring "tomato" AND "oregano"
    When I send Drone on an errand to get "salsa"
    Then a Drone arrives "Dock" with their backpack containing
      | salsa |


  Scenario: Drone is sent to get a simple errand that depends on a
    complex one but can't be done
    Given a shopless DroneTown
    And 1 shop opens providing free "tomato"
    And 1 shop opens providing "salsa" requiring "tomato" AND "oregano"
    When I send Drone on an errand to get "salsa"
    Then a Drone arrives in "Junkyard" with their passport having
      | salsa   |         |
      | oregano |         |
      | tomato  | tomato0 |


  Scenario: Drone is sent to get an errand-with-ampersand containing
    a just a complex item
    Given a shopless DroneTown
    And 1 shop opens providing free "tomato"
    And 1 shop opens providing free "oregano"
    And 1 shop opens providing "salsa" requiring "tomato" AND "oregano"
    And 1 shop opens providing free "pasta"
    When I send Drone on an errand to get "pasta" AND "salsa"
    When a Drone arrives "Dock" with their backpack containing "pasta" AND "salsa"


  Scenario: Drone is sent to get an errand-with-ampersand containing
    a just a complex item but can't be done
    Given a shopless DroneTown
    And 1 shop opens providing free "oregano"
    And 1 shop opens providing "salsa" requiring "tomato" AND "oregano"
    And 1 shop opens providing free "pasta"
    When I send Drone on an errand to get "pasta" AND "salsa"
    Then a Drone arrives in "Junkyard" with their passport having
      | salsa          |          |
      | oregano        |          |
      | tomato         |          |
      | pasta          | pasta0   |


  Scenario: Drone is sent to get an errand-with-ampersand that can be
    done and there are multiple tomato shops
    Given a shopless DroneTown
    And 2 shop opens providing free "tomato"
    And 1 shop opens providing free "oregano"
    When I send Drone on an errand to get "tomato" AND "oregano"
    Then a Drone arrives "Dock" with their passport having
      | oregano | oregano0 |
      | tomato  | tomato0  |
    And a Drone arrives "Dock" with their passport having
      | oregano | oregano0 |
      | tomato  | tomato1  |


  Scenario: Drone is sent to get an errand-with-ampersand that can be
    done and there are multiple tomato and oregano shops
    Given a shopless DroneTown
    And 2 shop opens providing free "tomato"
    And 2 shop opens providing free "oregano"
    When I send Drone on an errand to get "tomato" AND "oregano"
    Then a Drone arrives "Dock" with their passport having
      | oregano | oregano0 |
      | tomato  | tomato0  |
    And a Drone arrives "Dock" with their passport having
      | oregano | oregano0 |
      | tomato  | tomato1  |
    And a Drone arrives "Dock" with their passport having
      | oregano | oregano1 |
      | tomato  | tomato1  |
    And a Drone arrives "Dock" with their passport having
      | oregano | oregano1 |
      | tomato  | tomato0  |


  Scenario: Drone is sent to get a simple errand that depends on a
    complex one and there are multiple salsa shops
    Given a shopless DroneTown
    And 1 shop opens providing free "tomato"
    And 1 shop opens providing free "oregano"
    And 3 shop opens providing "salsa" requiring "tomato" AND "oregano"
    When I send Drone on an errand to get "salsa"
    Then a Drone arrives "Dock" with their passport having
      | salsa          | salsa0   |
      | oregano        | oregano0 |
      | tomato         | tomato0  |
    And a Drone arrives "Dock" with their passport having
      | salsa          | salsa1   |
      | oregano        | oregano0 |
      | tomato         | tomato0  |
    And a Drone arrives "Dock" with their passport having
      | salsa          | salsa2   |
      | oregano        | oregano0 |
      | tomato         | tomato0  |
