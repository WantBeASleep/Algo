#include <bits/stdc++.h>
 
using namespace std;
 
int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int q;
    cin >> q;
    for (int gl = 0; gl < q; gl++) {
        string s;
        cin >> s;
        vector<int> z(s.length(), 0);
        int l = 0, r = 0;
        for (int i = 1; i < s.length(); i++) {
            if (i < r) z[i] = min(z[i - l], r - i);
            while (i + z[i] < s.length() and s[z[i]] == s[i + z[i]]) z[i]++;
            if (r < i + z[i]) {
                r = i + z[i];
                l = i;
            }
        }
 
        vector<int> ans(z.size(), 1);
        for (int i = 0; i < z.size(); i++) {
            if (z[i] > 0) ans[z[i] - 1]++;
        }
        for (int i = ans.size() - 1 - 1; i >= 0; i--) {
            ans[i] += ans[i + 1] - 1;
        }
        for (int i = 0; i < ans.size(); i++) {
            cout << ans[i] << " ";
        }
        cout << endl;
    }
    return 0;
}