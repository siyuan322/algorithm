package main

// https://leetcode.cn/problems/course-schedule/description/?envType=daily-question&envId=2023-09-09

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 关键：寻找入度为 0 的节点
	// 构建邻接链表
	// 记录每个节点的后置课程
	grids := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, p := range prerequisites {
		// p[1]是 p[0]的前置课程
		grids[p[1]] = append(grids[p[1]], p[0])
		// 计算 p[0] 的入度
		indegree[p[0]]++
	}
	// 将入度为 0 的节点加入队列中
	queue := []int{}
	for idx, in := range indegree {
		if in == 0 {
			queue = append(queue, idx)
		}
	}

	for len(queue) > 0 {
		cur := queue[0] // 当前课程
		queue = queue[1:]

		// 已修一个课程
		numCourses--
		// 将这个节点指向的入度都-1
		for _, n := range grids[cur] {
			indegree[n]--
			if indegree[n] == 0 {
				queue = append(queue, n)
			}
		}
	}
	return numCourses == 0
}
