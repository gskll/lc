// https://leetcode.com/problems/car-fleet/
package main

import "slices"

// sort the cars based on position
// it's a bunch of linear relationships with intersections
// calculate time to target for each car
// if car further arrives after or same time as car nearer target - become fleet
// car further slows down at some point as can't pass car nearer

func carFleet(target int, position []int, speed []int) int {
	var fleets []float32
	cars := make([][2]int, 0, len(position))
	for i := range position {
		cars = append(cars, [2]int{position[i], speed[i]})
	}
	slices.SortFunc(cars, func(car1, car2 [2]int) int {
		return car2[0] - car1[0]
	})

	for _, car := range cars {
		timeUntilTarget := float32(target-car[0]) / float32(car[1])
		if len(fleets) == 0 {
			fleets = append(fleets, timeUntilTarget)
			continue
		}

		prevCarTimeUntilTarget := fleets[len(fleets)-1]
		if prevCarTimeUntilTarget < timeUntilTarget {
			fleets = append(fleets, timeUntilTarget)
		}
	}

	return len(fleets)
}
