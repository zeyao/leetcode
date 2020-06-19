public class LongestDuplicateSubstr {
    
    //Rolling hash
    //binary search for length that contains duplicate string,
    //if len l contains duplicate string, then we store it and no need search for 0 to l - 1;
    //if len l not contains dulpcatie string, then len with l+ 1 to len also not valid;
    //for each length doing a rolling hash

   class Solution {
    
    int base = 26;
    int mod = 1000000007;
    
    public String longestDupSubstring(String S) {
        String res = "";
        int left = 2;
        int right = S.length()-1;
        
        int maxLen = 0;
        int endIndex = 0;
        
        while (left <= right) {
            boolean find = false;
            int mid = (right - left) / 2 + left;
            
            //search for len with mid;
            Map<Long, List<Integer>> map = new HashMap<>();//use list as handle hash collison, store the end index for that string with the key hash
            //rolling hash for the length with mid;
            long pow = 1;
            long hash = 0;
            for (int i = 0 ; i < mid; i++) {
                hash = (base * hash + (S.charAt(i) - 'a')) % mod;
                pow = pow * base % mod;
            }
            map.put(hash, new ArrayList<>());
            map.get(hash).add(mid-1);
            for (int i = mid; i < S.length(); i++) {
                hash = (hash * base + (S.charAt(i) - 'a')) % mod;
                hash = ((hash - (S.charAt(i-mid) - 'a') * pow % mod) + mod) % mod;
                if (!map.containsKey(hash)) {
                    map.put(hash, new ArrayList<>());
                }
                else {
                    for (int j : map.get(hash)) {
                        // j is end index of the collision, check is it the new string same with what store in hashtable
                        if (compare(i, j, S, mid)) {
                            maxLen = mid;
                            endIndex = i;
                            find = true;
                            break;
                        }
                    }
                    if (find) {
                        break;
                    }
                }
                map.get(hash).add(i);
            }
            if (find) {
                left = mid + 1;
            }
            else {
                right = mid - 1;
            }
        }
        return S.substring(endIndex - maxLen + 1, endIndex + 1);
        
    }
    
    private boolean compare(int i, int j, String S, int len) {
        for (int a = 0 ; a < len; a++) {
            if (S.charAt(i-a) != S.charAt(j-a)) {
                return false;
            }
        }
        return true;
    }
    
}
}