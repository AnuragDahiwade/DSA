class Solution {
    public void sortStack(Stack<Integer> st) {
        // code here
        helper(st);
    }
    
    public void helper(Stack<Integer> st) {
        if(st.isEmpty()) {
            return;
        }
        int val = st.pop();
        helper(st);
        putAtCorrectPlace(st, val);
    }
    
    
    public void putAtCorrectPlace(Stack<Integer> st, int num) {
        if(st.isEmpty()) {
            st.push(num);
            return;
        }
        
        int val = st.pop();
        
        if(val <= num) {
            st.push(val);
            st.push(num);
            return;
        }
        putAtCorrectPlace(st, num);
        st.push(val);
    }
}