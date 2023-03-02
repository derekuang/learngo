package main

import (
	"fmt"
	"strconv"
	"strings"
)

// result stores the parsed result for a domain
type result struct {
	domain string
	visits int
	// add more metrics if needed
}

// parser keep tracks of the parsing
type parser struct {
	sum     map[string]result // metrics per domain
	domains []string          // unique domain names
	total   int               // total visits for all domains
	lines   int               // number of parsed lines (for the error messages)
}

func newParser() parser {
	return parser{sum: make(map[string]result)}
}

func parse(p parser, input string) (parsed result, err error) {
	fields := strings.Fields(input)
	if len(fields) != 2 {
		err = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return
	}

	domain := fields[0]

	visits, err := strconv.Atoi(fields[1])
	if visits < 0 || err != nil {
		err = fmt.Errorf("wrong input: %q (line #%d)", fields[1], p.lines)
		return
	}

	parsed.domain = domain
	parsed.visits = visits

	return
}

func update(p parser, parsed result) parser {
	domain := parsed.domain
	visits := parsed.visits

	// Collect the unique domains
	if _, ok := p.sum[domain]; !ok {
		p.domains = append(p.domains, domain)
	}

	// Keep track of total and per domain visits
	p.total += visits

	// You cannot assign to composite values
	// p.sum[domain].visits += visits

	// create and assign a new copy of `visit`
	p.sum[domain] = result{
		domain: domain,
		visits: visits + p.sum[domain].visits,
	}

	return p
}
