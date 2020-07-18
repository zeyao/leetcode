import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.*;

//https://leetcode.com/problems/merge-k-sorted-lists/ LeetCode 23 变种 if input is array

public class MergeKSortedArray {
    public static void main(String[] args) {	
        int[][] mat = new int[][]{
            {1,4,5},
            {1,3,4},
            {2,6}
        };
        System.out.println(sortKList(mat));
        
        mat = new int[][] {
            {1,2,3},
            {90,100,200,300},
            {78,79},
            {4,5,6,7,8},
            {4,5,6,7,8},
        };
        System.out.println(sortKList(mat));
    }

    private static List<Integer> sortKList(int[][] mat) {
        List<Integer> res = new ArrayList<>();
        PriorityQueue<Node> pq = new PriorityQueue<>((a, b) -> a.val - b.val);
        for (int i = 0; i < mat.length; i++) {
            if (mat[i].length > 0) {
                pq.offer(new Node(mat[i][0], i, 0));
            }
        }
        while (!pq.isEmpty()) {
            Node node = pq.poll();
            res.add(node.val);
            int i = node.row;
            int j = node.col+1;
            if (j < mat[i].length) {
                pq.offer(new Node(mat[i][j], i, j));
            }
        }
        return res;
    }

    static class Node {
        int val;
        int row;
        int col;
        public Node(int val, int row, int col) {
            this.val = val;
            this.row = row;
            this.col = col;
        }
    }
}