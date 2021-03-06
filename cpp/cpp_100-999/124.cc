// C++
// 124
// https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
// 通过深度优先遍历，算到经过每个节点的max路径，如果比全局的结果大，则替换
// 每个树可以分为根节点，左路最大路径和，右路最大路径和

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */

class Solution {
public:
    int res = INT_MIN;

    int visit(TreeNode* root){
        if(root == nullptr) return 0;
        int le = INT_MIN, ri = INT_MIN;
        le = max(visit(root->left) , 0);
        ri = max(visit(root->right), 0);
        int me = root->val + le + ri;
        res = max(res, me);
        return root->val + max(le,ri);
    }

    int maxPathSum(TreeNode* root) {
        visit(root);
        return res;
    }
};