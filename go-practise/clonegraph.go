/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    
    visited:=make(map[*Node]*Node)
    return dfs(node, visited)
}

func dfs(node *Node , visited map[*Node]*Node) *Node{
   if node == nil {
    return nil
   }
    
    if v , ok := visited[node]; ok {
        return v
    }
    clone:=&Node{Val:node.Val}
    visited[node] = clone
    for _ , neigh := range node.Neighbors{
      clone.Neighbors = append(clone.Neighbors , dfs(neigh , visited))
    }

    return clone
}