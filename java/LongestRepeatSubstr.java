public class LongestRepeatSubstr {
    int mod = 1000000007;


    private int hash(String str) {
        int hash = 0;
        int len = str.length();
        int i = 0;
        for (int j = len; j > 0; j--) {
            hash += ((str.charAt(i) - 'a') * Math.pow(10, j)) % mod;
            i++;
        }
        return hash;
    }
}

