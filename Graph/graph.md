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
