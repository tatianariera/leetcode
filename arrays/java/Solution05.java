package arrays.java;

public class Solution05 {

    public int singleNumber(int[] nums) {

        int resultado = 0;

        for (int num : nums) {
            resultado ^= num;
        }

        return resultado;
    }
}
