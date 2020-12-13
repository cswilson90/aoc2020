package day13

import (
	"strconv"
	"strings"
)

const MaxInt = int(^uint(0) >> 1)

type Notes struct {
	StartTime   int
	ActiveBuses []int
}

type Departure struct {
	Time  int
	BusNo int
}

// Parses note input into a Notes object
func ParseBusNotes(notes []string) (*Notes, error) {
	startTime, err := strconv.Atoi(notes[0])
	if err != nil {
		return nil, err
	}

	allBuses := strings.Split(notes[1], ",")
	activeBues := make([]int, 0)
	for _, v := range allBuses {
		if v == "x" {
			continue
		}
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		activeBues = append(activeBues, num)
	}

	return &Notes{
		StartTime:   startTime,
		ActiveBuses: activeBues,
	}, nil
}

// Finds the next bus departure for a given set of notes
func NextBusDeparture(notes *Notes) *Departure {
	nextTime := MaxInt
	nextBus := 0

	for _, busNum := range notes.ActiveBuses {
		minsSinceBus := notes.StartTime % busNum
		nextBusTime := notes.StartTime
		if minsSinceBus != busNum {
			nextBusTime = notes.StartTime + (busNum - minsSinceBus)
		}

		if nextBusTime < nextTime {
			nextTime = nextBusTime
			nextBus = busNum
		}
	}

	return &Departure{
		Time:  nextTime,
		BusNo: nextBus,
	}
}

// Get the time when the bus departures match the given schedule
func GetFirstTimeOfPattern(schedule string) (int, error) {
	busList := strings.Split(schedule, ",")
	busNums := make([]int, len(busList))

	for i, v := range busList {
		// An 'x' bus can be represented by ID 1 so it matches any time
		if v == "x" {
			busNums[i] = 1
		} else {
			num, err := strconv.Atoi(v)
			if err != nil {
				return 0, err
			}
			busNums[i] = num
		}
	}

	timeStamp := 0
	for {
		// Keep going until the whole schedule matches
		matches := patternMatches(busNums, timeStamp)
		if matches >= len(busNums) {
			break
		}

		// We can skip forward by the buses which match multiplied together
		// as we know that's the next time that number of buses will match again
		timeJump := busNums[0]
		for i := 1; i < matches; i++ {
			timeJump *= busNums[i]
		}
		timeStamp += timeJump
	}

	return timeStamp, nil
}

// The number of buses in the schedule which match the given start time
func patternMatches(busNums []int, time int) int {
	for i, v := range busNums {
		if (time+i)%v != 0 {
			return i
		}
	}
	return len(busNums)
}
