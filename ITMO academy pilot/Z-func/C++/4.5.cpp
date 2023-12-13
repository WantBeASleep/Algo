#include <bits/stdc++.h>
 
using namespace std;
 
int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    string s;
    cin >> s;
    int n = s.length();
    
    s = s + "$" + s;
    reverse(s.begin() + n + 1, s.end());
 
    vector<int> z(s.length(), 0);
    int l = 0, r = 0;
    for (int i = 1; i < s.length(); i++) {
        if (i < r) z[i] = min(z[i - l], r - i);
        while (z[i] + i < s.length() and s[z[i]] == s[i + z[i]]) z[i]++;
        if (r < i + z[i]) {
            r = i + z[i];
            l = i;
        }
    }
 
    for (int i = n + 1; i < s.length(); i++) {
        if (z[i] + i == s.length()) {
            cout << z[i] << endl;
            break;
        }
    }
   
    return 0;
}