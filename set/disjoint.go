package set

type DisjointSet struct {
	rank  []int
	array []int
}

// O(N)
func NewDisjointSet(size int) *DisjointSet {
	ds := DisjointSet{}
	for i := 0; i < size; i++ {
		ds.array = append(ds.array, i)
		ds.rank = append(ds.rank, 1)
	}
	return &ds
}

func (s *DisjointSet) Array() []int {
	return s.array
}

// O(N) worst
func (ds *DisjointSet) Connected(x, y int) bool {
	return ds.slowFind(x) == ds.slowFind(y)
}

// O(log N) - path compression (average case, worst is still O(N))
func (ds *DisjointSet) slowFind(x int) int {
	if x == ds.array[x] {
		return x
	}
	ds.array[x] = ds.slowFind(ds.array[x])
	return ds.array[x]
}

// Optimisation for union by rank: now Union is O(log(N)) because we always
// choose the longest tree branch root as the root node so we are unable to
// reach worst case O(N) scenario where we get 0-1-2-3-4-5-6 tree
func (ds *DisjointSet) Union(x, y int) {
	rootX := ds.slowFind(x)
	rootY := ds.slowFind(y)

	if rootX == rootY {
		return
	}

	if ds.rank[rootX] > ds.rank[rootY] {
		ds.array[rootY] = rootX
	} else if ds.rank[rootX] < ds.rank[rootY] {
		ds.array[rootX] = rootY
	} else {
		ds.array[rootY] = rootX
		ds.rank[rootX] += 1
	}
}
