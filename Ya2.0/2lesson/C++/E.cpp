#include <iostream>
#include <vector>

using namespace std;

int ans(vector<int>& a) {
    int count = 0, max = 0;
    for (int i = 0; i != a.size(); ++i) {
        count += a[i];
        if (a[i] > a[max]) max = i;
    }

    return count - a[max];
}

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int n;
    cin >> n;
    vector<int> a(n);
    for (int i = 0; i != n; ++i) cin >> a[i];

    cout << ans(a) << endl;
    return 0;
}