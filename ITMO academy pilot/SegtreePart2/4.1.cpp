#include <iostream>
#include <vector>

using namespace std;

typedef long long ll;

class Segtree{
    public:
        Segtree(int n) {
            init(n);
        }

        void modify(int l, int r, int v, int type) {
            modify(l, r, v, 0, 0, size, type);
        }

        ll calc(int l, int r) {
            return calc(l, r, 0, 0, size);
        }

    private:

        struct node {
            ll val;
            ll f;
            bool op;
            int type;
        };

        vector<node> a;
        int size;

        void init(int n) {
            size = 1;
            while (size < n) size *= 2;
            a.assign(2 * size - 1, {0, 0, false, 0});
        }

        ll op_mod1(ll v) {
            return v;
        }

        ll op_mod2(ll a, ll v) {
            return a + v;
        }

        ll op_calc(ll a, ll b) {
            return a + b;
        }

        void propagate(int x, int lx, int rx) {
            if (!a[x].op or (rx == lx + 1)) return;

            int half = (rx - lx) / 2;

            // присваивание
            if (a[x].type == 1) {
                a[2 * x + 1].f = op_mod1(a[x].f);
                a[2 * x + 2].f = op_mod1(a[x].f);

                a[2 * x + 1].val = op_mod1(a[x].f * half);
                a[2 * x + 2].val = op_mod1(a[x].f * half);

                a[x].f = 0;
                a[x].op = false;
                a[x].type = 0;

                a[2 * x + 1].op = true;
                a[2 * x + 2].op = true;
                
                a[2 * x + 1].type = 1;
                a[2 * x + 2].type = 1;

            } else if (a[x].type == 2) {
                
                if (a[2 * x + 1].op and a[2 * x + 1].type == 1) {
                    a[2 * x + 1].f = op_mod2(a[2 * x + 1].f, a[x].f);
                    a[2 * x + 1].type = 1;
                }
                else if (a[2 * x + 1].op and a[2 * x + 1].type == 2) {
                    a[2 * x + 1].f = op_mod2(a[2 * x + 1].f, a[x].f);
                    a[2 * x + 1].type = 2;
                } else {
                    a[2 * x + 1].f = op_mod1(a[x].f);
                    a[2 * x + 1].type = 2;  
                }

                if (a[2 * x + 2].op and a[2 * x + 2].type == 1) {
                    a[2 * x + 2].f = op_mod2(a[2 * x + 2].f, a[x].f);
                    a[2 * x + 2].type = 1;
                }
                else if (a[2 * x + 2].op and a[2 * x + 2].type == 2) {
                    a[2 * x + 2].f = op_mod2(a[2 * x + 2].f, a[x].f);
                    a[2 * x + 2].type = 2;
                } else {
                    a[2 * x + 2].f = op_mod1(a[x].f);
                    a[2 * x + 2].type = 2;  
                }   
                    
                
                a[2 * x + 1].val = op_mod2(a[2 * x + 1].val, a[x].f * half);
                a[2 * x + 2].val = op_mod2(a[2 * x + 2].val, a[x].f * half);

                a[x].f = 0;
                a[x].op = false;
                a[x].type = 0;

                a[2 * x + 1].op = true;
                a[2 * x + 2].op = true;
            }
        }

        void modify(int l, int r, ll v, int x, int lx, int rx, int type) {
            propagate(x, lx, rx);
            if (rx <= l or lx >= r) return;
            if (lx >= l and rx <= r) {
                if (type == 1) {
                    a[x].val = op_mod1(v * (rx - lx));
                    a[x].f = v;
                    a[x].op = true;
                    a[x].type = 1;
                } else {
                    a[x].val = op_mod2(a[x].val, v * (rx - lx));
                    if (a[x].op and a[x].type == 1) {
                        a[x].f = op_mod2(a[x].f, v);
                        a[x].type = 1;
                    } else if (a[x].op and a[x].type == 2) {
                        a[x].f = op_mod2(a[x].f, v);
                        a[x].type = 2;
                    } else {
                        a[x].f = op_mod1(v);
                        a[x].type = 2;
                    }

                    a[x].op = true;
                }
                return;
            }

            int m = (lx + rx) / 2;
            modify(l, r, v, 2 * x + 1, lx, m, type);
            modify(l, r, v, 2 * x + 2, m, rx, type);

            a[x].val = op_calc(a[2 * x + 1].val, a[2 * x + 2].val);
        }

        ll calc(int l, int r, int x, int lx, int rx) {
            if (rx <= l or lx >= r) return 0;
            if (lx >= l and rx <= r) return a[x].val;

            int m = (lx + rx) / 2;  
            
            ll left = calc(l, r, 2 * x + 1, lx, m);
            ll right = calc(l, r, 2 * x + 2, m, rx);

            if (!a[x].op) return op_calc(left, right);
            else if (a[x].type == 1) return op_mod1(a[x].f * (min(r, rx) - max(l, lx)));
            else return op_mod2(op_calc(left, right), a[x].f * (min(r, rx) - max(l, lx)));
        }
};

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);

    int n, m;
    cin >> n >> m;
    Segtree sg(n);
    while (m--) {
        int type;
        cin >> type;
        if (type == 1 or type == 2) {
            int l, r, v;
            cin >> l >> r >> v;
            sg.modify(l, r, v, type);
        } else {
            int l, r;
            cin >> l >> r;
            cout << sg.calc(l, r) << endl;
        }
    }
}