
class Solution {
  long countTriplets(int n, int sum, long arr[]) {
    Arrays.sort(arr);

    long count = 0;

    for (int i = 0; i < n - 2; i++) {
      int l = i + 1, r = n - 1;

      while (l < r) {
        long s = arr[i] + arr[l] + arr[r];

        if (s < sum) {
          count += (r - l);
          l++;
        } else {
          r--;
        }
      }
    }
    return count;
  }
}
