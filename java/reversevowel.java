import java.util.Arrays;

import java.util.*;

public class reversevowel {
    public String reverseVowels(String s) {
        Set<Character> set = new HashSet<>(Arrays.asList('a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'));
        char[] arr = s.toCharArray();
        int left = 0;
        int right = arr.length -1;
        while (left < right) {
            while (left < right && !set.contains(arr[left])) {
                left++;
            }
            while (left < right && !set.contains(arr[right])) {
                right--;
            }
            if (left < right) {
                char tmp = arr[left];
                arr[left] = arr[right];
                arr[right] = tmp;
                left++;
                right--;
            }
        }
        return new String(arr);
    } 
}