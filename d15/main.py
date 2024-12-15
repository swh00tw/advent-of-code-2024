"""
COPY FROM: https://www.reddit.com/r/adventofcode/comments/1hele8m/comment/m27qmzw/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
"""


DELTAS = {"v": (1, 0), "^": (-1, 0), ">": (0, 1), "<": (0, -1)}


# -------------------------------------------------------------------------------------
def read_file(filename: str, double_wide: bool) -> list[list[str]]:
    """Read in the input file and return grid."""
    with open(filename, "rt", encoding="utf-8") as fh:
        lines = [line.strip() for line in fh.readlines()]  # list of lists
    if double_wide:
        widen = {"#": "##", "O": "[]", ".": "..", "@": "@."}
        grid = [[] for _ in range(len(lines))]
        for i, line in enumerate(lines):
            for char in line:
                grid[i].extend(list(widen[char]))
        return grid

    return [list(line) for line in lines]


# -------------------------------------------------------------------------------------
def read_moves(filename: str) -> str:
    """Read in the moves file and return one continuous string of moves."""
    with open(filename, "rt", encoding="utf-8") as fh:
        return "".join([line.strip() for line in fh.readlines()])


# -------------------------------------------------------------------------------------
def find_start(grid: list[list[int]]) -> tuple[int, int]:
    """Find the starting position and return a coordinate tuple."""
    for row, line in enumerate(grid):
        if "@" in line:
            return (row, line.index("@"))


# -------------------------------------------------------------------------------------
def move_and_push(grid, start_r, start_c, dr, dc) -> tuple[int, int]:
    """Determines if the move is valid. If so, pushes all of the boxes that are in the
    way of the direction of travel. Returns the next position of the robot."""
    stack = []
    path = [(start_r, start_c)]
    visited = set()
    while path:
        r, c = path.pop()
        if (r, c) in visited or grid[r][c] == ".":
            continue
        visited.add((r, c))
        if grid[r][c] == "#":
            return (start_r, start_c)
        stack.append(((grid[r][c], r, c)))
        path.append((r + dr, c + dc))
        if grid[r][c] == "[":
            path.append((r, c + 1))
        if grid[r][c] == "]":
            path.append((r, c - 1))

    if dr > 0:
        stack.sort(key=lambda path: path[1])
    if dr < 0:
        stack.sort(key=lambda path: -path[1])
    if dc > 0:
        stack.sort(key=lambda path: path[2])
    if dc < 0:
        stack.sort(key=lambda path: -path[2])

    while stack:
        char, old_r, old_c = stack.pop()
        grid[old_r + dr][old_c + dc] = char
        grid[old_r][old_c] = "."

    return (start_r + dr, start_c + dc)


# -------------------------------------------------------------------------------------
def solve(map_file: str, moves_file: str, double_wide: bool) -> int:
    """Moves the robot through the warehouse based on the sequence of moves provided.
    Calculates the sum of the 'GPS coordinates' of the final box positions."""
    grid = read_file(map_file, double_wide)
    moves = read_moves(moves_file)
    r, c = find_start(grid)
    for move in moves:
        dr, dc = DELTAS[move]
        r, c = move_and_push(grid, r, c, dr, dc)

    # get "GPS" Coordinates
    coordinate_sum = 0
    for r, row in enumerate(grid):
        for c, location in enumerate(row):
            if location in {"[", "O"}:
                coordinate_sum += 100 * r + c

    print("\n".join("".join(cell for cell in row) for row in grid))

    return coordinate_sum


# -------------------------------------------------------------------------------------
if __name__ == "__main__":
    print("Part 1:", solve("input", "moves", double_wide=False))
    print("\nPart 2:", solve("input", "moves", double_wide=True))