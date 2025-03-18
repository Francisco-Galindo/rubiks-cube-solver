package main

import (
	"fmt"
)

type CubePos [6]uint64

type FullMove uint8

const (
	Ux1 FullMove = iota
	Ux2
	Ux3
	Rx1
	Rx2
	Rx3
	Fx1
	Fx2
	Fx3
	Dx1
	Dx2
	Dx3
	Lx1
	Lx2
	Lx3
	Bx1
	Bx2
	Bx3
	Invalid
)

type Move uint8

const (
	U Move = iota
	R
	F
	D
	L
	B
)

type Corn uint8

const (
	URF Corn = iota
	UFL
	ULB
	UBR
	DFR
	DLF
	DBL
	DRB
)

type CornerCubie struct {
	c Corn
	o uint8
}

var cornerCubieMoves = [6][8]CornerCubie{
	{
		{c: UBR, o: 0},
		{c: URF, o: 0},
		{c: UFL, o: 0},
		{c: ULB, o: 0},
		{c: DFR, o: 0},
		{c: DLF, o: 0},
		{c: DBL, o: 0},
		{c: DRB, o: 0},
	},
	{
		{c: DFR, o: 2},
		{c: UFL, o: 0},
		{c: ULB, o: 0},
		{c: URF, o: 1},
		{c: DRB, o: 1},
		{c: DLF, o: 0},
		{c: DBL, o: 0},
		{c: UBR, o: 2},
	},
	{
		{c: UFL, o: 1},
		{c: DLF, o: 2},
		{c: ULB, o: 0},
		{c: UBR, o: 0},
		{c: URF, o: 2},
		{c: DFR, o: 1},
		{c: DBL, o: 0},
		{c: DRB, o: 0},
	},
	{
		{c: URF, o: 0},
		{c: UFL, o: 0},
		{c: ULB, o: 0},
		{c: UBR, o: 0},
		{c: DLF, o: 0},
		{c: DBL, o: 0},
		{c: DRB, o: 0},
		{c: DFR, o: 0},
	},
	{
		{c: URF, o: 0},
		{c: ULB, o: 1},
		{c: DBL, o: 2},
		{c: UBR, o: 0},
		{c: DFR, o: 0},
		{c: UFL, o: 2},
		{c: DLF, o: 1},
		{c: DRB, o: 0},
	},
	{
		{c: URF, o: 0},
		{c: UFL, o: 0},
		{c: UBR, o: 1},
		{c: DRB, o: 2},
		{c: DFR, o: 0},
		{c: DLF, o: 0},
		{c: ULB, o: 2},
		{c: DBL, o: 1},
	},
}

type Edge uint8

const (
	UR Edge = iota
	UF
	UL
	UB
	DR
	DF
	DL
	DB
	FR
	FL
	BL
	BR
)

type EdgeCubie struct {
	e Edge
	o uint8
}

var edgeCubieMoves = [6][12]EdgeCubie{
	{{e: UB, o: 0}, {e: UR, o: 0}, {e: UF, o: 0}, {e: UL, o: 0}, {e: DR, o: 0}, {e: DF, o: 0}, //U
		{e: DL, o: 0}, {e: DB, o: 0}, {e: FR, o: 0}, {e: FL, o: 0}, {e: BL, o: 0}, {e: BR, o: 0}},
	{{e: FR, o: 0}, {e: UF, o: 0}, {e: UL, o: 0}, {e: UB, o: 0}, {e: BR, o: 0}, {e: DF, o: 0}, //R
		{e: DL, o: 0}, {e: DB, o: 0}, {e: DR, o: 0}, {e: FL, o: 0}, {e: BL, o: 0}, {e: UR, o: 0}},
	{{e: UR, o: 0}, {e: FL, o: 1}, {e: UL, o: 0}, {e: UB, o: 0}, {e: DR, o: 0}, {e: FR, o: 1}, //F
		{e: DL, o: 0}, {e: DB, o: 0}, {e: UF, o: 1}, {e: DF, o: 1}, {e: BL, o: 0}, {e: BR, o: 0}},
	{{e: UR, o: 0}, {e: UF, o: 0}, {e: UL, o: 0}, {e: UB, o: 0}, {e: DF, o: 0}, {e: DL, o: 0}, //D
		{e: DB, o: 0}, {e: DR, o: 0}, {e: FR, o: 0}, {e: FL, o: 0}, {e: BL, o: 0}, {e: BR, o: 0}},
	{{e: UR, o: 0}, {e: UF, o: 0}, {e: BL, o: 0}, {e: UB, o: 0}, {e: DR, o: 0}, {e: DF, o: 0}, //L
		{e: FL, o: 0}, {e: DB, o: 0}, {e: FR, o: 0}, {e: UL, o: 0}, {e: DL, o: 0}, {e: BR, o: 0}},
	{{e: UR, o: 0}, {e: UF, o: 0}, {e: UL, o: 0}, {e: BR, o: 1}, {e: DR, o: 0}, {e: DF, o: 0}, //B
		{e: DL, o: 0}, {e: BL, o: 1}, {e: FR, o: 0}, {e: FL, o: 0}, {e: UB, o: 1}, {e: DB, o: 1}},
}

type Perm struct {
	corns [8]CornerCubie
	edges [12]EdgeCubie
}

func makeMove(p Perm, m Move) Perm {
	var prod Perm

	for co := range DRB + 1 {
		prod.corns[co].c = p.corns[cornerCubieMoves[m][co].c].c
		oriA := p.corns[cornerCubieMoves[m][co].c].o
		oriB := cornerCubieMoves[m][co].o
		prod.corns[co].o = (oriA + oriB) % 3
	}

	for ed := range BR + 1 {
		prod.edges[ed].e = p.edges[edgeCubieMoves[m][ed].e].e
		oriA := p.edges[edgeCubieMoves[m][ed].e].o
		oriB := edgeCubieMoves[m][ed].o
		prod.edges[ed].o = (oriA + oriB) % 2
	}

	return prod
}

func makeFullMove(p Perm, m FullMove) Perm {
	switch m {
	case Ux1:
		p = makeMove(p, U)
	case Ux2:
		p = makeMove(p, U)
		p = makeMove(p, U)
	case Ux3:
		p = makeMove(p, U)
		p = makeMove(p, U)
		p = makeMove(p, U)
	case Rx1:
		p = makeMove(p, R)
	case Rx2:
		p = makeMove(p, R)
		p = makeMove(p, R)
	case Rx3:
		p = makeMove(p, R)
		p = makeMove(p, R)
		p = makeMove(p, R)
	case Fx1:
		p = makeMove(p, F)
	case Fx2:
		p = makeMove(p, F)
		p = makeMove(p, F)
	case Fx3:
		p = makeMove(p, F)
		p = makeMove(p, F)
		p = makeMove(p, F)
	case Dx1:
		p = makeMove(p, D)
	case Dx2:
		p = makeMove(p, D)
		p = makeMove(p, D)
	case Dx3:
		p = makeMove(p, D)
		p = makeMove(p, D)
		p = makeMove(p, D)
	case Lx1:
		p = makeMove(p, L)
	case Lx2:
		p = makeMove(p, L)
		p = makeMove(p, L)
	case Lx3:
		p = makeMove(p, L)
		p = makeMove(p, L)
		p = makeMove(p, L)
	case Bx1:
		p = makeMove(p, B)
	case Bx2:
		p = makeMove(p, B)
		p = makeMove(p, B)
	case Bx3:
		p = makeMove(p, B)
		p = makeMove(p, B)
		p = makeMove(p, B)
	}

	return p
}

func indexCornOri(p Perm) uint16 {
	var s uint16 = 0
	for co := range DRB {
		s = 3*s + uint16(p.corns[co].o)
	}

	return s
}

func indexEdgeOri(p Perm) uint16 {
	var s uint16 = 0
	for ed := range BR {
		s = 2*s + uint16(p.edges[ed].o)
	}

	return s
}

func indexPermCord(p Perm) int {
	t := 0
	for i := DRB; i > URF; i-- {
		s := 0
		for j := i - 1; j >= URF; j-- {
			if p.corns[j].c > p.corns[i].c {
				s++
			}
			if j == 0 {
				break
			}
		}
		t = (t + s) * int(i)
	}
	return t
}

func indexEdgeCord(p Perm) int {
	t := 0
	for i := BR; i > UR; i-- {
		s := 0
		for j := i - 1; j >= UR; j-- {
			if p.edges[j].e > p.edges[i].e {
				s++
			}
			if j == 0 {
				break
			}
		}
		t = (t + s) * int(i)
	}
	return t
}

func kociemba(p Perm) {
	for d := range 128 {
		res := phase1search(p, uint64(d+1), Invalid)
		fmt.Println(res)
		if res {
			fmt.Println("FINAL DEPTH", d)
			return
		}
	}
}

func phase1search(p Perm, depth uint64, lastMove FullMove) bool {
	if solved(p) {
		fmt.Println("YEEEEEEEA")
		return true
	}
	if depth > 0 {
		for move := range Bx3 + 1 {
			if move == lastMove {
				continue
			}
			newP := p
			newP = makeFullMove(newP, move)
			res := phase1search(newP, depth-1, move)
			if res {
				fmt.Println(move)
				return true
			}
		}
	}
	return false
}

func phase2start(p Perm, currDepth int) {
	for d := range 128 - currDepth {
		phase2search(p, uint64(d), Invalid)
	}
}

func phase2search(p Perm, depth uint64, lastMove FullMove) bool {
	if solved(p) {
		fmt.Println("YEEEEEEEA")
		return true
	}
	if depth > 0 {
		for move := range Bx3 + 1 {
			if move == lastMove {
				continue
			}
			newP := p
			newP = makeFullMove(newP, move)
			res := phase2search(newP, depth-1, move)
			if res {
				fmt.Println(move)
				return true
			}
		}
	}
	return false
}

func solved(p Perm) bool {
	return indexCornOri(p) == 0 && indexPermCord(p) == 0
}

func main() {
	var p Perm = Perm{corns: cornerCubieMoves[R], edges: edgeCubieMoves[R]}
	// fmt.Println(indexCornOri(p), indexPermCord(p))
	// p = makeMove(p, R)
	// fmt.Println(indexCornOri(p), indexPermCord(p))
	// p = makeMove(p, R)
	// fmt.Println(indexCornOri(p), indexPermCord(p))
	// p = makeMove(p, R)
	// fmt.Println(indexCornOri(p), indexPermCord(p))
	// fmt.Println(p.corns, p.edges)
	// // R U R' U'
	//
	p = makeFullMove(p, Rx1)
	p = makeFullMove(p, Ux1)
	p = makeFullMove(p, Rx3)
	p = makeFullMove(p, Ux3)

	p = makeFullMove(p, Rx1)
	p = makeFullMove(p, Ux1)
	// p = makeFullMove(p, Rx3)
	// p = makeFullMove(p, Ux3)

	kociemba(p)
}
