package day12

func bfs(w World, part2 bool) (int, error) {
	// Keep track of how many steps from every visited tile back to the start.
	dist := map[*Tile]int{w.To(): 0}

	// The working queue of visited tiles.
	queue := []*Tile{w.To()}

	// Keep track of the shortest route for part 2.
	var shortest *Tile

	for len(queue) > 0 {
		// Pop the first tile off the front of the queue.
		cur := queue[0]
		queue = queue[1:]

		// Set shortest if it is not already set.
		if cur.Kind == 'a' && shortest == nil {
			shortest = cur
		}

		// The neighbours are referenced quite a lot below, so call this method just the once up here.
		neighbours := cur.neighbours()

		// Do these in index order each time to honour the up/down/left/right direction sequencing.
		for index := range neighbours {
			if _, seen := dist[neighbours[index]]; !seen {
				dist[neighbours[index]] = dist[cur] + 1
				queue = append(queue, neighbours[index])
			}
		}
	}

	if !part2 {
		return dist[w.From()], nil
	} else {
		return dist[shortest], nil
	}
}
