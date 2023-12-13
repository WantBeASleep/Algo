#include <bits/stdc++.h>

using namespace std;

typedef long long ll;

class Segtree {
    public:
        Segtree(int n) {
            init(n);
        }

        void set(int i, int v) {
            set(i, v, 0, 0, size);
        }

        ll get(int i) {
            return get(i, 0, 0, size);
        }
        
    private:
        vector<ll> a;
        int size;

        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, 0);
        }

        void set(int i, int v, int x, int lx, int rx) {
            if (rx == lx + 1) {
                a[x] += v;
                return;
            }

            int m = (lx + rx) / 2;
            if (i < m) set(i, v, 2 * x + 1, lx, m);
            else set(i, v, 2 * x + 2, m, rx);
            a[x] = a[2 * x + 1] + a[2 * x + 2];
        }

        ll get(int i, int x, int lx, int rx) {
            if (lx > i) return 0;
            if (rx <= i + 1) return a[x];

            int m = (lx + rx) / 2;
            ll lsum = get(i, 2 * x + 1, lx, m);
            ll rsum = get(i, 2 * x + 2, m, rx);
            return lsum + rsum;
        }
};

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int n, m;
    cin >> n >> m;
    Segtree a(n);
    while (m--) {
        int type;
        cin >> type;
        if (type == 1) {
            int l, r, v;
            cin >> l >> r >> v;
            a.set(l, v);
            if (r < n) a.set(r, -1 * v);
        } else {
            int i;
            cin >> i;
            cout << a.get(i) << "\n";
        }
    }
    return 0;
}