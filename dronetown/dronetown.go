/*
 * Copyright 2021 Banco Bilbao Vizcaya Argentaria, S.A.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dronetown

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
)

type DroneTown struct {
	capacity         sync.WaitGroup
	locationRegistry chan Location
	roadToTower      chan Drone
	roadToJunkyard   chan Drone
	Junkyard         []Drone
	Dock             []Drone
}

// Start all we nedd
func (dt *DroneTown) Init() {
	dt.locationRegistry = make(chan Location)
	dt.roadToTower = make(chan Drone)
	dt.roadToJunkyard = make(chan Drone)
	go dt.controlTower()
	go dt.JunkyardField()
	dt.ANDrés()
	dt.HéctOR()
}

func (dt *DroneTown) OpenDock(items []string, backpack []string) {
	log.Println("[Dock] Drone Dock Open")

	id := "dock"
	pathToDock := make(chan Drone)

	log.Println("[Dock] checking with control tower")
	dt.locationRegistry <- Address{
		ID:   id,
		Road: pathToDock,
	}

	dt.capacity.Add(1)
	b := List2Stack(append([]string{"go", id}, items...))
	drone := Drone{
		ID:       "Drone",
		Backpack: backpack,
		Brain:    b,
	}

	go func() {
		log.Println("[Dock] Drone send on a mission")
		dt.roadToTower <- drone
		for {
			drone := <-pathToDock
			log.Printf("[Dock] Drone back from mission: %v", drone)
			drone.Passport = buildPassport(drone)
			log.Printf("[Dock] let's see drone log: %v", drone.Passport)
			dt.Dock = append(dt.Dock, drone)
			dt.capacity.Done()
		}
	}()

	dt.capacity.Wait()
}

// TODO: move to their own file
func (dt *DroneTown) OpenShop(id, provides string, requires []string) {
	roadToShop := make(chan Drone)

	log.Printf("[%s] just open a shop providing %s", id, provides)
	dt.locationRegistry <- Shop{
		ID:   id,
		Prov: provides,
		Road: roadToShop,
	}
	goodByeDrone := func(drone Drone, b Brain, s Stack) {
		drone.Brain = rebuildDroneBrain(b, s)
		go func(droneSaliente Drone) {
			log.Printf("[%s] Drone just departured ...", id)
			dt.roadToTower <- clonaDrone(droneSaliente)
		}(drone)
	}

	buy := func(drone Drone) Action {
		return func(brain Brain, stack Stack) error {
			_, _ = stack.Pop()     //TODO use ok, is necessary
			item, _ := stack.Pop() //TODO use ok
			log.Printf("[%s] Drone landing search for %s", id, item)
			if item != provides {
				goodByeDrone(drone, brain, stack)
			}

			ok := hasMissingDeps(drone.Backpack, requires)
			if ok {
				log.Printf("[%s] adding deps: %v", id, requires)
				r := List2Stack(requires)
				for {
					i, cont := r.Pop()
					if !cont {
						break
					}
					stack.Push(i)
				}
				stack.Push("fetch")
				stack.Push(id)
				stack.Push("go")
				stack.Push(item)
				stack.Push(id)
				stack.Push("buy") //this restore our current instruction do
				log.Printf("[%s] addDeps? newStack: %v", id, stack)
			} else {
				// I have everything needed
				log.Printf("[%s] Stack: %v", id, stack)
				drone.Backpack = append(drone.Backpack, item)
				stack.Push(item)
				stack.Push(id)
				stack.Push("buyOK")
				log.Printf("[%s] buyOK? newStack: %v", id, stack)
			}
			goodByeDrone(drone, brain, stack)
			return nil
		}
	}

	go func() {
		for {
			drone := <-roadToShop
			log.Printf("[%s] Drone landing: %#v", id, drone)

			sel := map[string]Action{
				"buy": buy(drone),
			}

			RunShop(sel, drone.Brain) //TODO send drone to info if EndOfStack
		}
	}()
}

func (dt *DroneTown) controlTower() {
	log.Println("[Tower] Control Tower is start operations")
	yellowPages := make(map[string][]string) //TODO how is called this in a tower
	callejero := make(map[string]chan Drone) //TODO rename to skyroads

	sendDroneTo := func(drone Drone, to chan Drone, b Brain, s Stack) {
		drone.Brain = rebuildDroneBrain(b, s)
		go func(droneSaliente Drone) {
			log.Print("[Tower] Drone departured...")
			to <- clonaDrone(droneSaliente)
		}(drone)
	}

	sendDroneToJunkyard := func(drone Drone, b Brain, s Stack) {
		log.Print("[Tower] Destination set to Junkyard.")
		sendDroneTo(drone, dt.roadToJunkyard, b, s)
	}

	fetch := func(drone Drone) Action {
		return func(b, s Stack) error {
			item, _ := s.Pop() //TODO use ok if necessary

			log.Printf("[Tower] fetch item: %q", item)
			ids, ok := yellowPages[item]
			if ok {
				log.Printf("[Tower] shops found: %q", ids)
				for i, id := range ids {
					if i > 0 {
						dt.capacity.Add(1)
					}
					road := callejero[id]
					newStack := s.Fork()
					newStack.Push(item)
					newStack.Push(id)
					newStack.Push("buy")
					log.Printf("[Tower] Destination set to %q.", id)
					sendDroneTo(drone, road, b, newStack)
				}
			} else {
				log.Printf("[Tower] Shops not found: %v", item)
				s.Push(item)
				s.Push("fetch")
				sendDroneToJunkyard(drone, b, s)
			}

			return nil // armstrong
		}
	}

	buy := func(drone Drone) Action {
		return func(b, s Stack) error {
			log.Println("[Tower] Drone returning but not obey... so to the junkyard")
			s.Push("buy") //restore the stack action
			sendDroneToJunkyard(drone, b, s)
			return nil
		}
	}

	goToShop := func(drone Drone) Action {
		return func(b, s Stack) error {
			id, _ := s.Pop()
			road, ok := callejero[id]
			log.Printf("[Tower] Drone looking for %q", id)
			if ok {
				log.Printf("[Tower] Drone heading destination set to %q.", id)
				sendDroneTo(drone, road, b, s)
			} else {
				log.Printf("[Tower] Not found %q, send to Junkyard", id)
				sendDroneToJunkyard(drone, b, s)
			}
			return nil
		}
	}

	for {
		select {
		case drone := <-dt.roadToTower:
			log.Printf("[Tower] Drone arriving: %#v", drone)

			sel := map[string]Action{
				"buy":   buy(drone),
				"fetch": fetch(drone),
				"go":    goToShop(drone),
			}

			deport := RunShop(sel, drone.Brain)
			if deport != nil {
				log.Print("[Tower] We don't know how to help this Drone...")
				sendDroneToJunkyard(drone, drone.Brain, Stack{})
			}

		case address := <-dt.locationRegistry:
			log.Printf("[Tower] new location added to DroneTown: %q", address.GetID())
			callejero[address.GetID()] = address.GetRoad()
			switch v := address.(type) {
			case Shop:
				log.Printf("\t[Tower] seem to sell: %q", v.Prov)
				ids, ok := yellowPages[v.Prov]
				if !ok {
					ids = make([]string, 0)
				}
				ids = append(ids, address.GetID())
				yellowPages[v.Prov] = ids
			}
			// log.Printf("[INFO] yellow pages: %#v", yellowPages) //debug pruposes
		}
	}
}

func (dt *DroneTown) JunkyardField() {
	log.Println("[Junkyard] is now open")
	for {
		drone := <-dt.roadToJunkyard
		log.Println("[Junkyard] A Drone has arrived")
		drone.Passport = buildPassport(drone)
		log.Printf("[Junkyard] let's see logs: %v", drone.Passport)
		dt.Junkyard = append(dt.Junkyard, drone)
		dt.capacity.Done()
	}
}

func (dt *DroneTown) ANDrés() {
	road := make(chan Drone)
	id := "ANDrés"

	log.Printf("[%s] now serving & ;)", id)
	dt.locationRegistry <- Shop{
		ID:   id,
		Prov: "&",
		Road: road,
	}

	goodByeDrone := func(drone Drone, b Brain, s Stack) {
		drone.Brain = rebuildDroneBrain(b, s)
		go func(droneSaliente Drone) {
			log.Printf("[%s] Drone departuring...", id)
			dt.roadToTower <- clonaDrone(droneSaliente)
		}(drone)
	}

	buy := func(drone Drone) Action {
		return func(b, s Stack) error {
			_, _ = s.Pop() // id
			_, _ = s.Pop() //& //TODO not &!?
			ns, _ := s.Pop()
			n, _ := strconv.Atoi(ns)

			missing := []string{}

			for i := 0; i < n; i++ {
				elm, _ := s.Pop()
				if !hasInBackPack(drone.Backpack, elm) {
					missing = append(missing, elm)
				}
			}

			for _, miss := range missing {
				s.Push(miss)
				s.Push("fetch")
			}
			goodByeDrone(drone, b, s)

			return nil
		}
	}

	go func() {
		for drone := range road {
			log.Printf("[%s] Drone arriving: %v", id, drone)
			sel := map[string]Action{
				"buy": buy(drone),
			}

			RunShop(sel, drone.Brain)

			/*
				TODO some cases to check
				fetch oregano
				fetch oregano

				fetch & 2 tomato oregano
				fetch tomato fetch oregano

				fetch | 2 tomato oregano
				fetch tomato
				fetch oregano

				fetch & 2 pasta | 2 tomate nata
				fetch pasta fetch tomate
				fetch pasta fetch nata

				fetch | 2 pizza & 2 pasta & 2 tomato oregano
				fetch pizza
				fetch pasta fetch tomato fetch oregano

				combo breaker!!
				fetch | 2 & 2 pasta & 2 tomato oregano pizza
				fetch pasta fetch tomato fetch oregano
				fetch pizza

				extra extra
				fetch | 2 pizza & 2 pasta & 2 tomato oregano fetch pañales
				fetch pizza fetch pañales
				fetch pasta fetch tomato fetch oregano fetch pañales

			*/
		}
	}()
}

func (dt *DroneTown) HéctOR() {
	road := make(chan Drone)
	id := "HéctOR"

	log.Printf("[%s] now serving | ;)", id)
	dt.locationRegistry <- Shop{
		ID:   id,
		Prov: "|",
		Road: road,
	}

	enviarDrone := func(drone Drone) {
		go func(droneSaliente Drone) {
			log.Printf("[%s] Drone departuring...", id)
			dt.roadToTower <- droneSaliente
		}(drone)
	}

	buy := func(drone Drone) Action {
		return func(b, s Stack) error {
			_, _ = s.Pop() // id
			_, _ = s.Pop() //| //TODO not |!?
			ns, _ := s.Pop()
			n, _ := strconv.Atoi(ns)

			log.Printf("[%s] what we found on this Drone: %s %s", id, b, s)
			items := []string{}
			for i := 0; i < n; i++ {
				elm, _ := s.Pop()
				items = append(items, elm)
			}

			hasMissing := hasMissingDeps(drone.Backpack, append([]string{"|", ns}, items...))

			if hasMissing {
				for i, elm := range items {
					d := clonaDrone(drone)
					if i > 0 {
						dt.capacity.Add(1)
						log.Printf("[%s] Using clonadrone from Aliexpress!", id)
					}
					if !hasInBackPack(drone.Backpack, elm) {
						log.Printf("[%s] %s needed! new errand added", id, elm)
						newStack := s.Fork()
						newStack.Push(elm)
						newStack.Push("fetch")

						d.Brain = rebuildDroneBrain(b, newStack)
						log.Printf("[%s] look good to me: %v", id, d)
						enviarDrone(d)
					} else {
						//this will never happend, any futfill errand will prevent it
						panic("you never must reach this place")
					}
				}
			} else {
				log.Printf("[%s] Turn arround little drone, no missing items on you", id)
				drone.Brain = rebuildDroneBrain(b, s)
				enviarDrone(drone)
			}

			// [tomate]
			// & 2 pasta tomate
			// go italiano0 fetch & 2 pasta tomate
			// go italiano0 buy andres & 2 pasta tomate
			// go italiano0 fetch pasta
			// go italiano0 buy pasta0 pasta
			// [tomate, pasta]
			// go italiano0 buyOK pasta0 pasta
			// TODO buyOK composed cases
			// & 2 pasta tomate ?? true

			return nil
		}
	}

	go func() {
		for drone := range road {
			log.Printf("[%s] Drone arriving: %v", id, drone)
			sel := map[string]Action{
				"buy": buy(drone),
			}

			RunShop(sel, drone.Brain)
		}
	}()
}

// Utilities TODO: refucktor
func hasInBackPack(backpack []string, target string) bool {
	for _, item := range backpack {
		if item == target {
			return true
		}
	}
	return false
}

func rebuildDroneBrain(
	currentBrain Stack,
	currentStack Stack,
) Stack {
	brain := make([]string, len(currentBrain.items))
	copy(brain, currentBrain.items)
	stack := make([]string, len(currentStack.items))
	copy(stack, currentStack.items)
	s := List2Stack(stack)
	b := List2Stack(brain)
	for {
		item, ok := s.Pop()
		if !ok {
			break
		}
		b.Push(item)
	}
	return b
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func clonaDrone(drone Drone) Drone {
	brain := make([]string, len(drone.Brain.items))
	copy(brain, drone.Brain.items)
	backpack := make([]string, len(drone.Backpack))
	copy(backpack, drone.Backpack)
	return Drone{
		ID:       RandStringRunes(21),
		Backpack: backpack,
		Brain:    List2Stack(brain),
	}
}

func buildPassport(drone Drone) []Stamp {
	//TODO build passport without use private items
	st := make([]Stamp, 0)
	log.Printf("[PASSPORT] Brain dump: %s", drone.Brain)
	for i, item := range drone.Brain.items {
		if item == "buyOK" {
			st = append(st, Stamp{
				ID:   drone.Brain.items[i+1],
				Item: drone.Brain.items[i+2],
			})
		} else if item == "buy" {
			st = append(st, Stamp{
				ID:   "",
				Item: drone.Brain.items[i+2],
			})
		} else if item == "fetch" {
			st = append(st, Stamp{
				ID:   "",
				Item: drone.Brain.items[i+1],
			})
		}
	}

	return st
}

func hasMissingDeps(backpack, req []string) bool {
	satisfied := false
	stack := Stack{}
	br := List2Stack(req)

	if len(backpack) == 0 && len(req) > 0 {
		return true
	}

	for {
		item, ok := br.Pop()
		if !ok {
			break
		}

		switch item {
		case "&":
			ns, _ := stack.Pop()
			n, _ := strconv.Atoi(ns)
			for i := 0; i < n; i++ {
				req, _ := stack.Pop()

				if !hasInBackPack(backpack, req) {
					return true
				}
			}
		case "|":
			ns, _ := stack.Pop()
			n, _ := strconv.Atoi(ns)
			for i := 0; i < n; i++ {
				req, _ := stack.Pop()

				if hasInBackPack(backpack, req) {
					return false
				}
			}
		default:
			stack.Push(item)
		}
	}

	return satisfied
}
