package main

import (
	"strconv"
	"strings"

	"assalielmehdi/adventofcode2023/util"
)

type part struct {
	x, m, a, s int
}

type rule struct {
	field string
	op    string
	val   int
	to    string
}

type workflow struct {
	name  string
	rules []rule
}

func toInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

func newRule(str string) rule {
	if !strings.Contains(str, ":") {
		return rule{to: str}
	}
	fields := strings.Split(str, ":")
	var r rule
	r.field = fields[0][0:1]
	r.op = fields[0][1:2]
	r.val = toInt(fields[0][2:])
	r.to = fields[1]
	return r
}

func newWorkflow(str string) workflow {
	fields := strings.Fields(str)
	var w workflow
	w.name = fields[0]
	w.rules = make([]rule, 0, len(fields)-1)
	for _, field := range fields[1:] {
		w.rules = append(w.rules, newRule(field))
	}
	return w
}

func newPart(str string) part {
	fields := strings.Fields(str)
	return part{
		x: toInt(strings.Split(fields[0], "=")[1]),
		m: toInt(strings.Split(fields[1], "=")[1]),
		a: toInt(strings.Split(fields[2], "=")[1]),
		s: toInt(strings.Split(fields[3], "=")[1]),
	}
}

func check(p part, r rule) bool {
	if len(r.op) == 0 {
		return true
	}
	var fieldVal int
	switch r.field {
	case "x":
		fieldVal = p.x
	case "m":
		fieldVal = p.m
	case "a":
		fieldVal = p.a
	case "s":
		fieldVal = p.s
	}
	if r.op == "<" {
		return fieldVal < r.val
	}
	return fieldVal > r.val
}

func partVal(p part) int {
	return p.x + p.m + p.a + p.s
}

func dfs1(p part, wn string, ws map[string]workflow) bool {
	if wn == "R" {
		return false
	}
	if wn == "A" {
		return true
	}
	w := ws[wn]
	for _, r := range w.rules {
		if check(p, r) {
			return dfs1(p, r.to, ws)
		}
	}
	return false
}

func solve1(sc *util.Scanner) any {
	ws := make(map[string]workflow, 0)
	for sc.HasNextLine() {
		line := sc.NextLine()
		if len(line) == 0 {
			break
		}
		w := newWorkflow(line)
		ws[w.name] = w
	}
	ps := make([]part, 0)
	for sc.HasNextLine() {
		ps = append(ps, newPart(sc.NextLine()))
	}
	answer := 0
	for _, p := range ps {
		if dfs1(p, "in", ws) {
			answer += partVal(p)
		}
	}
	return answer
}

type partInterval struct {
	x, m, a, s util.Pair[int, int]
}

func fix(p partInterval, r rule, inverse bool) partInterval {
	if len(r.op) == 0 {
		return p
	}
	var fieldInterval util.Pair[int, int]
	switch r.field {
	case "x":
		fieldInterval = p.x
	case "m":
		fieldInterval = p.m
	case "a":
		fieldInterval = p.a
	case "s":
		fieldInterval = p.s
	}
	if !inverse && r.op == "<" && fieldInterval.Second >= r.val {
		fieldInterval.Second = r.val - 1
	} else if !inverse && r.op == ">" && fieldInterval.First <= r.val {
		fieldInterval.First = r.val + 1
	} else if inverse && r.op == "<" && fieldInterval.First < r.val {
		fieldInterval.First = r.val
	} else if inverse && r.op == ">" && fieldInterval.Second > r.val {
		fieldInterval.Second = r.val
	}
	switch r.field {
	case "x":
		p.x = fieldInterval
	case "m":
		p.m = fieldInterval
	case "a":
		p.a = fieldInterval
	case "s":
		p.s = fieldInterval
	}
	return p
}

func isIntervalValid(i util.Pair[int, int]) bool {
	return i.First <= i.Second
}

func isPartIntervalValid(p partInterval) bool {
	return isIntervalValid(p.x) && isIntervalValid(p.m) && isIntervalValid(p.a) && isIntervalValid(p.x)
}

func dfs2(p partInterval, wn string, ws map[string]workflow, path []string) []partInterval {
	if !isPartIntervalValid(p) {
		return []partInterval{}
	}
	if wn == "R" {
		return []partInterval{}
	}
	if wn == "A" {
		return []partInterval{p}
	}
	w := ws[wn]
	intervals := make([]partInterval, 0)
	for _, r := range w.rules {
		intervals = append(intervals, dfs2(fix(p, r, false), r.to, ws, append(path, r.to))...)
		p = fix(p, r, true)
	}
	return intervals
}

func combin(p partInterval) int64 {
	return int64(p.x.Second-p.x.First+1) * int64(p.m.Second-p.m.First+1) * int64(p.a.Second-p.a.First+1) * int64(p.s.Second-p.s.First+1)
}

func solve2(sc *util.Scanner) any {
	ws := make(map[string]workflow, 0)
	for sc.HasNextLine() {
		line := sc.NextLine()
		if len(line) == 0 {
			break
		}
		w := newWorkflow(line)
		ws[w.name] = w
	}
	intervals := dfs2(partInterval{
		x: util.Pair[int, int]{First: 1, Second: 4000},
		m: util.Pair[int, int]{First: 1, Second: 4000},
		a: util.Pair[int, int]{First: 1, Second: 4000},
		s: util.Pair[int, int]{First: 1, Second: 4000},
	}, "in", ws, []string{"in"})
	answer := int64(0)
	for _, interval := range intervals {
		answer += combin(interval)
	}
	return answer
}

func main() {
	util.RunAll("Day 19 - Part 1", solve1)
	util.RunAll("Day 19 - Part 2", solve2)
}
