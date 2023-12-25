package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Swop/advent-of-code/pkg/runner"
	"github.com/gammazero/deque"
)

func main() {
	runner.Run(part1, func(_ []string) any { return 0 })
}

func part1(input []string) any {
	// For this puzzle, we'll use a graph visualisation of the input, and we'll manually find the min-cut of the graph
	// (3 edges), either for the real input or for the sample input.
	// (see documentation of generateGraphVisualization function below).
	// The following line (commented by default) has been used to generate a graphviz graph file, which can be used
	// to manually find the three edges (min-cut), which is used in the resolvePart1 function below.
	//
	// generateGraphVisualization(parseInput(input))
	//
	return resolvePart1(
		input,
		// the three handpicked edges to remove (from real input). For the sample input, see main_test.go.
		[3][2]string{{"vtt", "fht"}, {"bbg", "kbr"}, {"czs", "tdk"}},
	)
}

type state struct {
	group int
	name  string
}

func resolvePart1(input []string, edgeToRemove [3][2]string) any {
	edges, revEdges := parseInput(input)
	q := deque.New[state]()

	// First, remove the three edges (min cut) from the graph.
	// (Those edges where handpicked, based on graphviz visualisation).
	// This is central to the solving, since this will split the graph into two groups of vertices, and we'll then
	// be able to count the number of vertices in each group easily with a simple graph traversal of each group.
	removeEdge := func(edges map[string][]string, n1, n2 string) {
		newEdgeList := make([]string, 0, len(edges[n1]))
		for _, e2 := range edges[n1] {
			if e2 != n2 {
				newEdgeList = append(newEdgeList, e2)
			}
		}
		edges[n1] = newEdgeList
	}
	for _, e := range edgeToRemove {
		// Remove both the edge and its reverse edge, for each of the three handpicked edges.
		removeEdge(edges, e[0], e[1])
		removeEdge(revEdges, e[1], e[0])
	}

	// Then we then start a traversal of the two resulting split graphs to count the number of vertices in each group.
	visited := map[string]struct{}{}
	verticesCount := [2]int{0, 0}
	q.PushBack(state{group: 0, name: edgeToRemove[0][0]})
	q.PushBack(state{group: 1, name: edgeToRemove[0][1]})
	for q.Len() > 0 {
		s := q.PopFront()
		if _, ok := visited[s.name]; ok {
			continue
		}
		visited[s.name] = struct{}{}
		verticesCount[s.group]++
		for _, e := range edges[s.name] {
			q.PushBack(state{group: s.group, name: e})
		}
		for _, e := range revEdges[s.name] {
			q.PushBack(state{group: s.group, name: e})
		}
	}
	return verticesCount[0] * verticesCount[1]
}

// generateGraphVisualization will generate a graphviz graph in the file graph.svg using the neato layout algorithm.
// This will easily highlight the two groups of vertices (and the three edges linking the two groups).
// This is not ideal, since this is not a code-based resolution of the problem, but this avoids having to write a
// complex algorithm to find the min-cut (like Karger's one), or having to rely on a 3rd-party graph library.
//
//nolint:unused
func generateGraphVisualization(edges map[string][]string) {
	dotStatements := "digraph D {\n"
	for src, dst := range edges {
		for _, d := range dst {
			dotStatements += fmt.Sprintf("%s -> %s\n", src, d)
		}
	}
	dotStatements += "}"

	f, err := os.Create("graph.svg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// use neato graphviz layout, which is better to highlight specificities of the graph
	cmd := exec.Command("dot", "-Tsvg", "-Kneato")
	var b bytes.Buffer
	cmd.Stdout = f
	cmd.Stderr = &b
	cmd.Stdin = strings.NewReader(dotStatements)
	err = cmd.Run()
	out := b.String()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "OUT: %s\n", out)
		panic(fmt.Errorf("error when running graphviz: %w", err))
	}

	fmt.Println("Graph generated in " + f.Name())
}

func parseInput(input []string) (map[string][]string, map[string][]string) {
	edges := make(map[string][]string)
	revEdges := make(map[string][]string)
	for _, line := range input {
		p := strings.Split(line, ":")
		p2 := strings.Fields(p[1])
		edges[p[0]] = append(edges[p[0]], p2...)
		// Graph is undirected, we need to consider reverse edges too.
		// We're storing those in a separate map to avoid overloading the graphviz visualisation with too many edges.
		// We'll however consider those during our traversal (to count the number of vertices in each group).
		for _, p3 := range p2 {
			revEdges[p3] = append(revEdges[p3], p[0])
		}
	}
	return edges, revEdges
}
