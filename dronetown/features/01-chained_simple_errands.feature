Feature: Can be a lot of shops on DroneTown, so we need to find the way


  Scenario Outline: TBD: Two single level chain
    Given a shopless DroneTown
    And <salsashops> shop opens providing "salsa" requiring "tomatoes"
    And <tomatoeshops> shop opens providing free "tomatoes"
    When I send Drone on an errand to get "salsa"
    Then <dock> Drones get dock
    And <junkyard> Drone gets to Junkyard

    Examples:
      | salsashops | tomatoeshops | dock | junkyard |
      | 1          | 0            | 0    | 1        |
      | 1          | 1            | 1    | 0        |
      | 2          | 0            | 0    | 2        |
      | 2          | 1            | 2    | 0        |
      | 2          | 2            | 4    | 0        |


  Scenario Outline: TBD: Three single level chain
    Given a shopless DroneTown
    And <pastashops> shop opens providing "pasta" requiring "salsa"
    And <salsashops> shop opens providing "salsa" requiring "tomatoes"
    And <tomatoeshops> shop opens providing free "tomatoes"
    When I send Drone on an errand to get "pasta"
    Then <dock> Drones get dock
    And <junkyard> Drone gets to Junkyard

    Examples:
      | pastashops | salsashops | tomatoeshops | dock | junkyard |
      | 0          | 0          | 0            | 0    | 1        |
      | 0          | 0          | 1            | 0    | 1        |
      | 0          | 0          | 2            | 0    | 1        |
      | 0          | 1          | 0            | 0    | 1        |
      | 0          | 1          | 1            | 0    | 1        |
      | 0          | 1          | 2            | 0    | 1        |
      | 0          | 2          | 0            | 0    | 1        |
      | 0          | 2          | 1            | 0    | 1        |
      | 0          | 2          | 2            | 0    | 1        |
      | 1          | 0          | 0            | 0    | 1        |
      | 1          | 0          | 1            | 0    | 1        |
      | 1          | 0          | 2            | 0    | 1        |
      | 1          | 1          | 0            | 0    | 1        |
      | 1          | 1          | 1            | 1    | 0        |
      | 1          | 1          | 2            | 2    | 0        |
      | 1          | 2          | 0            | 0    | 2        |
      | 1          | 2          | 1            | 2    | 0        |
      | 1          | 2          | 2            | 4    | 0        |
      | 2          | 0          | 0            | 0    | 2        |
      | 2          | 0          | 1            | 0    | 2        |
      | 2          | 0          | 2            | 0    | 2        |
      | 2          | 1          | 0            | 0    | 2        |
      | 2          | 1          | 1            | 2    | 0        |
      | 2          | 1          | 2            | 4    | 0        |
      | 2          | 2          | 0            | 0    | 4        |
      | 2          | 2          | 1            | 4    | 0        |
      | 2          | 2          | 2            | 8    | 0        |


  Scenario: TBD: Several shops do not make a Drone to be lost
    the product offered by the last level
    Given a shopless DroneTown
    And 2 shop opens providing "salsa" requiring "tomatoes"
    And 1 shop opens providing free "tomatoes"
    When I send Drone on an errand to get "tomatoes"
    Then 1 Drones get dock
    And 0 Drone gets to Junkyard
