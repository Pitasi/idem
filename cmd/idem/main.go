package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	antonio := GetUser("antonio")
	roberto := GetUser("roberto")
	edoardo := GetUser("edoardo")

	antonio.Like("a")

	roberto.Like("b")
	roberto.Like("a")

	edoardo.Like("a")
	edoardo.Like("b")
	edoardo.Like("c")

	suggestions := roberto.Recommend()

	fmt.Println("Suggestions:")
	for _, s := range suggestions {
		fmt.Printf("- %s (score %d)\n", s.Item, s.Count)
	}
}

type ItemCount struct {
	Item  interface{}
	Count int
}

type ItemCounter struct {
	counts map[interface{}]int
	keys   []interface{}
}

func NewItemCounter() *ItemCounter {
	return &ItemCounter{
		counts: make(map[interface{}]int),
	}
}

func (s *ItemCounter) Inc(item interface{}) {
	s.IncBy(item, 1)
}

func (s *ItemCounter) IncBy(item interface{}, by int) {
	_, found := s.counts[item]
	if !found {
		s.keys = append(s.keys, item)
	}
	s.counts[item] = by
}

func (s *ItemCounter) Items() []ItemCount {
	res := make([]ItemCount, len(s.keys))
	for i, k := range s.keys {
		res[i] = ItemCount{k, s.counts[k]}
	}
	return res
}

func (s *ItemCounter) Sort() {
	sort.Sort(s)
}

func (s *ItemCounter) Len() int {
	return len(s.keys)
}

func (s *ItemCounter) Less(i, j int) bool {
	return s.counts[s.keys[i]] > s.counts[s.keys[j]]
}

func (s *ItemCounter) Swap(i, j int) { s.keys[i], s.keys[j] = s.keys[j], s.keys[i] }

////////////////////////////////////////////////////////////////////

var g = NewGraph()

type Node interface{}

type Graph struct {
	Relations map[Node]map[Node]int
}

func NewGraph() *Graph {
	return &Graph{
		Relations: make(map[Node]map[Node]int),
	}
}

// TODO use only half the memory
// O(1)
func (g *Graph) Relate(a, b Node) {
	if a == b {
		return
	}
	if _, f := g.Relations[a]; !f {
		g.Relations[a] = make(map[Node]int)
	}
	if _, f := g.Relations[b]; !f {
		g.Relations[b] = make(map[Node]int)
	}
	g.Relations[a][b]++
	g.Relations[b][a]++
}

//
// O(1)
func (g *Graph) FindRelations(a Node) map[Node]int {
	return g.Relations[a]
}

func (g *Graph) String() string {
	nodes := make([]string, len(g.Relations))
	var i int
	for n, nRelations := range g.Relations {
		relations := make([]string, len(nRelations))
		var j int
		for m, score := range nRelations {
			relations[j] = fmt.Sprintf("%s => %s (score %d)", n, m, score)
			j++
		}
		nodes[i] = strings.Join(relations, "\n")
		i++
	}
	return strings.Join(nodes, "\n")
}

////////////////////////////////////////////////////////////////////

var users = make(map[interface{}]*User)

func GetUser(id interface{}) *User {
	if _, found := users[id]; !found {
		users[id] = NewUser()
	}
	return users[id]
}

type User struct {
	likes map[Node]interface{}
}

func NewUser() *User {
	return &User{
		likes: make(map[Node]interface{}),
	}
}

func (u *User) Like(thing Node) {
	for other := range u.likes {
		g.Relate(thing, other)
	}
	u.likes[thing] = struct{}{}
}

func (u *User) IsLiking(thing interface{}) bool {
	return u.likes[thing] != nil
}

// O(n*n)
func (u *User) Recommend() []ItemCount {
	scores := NewItemCounter()

	for like := range u.likes {
		related := g.FindRelations(like)
		for other, score := range related {
			if u.IsLiking(other) {
				continue
			}
			scores.IncBy(other, score)
		}
	}

	scores.Sort()

	res := make([]ItemCount, scores.Len())
	for i, v := range scores.Items() {
		res[i] = v
	}
	return res
}
