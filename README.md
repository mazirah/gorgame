# gorgame

Green VS Red
Green VS Red is a game played on a 2D grid.

Each cell of this grid can be either green (represented by 1) or red (represented by 0). The game always recieves an initial
state of the grid which is called Generation Zero. After that a set of 4 rules are applied across the grid and those rule 
for the next generation

The rules are as follows:
1. Each red cell that is surrounded by exactly 3 or exactly 6 green cells will also become greenin the next generation
2. A red cell will stay red in the next generation if it has either   0, 1, 2, 4, 5, 7 or 8 green neighbours
3. Each green cell surrounded by up 0, 1, 4, 5, 7, or 8 green neighbours will become red in the next generation
4. A green cell will stay green in the next generation if it has either 2, 3, or 6 green neighbours
