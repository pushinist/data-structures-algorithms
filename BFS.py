from collections import deque

SOURCE: int = 0


def convert_adjacency_matrix_to_matrix_list(path_to_matrix: str) -> list[list[int]]:
    with open(path_to_matrix) as matrix_file:
        adjacency_list: list[list[int]] = []
        new_raw: list = matrix_file.readline().split()
        new_raw = [int(i) for i in new_raw]
        adjacency_matrix: list = []
        raw_number: int = 0
        while new_raw != []:
            adjacency_matrix.append(new_raw)
            for column_number in range(len(adjacency_matrix)):
                if adjacency_matrix[column_number][raw_number] == 1:
                    adjacency_list.append([column_number, raw_number])
            raw_number += 1
            new_raw = matrix_file.readline().split()
            new_raw = [int(i) for i in new_raw]
        return adjacency_list


def bfs(
    graph: list[list[int]], source: int, parent: list[int], distance: list[float]
) -> None:
    q = deque()
    distance[source] = 0
    q.append(source)

    while q:
        node: int = q.popleft()

        for neighbor in graph[node]:
            if distance[neighbor] == float("inf"):
                parent[neighbor] = node
                distance[neighbor] = distance[node] + 1
                q.append(neighbor)


def print_shortest_distance(graph: list[list[int]], source: int, vertices: int) -> None:
    parent: list[int] = [-1] * vertices

    distance: list[int] = [float("inf")] * vertices

    bfs(graph, source, parent, distance)

    for i in range(vertices - 1):
        if i != source:
            path: list[int] = []
            current_node: int = i
            path.append(current_node)
            while parent[current_node] != -1:
                path.append(parent[current_node])
                current_node = parent[current_node]
            for i in range(len(path) - 1, -1, -1):
                if i != 0:
                    print(path[i], end=" -> ")
                else:
                    print(path[i])


edges: list[list[int]] = convert_adjacency_matrix_to_matrix_list("adjacency_matrix")
vertices: int = len(edges)
graph: list[list[int]] = [[] for _ in range(vertices)]
for edge in edges:
    graph[edge[0]].append(edge[1])
    graph[edge[1]].append(edge[0])
print_shortest_distance(graph, SOURCE, vertices)
