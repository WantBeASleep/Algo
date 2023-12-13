#include <bits/stdc++.h>

using namespace std;

#define LMAX 1e9 + 1

typedef pair<int, int> p;

const p LMAXP = p(LMAX, 1);

p min(p s1, p s2) {
    if (s1.first == s2.first) return p(s1.first, s1.second + s2.second);
    if (s1.first < s2.first) return s1;
    else return s2;
}

class segtree {
    public:
        vector<p> a;
        int size;

        void build(vector<int>& data) {
            init(data.size());
            build(data, 0, 0, size);
        }

        void set(int i, int v) {
            set(i, v, 0, 0, size);
        }

        p min(int l, int r) {
            return min(l, r, 0, 0, size);
        }
    
    private:
        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, LMAXP);
        }

        void build(vector<int>& data, int x, int lx, int rx) {
            if (rx == lx + 1) {
                if (lx < data.size()) a[x] = p(data[lx], 1);
                return; 
            }
            int m = (lx + rx) / 2;
            build(data, 2 * x + 1, lx, m);
            build(data, 2 * x + 2, m, rx);
            a[x] = ::min(a[2 * x + 1], a[2 * x + 2]);
        }

        void set(int i, int v, int x, int lx, int rx) {
            if (rx == lx + 1) {
                a[x] = p(v, 1);
                return;
            }

            int m = (lx + rx) / 2;
            if (i < m) set(i, v, 2 * x + 1, lx, m);
            else set(i, v, 2 * x + 2, m, rx);
            a[x] = ::min(a[2 * x + 1], a[2 * x + 2]);
        }

        p min(int l, int r, int x, int lx, int rx) {
            if (lx >= r or rx <= l) return LMAXP;
            if (lx >= l and rx <= r) return a[x];

            
            int m = (lx + rx) / 2;
            p lmin = min(l, r, 2 * x + 1, lx, m);
            p rmin = min(l, r, 2 * x + 2, m, rx);
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

            p ans = sg.min(l, r);
            cout << ans.first << " " << ans.second << "\n";
        }
    }  
    return 0;
}