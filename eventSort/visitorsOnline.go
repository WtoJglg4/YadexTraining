package eventsort

import (
	"fmt"
	"github/YandexTraining/mathem"
)

const (
	userIn  = -1
	userOut = 1
)

func sortEvents(pevents *[][]int) {
	events := *pevents
	for i := 0; i < len(events)-1; i++ {
		for j := i + 1; j < len(events); j++ {
			if events[i][0] > events[j][0] {
				events[i], events[j] = events[j], events[i]
			}
			if events[i][0] == events[j][0] && events[i][1] > events[j][1] {
				events[i][1], events[j][1] = events[j][1], events[i][1]
			}
		}
	}
}

func MaxVisitorsOnline(n int, tin, tout []int) int {
	events := make([][]int, 2*n)
	eventIndex := 0
	for i := 0; i < 2*n; i += 2 {
		events[i] = []int{tin[eventIndex], userIn}
		events[i+1] = []int{tout[eventIndex], userOut}
		eventIndex++
	}
	sortEvents(&events)

	online, maxOnline := 1, 1
	for i := 1; i < len(events); i++ {
		if events[i][1] == userIn {
			online++
		} else {
			online--
		}
		maxOnline = mathem.Max(online, maxOnline)
	}
	fmt.Println(events)
	return maxOnline
}
