class Solution {
  public:
    void segregate0and1(vector<int> &arr) {
        // code here
        
        int l = 0, r = arr.size() - 1;
        
        while(l < r) {
            
            while(arr[l] == 0 && l < r) {
                l++;
            }
            
            while(arr[r] == 1 && l < r) {
                r--;
            }
            
            if(l < r ) {
                int temp = arr[l];
                arr[l] = arr[r];
                arr[r] = temp;
                
                l++;
                r--;
            }
            
        }
    }
};