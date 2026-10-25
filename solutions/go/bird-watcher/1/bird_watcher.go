package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var counted int 
    for i := 0; i<=len(birdsPerDay)-1; i++{
        counted+=birdsPerDay[i]
    }
    return counted
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	start := (week - 1) * 7
	end := start + 7
	return TotalBirdCount(birdsPerDay[start:end])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	evenDays := len(birdsPerDay)/2
    for i:=0; i<evenDays; i++{
        birdsPerDay[i*2]+=1
    }
    return birdsPerDay
}
