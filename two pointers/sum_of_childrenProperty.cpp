class Node {
public:
    int data;
    Node* left;
    Node* right;

    Node(int val) {
        data = val;
        left = nullptr;
        right = nullptr;
    }
};

// https://www.geeksforgeeks.org/problems/children-sum-parent/1

class Solution {
  public:
    bool isSumProperty(Node *root) {
        // code here
        
        if(root == NULL){
            return true;
        }
        
        if(root->left == NULL && root->right == NULL){
            return true;
        }
        
        int l = 0, r = 0;
        
        if(root->left) {
            l = root->left->data;
        }
        if(root->right) {
            r = root->right->data;
        }
        
        if(root->data != l + r) {
            return false;
        }
        
        return isSumProperty(root->left) && isSumProperty(root->right);
    }
};