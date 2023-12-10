#include <bits/stdc++.h>

using namespace std;

typedef tuple<int, int, int, int> t;

const t uno = t(1, 0, 0, 1);

class Segtree {
    public:
        Segtree(vector<t>& data, int mod) {
            init(data.size(), mod);
            build(data, 0, 0, size);
        }

        t multrange(int l, int r) {
            return multrange(l, r, 0, 0, size);
        }
    private:
        vector<t> a;
        int size;
        int mod;
        /*
            [0, 1]
            [2, 3]
        */
        t mult(t i, t j) {
            int i0, i1, i2, i3;
            int j0, j1, j2, j3;
            int r0, r1, r2, r3;


            tie(i0, i1, i2, i3) = i;
            tie(j0, j1, j2, j3) = j;

            r0 = (i0 * j0 + i1 * j2) % mod;
            r1 = (i0 * j1 + i1 * j3) % mod;
            r2 = (i2 * j0 + i3 * j2) % mod;
            r3 = (i2 * j1 + i3 * j3) % mod;

            return t(r0, r1, r2, r3);
        }

        void init(int n, int mod) {
            size = 1;
            this->mod = mod;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, uno);
        }

        void build(vector<t>& data, int x, int lx, int rx) {
            if (rx == lx + 1) {
                if (lx < data.size()) a[x] = data[lx];
                return;
            }

            int m = (lx + rx) / 2;
            build(data, 2 * x + 1, lx, m);
            build(data, 2 * x + 2, m, rx);
            a[x] = mult(a[2 * x + 1], a[2 * x + 2]);
        }

        t multrange(int l, int r, int x, int lx, int rx) {
            if (lx >= r or rx <= l) return uno;
            if (lx >= l and rx <= r) return a[x];

            int m = (lx + rx) / 2;
            t left = multrange(l, r, 2 * x + 1, lx, m);
            t right = multrange(l, r, 2 * x + 2, m, rx);
            return mult(left, right);
        }
};

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int r, n, m;
    cin >> r >> n >> m;
    vector<t> a(n);
    for (int i = 0; i < n; i++) {
        int i0, i1, i2, i3;
        cin >> i0 >> i1 >> i2 >> i3;
        a[i] = t(i0, i1, i2, i3);
    }

    Segtree sg(a, r);
    while (m--) {
        int l, r;
        cin >> l >> r;
        l--;
        t res = sg.multrange(l, r);

        int a1, a2, a3, a4;
        tie(a1, a2, a3, a4) = res;
        cout << a1 << " " << a2 << "\n" << a3 << " " << a4 << "\n" << "\n";
    }
    return 0;
}