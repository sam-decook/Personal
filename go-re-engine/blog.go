package bre

import (
	"fmt"
	"strconv"
	"strings"
)

/* ------------ Parsing the regular expression ------------ */

type tokenType uint8

const (
	group tokenType = iota
	bracket
	or
	repeat
	literal
	groupUncaptured
)

const repeatInfinity = -1

type token struct {
	tokenType tokenType
	value     interface{}
}

type repeatPayload struct {
	min,
	max int
	value token
}

type parseContext struct {
	pos    int
	tokens []token
}

func parse(regex string) *parseContext {
	ctx := &parseContext{
		tokens: []token{},
	}

	for ; ctx.pos < len(regex); ctx.pos++ {
		process(regex, ctx)
	}

	return ctx
}

func process(regex string, ctx *parseContext) {
	ch := regex[ctx.pos]
	switch ch {
	case '(':
		groupCtx := &parseContext{
			pos:    ctx.pos,
			tokens: []token{},
		}

		parseGroup(regex, groupCtx)

		ctx.tokens = append(ctx.tokens, token{
			tokenType: group,
			value:     groupCtx.tokens,
		})

	case '[':
		parseBracket(regex, ctx)

	case '|':
		parseOr(regex, ctx)

	case '*', '?', '+':
		parseRepeat(regex, ctx)

	case '{':
		parseRepeatSpecified(regex, ctx)

	default:
		ctx.tokens = append(ctx.tokens, token{
			tokenType: literal,
			value:     ch,
		})
	}
}

func parseGroup(regex string, ctx *parseContext) {
	ctx.pos += 1 //advance past LPAREN
	for ; regex[ctx.pos] != ')'; ctx.pos++ {
		process(regex, ctx)
	}
}

func parseBracket(regex string, ctx *parseContext) {
	ctx.pos += 1 //advance past LBRACKET
	var literals []string
	for ; regex[ctx.pos] != ']'; ctx.pos++ {
		ch := regex[ctx.pos]

		if ch == '-' {
			next := regex[ctx.pos+1]
			prev := literals[len(literals)-1][0]
			literals[len(literals)-1] = fmt.Sprintf("%c%c", prev, next)
			ctx.pos += 1
		} else {
			literals = append(literals, fmt.Sprintf("%c", ch))
		}
	}

	literalsSet := map[uint8]struct{}{}

	for _, lit := range literals {
		for i := lit[0]; i <= lit[len(lit)-1]; i++ {
			literalsSet[i] = struct{}{}
		}
	}

	ctx.tokens = append(ctx.tokens, token{
		tokenType: bracket,
		value:     literalsSet,
	})
}

func parseOr(regex string, ctx *parseContext) {
	rhsCtx := &parseContext{
		pos:    ctx.pos,
		tokens: []token{},
	}

	rhsCtx.pos += 1 //advance past PIPE
	for ; rhsCtx.pos < len(regex) && regex[rhsCtx.pos] != ')'; rhsCtx.pos++ {
		process(regex, rhsCtx)
	}

	left := token{
		tokenType: groupUncaptured,
		value:     ctx.tokens,
	}
	right := token{
		tokenType: groupUncaptured,
		value:     rhsCtx.tokens,
	}

	ctx.tokens = []token{{
		tokenType: or,
		value:     []token{left, right},
	}}

	ctx.pos = rhsCtx.pos
}

func parseRepeat(regex string, ctx *parseContext) {
	var min, max int
	switch regex[ctx.pos] {
	case '*':
		min = 0
		max = repeatInfinity
	case '?':
		min = 0
		max = 1
	case '+':
		min = 1
		max = repeatInfinity
	}

	last := len(ctx.tokens) - 1
	lastToken := ctx.tokens[last]

	ctx.tokens[last] = token{
		tokenType: repeat,
		value: repeatPayload{
			min:   min,
			max:   max,
			value: lastToken,
		},
	}
}

func parseRepeatSpecified(regex string, ctx *parseContext) {
	start := ctx.pos + 1 //advance past LCURLY

	for regex[ctx.pos] != '}' {
		ctx.pos++
	}

	boundariesStr := regex[start:ctx.pos]
	pieces := strings.Split(boundariesStr, ",")

	var min, max int
	if len(pieces) == 1 {
		value, err := strconv.Atoi(pieces[0])
		if err != nil {
			panic(err.Error())
		}

		min = value
		max = value
	} else if len(pieces) == 2 {
		value, err := strconv.Atoi(pieces[0])
		if err != nil {
			panic(err.Error())
		}
		min = value

		if pieces[1] == "" {
			max = repeatInfinity
		} else if value, err := strconv.Atoi(pieces[1]); err != nil {
			panic(err.Error())
		} else {
			max = value
		}
	} else {
		panic(fmt.Sprintf("There must be either 1 or 2 values specified for the quantifier: provided '%s'", boundariesStr))
	}

	lastToken := ctx.tokens[len(ctx.tokens)-1]
	ctx.tokens[len(ctx.tokens)-1] = token{
		tokenType: repeat,
		value: repeatPayload{
			min:   min,
			max:   max,
			value: lastToken,
		},
	}
}

/* ------------ Building the state machine ------------ */
const epsilon uint8 = 0

type state struct {
	start,
	terminal bool
	transitions map[uint8][]*state //why not char instead of uint8?
}

func toNfa(ctx *parseContext) (start *state) {
	startState, endState := tokenToNfa(&ctx.tokens[0])

	for i := 1; i < len(ctx.tokens); i++ {
		startNext, endNext := tokenToNfa(&ctx.tokens[i])

		endState.transitions[epsilon] = append(
			endState.transitions[epsilon],
			startNext,
		)

		endState = endNext
	}

	start = &state{
		transitions: map[uint8][]*state{
			epsilon: {startState},
		},
		start: true,
	}

	end := &state{
		transitions: map[uint8][]*state{},
		terminal:    true,
	}

	endState.transitions[epsilon] = append(
		endState.transitions[epsilon],
		end,
	)

	return start
}

func tokenToNfa(t *token) (start, end *state) {
	start = &state{
		transitions: map[uint8][]*state{},
	}
	end = &state{
		transitions: map[uint8][]*state{},
	}

	switch t.tokenType {
	case literal:
		ch := t.value.(uint8) //type assertion
		start.transitions[ch] = []*state{end}

	case or:
		values := t.value.([]token)
		left := values[0]
		right := values[1]

		s1, e1 := tokenToNfa(&left)
		s2, e2 := tokenToNfa(&right)

		start.transitions[epsilon] = []*state{s1, s2}
		e1.transitions[epsilon] = []*state{end}
		e2.transitions[epsilon] = []*state{end}

	case bracket:
		literals := t.value.(map[uint8]struct{})

		for lit := range literals {
			start.transitions[lit] = []*state{end}
		}

	case group, groupUncaptured:
		tokens := t.value.([]token)
		s, end := tokenToNfa(&tokens[0]) //repeated code starting here
		start = s

		for i := 1; i < len(tokens); i++ {
			startNext, endNext := tokenToNfa(&tokens[i])

			end.transitions[epsilon] = append(
				end.transitions[epsilon],
				startNext,
			)

			end = endNext
		}

	case repeat:
		p := t.value.(repeatPayload)

		if p.min == 0 { //payload is optional, can skip to end
			start.transitions[epsilon] = []*state{end}
		}

		var copyCount int

		if p.max == repeatInfinity {
			if p.min == 0 {
				copyCount = 1
			} else {
				copyCount = p.min
			}
		} else {
			copyCount = p.max
		}

		from, to := tokenToNfa(&p.value)
		start.transitions[epsilon] = append(start.transitions[epsilon], from)

		for i := 2; i <= copyCount; i++ {
			s, e := tokenToNfa(&p.value)
			to.transitions[epsilon] = append(to.transitions[epsilon], s)

			from = s
			to = e

			if i > p.min {
				s.transitions[epsilon] = append(s.transitions[epsilon], end)
			}
		}

		to.transitions[epsilon] = append(to.transitions[epsilon], end)

		if p.max == repeatInfinity {
			end.transitions[epsilon] = append(end.transitions[epsilon], from)
		}
	}

	return start, end
}

/* ------------ Running the state machine ------------ */
const (
	startOfText uint8 = 1
	endofText   uint8 = 2
)

func getChar(input string, pos int) uint8 {
	switch {
	case pos >= len(input):
		return endofText
	case pos < 0:
		return startOfText
	default:
		return input[pos]
	}
}

func (s *state) check(input string, pos int) bool {
	ch := getChar(input, pos)

	if ch == endofText && s.terminal {
		return true
	}

	if states := s.transitions[ch]; len(states) > 0 {
		nextState := states[0]
		if nextState.check(input, pos+1) {
			return true
		}
	}

	for _, state := range s.transitions[epsilon] {
		if state.check(input, pos) {
			return true
		}

		if ch == startOfText && state.check(input, pos+1) {
			return true
		}
	}

	return false
}
