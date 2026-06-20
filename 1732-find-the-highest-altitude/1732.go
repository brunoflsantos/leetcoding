// Link: https://leetcode.com/problems/find-the-highest-altitude
package findthehighestaltitude

import "math"

// Complexity: O(n)
func largestAltitude(gain []int) int {
	highest := 0
	altitude := 0
	for i := range gain {
		altitude = altitude + gain[i]
		highest = int(math.Max(float64(highest), float64(altitude)))
	}
	return highest
}
