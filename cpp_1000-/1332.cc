// C++
// leetcode 1332
// https://leetcode-cn.com/problems/remove-palindromic-subsequences/

class Solution {
public:
    int removePalindromeSub(string s) {
    	for(int i = 0, j = s.size() - 1; i< j; i++, j--) {
    		if(s[i] != s[j]) return 2;
    	}
    	return 1;
    }
};