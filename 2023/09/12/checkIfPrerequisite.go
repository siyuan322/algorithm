package main

// floyd算法
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	// 矩阵，记录节点 i 到节点 j 是否可达
	f := make([][]bool, numCourses)
	for i := range f {
		f[i] = make([]bool, numCourses)
	}
	// 记录 pre i 到 j 可达
	for _, p := range prerequisites {
		f[p[0]][p[1]] = true
	}
	// 对于中间节点 k，枚举起始点 i 到终点 j 是否可达
	for k := 0; k < numCourses; k++ {
		for i := 0; i < numCourses; i++ {
			for j := 0; j < numCourses; j++ {
				f[i][j] = f[i][j] || (f[i][k] && f[k][j])
			}
		}
	}

	res := make([]bool, len(queries))
	for i, q := range queries {
		res[i] = f[q[0]][q[1]]
	}
	return res
}
