package aoc

import (
	"log"
	"strconv"
)

type operation string

const (
	jmp operation = "jmp"
	acc operation = "acc"
	nop operation = "nop"
)

type instruction struct {
	op  operation
	arg int
}

func accumulatorValueAtInifiniteLoop(reader lineReader) int {
	instructions := parseInstructions(reader)
	acc, _, _ := runProgram(instructions)
	return acc
}

func runProgram(instructions []instruction) (int, bool, []int) {
	vi := make(map[int]struct{}, 0)
	il := len(instructions)
	pc := 0
	sum := 0
	finished := false
	pcs := []int{}
	for {
		if pc >= il {
			finished = true
			break
		}
		ci := instructions[pc]
		if _, ok := vi[pc]; ok {
			break
		}
		vi[pc] = struct{}{}
		switch ci.op {
		case nop:
			pc++
		case acc:
			sum += ci.arg
			pc++
		case jmp:
			pc += ci.arg
		}
		pcs = append(pcs, pc)
	}
	return sum, finished, pcs
}

func adjustProgramAndRunToEnd(reader lineReader) int {
	instructions := parseInstructions(reader)
	_, _, reccurence := runProgram(instructions)
	last := reccurence[len(reccurence)-1]
	for i, v := range reccurence {
		if v == last {
			reccurence = reccurence[i:]
			break
		}
	}

	for _, mpc := range reccurence {
		i := instructions[mpc]
		if i.op == nop {
			i.op = jmp
		} else if i.op == jmp {
			i.op = nop
		} else {
			continue
		}
		modified := make([]instruction, len(instructions))
		copy(modified, instructions)
		modified[mpc] = i
		acc, finished, _ := runProgram(modified)
		if finished {
			return acc
		}
	}
	panic("no modification works")
}

func parseInstructions(reader lineReader) []instruction {
	instructions := make([]instruction, 0)
	for {
		line, eof := reader.readLine()
		op := line[0:3]
		arg, err := strconv.ParseInt(line[4:], 10, 32)
		if err != nil {
			log.Fatalf("unable to parse int %v", err)
		}
		instr := instruction{
			op:  operation(op),
			arg: int(arg),
		}
		instructions = append(instructions, instr)
		if eof {
			break
		}
	}
	return instructions
}
