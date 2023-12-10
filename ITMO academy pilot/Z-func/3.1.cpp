#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    string s;
    cin >> s;
    vector<int> z(s.length(), 0);
    int l = 0; int r = 0;

    for (int i = 1; i < s.length(); i++) {
        if (r > i) z[i] = min(z[i - l], r - i);
        while (i + z[i] < s.length() and s[z[i]] == s[i + z[i]]) z[i]++;
        if (r < i + z[i]) {
            r = i + z[i];
            l = i;
        }
    }

    for (int i = 0; i < s.length(); i++) cout << z[i] << " ";
    cout << endl;
    return 0;
}