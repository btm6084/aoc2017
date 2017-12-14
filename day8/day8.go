package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// Instruction represents a single instruction
type Instruction struct {
	Reg   string
	Op    string
	OpVal int
	IfReg string
	IfOp  string
	IfVal int
}

// Register is a representation of a single register
type Register struct {
	Value int
}

// Registers is a bank of individual registers
type Registers struct {
	registers    map[string]Register
	initialized  bool
	historicHigh int
}

// NewRegisterBank creates and initializes a new register bank.
func NewRegisterBank() *Registers {
	var r Registers
	r.registers = make(map[string]Register)
	r.historicHigh = 0

	return &r
}

// Get retrieves the value of a given register
func (r *Registers) Get(reg string) int {
	return r.registers[reg].Value
}

// Set sets the value of a given register
func (r *Registers) Set(reg string, value int) {
	r.registers[reg] = Register{Value: value}
}

// SetHistoricHigh sets a new high value if it's the highest seen at any time.
func (r *Registers) SetHistoricHigh(value int) {
	if r.historicHigh < value {
		r.historicHigh = value
	}
}

// GetHistoricHigh returns the highest value seen in any register.
func (r *Registers) GetHistoricHigh() int {
	return r.historicHigh
}

// RunInstructionSet runs a set of instructions.
func (r *Registers) RunInstructionSet(instructions []Instruction) {
	for _, i := range instructions {
		r.RunInstruction(i)
		r.SetHistoricHigh(r.FindLargestValue())
	}
}

// RunInstruction runs a single instruction.
func (r *Registers) RunInstruction(i Instruction) {
	doOp := false

	switch i.IfOp {
	case ">":
		doOp = r.Get(i.IfReg) > i.IfVal
	case ">=":
		doOp = r.Get(i.IfReg) >= i.IfVal
	case "<":
		doOp = r.Get(i.IfReg) < i.IfVal
	case "<=":
		doOp = r.Get(i.IfReg) <= i.IfVal
	case "==":
		doOp = r.Get(i.IfReg) == i.IfVal
	case "!=":
		doOp = r.Get(i.IfReg) != i.IfVal
	default:
		fmt.Println(spew.Sdump(i))
		fmt.Printf("Unset Op: %s\n", i.IfOp)
	}

	if doOp {
		switch i.Op {
		case "inc":
			r.Set(i.Reg, r.Get(i.Reg)+i.OpVal)
		case "dec":
			r.Set(i.Reg, r.Get(i.Reg)-i.OpVal)
		}
	}
}

// FindLargestValue finds the largest value in any register.
func (r *Registers) FindLargestValue() int {
	max := 0

	for _, v := range r.registers {
		if v.Value > max {
			max = v.Value
		}
	}

	return max
}

func main() {
	instructions := loadFromFile("./input.txt")

	r := NewRegisterBank()
	r.RunInstructionSet(instructions)

	fmt.Printf("Part 1: Largest Value: %d\n", r.FindLargestValue())
	fmt.Printf("Part 2: Largest Historic Value: %d\n", r.GetHistoricHigh())

}

func loadFromFile(filename string) []Instruction {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(bytes) <= 0 {
		fmt.Println("No Input")
		os.Exit(1)
	}

	lines := strings.Split(string(bytes), "\n")
	return buildInstructions(lines)
}

func buildInstructions(lines []string) []Instruction {
	var instructions []Instruction

	for _, l := range lines {
		var i Instruction
		n, err := fmt.Sscanf(l, "%s %s %d if %s %s %d", &i.Reg, &i.Op, &i.OpVal, &i.IfReg, &i.IfOp, &i.IfVal)
		if err != nil || n < 6 {
			continue
		}

		instructions = append(instructions, i)
	}

	return instructions
}
