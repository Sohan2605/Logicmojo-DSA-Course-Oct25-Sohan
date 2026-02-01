# 📘 Graph DSA  (Java Oriented)

---
1️⃣ Must-Remember Concepts (Non-negotiable)

You should be able to say these out loud in an interview.

🔹 Graph Basics

Graph = nodes + edges

Directed vs Undirected

Weighted vs Unweighted

Cyclic vs Acyclic

Tree = special graph


🔹 Graph Representation
Map<Integer, List<Integer>> graph;

Be ready to explain:
Why adjacency list over matrix
Space: O(V + E)


2️⃣ DFS & BFS Templates (MOST IMPORTANT)
🔵 DFS Template (Memorize)
void dfs(int node) {
    visited[node] = true;

    for (int neighbor : graph.get(node)) {
        if (!visited[neighbor]) {
            dfs(neighbor);
        }
    }
}


BFS Template (Memorize)
Queue<Integer> q = new LinkedList<>();
q.offer(start);
visited[start] = true;

while (!q.isEmpty()) {
    int node = q.poll();

    for (int neighbor : graph.get(node)) {
        if (!visited[neighbor]) {
            visited[neighbor] = true;
            q.offer(neighbor);
        }
    }
}



# Graph Basics

## 1. Degree of a Node

### 🔹 Degree (Undirected Graph)

* **Degree of a node** = Number of edges connected to it

Example:

```
A --- B --- C
```

* Degree(A) = 1
* Degree(B) = 2
* Degree(C) = 1

📌 **Important Rule**:

* Sum of degrees of all nodes = **2 × number of edges**

---

## 2. Indegree and Outdegree (Directed Graph)

### 🔹 Indegree

* Number of **incoming edges** to a node

### 🔹 Outdegree

* Number of **outgoing edges** from a node

Example:

```
A → B → C
↑         ↓
└─────────┘
```

* Indegree(A) = 1, Outdegree(A) = 1
* Indegree(B) = 1, Outdegree(B) = 1
* Indegree(C) = 1, Outdegree(C) = 1

📌 **Important Rule**:

* Sum of indegrees = Sum of outdegrees = number of edges

---

## 3. Degree in Special Graphs

### 🔹 Self-loop

* Edge from a node to itself
* Adds:

  * **2 to degree** (undirected)
  * **1 indegree + 1 outdegree** (directed)

### 🔹 Parallel Edges

* Multiple edges between same nodes
* Each edge counts separately in degree

---

## 4. Types of Graphs

### 🔹 Undirected Graph

* Edges are **bidirectional**

```
A --- B
```

* Degree concept only

---

### 🔹 Directed Graph (Digraph)

* Edges have direction

```
A → B
```

* Uses indegree & outdegree

---

### 🔹 Weighted Graph

* Each edge has a weight/cost

```
A --5--> B
```

* Used in shortest path problems

---

### 🔹 Unweighted Graph

* All edges have equal weight
* BFS works for shortest path

---

### 🔹 Cyclic Graph

* Contains at least one cycle

Example:

```
A → B → C → A
```

---

### 🔹 Acyclic Graph

* No cycles

Example:

* Tree
* DAG (Directed Acyclic Graph)

---

### 🔹 DAG (Directed Acyclic Graph)

* Directed graph with no cycles
* Used in:

  * Topological sort
  * Course scheduling

---

### 🔹 Connected Graph

* Every node is reachable from every other node

---

### 🔹 Disconnected Graph

* Graph has more than one connected component

---

### 🔹 Complete Graph

* Every node connected to every other node
* For `n` nodes:

  * Edges = `n(n-1)/2`

---

### 🔹 Bipartite Graph

* Nodes can be colored using **2 colors**
* No two adjacent nodes have same color

Used in:

* Matching problems

---

### 🔹 Tree

* Connected + Acyclic
* Edges = Nodes - 1

---

### 🔹 Forest

* Collection of trees

---

## 5. Interview Quick Notes

* BFS → shortest path (unweighted)
* DFS → cycles, components
* Indegree array → Topological Sort
* Degree concept → graph properties

---

## 6. Common LeetCode Usage

* Indegree → Course Schedule
* Degree → Graph validation
* Directed graph → Dependencies
* Undirected graph → Islands, components

---

✅ **These concepts are foundational for all graph proble


Below are Java solutions + short explanations for each of the graph problems you listed.
I’ll keep them interview-friendly and clear, not too long but complete.

1. Graph Traversal (BFS & DFS)
BFS – Breadth First Search

Use Queue. Visit level by level.

import java.util.*;

class GraphBFS {
    public static void bfs(int start, List<List<Integer>> adj, int n) {
        boolean[] visited = new boolean[n];
        Queue<Integer> q = new LinkedList<>();

        q.add(start);
        visited[start] = true;

        while (!q.isEmpty()) {
            int node = q.poll();
            System.out.print(node + " ");

            for (int nei : adj.get(node)) {
                if (!visited[nei]) {
                    visited[nei] = true;
                    q.add(nei);
                }
            }
        }
    }
}


Time: O(V + E)

DFS – Depth First Search

Use Recursion / Stack

class GraphDFS {
    public static void dfs(int node, List<List<Integer>> adj, boolean[] visited) {
        visited[node] = true;
        System.out.print(node + " ");

        for (int nei : adj.get(node)) {
            if (!visited[nei]) {
                dfs(nei, adj, visited);
            }
        }
    }
}

2. Clone Graph (LeetCode 133)

Idea: Use HashMap to map original → cloned node.

class Node {
    public int val;
    public List<Node> neighbors;
    public Node(int val) { this.val = val; neighbors = new ArrayList<>(); }
}

class Solution {
    public Node cloneGraph(Node node) {
        if (node == null) return null;
        Map<Node, Node> map = new HashMap<>();
        return dfs(node, map);
    }

    private Node dfs(Node node, Map<Node, Node> map) {
        if (map.containsKey(node)) return map.get(node);

        Node copy = new Node(node.val);
        map.put(node, copy);

        for (Node nei : node.neighbors) {
            copy.neighbors.add(dfs(nei, map));
        }
        return copy;
    }
}


Concept: Avoid infinite loop using map.

3. Rotting Oranges (Multi-Source BFS)

Idea: All rotten oranges start BFS simultaneously.

class Solution {
    public int orangesRotting(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        Queue<int[]> q = new LinkedList<>();
        int fresh = 0, minutes = 0;

        for (int i=0;i<m;i++) {
            for (int j=0;j<n;j++) {
                if (grid[i][j] == 2) q.add(new int[]{i,j});
                if (grid[i][j] == 1) fresh++;
            }
        }

        int[][] dirs = {{1,0},{-1,0},{0,1},{0,-1}};

        while (!q.isEmpty() && fresh > 0) {
            int size = q.size();
            minutes++;

            for (int s=0;s<size;s++) {
                int[] cur = q.poll();
                for (int[] d : dirs) {
                    int x = cur[0]+d[0], y = cur[1]+d[1];
                    if (x>=0 && y>=0 && x<m && y<n && grid[x][y]==1) {
                        grid[x][y] = 2;
                        fresh--;
                        q.add(new int[]{x,y});
                    }
                }
            }
        }
        return fresh==0 ? minutes : -1;
    }
}

4. Longest Increasing Path in Matrix

Idea: DFS + Memoization (DP)

class Solution {
    int[][] dirs = {{1,0},{-1,0},{0,1},{0,-1}};

    public int longestIncreasingPath(int[][] matrix) {
        int m = matrix.length, n = matrix[0].length;
        int[][] dp = new int[m][n];
        int ans = 0;

        for (int i=0;i<m;i++)
            for (int j=0;j<n;j++)
                ans = Math.max(ans, dfs(matrix, i, j, dp));

        return ans;
    }

    private int dfs(int[][] mat, int r, int c, int[][] dp) {
        if (dp[r][c] != 0) return dp[r][c];

        int max = 1;
        for (int[] d : dirs) {
            int nr = r+d[0], nc = c+d[1];
            if (nr>=0 && nc>=0 && nr<mat.length && nc<mat[0].length &&
                mat[nr][nc] > mat[r][c]) {
                max = Math.max(max, 1 + dfs(mat, nr, nc, dp));
            }
        }
        return dp[r][c] = max;
    }
}

5. Dijkstra’s Algorithm

Idea: Min Heap (PriorityQueue)

class Dijkstra {
    static int[] dijkstra(int V, List<List<int[]>> adj, int src) {
        int[] dist = new int[V];
        Arrays.fill(dist, Integer.MAX_VALUE);
        dist[src] = 0;

        PriorityQueue<int[]> pq = new PriorityQueue<>(Comparator.comparingInt(a -> a[1]));
        pq.add(new int[]{src, 0});

        while (!pq.isEmpty()) {
            int[] cur = pq.poll();
            int node = cur[0], d = cur[1];

            if (d > dist[node]) continue;

            for (int[] nei : adj.get(node)) {
                int next = nei[0], w = nei[1];
                if (dist[node] + w < dist[next]) {
                    dist[next] = dist[node] + w;
                    pq.add(new int[]{next, dist[next]});
                }
            }
        }
        return dist;
    }
}

6. Network Delay Time (LeetCode 743)

Just Dijkstra variant

class Solution {
    public int networkDelayTime(int[][] times, int n, int k) {
        List<List<int[]>> adj = new ArrayList<>();
        for (int i=0;i<=n;i++) adj.add(new ArrayList<>());

        for (int[] t : times)
            adj.get(t[0]).add(new int[]{t[1], t[2]});

        int[] dist = new int[n+1];
        Arrays.fill(dist, Integer.MAX_VALUE);
        dist[k] = 0;

        PriorityQueue<int[]> pq = new PriorityQueue<>(Comparator.comparingInt(a -> a[1]));
        pq.add(new int[]{k,0});

        while (!pq.isEmpty()) {
            int[] cur = pq.poll();
            int node = cur[0], d = cur[1];
            if (d > dist[node]) continue;

            for (int[] nei : adj.get(node)) {
                if (dist[node] + nei[1] < dist[nei[0]]) {
                    dist[nei[0]] = dist[node] + nei[1];
                    pq.add(new int[]{nei[0], dist[nei[0]]});
                }
            }
        }

        int max = 0;
        for (int i=1;i<=n;i++) {
            if (dist[i] == Integer.MAX_VALUE) return -1;
            max = Math.max(max, dist[i]);
        }
        return max;
    }
}

7. Bellman–Ford

Handles Negative Weights

class BellmanFord {
    static int[] bellmanFord(int V, int[][] edges, int src) {
        int[] dist = new int[V];
        Arrays.fill(dist, Integer.MAX_VALUE);
        dist[src] = 0;

        for (int i=0;i<V-1;i++) {
            for (int[] e : edges) {
                int u=e[0], v=e[1], w=e[2];
                if (dist[u] != Integer.MAX_VALUE && dist[u]+w < dist[v]) {
                    dist[v] = dist[u]+w;
                }
            }
        }

        // Negative cycle check
        for (int[] e : edges) {
            if (dist[e[0]] + e[2] < dist[e[1]]) {
                System.out.println("Negative Cycle Detected");
            }
        }
        return dist;
    }
}