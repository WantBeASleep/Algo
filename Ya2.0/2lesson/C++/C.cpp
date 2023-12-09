#include <iostream>

using namespace std;

int cost(string &a) {
    int ans = 0;
    for (int i = 0; i < a.size() / 2; ++i) {
        if (a[i] != a[a.size() - 1 - i]) ++ans;
    }
    return ans;
}

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    string a;
    cin >> a;
    cout << cost(a) << endl;
    return 0;
}