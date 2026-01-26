
class deleteMidFromStack {
    // Function to delete middle element of a stack.
    public void deleteMid(Stack<Integer> s) {
        // code here
        
        int n = s.size();
        n = n / 2;
        helper(s, n);
        
        // Stack<Integer> temp = new Stack<>();
        
        // for(int i=0; i<n; i++) {
        //     temp.push(s.pop());    
        // }
        
        // for(int i=0; i<n-1; i++){
        //     if(i==0){
        //         temp.pop();
        //         continue;
        //     }
        //     s.push(temp.pop());
        // }
        // s.push(temp.pop());
    }
    
    public void helper(Stack<Integer> s, int n) {
        if(n==0) {
            s.pop();
            return;
        }
        
        int val = s.pop();
        helper(s, n - 1);
        s.push(val);   
    }
}