package main

import "slices"

// https://leetcode.com/problems/car-fleet/

// There are n cars at given miles away from the starting mile 0, traveling to reach the mile target.
//
// You are given two integer array position and speed, both of length n, where position[i] is the starting mile of the ith car and speed[i] is the speed of the ith car in miles per hour.
//
// A car cannot pass another car, but it can catch up and then travel next to it at the speed of the slower car.
//
// A car fleet is a car or cars driving next to each other. The speed of the car fleet is the minimum speed of any car in the fleet.
//
// If a car catches up to a car fleet at the mile target, it will still be considered as part of the car fleet.
//
// Return the number of car fleets that will arrive at the destination.
//
// n == position.length == speed.length
// 1 <= n <= 10^5
// 0 < target <= 10^6
// 0 <= position[i] < target
// All the values of position are unique.
// 0 < speed[i] <= 10^6

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
