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
