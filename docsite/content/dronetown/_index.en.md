---
title: "Dronetown"
date: 2021-10-28T09:05:40+02:00
draft: true
---

DroneTown is a support metaphor to help to understand Overlord system needs, because its level of abstraction can cause uncertainty in the implementation details.

## Metaphore cast

Within DroneTown we have several actors that will represent parts of a hypothetical real Overlord system. To better understand the features we should black-box its internals until we understand the system’s core


### DroneTown
The city itself is an integral part of the metaphor, it represents the system as a whole without worrying about its internal organization.

All DroneTown members who are represented as a building/shop act in a similar way. When a drone comes to them, they sequentially process their brain’s content until they find an order they understand. At that moment they carry out their task. And if they don't find an order they understand, they send the drone to a default site (which is usually the default control tower)


### Dock
The drones leave the Dock to carry out their errands. These drones only have two orders in their brains:

 * To obtain their target information 
 * Return to the Dock

If a drone has no target, the tower will attend its second instruction and send it back to the Dock.

If, on the contrary, it does not have the instruction to return to the Dock, the drone will try to obtain its target, and then end up in the Junkyard whether it succeeds or not.


### Junkyard
It is the place where the drones arrive if they do not have a set of correct orders in the brain or if they cannot fulfill their objective.

Inspecting these drones can reveal unknown requirements from the drone's objectives, so they return information about their journey even if it was not asked for.

### Torre
The tower is a structural DroneTown building, it has two responsibilities:

 * Record all DroneTown addresses
 * Register all DroneTown stores

If the tower finds an instruction to send the drone to a specific address, it simply deletes this instruction and sends the drone to that address. This is the common mechanism used by other shops when they have pending processes

If the tower finds an instruction to obtain a certain item, it searches the site among its addresses, and sends it to that address with an instruction to buy.

It is possible that there is more than one store that supplies a certain product, so in that case the tower will dispatch as many drones as there are stores to ensure that all possibilities are explored. Each of these drones is a complete clone of the one that entered the tower, including its belongings and instructions.

It is also possible that a drone is targeting an item that is not registered in the tower. In this case the tower will send the drone to the Junkyard. The tower is the only building that will send a drone to the Junkyard when it doesn't know how to proceed. The rest of the buildings send the drone to the tower if they don't know what to do.

### Shop
All stores work in a similar way. They can carry out three basic operations:
 
 * In the event that a drone arrives, and they have everything it needs, they give the item the drone is looking for and change the purchase operation for a successful purchase.
 * In the event that a drone arrives, but they lack what it needs, they will add an instruction to return to the store and then another instruction to get what they lack.
 * In any case they always send the drone back to the tower.

## Special transactions

Inside Drone Town, and without breaking its rules, there are some curious shops that allow you to control the most complex errands. In the case that a store, for example, can obtain an item from others, we need to tell the drone that it can obtain one or the other. Similar to a logical OR. It is also possible that a store requires more than one item to carry out the transaction, so in the same way, we need to tell the drone that we need both one item and the next. Similar to an AND.

To solve these complexities there are two special stores

### ANDrés Store
This store takes care of the orders and requirements that require a logical AND.
When a drone arrives looking for its item (the AND), they examine the AND members. For each of the items, it examines whether the drone already has them or not. And in the case of not having them, it adds an instruction to buy that product.


Always, before sending the drone back to the tower, it removes any traces of its activity to keep things simple.

### HectOR Store
In this case the store works in a similar way. It inspects the backpack of the drone, if it already has any of the items you need, erase the trace of the store and send it to the tower.

If it doesn't have any of the items, HectOR will dispatch as many cloned drones as items in the instruction. This ensures that they will have the maximum number of chances to complete the errand successfully. This operation is similar to the one carried out by the tower when there are several stores, but in the case that there are several requirements.
