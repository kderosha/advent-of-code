package updates

import (
	"log/slog"
    "fmt"

	"github.com/dominikbraun/graph"
	"github.com/kderosha/advent-of-code/input"
)

type Puzzle struct {
    rules []*Rule
    updates []*update
}

type Rule struct {
    x,y int
}

func New(pi *input.PuzzleInput) *Puzzle {
    p := &Puzzle{
        rules: make([]*Rule, 0),
        updates: make([]*update, 0),
    }

    lineItems := pi.LineItems()
    counter := 0
    for _, li := range lineItems {
        if string(li) == "" {
            counter++
            break
        }
        matches := li.ConvertToNumbersArray()
        p.rules = append(p.rules, &Rule{int(matches[0]), int(matches[1])})
        counter++
    }
    slog.Info("Done processing rules", "counter", counter)
    
    for i := counter; i < len(lineItems); i++ {
        p.AddUpdate(lineItems[i])
    }

    return p
}

func (p *Puzzle) AddUpdate(li input.LineItem) {
    matches := li.ConvertToNumbersArray()
    u := &update{u:matches, g:graph.New(graph.IntHash, graph.Directed(), graph.PreventCycles()), rules:make([]*Rule,0), correct:false}
    u.AddRulesPertainingToUpdate(p.rules)
    u.FormGraph()
    p.updates = append(p.updates, u)
}

type update struct {
    u []int64
    rules []*Rule
    correct bool
    g graph.Graph[int, int]
}

func (u *update) AddRulesPertainingToUpdate(rules []*Rule) {
    for _, n := range u.u {
        for _, rule := range rules{
            if rule.x == int(n) {
                u.rules = append(u.rules, rule)
            }
        }
    }
}

func (u *update) FormGraph() {
    for _, rule := range u.rules {
        err := u.g.AddVertex(int(rule.x))
        if err != nil && err != graph.ErrVertexAlreadyExists {
            panic(err)
        }
        err = u.g.AddVertex(rule.y)
        if err != nil && err != graph.ErrVertexAlreadyExists {
            panic(err)
        }
        err = u.g.AddEdge(rule.x, rule.y)
        if err != nil && err == graph.ErrEdgeCreatesCycle {
            panic(err)
        }
    }

}

func (u *update) Evaluate() {
    order, err := graph.TopologicalSort(u.g)
    if err != nil {
        panic(err)
    }
    slog.Info("Evaluating update", "update", u.u, "order", order)
    // for each number before current index
    for idx, n := range u.u {
        currentTopoIndex := findIndex(int(n), order)
        // make sure the any numbers don't come after this number in the topo sort order
        for i := idx - 1; i >= 0; i-- {
            // Check to make sure the number in the update at index i doesn't come after current n
            prevTopoIndex := findIndex(int(u.u[i]), order)
            if prevTopoIndex > currentTopoIndex {
                u.correct = false
                return
            }
        }
    }
    u.correct = true
    return
}

func findIndex(n int, order []int) int {
    for idx, compare := range order {
        if n == compare {
            return idx
        }
    }
    return -1
}

func findIndexInt64(n int, order []int64) int {
    for idx, compare := range order {
        if int64(n) == compare {
            return idx
        }
    }
    return -1
}

func (u *update) Correct() bool {
    slog.Info("Is update correct", "update", u)
    return u.correct
}

func (u *update) CorrectTheOrdering() {
    updatedArray := make([]int64, 0)
    topoSort, err := graph.TopologicalSort(u.g)
    if err != nil {
        panic(err)
    }
    for _, n := range topoSort {
        if findIndexInt64(n, u.u) != -1 {
            updatedArray = append(updatedArray, int64(n))
        }
    }
    u.u = updatedArray
}

func (u *update) String() string {
    return fmt.Sprintf("%+v, %v", u.u, u.correct)
}

// []arr{1,2,3,4} length = 4, length / 2 == 2
// []arr{1,2,3,4,5} length = 5, length / 2 == 2,
func (u *update) FindMiddle() int {
    length := len(u.u)
    return int(u.u[length / 2])
}

func (p *Puzzle) SolutionOne() int {
    answer := 0
    for _, update := range p.updates {
        update.Evaluate()
        if update.Correct() {
            n := update.FindMiddle()
            answer += n
        }  
    }
    return answer
}

func (p *Puzzle) SolutionTwo() int {
    answer := 0
    for _, update := range p.updates {
        update.Evaluate()
        if !update.Correct() {
            slog.Info("Correcting the ordering", "before", update.u)
            update.CorrectTheOrdering()
            update.Evaluate()
            slog.Info("Is new order correct", "correct", update.Correct())
            if update.Correct() {
                answer += update.FindMiddle()
            }
        }
    }
    return answer
}
