#include <bits/stdc++.h>

using namespace std;

typedef long long ll;

class segtree {
    public:
        vector<ll> a;
        int size;

        void build(vector<int>& data) {
            init(data.size());
            build(data, 0, 0, size);
        }

        void set(int i, int v) {
            set(i, v, 0, 0, size);
        }

        ll sum(int l, int r) {
            return sum(l, r, 0, 0, size);
        }
    
    private:
        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, 0);
        }

        void build(vector<int>& data, int x, int lx, int rx) {
            if (rx == lx + 1) {
                if (lx < data.size()) a[x] = data[lx];
                return; // !
            }
            int m = (lx + rx) / 2;
            build(data, 2 * x + 1, lx, m);
            build(data, 2 * x + 2, m, rx);
            a[x] = a[2 * x + 1] + a[2 * x + 2];
        }

        void set(int i, int v, int x, int lx, int rx) {
            if (rx == lx + 1) {
                a[x] = v;
                return;
            }

            int m = (lx + rx) / 2;
            if (i < m) set(i, v, 2 * x + 1, lx, m);
            else set(i, v, 2 * x + 2, m, rx);
            a[x] = a[2 * x + 1] + a[2 * x + 2];
        }

        ll sum(int l, int r, int x, int lx, int rx) {
            if (lx >= r or rx <= l) return 0;
            if (lx >= l and rx <= r) return a[x];

            
            int m = (lx + rx) / 2;
            ll lsum = sum(l, r, 2 * x + 1, lx, m);
            ll rsum = sum(l, r, 2 * x + 2, m, rx);
            return lsum + rsum;
        }
    
};

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int n, m;
    cin >> n >> m;
    vector<int> a(n);
    for (int i = 0; i < n; i++) cin >> a[i];

    segtree sg;
    sg.build(a);

    while (m--) {
        int type;
        cin >> type;
        if (type == 1) {
            int i, v;
            cin >> i >> v;
            sg.set(i, v);
        } else {
            int l, r;
            cin >> l >> r;
            cout << sg.sum(l, r) << "\n";
        }
    }
    return 0;
}