import networkx as nx

G = nx.Graph() # F it, make the graph global so we dont have to build it twice
def find_triangles_with_t(connections: list[str]) -> int:
    for connection in connections:
        a, b = connection.split('-')
        G.add_edge(a, b)
        
    triangles = [clique for clique in nx.enumerate_all_cliques(G) if len(clique) == 3]
    return len([triangle for triangle in triangles if any(node.startswith('t') for node in triangle)])

def part1(input: list[str]) -> int:
    return find_triangles_with_t(input)

def find_lan_party_password() -> str:    
    cliques = list(nx.find_cliques(G))
    password = ','.join(sorted(max(cliques, key=len)))
    return password

def part2() -> str:
    return find_lan_party_password()

def main():
    input = open("input.txt", "r").read().splitlines()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2()}")

if __name__ == "__main__":
    main()