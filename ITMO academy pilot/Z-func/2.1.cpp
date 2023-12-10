#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    string str;
    cin >> str;
    vector<int> z(str.length(), 0);

    for (int i = 1; i < str.size(); i++) {
        while (i + z[i] < str.size() and str[z[i]] == str[z[i] + i]) z[i]++;
    }

    for (int i = 0; i < z.size(); i++) cout << z[i] << " ";
    cout << endl;
    

    return 0;
}