// Link: https://leetcode.com/problems/find-the-highest-altitude
package findthehighestaltitude

import "math"

// Complexity: O(n)
func largestAltitude(gain []int) int {
	highest := 0
	altitudes := make([]int, len(gain)+1)
	altitudes[0] = 0
	for i := range gain {
		j := i + 1
		if j <= len(gain) {
			altitudes[j] = altitudes[i] + gain[i]
			highest = int(math.Max(float64(highest), float64(altitudes[j])))
		}
	}
	return highest
}
