Feature: When Drone is asked to get an errand with and |, she has to
  visit HectOR's store in order to get it


  Scenario: Drone already has everthing required for their errand-with-pipe
    Given a shopless DroneTown
    And Drone's backpack contains
      | tomato |
    When I send Drone on an errand to get "tomato" OR "oregano"
    Then 1 Drones get dock
    And 0 Drone gets to Junkyard


  Scenario: Drone already has everthing required for their errand-with-pipe
    Given a shopless DroneTown
    And Drone's backpack contains
      | oregano |
    When I send Drone on an errand to get "tomato" OR "oregano"
    Then 1 Drones get dock
    And 0 Drone gets to Junkyard


  Scenario: Drone already has everthing required for their errand-with-pipe
    Given a shopless DroneTown
    And Drone's backpack contains
      | tomato  |
      | oregano |
    When I send Drone on an errand to get "tomato" OR "oregano"
    Then 1 Drones get dock
    And 0 Drone gets to Junkyard


  Scenario: Drone already is missing both items required for their errand-with-pipe
            and there is no hope to find it
    Given a shopless DroneTown
    When I send Drone on an errand to get "tomato" OR "oregano"
    Then a Drone arrives "Junkyard" with their passport having
      | tomato  |  |
    And a Drone arrives "Junkyard" with their passport having
      | oregano |  |


  Scenario: Drone is sent to get an errand-with-pipe that can be done
    Given a shopless DroneTown
    And 1 shop opens providing free "tomato"
    And 1 shop opens providing free "oregano"
    When I send Drone on an errand to get "tomato" OR "oregano"
    Then a Drone arrives "Dock" with their passport having
      | tomato          | tomato0 |
    And a Drone arrives "Dock" with their passport having
      | oregano         | oregano0 |


  Scenario: Drone is sent to get an errand-with-pipe that can be done
    Given a shopless DroneTown
    And 1 shop opens providing free "tomato"
    When I send Drone on an errand to get "tomato" OR "oregano"
    Then a Drone arrives "Dock" with their passport having
      | tomato          | tomato0 |
    And a Drone arrives in "Junkyard" with their passport having
      | oregano         |        |


  Scenario: Drone is sent to get an errand-with-pipe that can be done
    Given a shopless DroneTown
    And 1 shop opens providing free "oregano"
    When I send Drone on an errand to get "tomato" OR "oregano"
    Then a Drone arrives "Dock" with their passport having
      | oregano         | oregano0 |
    And a Drone arrives in "Junkyard" with their passport having
      | tomato          |        |


  Scenario: Drone is sent to get an errand-with-pipe that can be done
    Given a shopless DroneTown
    And 1 shop opens providing free "tomato"
    And 1 shop opens providing free "oregano"
    When I send Drone on an errand to get "tomato" OR "oregano"
    Then a Drone arrives "Dock" with their passport having
      | tomato          | tomato0 |
    And a Drone arrives "Dock" with their passport having
      | oregano         | oregano0 |


  Scenario: Drone is sent to get an errand-with-pipe that can be done
    Given a shopless DroneTown
    And 1 shop opens providing free "tomato"
    And 1 shop opens providing free "oregano"
    When I send Drone on an errand to get "tomato" OR "oregano"
    Then a Drone arrives "Dock" with their passport having
      | tomato          | tomato0 |
    And a Drone arrives "Dock" with their passport having
      | oregano         | oregano0 |
