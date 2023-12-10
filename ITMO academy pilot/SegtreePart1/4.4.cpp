/*
    РЕШЕНО КАК ГОВНО, НО МНЕ БЫЛО ОООЧЧЕНЬ ЛЕНЬ ПЕРЕПИСЫВАТЬ, ЗНАЯ ЧТО ТУТ ВСЕ РАВНО РЕШЕНИЕ ГОВНА, 
    ДЕРЕВА НА UNORDERED_SET. РЕШЕНИЕ ОТ АВГУСТА.
*/

#include <bits/stdc++.h>

using namespace std;

typedef unordered_set<int> ms;

class segtree {
    public:
        vector<ms> a;
        int size;

        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, {});
        }

        void build(vector<int>& _a) {
            build(_a, 0, 0, size);
        }

        void build(vector<int>& _a, int x, int lx, int rx) {
            if (rx == lx + 1) {
                if (lx < _a.size()) a[x].insert(_a[lx]);
                return;
            }
            int m = (lx + rx) / 2;
            build(_a, 2 * x + 1, lx, m);
            build(_a, 2 * x + 2, m, rx);
            a[x] = a[2 * x + 1];
            a[x].insert(a[2 * x + 2].begin(), a[2 * x + 2].end());
        }

        void set(int i, int v) {
            set(i, v, 0, 0, size);
        }

        int set(int i, int v, int x, int lx, int rx) {
            if (rx == lx + 1) {
                int del = *a[x].begin();
                a[x].clear();
                a[x].insert(v);
                return del;
            }
            int del;
            int m = (lx + rx) / 2;
            if (i < m) del = set(i, v, 2 * x + 1, lx, m);
            else del = set(i, v, 2 * x + 2, m, rx);
            if (!(a[2 * x + 1].count(del) or a[2 * x + 2].count(del))) a[x].erase(del);
            a[x].insert(v);
            return del;
        }

        int uni(int l, int r) {
            ms ans = {};
            uni(l, r, 0, 0, size, ans);
            return ans.size();
        }

        void uni(int l, int r, int x, int lx, int rx, ms& ans) {
            if (lx >= r or rx <= l) return;
            if (lx >= l and rx <= r) return ans.insert(a[x].begin(), a[x].end());
            int m = (lx + rx) / 2;
            uni(l, r, 2 * x + 1, lx, m, ans);
            uni(l, r, 2 * x + 2, m, rx, ans);
        }
};



int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    srand(time(NULL));
    int n, q;
    cin >> n >> q;
    vector<int> d(n);
    for (int i = 0; i < n; i++) cin >> d[i];

    segtree sg;
    sg.init(n);
    sg.build(d);

    for (int i = 0; i < q; i++) {
        int type;
        cin >> type;

        if (type == 1) {
            int l, r;
            cin >> l >> r;
            l--;

            cout << sg.uni(l, r) << endl;
        } else {
            int i, v;
            cin >> i >> v;
            i--;

            sg.set(i, v);
        }
    }

    return 0;
}