---
titlepage: true
title: "Report 02"
subtitle: "Rubik's Cube"
author: [Francisco Galindo Mena, Rodrigo García Peñafort, Andrea Saldaña Navarrete, Gustavo Santana Sánchez]
date: "March, 2025"
toc: true
lang: "en"
listings-disable-line-numbers: true
bibliography: "fuentes.bib"
nocite: |
  @*
...

# Abstract
This paper explores the implementation of Kociemba's Algorithm for solving the Rubik's Cube efficiently,
leveraging a cubie-level representation. The two-phase approach of Kociemba's method significantly reduces
the search space, allowing for near-optimal solutions in under 20 moves. We discuss the advantages
of the cubie-level representation in terms of computational efficiency, memory usage, and move application.

# Goals
* Implement an optimized Rubik's Cube solver using Kociemba’s Algorithm.
* Utilize a cubie-level representation to optimize state transitions.
* Improve computational efficiency compared to standard facelet-based representations.
* Analyze the effectiveness of heuristic-based search methods, such as IDA*.

# Introduction
The Rubik's Cube is a well-known combinatorial puzzle with over 43 quintillion possible states. Efficiently
solving the cube has been a long-standing challenge in algorithmic problem-solving. While brute-force approaches are
infeasible, heuristic-based methods such as Kociemba's Algorithm provide an effective way to find near-optimal
solutions. The core idea of Kociemba's method is to break the problem into two phases: first reducing the cube
to a predefined subgroup (G1), and then solving it completely in an optimized manner. This method greatly reduces
computational complexity and allows solutions close to the theoretical lower bound of 20 moves (God's Number).

A key aspect of our implementation is the cubie-level representation, which encodes the cube's state in terms
of piece positions and orientations rather than individual face colors. This representation enables efficient
move applications and state transitions, making it ideal for fast search algorithms like IDA*.

# Development
## Rubik's Cube Representation
The first challenge was representing the Rubik's Cube. Traditional Rubik's Cube solvers often rely on a facelet
representation, which stores colors at each position on the cube. This approach, while intuitive, is
inefficient for advanced solving algorithms. Instead, a cubie-level representation is used, where the cube's
state is described in terms of the position and orientation of its 8 corner pieces and 12 edge pieces. This
method allows for a more efficient representation of moves and enables quicker lookup operations. The state of
a cubie-level representation can be stored as two separate permutations: one for edges and another
for corners. Additionally, each piece's orientation is stored separately. This reduces memory usage
and speeds up move calculations.

## Kociemba's Algorithm Implementation
### Phase 1: Reducing the Cube to Subgroup G1
The first phase aims to transform the cube into a specific subset of states known as the G1 group.
This means ensuring that:
- All edge pieces are oriented correctly.
- The cube reaches a reduced subset where only a subset of allowed moves is required for solving.
- This phase utilizes precomputed lookup tables to guide moves efficiently, reducing the number of steps needed to reach G1.


### Phase 2: Solving from G1 to the Solved State
Once the cube is in the G1 group, a second search is performed to find the shortest solution to reach the solved state.
This step benefits from additional heuristics and pruning techniques, ensuring that an optimal or near-optimal
solution is found quickly.

## Search Strategy and Optimization
- IDA Algorithm:* Since an exhaustive search is computationally expensive, Kociemba's Algorithm relies on
Iterative Deepening A* (IDA*), an improved variant of A* that uses iterative deepening to limit memory
consumption while maintaining optimality.

- Heuristic Function: A key part of IDA* is the heuristic function, which estimates the minimum number
of moves required to solve the cube from a given state. This is based on pruning tables.

- Pruning Tables: To accelerate the search, pruning tables are used to eliminate redundant paths and
avoid unnecessary computations. These tables store distances to specific states and guide the search
more effectively.

# Results

# Conclusions

# Bibliography
