import java.util.*;


public class combinationSum {
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        List<List<Integer>> res = new ArrayList<>();
        dfs(res, new ArrayList<>(), target, 0, candidates);
        return res;
    }
    
    private void dfs(List<List<Integer>> res, List<Integer> innerList, int target, int start, int[] candidates) {
        if (target == 0) {
            res.add(new ArrayList<>(innerList));
            return;
        }
        if (target < 0) {
            return;
        }
        for (int i = start; i < candidates.length; i++) {
            innerList.add(candidates[i]);
            dfs(res, innerList, target - candidates[i], i, candidates);
            innerList.remove(innerList.size()-1);
        }
    }
}

