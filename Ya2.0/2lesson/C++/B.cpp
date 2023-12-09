#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

int secondrun(vector<int>& a, int stpos, int endpos, int ls, int rs);

int maxdist(vector<int>& a) {
    int lastpos = -1;
    int ans = 0;// никому 20 шагов точно идти не прийдется

    for (int i = 0; i != 10; ++i) {
        if (a[i] == 2) {
            if (lastpos == -1) {
                ans = max(ans, secondrun(a, 0, i, -100, i));
            } else {
                ans = max(ans, secondrun(a, lastpos + 1, i, lastpos, i));
            }
            lastpos = i;
        }
    }

    ans = max(ans, secondrun(a, lastpos + 1, 10, lastpos, 100));
    return ans;
}

int secondrun(vector<int>& a, int stpos, int endpos, int ls, int rs) {
    int maxdist = 0;
    for (int i = stpos; i != endpos; ++i) {
        if (a[i] == 1) {
            maxdist = max(maxdist, min(i - ls, rs - i));
        }
    }
    return maxdist;
}

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    vector<int> a(10);
    for (int i = 0; i != 10; ++i) cin >> a[i];
    cout << maxdist(a) << endl;
    return 0;
}