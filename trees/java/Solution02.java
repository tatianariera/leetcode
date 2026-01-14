
public class Solution02 {

    public class TreeNode {

        TreeNode left;
        TreeNode right;
        int val;
    }

    public boolean isValidBST(TreeNode root) {
        return helper(root, null, null);
    }

    public boolean helper(TreeNode node, Integer min, Integer max) {
        if (node == null) {
            return true;
        }

        if ((min != null && node.val <= min) || (max != null && node.val >= max)) {
            return false;
        }

        return helper(node.left, min, node.val) && helper(node.right, node.val, max);
    }
}
