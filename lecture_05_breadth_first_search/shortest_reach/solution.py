# num of testcases
t = int(input().strip())
# 4 2
# 1 2
# 1 3
# 1
def calculate_cost(nodes, travel_edges):    
    # the queue
    new_edges = []
    weight = 6
    current_cost = 0
    # initial condition, the start already given
    # travel_edges is an array
    while travel_edges:
        for e in travel_edges:
            id, cost, edges = nodes[e-1]
            # the node is not visited
            if cost == -1:
                # append all the connected nodes to the current node to the queue
                new_edges = new_edges + edges
                cost = current_cost
                nodes[e-1] = (id, cost, edges)
        # start the while loop again 
        travel_edges = new_edges
        # delete the temp queue
        new_edges = []
        # update the cost
        current_cost += weight
                
    return nodes
    
for test in range(t):
    # num of nodes, edges
    n, m = [int(i) for i in input().strip().split()]
    
    # graph, also the path
    nodes = [(i, -1, []) for i in range(1, n+1)]

    for edge in range(m):
        # u, v
        start, end = [int(i) for i in input().strip().split()]
        # node start from 0, get the pev value
        id, cost, edges = nodes[start-1]
        # graph[u] = append(graph[u], v), add more to the graph
        nodes[start-1] = (id, cost, edges + [end])
        # graph[v] = append(graph[v], u), for the opposite node
        id, cost, edges = nodes[end-1]
        nodes[end-1] = (id, cost, edges + [start])
    # start 
    s = int(input().strip())
    
    nodes = calculate_cost(nodes, [s])
    
    weights = ' '.join([str(node[1]) for node in nodes if node[1] != 0])
    
    print(weights)