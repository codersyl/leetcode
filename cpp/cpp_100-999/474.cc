// C++
// leetcode 474
// https://leetcode-cn.com/problems/ones-and-zeroes/
// 


class Solution {
public:

    int findMaxForm(vector<string>& strs, int m, int n) {
        int size = strs.size();
        int cnt[size][2];
        for(int i = 0; i < size; i++){
            int zero = 0, one = 0;
            for(char c : strs[i]) {
                if(c== '0')
                    zero++;
                else
                    one++;
            }
            cnt[i][0] = zero;
            cnt[i][1] = one;
        }   // 预处理完毕

        // f[i][j]表示不超过i个0与j个1能得到的个数
        int f[m+1][n+1];
        // 初始化为0
        memset(f, 0, sizeof(f));
        for(int k = 0; k < size; k++){
            int zero = cnt[k][0], one = cnt[k][1];
            for(int i = m; i >= 0; i--){
                for(int j = n; j >= 0; j--){
                    // 不要第k件
                    int a =f[i][j];
                    // 要第k件
                    int b = (i >= zero && j >= one) ? f[i - zero][j - one] + 1 : 0;
                    f[i][j] = max(a,b);
                }
            }
        }
        return f[m][n];
    }
};