package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2(inputs []string) (int, int, error) {
	var part1, part2 int
	var instances []inst
	for _, input := range inputs {
		i, err := newInst(input)
		if err != nil {
			return part1, part2, err
		}
		instances = append(instances, i)
	}

	for _, inst := range instances {
		if validateCount(inst) {
			part1 = part1 + 1
		}
		if validatePosition(inst) {
			part2 = part2 + 1
		}
	}

	return part1, part2, nil
}

type policy struct {
	lower  int
	upper  int
	letter string
}

func newPolicy(p string) (policy, error) {
	s := strings.Split(p, " ")     // range is at 0, letter is at 1
	ss := strings.Split(s[0], "-") // lower is at 0, upper is at 1
	lower, err := strconv.Atoi(ss[0])
	if err != nil {
		return policy{}, fmt.Errorf("error parsing lower bound in policy: %s", err.Error())
	}
	upper, err := strconv.Atoi(ss[1])
	if err != nil {
		return policy{}, fmt.Errorf("error parsing upper bound in policy: %s", err.Error())
	}
	return policy{
		lower:  lower,
		upper:  upper,
		letter: s[1],
	}, nil
}

type inst struct {
	password string
	policy   policy
}

func newInst(line string) (inst, error) {
	s := strings.Split(line, ":")
	p, err := newPolicy(s[0])
	if err != nil {
		return inst{}, fmt.Errorf("failed to create instance: %s", err.Error())
	}
	return inst{
		password: s[1],
		policy:   p,
	}, nil
}

// validates the part 1 policy style
func validateCount(inst inst) bool {
	var count int
	var c string
	for _, r := range inst.password {
		c = string(r)
		if c == inst.policy.letter {
			count = count + 1
		}
	}
	return count >= inst.policy.lower && count <= inst.policy.upper
}

// validates the part 2 policy style
func validatePosition(inst inst) bool {
	return (string(inst.password[inst.policy.lower]) == inst.policy.letter && string(inst.password[inst.policy.upper]) != inst.policy.letter) || (string(inst.password[inst.policy.lower]) != inst.policy.letter && string(inst.password[inst.policy.upper]) == inst.policy.letter)
}
