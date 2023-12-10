#include <bits/stdc++.h>

using namespace std;

class Segtree{
    public:
        Segtree(vector<int>& data) {
            init(data.size());
            build(data, 0, 0, size);
        }

        void set(int i, int v) {
            set(i, v, 0, 0, size);
        }

        int getidx(int v) {
            return getidx(v, 0, 0, size);
        }

    private:
        vector<int> a;
        int size;

        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, 0);
        }

        void build(vector<int>& data, int x, int lx, int rx) {
            if (rx == lx + 1) {
                if (lx < data.size()) a[x] = data[lx];
                return;
            }

            int m = (lx + rx) / 2;
            build(data, 2 * x + 1, lx, m);
            build(data, 2 * x + 2, m, rx);
            a[x] = ::max(a[2 * x + 1], a[2 * x + 2]);
        }

        void set(int i, int v, int x, int lx, int rx) {
            if (rx == lx + 1) {
                a[x] = v;
                return;
            }

            int m = (lx + rx) / 2;
            if (i < m) set(i, v, 2 * x + 1, lx, m);
            else set(i, v, 2 * x + 2, m, rx);
            a[x] = ::max(a[2 * x + 1], a[2 * x + 2]);
        }

        int getidx(int v, int x, int lx, int rx) {
            if (a[x] < v) return -1;
            if (rx == lx + 1) {
                return lx;
            }

            int m = (lx + rx) / 2;
            if (a[2 * x + 1] >= v) return getidx(v, 2 * x + 1, lx, m);
            else return getidx(v, 2 * x + 2, m, rx);
        }
};

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int n, m;
    cin >> n >> m;
    vector<int> a(n);
    for (int i = 0; i < n; i++) cin >> a[i];
    Segtree sg(a);

    while (m--) {
        int type;
        cin >> type;
        if (type == 1) {
            int i, v;
            cin >> i >> v;
            sg.set(i, v);
        } else {
            int x;
            cin >> x;
            cout << sg.getidx(x) << "\n";
        }
    }
    return 0;
}