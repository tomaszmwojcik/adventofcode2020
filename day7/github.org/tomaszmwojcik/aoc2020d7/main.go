package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
light red bags contain 1 bright white bag, 2 muted yellow bags.
faded blue bags contain no other bags.
*/

var bags map[string]*Node
var visitedBags map[string]bool

// BagRule represents single bag rule
type BagRule struct {
	node  *Node
	count int
}

// Node represents a bag
type Node struct {
	Name string
	Prev []BagRule
	Next []BagRule
}

// ParseBag parses bagDesc, returns bag name and bag count
// bagDesc might contain e.g. "light red bags  ", " 2 muted yellow bags."
func ParseBag(bagDesc string) (bagName string, count int) {
	bagDesc = strings.Trim(bagDesc, ". ")
	bagDesc = bagDesc[0:strings.LastIndex(bagDesc, "bag")]
	countOp, err := strconv.Atoi(strings.Split(bagDesc, " ")[0])
	if err == nil {
		count = countOp
	}
	bagName = strings.Trim(bagDesc, "0123456789 ,.")
	return
}

// GetBag potentially initializes and returns a node associated with the bag
func GetBag(bagName string) *Node {
	bag, isPresent := bags[bagName]
	if !isPresent {
		bag = &Node{bagName, make([]BagRule, 0), make([]BagRule, 0)}
		bags[bagName] = bag
		//fmt.Printf("Added %s bag\n", bagName)
	}
	return bag
}

func countParents(bag *Node) int {
	count := 1 // for this bag itself
	for i := 0; i < len(bag.Prev); i++ {
		prevBag := bag.Prev[i].node
		//fmt.Printf("Checking %s\n", prevBag.Name)
		_, isPresent := visitedBags[prevBag.Name]
		if !isPresent {
			visitedBags[prevBag.Name] = true
			count += countParents(prevBag)
		}
	}
	return count
}

func countInsideBags(bag *Node) int {
	count := 1
	for i := 0; i < len(bag.Next); i++ {
		insideBag := bag.Next[i].node
		insideBagCount := bag.Next[i].count
		//fmt.Printf("Checking %s\n", prevBag.Name)
		count += insideBagCount * countInsideBags(insideBag)
	}
	return count
}

func main() {
	bags = make(map[string]*Node)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}
		split := strings.Split(line, "contain")
		bagName, _ := ParseBag(split[0])
		bag := GetBag(bagName)
		if strings.Index(split[1], "no other bags") == -1 {
			contents := strings.Split(split[1], ",")
			for i := 0; i < len(contents); i++ {
				insideBagName, insideBagCount := ParseBag(contents[i])
				insideBag := GetBag(insideBagName)
				insideBag.Prev = append(insideBag.Prev, BagRule{bag, insideBagCount})
				bag.Next = append(bag.Next, BagRule{insideBag, insideBagCount})
			}
		}
		if err != nil {
			break
		}
	}
	bag, _ := bags["shiny gold"]
	//fmt.Printf("%v\n", bag)
	visitedBags = make(map[string]bool)
	parentsCount := countParents(bag) - 1 //excluding shiny gold bag
	fmt.Printf("There are %d bag colors that can contain a shiny gold bag.\n", parentsCount)
	insideCount := countInsideBags(bag) - 1 //excluding shiny gold bag
	fmt.Printf("There must be %d bags inside a shiny gold bag.\n", insideCount)
}
