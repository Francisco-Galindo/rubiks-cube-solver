# rubiks-cube-solver
A solver for Rubik's Cube, using A*, written in Go

# Abstract
This paper explores the implementation of Kociemba's Algorithm for solving the Rubik's Cube efficiently, 
leveraging a *is replaced by* representation. The two-phase approach of Kociemba's method significantly reduces 
the search space, allowing for near-optimal solutions in under 20 moves.

# Goals
* Implement an optimized Rubik's Cube solver using Kociemba’s Algorithm.
* Improve computational efficiency compared to standard facelet-based representations.
* Analyze the effectiveness of heuristic-based search methods, such as IDA*.

# Introduction
The Rubik's Cube is a well-known combinatorial puzzle with over 43 quintillion possible states. Efficiently 
solving the cube has been a long-standing challenge in algorithmic problem-solving. While brute-force approaches are 
infeasible, heuristic-based methods such as Kociemba's Algorithm provide an effective way to find near-optimal 
solutions. The core idea of Kociemba's method is to break the problem into two phases: first reducing the cube 
to a predefined subgroup (G1), and then solving it completely in an optimized manner. This method greatly reduces 
computational complexity and allows solutions close to the theoretical lower bound of 20 moves (God's Number).

# Development
## Rubik's Cube Representation
The first challenge was representing the Rubik's Cube. Traditional Rubik's Cube solvers often rely on a facelet 
representation, which stores colors at each position on the cube. This approach, while intuitive, is 
inefficient for advanced solving algorithms. Instead, a *is replaced by* representation is used, which encodes 
the cube's state in two lists. The first list stores the corners along with their orientation. Likewise, 
the second list stores the edges and their orientation. Each move is a permutation of these lists. This 
representation enables efficient  move application and state transitions, making it ideal for fast 
search algorithms like IDA*.

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
- IDA* Algorithm: Since an exhaustive search is computationally expensive, Kociemba's Algorithm relies on
Iterative Deepening A* (IDA*), an improved variant of A* that uses iterative deepening to limit memory
consumption while maintaining optimality.

- Heuristic Function: A key part of IDA* is the heuristic function, which estimates the minimum number
of moves required to solve the cube from a given state. This is based on pruning tables.

- Pruning Tables: To accelerate the search, pruning tables are used to eliminate redundant paths and
avoid unnecessary computations. These tables store distances to specific states and guide the search
more effectively.

# Results

# Conclusions
The A* algorithm is very efficient for problems where a fixed graph is available and does not change over time. However, when
applying it to solving the Rubik's Cube, challenges arose due to the vast number of possible states the cube can take. To tackle
this issue, optimizations were necessary to reduce the search space and improve the algorithm's performance.

One of the key factors in achieving this was the way the cube was represented, as an efficient representation made it easier to
apply the Kociemba algorithm. The latter allowed for finding optimal solutions in fewer moves by leveraging a more effective
heuristic search. Thanks to these optimizations, the project’s objectives were successfully met, demonstrating that the
combination of a well-structured problem representation and advanced search strategies can make algorithms like A* viable even for
problems with large search spaces.

# Bibliography




