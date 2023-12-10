#include <iostream>

#include <vector>
#include <deque>
#include <unordered_set>
#include <string>
#include <utility>

using namespace std;

typedef pair<int, int> pr;

vector<int> dx = {1, 1, -1, -1, 2, 2, -2, -2};
vector<int> dy = {2, -2, 2, -2, 1, -1, 1, -1};

string ptos(pair<int, int> p);

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);

    int n, m, s, t, q;
    cin >> n >> m >> s >> t >> q;

    vector<vector<bool>> visit(n + 1);
    
    for (int i = 1; i != n + 1; i++) {
        visit[i].assign(m + 1, false);
    }

    unordered_set<string> sq(q);

    for (int i = 0; i != q; i++) {
        int a, b;
        cin >> a >> b;
        sq.insert(to_string(a) + ":" + to_string(b));
    }

    deque<pair<int, int>> qe;
    qe.push_back(pr(s, t));
    visit[s][t] = true;
    
    int len = 1;
    int cntfly = 0;

    if (sq.count(ptos(pr(s, t)))) {
        cntfly++;
    }

    int ans = 0;
    while (!qe.empty()) {
        int oldS = qe.size();

        for (int i = 0; i != oldS; ++i) {
            pr cur = qe.front();

            for (int j = 0; j != 8; j++) {
                pr nxt = pr(cur.first + dx[j], cur.second + dy[j]);
                // вылет
                if (nxt.first < 1 || nxt.first > n || nxt.second < 1 || nxt.second > m) {
                    continue;
                }

                // тут нужно проверить посещали ли мы уже эту вершину, в голову лезет одно говно, 
                // наиболее просто наверное ебнуть массивчик булеаноv
                if (!visit[nxt.first][nxt.second]) {
                    if (sq.count(ptos(nxt))) {
                        ans += len;
                        cntfly++;
                    }

                    qe.push_back(nxt);
                    visit[nxt.first][nxt.second] = true;
                }
            }

            qe.pop_front();
        }

        len++;
    }

    // cout << cntfly << endl;
    
    if (cntfly != q) {
        cout << -1 << endl;
    } else {
        cout << ans << endl;
    }

    return 0;
}

string ptos(pair<int, int> p) {
    return to_string(p.first) + ":" + to_string(p.second);
}