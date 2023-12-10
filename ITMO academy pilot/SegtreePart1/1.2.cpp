#include <bits/stdc++.h>

using namespace std;

#define LMAX 1e9 + 1

class segtree {
    public:
        vector<int> a;
        int size;

        void build(vector<int>& data) {
            init(data.size());
            build(data, 0, 0, size);
        }

        void set(int i, int v) {
            set(i, v, 0, 0, size);
        }

        int min(int l, int r) {
            return min(l, r, 0, 0, size);
        }
    
    private:
        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, LMAX);
        }

        void build(vector<int>& data, int x, int lx, int rx) {
            if (rx == lx + 1) {
                if (lx < data.size()) a[x] = data[lx];
                return; 
            }
            int m = (lx + rx) / 2;
            build(data, 2 * x + 1, lx, m);
            build(data, 2 * x + 2, m, rx);
            a[x] = ::min(a[2 * x + 1], a[2 * x + 2]);
        }

        void set(int i, int v, int x, int lx, int rx) {
            if (rx == lx + 1) {
                a[x] = v;
                return;
            }

            int m = (lx + rx) / 2;
            if (i < m) set(i, v, 2 * x + 1, lx, m);
            else set(i, v, 2 * x + 2, m, rx);
            a[x] = ::min(a[2 * x + 1], a[2 * x + 2]);
        }

        int min(int l, int r, int x, int lx, int rx) {
            if (lx >= r or rx <= l) return LMAX;
            if (lx >= l and rx <= r) return a[x];

            
            int m = (lx + rx) / 2;
            int lmin = min(l, r, 2 * x + 1, lx, m);
            int rmin = min(l, r, 2 * x + 2, m, rx);
            return ::min(lmin, rmin);
        }
    
};

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int n, m;
    cin >> n >> m;

    vector<int> data(n);
    for (int i = 0; i < n; i++) cin >> data[i];

    segtree sg;
    sg.build(data);

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

            cout << sg.min(l, r) << "\n";
        }
    }  
    return 0;
}