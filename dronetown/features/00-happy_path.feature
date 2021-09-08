Feature: First herrands must be easy to test it up

  Scenario Outline: We send a Drone on an impossible mission
    Given a shopless DroneTown
    When I send Drone on an errand to get "<item>"
    Then 0 Drones get dock
    And 1 Drone gets to Junkyard

    Examples:
      | item     |
      | tomato   |
      | oregano  |
      | a beamer |


  Scenario Outline: We send Drone to do an easy task
    Given a shopless DroneTown
    And 1 shop opens providing free "<item>"
    When I send Drone on an errand to get "<item>"
    Then 1 Drone gets dock
    And 0 Drones get to Junkyard

    Examples:
      | item     |
      | tomato   |
      | oregano  |
      | a beamer |


  Scenario Outline: We send Drone Hunt on an impossible mission
    Given a shopless DroneTown
    And 1 shop opens providing free "unicorns"
    When I send Drone on an errand to get "<item>"
    Then 0 Drones get dock
    And 1 Drone gets to Junkyard

    Examples:
      | item     |
      | tomato   |
      | oregano  |
      | a beamer |


  Scenario Outline: We send Drone to do an easy task to a saturated market
    Given a shopless DroneTown
    And 42 shops open providing free "<item>"
    When I send Drone on an errand to get "<item>"
    Then 42 Drones get dock
    And 0 Drones get to Junkyard

    Examples:
      | item     |
      | tomato   |
      | oregano  |
      | a beamer |
