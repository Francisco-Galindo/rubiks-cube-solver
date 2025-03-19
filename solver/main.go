package main

import (
	"fmt"
	"slices"
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
	InvalidFullMove
)

// 5
// 2
// 3
// 11
// 5
// 2
// 3
// 11
// 14

type Move uint8

const (
	U Move = iota
	R
	F
	D
	L
	B
	InvalidMove
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
	InvalidCorn
)

type CornerCubie struct {
	c Corn
	o uint8
}

var cornerCubieMoves = [7][8]CornerCubie{
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
	{
		{c: URF, o: 0},
		{c: UFL, o: 0},
		{c: ULB, o: 0},
		{c: UBR, o: 0},
		{c: DFR, o: 0},
		{c: DLF, o: 0},
		{c: DBL, o: 0},
		{c: DRB, o: 0},
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
	InvalidEdge
)

type EdgeCubie struct {
	e Edge
	o uint8
}

var edgeCubieMoves = [7][12]EdgeCubie{
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
	{{e: UR, o: 0}, {e: UF, o: 0}, {e: UL, o: 0}, {e: UB, o: 0}, {e: DR, o: 0}, {e: DF, o: 0}, //TROLl
		{e: DL, o: 0}, {e: DB, o: 0}, {e: FR, o: 0}, {e: FL, o: 0}, {e: BL, o: 0}, {e: BR, o: 0}},
}

type Perm struct {
	corns [8]CornerCubie
	edges [12]EdgeCubie
}

type QueueElement struct {
	p          Perm
	parentMove FullMove
	level      uint8
}

var prune1 = make([]uint, 0)
var prune2 = make([]uint, 0)
var prune3 = make([]uint, 0)
var phase2ForbiddenMoves = []FullMove{Rx1, Rx3, Lx1, Lx3, Fx1, Fx3, Bx1, Bx3}

var maxLength = 40

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

func indexCornPerm(p Perm) int {
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

func indexEdgePerm(p Perm) int {
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

func genPrune1() []uint {
	var count uint16 = 0
	var root Perm = Perm{
		corns: cornerCubieMoves[InvalidMove],
		edges: edgeCubieMoves[InvalidMove],
	}
	var q = make([]QueueElement, 0)
	var pruneTable = make([]uint, 2187)

	q = append(q, QueueElement{p: root, parentMove: InvalidFullMove, level: 0})

	for count < 2186 && len(q) > 0 {
		curr := q[0]
		p := curr.p
		pHash := indexCornOri(p)
		if pruneTable[pHash] == 0 && pHash != 0 {
			count++
			pruneTable[pHash] = uint(curr.level)
			fmt.Printf("%v\t%v\n", pHash, curr.level)
		}
		q = q[1:]

		for newMove := range Bx3 + 1 {
			if curr.parentMove/3 == newMove/3 {
				continue
			}
			p = makeFullMove(p, newMove)
			q = append(q, QueueElement{p: p, parentMove: newMove, level: curr.level + 1})
		}
	}

	return pruneTable
}

func genPrune2() []uint {
	var count uint16 = 0
	var root Perm = Perm{
		corns: cornerCubieMoves[InvalidMove],
		edges: edgeCubieMoves[InvalidMove],
	}
	var q = make([]QueueElement, 0)
	var pruneTable = make([]uint, 2048)

	q = append(q, QueueElement{p: root, parentMove: InvalidFullMove, level: 0})

	for count < 1799 && len(q) > 0 {
		curr := q[0]
		p := curr.p
		pHash := indexEdgeOri(p)
		if pruneTable[pHash] == 0 && pHash != 0 {
			count++
			pruneTable[pHash] = uint(curr.level)
		}
		q = q[1:]

		for newMove := range Bx3 + 1 {
			if curr.parentMove/3 == newMove/3 {
				continue
			}
			p = makeFullMove(p, newMove)
			q = append(q, QueueElement{p: p, parentMove: newMove, level: curr.level + 1})
		}
	}

	return pruneTable
}

func genPrune3() []uint {
	var count uint16 = 0
	var root Perm = Perm{
		corns: cornerCubieMoves[InvalidMove],
		edges: edgeCubieMoves[InvalidMove],
	}
	var q = make([]QueueElement, 0)
	var pruneTable = make([]uint, 40320)

	q = append(q, QueueElement{p: root, parentMove: InvalidFullMove, level: 0})

	for count < 40319 && len(q) > 0 {
		curr := q[0]
		p := curr.p
		pHash := indexCornPerm(p)
		if pruneTable[pHash] == 0 && pHash != 0 {
			count++
			// fmt.Println(count)
			pruneTable[pHash] = uint(curr.level)
		}
		q = q[1:]

		// for _, newMove := range []FullMove{Ux1, Dx1, Rx2, Lx2, Fx2, Bx2} {
		for newMove := range Bx3 + 1 {
			if curr.parentMove/3 == newMove/3 {
				continue
			}
			p = makeFullMove(p, newMove)
			q = append(q, QueueElement{p: p, parentMove: newMove, level: curr.level + 1})
		}
	}

	return pruneTable
}

func genPrune4() []uint {
	var count uint = 0
	var root Perm = Perm{
		corns: cornerCubieMoves[InvalidMove],
		edges: edgeCubieMoves[InvalidMove],
	}
	var q = make([]QueueElement, 0)
	var pruneTable = make([]uint, 479_001_600)

	q = append(q, QueueElement{p: root, parentMove: InvalidFullMove, level: 0})

	for count < 479_001_599 && len(q) > 0 {
		curr := q[0]
		p := curr.p
		pHash := indexEdgePerm(p)
		if pruneTable[pHash] == 0 && pHash != 0 {
			count++
			pruneTable[pHash] = uint(curr.level)
			fmt.Printf("%v\t%v\n", pHash, curr.level)
		}
		q = q[1:]

		// for _, newMove := range []FullMove{Ux1, Dx1, Rx2, Lx2, Fx2, Bx2} {
		for newMove := range Bx3 + 1 {
			if curr.parentMove/3 == newMove/3 {
				continue
			}
			p = makeFullMove(p, newMove)
			q = append(q, QueueElement{p: p, parentMove: newMove, level: curr.level + 1})
		}
	}

	return pruneTable
}

func kociemba(p Perm) {
	for d := range maxLength + 1 {
		res := phase1search(p, uint(d+1), InvalidFullMove)
		// fmt.Println(res)
		if res {
			fmt.Println("FINAL DEPTH", d)
			return
		}
	}
}

func phase1search(p Perm, depth uint, lastMove FullMove) bool {
	if depth == 0 && phase1GoalReached(p) && slices.Index(phase2ForbiddenMoves, lastMove) != 0 {
		return phase2start(p, depth)
	} else if depth > 0 {
		for move := range Bx3 + 1 {
			pHash := indexCornOri(p)
			pHash2 := indexEdgeOri(p)
			if max(prune1[pHash], prune2[pHash2]) > depth || move/3 == lastMove/3 {
				// if  move/3 == lastMove/3 {
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

func phase2start(p Perm, currDepth uint) bool {
	for d := range min(uint(maxLength)-currDepth, 11) {
		res := phase2search(p, d, InvalidFullMove)
		// fmt.Println("phase 2", res, currDepth)
		maxLength = int(currDepth) - 1
		if res {
			return true
		}
	}
	return false
}

func phase2search(p Perm, depth uint, lastMove FullMove) bool {
	if isSolved(p) {
		fmt.Println("YEEEEEEEA")
		return true
	}
	if depth > 0 {
		for move := range Bx3 + 1 {
			pHash := indexCornPerm(p)
			if slices.Index(phase2ForbiddenMoves, move) != 0 || prune3[pHash] > depth || move/3 == lastMove/3 {
				// if move/3 == lastMove/3 {
				// if prune2[pHash] > depth {
				// 	fmt.Println("LOL")
				// }
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

func phase1GoalReached(p Perm) bool {
	return indexCornOri(p) == 0 &&
		indexEdgeOri(p) == 0
}

func isSolved(p Perm) bool {
	return indexCornOri(p) == 0 &&
		indexCornPerm(p) == 0 &&
		indexEdgePerm(p) == 0 &&
		indexEdgeOri(p) == 0
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

	p = makeFullMove(p, Rx3)

	// p = makeFullMove(p, Rx1)
	// p = makeFullMove(p, Ux1)
	// p = makeFullMove(p, Rx3)
	// p = makeFullMove(p, Ux3)
	//
	// p = makeFullMove(p, Rx1)
	// p = makeFullMove(p, Ux1)
	// p = makeFullMove(p, Rx3)
	// p = makeFullMove(p, Ux3)

	// R U R' D R U R' D L F R B R' D L F' R U R' D
	p = makeFullMove(p, Rx1)
	p = makeFullMove(p, Ux1)
	p = makeFullMove(p, Rx3)
	p = makeFullMove(p, Dx1)
	p = makeFullMove(p, Rx1)
	p = makeFullMove(p, Ux1)
	p = makeFullMove(p, Rx3)
	p = makeFullMove(p, Dx1)
	p = makeFullMove(p, Lx1)
	p = makeFullMove(p, Fx1)
	p = makeFullMove(p, Rx1)
	// p = makeFullMove(p, Bx1)
	// p = makeFullMove(p, Rx3)
	// p = makeFullMove(p, Dx1)
	// p = makeFullMove(p, Lx1)
	// p = makeFullMove(p, Fx3)
	// p = makeFullMove(p, Rx1)
	// p = makeFullMove(p, Ux1)
	// p = makeFullMove(p, Rx3)
	// p = makeFullMove(p, Dx1)

	// prune1 = genPrune1()
	// prune2 = genPrune2()
	// prune3 = genPrune3()
	// genPrune1()
	genPrune4()

	// for i, v := range genPrune1() {
	// 	fmt.Printf("%v\t%v\n", i, v)
	// }

	fmt.Println("HOLA")
	// kociemba(p)
}
