package dronetown

import (
	"fmt"
	"log"

	"github.com/cucumber/godog"
)

type DroneTownTest struct {
	dt DroneTown
	d  Drone
}

func (dtt *DroneTownTest) aShoplessDroneTown() error {
	dtt.dt = DroneTown{}
	dtt.dt.Init()
	return nil
}

func (dtt *DroneTownTest) droneGetsToJunkyard(dronecount int) error {
	if len(dtt.dt.Junkyard) != dronecount {
		return fmt.Errorf(
			"%d Drones arrived Junkyard, but %d expected (%d at dock).",
			len(dtt.dt.Junkyard),
			dronecount,
			len(dtt.dt.Dock),
		)
	}
	return nil
}

func (dtt *DroneTownTest) droneGetsDock(dronecount int) error {
	if len(dtt.dt.Dock) != dronecount {
		return fmt.Errorf(
			"%d Drones at dock, but %d expected (%d at Junkyard).",
			len(dtt.dt.Dock),
			dronecount,
			len(dtt.dt.Junkyard),
		)
	}
	return nil
}

func (dtt *DroneTownTest) iSendDroneOnAnErrandToGet(item string) error {
	dtt.dt.OpenDock([]string{"fetch", item}, dtt.d.Backpack)

	return nil
}

func (dtt *DroneTownTest) shopsOpenProvidingFree(shopCount int, item string) error {

	for i := 0; i < shopCount; i++ {
		id := fmt.Sprintf("%s%d", item, i)
		dtt.dt.OpenShop(id, item, []string{})
	}

	return nil
}

func (dtt *DroneTownTest) shopOpensProvidingRequiring(qty int, provides, requires string) error {

	for i := 0; i < qty; i++ {
		id := fmt.Sprintf("%s%d", provides, i)
		dtt.dt.OpenShop(id, provides, []string{requires})
	}

	return nil
}

func (dtt *DroneTownTest) shopOpensProvidingRequiringAND(qty int, provides, requires1, requires2 string) error {

	for i := 0; i < qty; i++ {
		id := fmt.Sprintf("%s%d", provides, i)
		dtt.dt.OpenShop(id, provides, []string{"&", "2", requires1, requires2})
	}

	return nil
}

func (dtt *DroneTownTest) aDroneArrivesInWithTheirPassportHaving(where string, items *godog.Table) error {
	var passportExpected []Stamp
	for i := 0; i < len(items.Rows); i++ {
		item := items.Rows[i].Cells[0].Value
		id := items.Rows[i].Cells[1].Value
		log.Printf("[DEBUG] Passport item: %q, id: %q", item, id)
		passportExpected = append(passportExpected, Stamp{
			Item: item,
			ID:   id,
		})
	}
	log.Printf("[TEST] passport to check %#v", passportExpected)

	var dronesToCheck []Drone
	if where == "Junkyard" {
		dronesToCheck = dtt.dt.Junkyard
	} else if where == "Dock" {
		dronesToCheck = dtt.dt.Dock
	} else {
		return fmt.Errorf("[TEST] Malformed TEST, no such place %q", where)
	}

	err := fmt.Errorf("[TEST] No Drones in %s", where)
	for _, droneToCheck := range dronesToCheck {
		log.Printf("Drone has: %#v\n\tWe expect: %#v", droneToCheck.Passport, passportExpected)
		err = doesDroneHaveThisPassport(droneToCheck, passportExpected)
		if err == nil {
			return nil
		} else {
			log.Printf("[TEST] This Drone is not we expect 'cause: %q", err)
		}
	}
	return err
}

func doesDroneHaveThisPassport(drone Drone, passportToCheck []Stamp) error {
	dronePassport := drone.Passport

	if len(dronePassport) != len(passportToCheck) {
		return fmt.Errorf(
			"[TEST] passport length mismatch: %d != %d",
			len(dronePassport),
			len(passportToCheck),
		)
	}

	for i, item := range dronePassport {
		expected := passportToCheck[i]

		if item.Item != expected.Item {
			return fmt.Errorf(
				"[TEST] Not the same item: expected: %q, actual: %q",
				expected.Item,
				item.Item,
			)
		} else if item.ID != expected.ID {
			return fmt.Errorf(
				"[TEST] Not the same ID: expected: %q, actual: %q",
				expected.ID,
				item.ID,
			)
		}
	}

	return nil
}

func (dtt *DroneTownTest) aDroneArrivesWithTheirBackpackContaining(where string, items *godog.Table) error {
	var backpackExpected []string
	for i := 0; i < len(items.Rows); i++ {
		item := items.Rows[i].Cells[0].Value
		backpackExpected = append(backpackExpected, item)
	}
	log.Printf("[TEST] backpack to check %#v", backpackExpected)

	var dronesToCheck []Drone
	if where == "Junkyard" {
		dronesToCheck = dtt.dt.Junkyard
	} else if where == "Dock" {
		dronesToCheck = dtt.dt.Dock
	} else {
		return fmt.Errorf("[TEST] Malformed TEST, no such place %q", where)
	}

	err := fmt.Errorf("[TEST] No Drones in %s", where)
	for _, droneToCheck := range dronesToCheck {
		log.Printf("Drone has: %#v\n\tWe expect: %#v", droneToCheck.Backpack, backpackExpected)
		has := doesDroneHaveThisInTheirBackpack(droneToCheck, backpackExpected)
		if has {
			return nil
		} else {
			log.Printf("[TEST] This Drone does not have what we expect")
		}
	}

	return err
}

func (dtt *DroneTownTest) dronesBackpackContains(items *godog.Table) error {
	var backpack []string
	for i := 0; i < len(items.Rows); i++ {
		item := items.Rows[i].Cells[0].Value
		backpack = append(backpack, item)
	}
	log.Printf("[TEST] backpack to fill %#v", backpack)

	dtt.d.Backpack = backpack

	return nil
}

func doesDroneHaveThisInTheirBackpack(drone Drone, backpackExpected []string) bool {
	for _, item := range backpackExpected {
		has := hasInBackPack(drone.Backpack, item)
		if !has {
			return has
		}
	}
	return true
}

func doesDroneHaveThisItemInTheirBackpack(drone Drone, item string) bool {
	return hasInBackPack(drone.Backpack, item)
}

func (dtt *DroneTownTest) aDroneArrivesWithTheirBackpackContainingAND(where, item1, item2 string) error {
	var dronesToCheck []Drone
	if where == "Junkyard" {
		dronesToCheck = dtt.dt.Junkyard
	} else if where == "Dock" {
		dronesToCheck = dtt.dt.Dock
	} else {
		return fmt.Errorf("[TEST] Malformed TEST, no such place %q", where)
	}

	var err1, err2 error
	err := fmt.Errorf("[TEST] No Drones in %s", where)
	for _, droneToCheck := range dronesToCheck {
		log.Printf("Drone has: %#v\n\tWe expect: %q AND %q", droneToCheck.Backpack, item1, item2)
		has1 := doesDroneHaveThisItemInTheirBackpack(droneToCheck, item1)
		has2 := doesDroneHaveThisItemInTheirBackpack(droneToCheck, item1)
		if has1 && has2 {
			return nil
		} else {
			log.Printf("[TEST] This Drone does not have what we expect 'cause: %q and %q", err1, err2)
		}
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return err
}

func (dtt *DroneTownTest) iSendDroneOnAnErrandToGetAND(item1, item2 string) error {
	dtt.dt.OpenDock([]string{"fetch", "&", "2", item1, item2}, dtt.d.Backpack)
	return nil
}

func (dtt *DroneTownTest) iSendDroneOnAnErrandToGetOR(item1, item2 string) error {
	dtt.dt.OpenDock([]string{"fetch", "|", "2", item1, item2}, dtt.d.Backpack)
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {

	dtt := &DroneTownTest{}

	ctx.Step(`^a shopless DroneTown$`, dtt.aShoplessDroneTown)
	ctx.Step(`^(\d+) Drones? gets? to Junkyard$`, dtt.droneGetsToJunkyard)
	ctx.Step(`^(\d+) Drones? gets? dock$`, dtt.droneGetsDock)
	ctx.Step(`^I send Drone on an errand to get "([^"]*)"$`, dtt.iSendDroneOnAnErrandToGet)
	ctx.Step(`^(\d+) shops? opens? providing free "([^"]*)"$`, dtt.shopsOpenProvidingFree)
	ctx.Step(`^(\d+) shop opens providing "([^"]*)" requiring "([^"]*)"$`, dtt.shopOpensProvidingRequiring)
	ctx.Step(`^a Drone arrives in "([^"]*)" with their passport having$`,
		dtt.aDroneArrivesInWithTheirPassportHaving)
	ctx.Step(`^a Drone arrives "([^"]*)" with their passport having$`,
		dtt.aDroneArrivesInWithTheirPassportHaving)

	ctx.Step(`^a Drone arrives "([^"]*)" with their backpack containing$`, dtt.aDroneArrivesWithTheirBackpackContaining)
	ctx.Step(`^Drone\'s backpack contains$`, dtt.dronesBackpackContains)
	ctx.Step(`^a Drone arrives "([^"]*)" with their backpack containing "([^"]*)" AND "([^"]*)"$`,
		dtt.aDroneArrivesWithTheirBackpackContainingAND)
	ctx.Step(`^I send Drone on an errand to get "([^"]*)" AND "([^"]*)"$`,
		dtt.iSendDroneOnAnErrandToGetAND)
	ctx.Step(`^(\d+) shop opens providing "([^"]*)" requiring "([^"]*)" AND "([^"]*)"$`,
		dtt.shopOpensProvidingRequiringAND)
	ctx.Step(`^I send Drone on an errand to get "([^"]*)" OR "([^"]*)"$`, dtt.iSendDroneOnAnErrandToGetOR)
}
