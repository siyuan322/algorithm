package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	// p[1]是 p[0]的前置课程
	// 构建邻接链表，记录每个节点的后置课程
	edges := make([][]int, numCourses)
	// 每个节点的入度
	indegree := make([]int, numCourses)
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
		indegree[p[0]]++
	}
	var queue []int
	for idx, in := range indegree {
		// 找到入度为 0 的课程
		if in == 0 {
			queue = append(queue, idx)
		}
	}
	// 记录每个
	result := make([]int, 0, numCourses)
	// 广度遍历
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		result = append(result, cur)
		numCourses--

		// 将 cur 的后置节点的入度-1
		for _, node := range edges[cur] {
			indegree[node]--
			if indegree[node] == 0 {
				queue = append(queue, node)
			}
		}
	}
	if numCourses != 0 {
		return []int{}
	}
	return result
}
