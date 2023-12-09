#include <iostream>
#include <vector>

using namespace std;

typedef long long ll;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int n, q;
    cin >> n >> q;
    vector<ll> data(n);
    vector<ll> preff(n + 1);

    for (int i = 0; i != n; ++i) {
        cin >> data[i];
    }

    for (int i = 1; i != n + 1; ++i) {
        preff[i] = preff[i - 1] + data[i - 1];
    }

    for (int i = 0; i != q; i++) {
        int l, r;
        cin >> l >> r;
        l--;
        cout << preff[r] - preff[l] << "\n";
    }
    return 0;
}