#include <bits/stdc++.h>

using namespace std; 

class Segtree {
    public:
        Segtree(int n) {
            init(n);
        }

        void set(int i, int v) {
            set(i, v, 0, 0, size);
        }

        int getsuff(int i) {
            return getsuff(i, 0, 0, size);
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
            else set (i, v, 2 * x + 2, m, rx);
            a[x] = a[2 * x + 1] + a[2 * x + 2];
        }

        int getsuff(int l, int x, int lx, int rx) {
            if (rx <= l) return 0;
            if (lx >= l) return a[x];

            int m = (lx + rx) / 2;
            return getsuff(l, 2 * x + 1, lx, m) + getsuff(l, 2 * x + 2, m, rx);
        }
};

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int n;
    cin >> n;
    vector<int> data(n);
    for (int i = 0; i < n; i++) cin >> data[i];

    Segtree sg(n);

    for (int num : data) {
        sg.set(num - 1, 1);
        cout << sg.getsuff(num) << " ";
    }
    cout << "\n";

    return 0;
}