#include <bits/stdc++.h>

using namespace std;

class Segtree {
    public:
        Segtree(int n) {
            init(2 * n);
        }

        void set(int i, int v) {
            set(i, v, 0, 0, size);
        }

        int sum(int l, int r) {
            return sum(l, r, 0, 0, size);
        }

    private:
        vector<int> a;
        int size;

        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, 0);
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

        int sum(int l, int r, int x, int lx, int rx) {
            if (rx <= l or lx >= r) return 0;
            if (lx >= l and rx <= r) return a[x];

            int m = (lx + rx) / 2;
            int lsum = sum(l, r, 2 * x + 1, lx, m);
            int rsum = sum(l, r, 2 * x + 2, m, rx);
            return lsum + rsum;
        }
};

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int n;
    cin >> n;
    vector<int> a(2 * n);
    for (int i = 0; i < 2 * n; i++) cin >> a[i];
    vector<int> spos(n, -1);
    vector<int> ans(n, 0);


    Segtree sg1(n);
    for (int i = 0; i < 2 * n; i++) {
        if (spos[a[i] - 1] == -1) {
            spos[a[i] - 1] = i;
            sg1.set(i, 1);
        } else {
            sg1.set(spos[a[i] - 1], 0);
            ans[a[i] - 1] += sg1.sum(spos[a[i] - 1] + 1, i);
        }
    }

    spos.assign(n, -1);
    reverse(a.begin(), a.end());
    Segtree sg2(n);
    for (int i = 0; i < 2 * n; i++) {
        if (spos[a[i] - 1] == -1) {
            spos[a[i] - 1] = i;
            sg2.set(i, 1);
        } else {
            sg2.set(spos[a[i] - 1], 0);
            ans[a[i] - 1] += sg2.sum(spos[a[i] - 1] + 1, i);
        }
    }

    for (int i = 0; i < n; i++) {
        cout << ans[i] << " ";
    }
    cout << "\n";

    return 0;
}