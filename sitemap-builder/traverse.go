package main

import (
	"fmt"
	"log"
)

type Queue []string

var (
	visited map[string]bool
)

func (q *Queue) enqueue(ele string) error {
	(*q) = append((*q), ele)
	return nil
}

func (q *Queue) dequeue() (string, error) {
	if len(*q) == 0 {
		return "", fmt.Errorf("cannot dequeue from an empty queue!")
	}
	res := (*q)[0]
	(*q) = (*q)[1:]
	return res, nil
}

// Traverse through the target URLs in breadth-first approach
func bfs(baseURL string, depth int) error {

	var q Queue
	q.enqueue(baseURL)
	visited = make(map[string]bool)

	visited[baseURL] = true
	var level, nodesLeftInLayer, nodesInNextLayer = 0, 1, 0

	for len(q) > 0 && level < depth {
		ele, err := q.dequeue()
		if err != nil {
			return err
		}

		nodesChan := make(chan string, 0)
		errChan := make(chan error, 0)
		go exploreNeighbors(ele, nodesChan, errChan)
		go func() {
			err = <-errChan
			log.Printf("Error: %+v at element %v\n", err, ele)
			return
		}()
		for neighbor := <-nodesChan; neighbor != ""; neighbor = <-nodesChan {
			if !visited[neighbor] {
				fmt.Println("Neighbor: ", neighbor)
				visited[neighbor] = true
				q.enqueue(neighbor)
				nodesInNextLayer++
			}
		}
		nodesLeftInLayer--

		if nodesLeftInLayer == 0 {
			nodesLeftInLayer = nodesInNextLayer
			nodesInNextLayer = 0
			level++
		}
	}

	return nil
}

// Explore the neighbors for that element
func exploreNeighbors(element string, neighborsChan chan string, errChan chan error) {
	defer close(neighborsChan)
	defer close(errChan)

	res, err := fetchResponse(element)
	if err != nil {
		errChan <- err
		return
	}

	links, err := parseResponse(res, element)
	if err != nil {
		errChan <- err
		return
	}

	for _, link := range links {
		neighborsChan <- link
	}

}

// To display / debug all the visited links
func displayLinks() {
	for link, _ := range visited {
		fmt.Println("Link: ", link)
	}
}
