// Link: https://leetcode.com/problems/pascals-triangle-ii
package pascaltriangle2

// Complexity: O(n^2)
func getRow(rowIndex int) []int {
	triangleMap := make([][]int, rowIndex+1)
	for i := range triangleMap {
		triangleMap[i] = make([]int, rowIndex+1)
		triangleMap[i][0], triangleMap[i][i] = 1, 1
		for j := 1; j < i; j++ {
			triangleMap[i][j] = triangleMap[i-1][j-1] + triangleMap[i-1][j]
		}
	}
	return triangleMap[rowIndex]
}
