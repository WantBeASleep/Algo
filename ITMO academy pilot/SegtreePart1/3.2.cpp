#include <bits/stdc++.h>

using namespace std;

class Segtree {
    public:
        Segtree(int n) {
            init(n);
            build(0, 0, size);
        }

        void set(int i, int v) {
            set(i, v, 0, 0, size);
        }

        int getnum(int k) {
            return getnum(k, 0, 0, size);
        }
        
    private:
        vector<int> a;
        int size;
        int rsize;

        void init(int n) {
            size = 1; rsize = n;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, 0);
        }

        void build(int x, int lx, int rx) {
            if (rx == lx + 1) {
                if (lx < rsize) a[x] = 1;
                return;
            }
            int m = (lx + rx) / 2;
            build(2 * x + 1, lx, m);
            build(2 * x + 2, m, rx);
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

        int getnum(int k, int x, int lx, int rx) {
            if (rx == lx + 1) return lx + 1;

            int m = (lx + rx) / 2;
            if (a[2 * x + 2] >= k) return getnum(k, 2 * x + 2, m, rx);
            else return getnum(k - a[2 * x + 2], 2 * x + 1, lx, m);
        }
};

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int n;
    cin >> n;
    vector<int> a(n);
    for (int i = 0; i < n; i++) cin >> a[i];
    vector<int> ans(n);

    Segtree sg(n);

    for (int i = n - 1; i >= 0; i--) {
        int findpos = a[i] + 1;
        int num = sg.getnum(findpos);
        ans[i] = num;
        sg.set(num - 1, 0);
    }

    for (int num : ans) {
        cout << num << " ";
    }
    cout << "\n";

    return 0;
}