#include <bits/stdc++.h>

using namespace std;

vector<int> zfunc(string s) {
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
    return z;
}

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    string t, p;
    int k;
    
    cin >> t >> p;
    cin >> k;

    int n = t.length(); 
    int m = p.length();
    
    string directs = p + "$" + t;
    reverse(p.begin(), p.end());
    reverse(t.begin(), t.end());
    string reverses = p + "$" + t;

    vector<int> zd = zfunc(directs);
    vector<int> zr = zfunc(reverses);

    vector<int> ans;
    for (int l = 0; l < n - m + 1; l++) {
        int a = zd[m + 1 + l];
        int b = zr[n - l + 1];
        if (a + b >= m - k) ans.push_back(l + 1);
    }
    
    cout << ans.size() << "\n";
    for (int num : ans) {
        cout << num << " ";
    }
    cout << "\n";
    
    return 0;
}