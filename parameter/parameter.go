package parameter

import (
    "image"
    "strings"
    "fmt"
)

type Rule struct {
    operation string
    params    [] string
}

type Pipe struct {
    rules []*Rule
}

func NewRule(params string) *Rule {
    return nil
}

func NewPipe(params string) *Pipe {
    p := &Pipe{
        rules: []*Rule{},
    }
    for i, rule := range strings.Split(params, ",") {
        fmt.Printf("Processing %d\n", i)
        newRule := NewRule(rule)
        p.rules = append(p.rules, newRule)
    }
    return p
}

func (p *Pipe) Transform(input image.Image) (image.Image, error) {
    return input, nil
}