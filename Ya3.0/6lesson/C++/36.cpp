#include <iostream>
#include <vector>
#include <deque>

using namespace std;

int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);

    int n;
    cin >> n;

    vector<vector<int>> g(n + 1);
    for (int i = 1; i != n + 1; ++i) {
        for (int j = 1; j != n + 1; ++j) {
            int edge;
            cin >> edge;
            if (edge == 1) {
                g[i].push_back(j);
            }
        }
    }

    int start, end ;
    cin >> start >> end;

    // так, ну тут можно впихнуть блядский if в очередь, но проще лишний вектор булов намутить
    vector<int> dist(n + 1, 0);
    vector<bool> visit(n + 1, false);
    dist[start] = 0;
    visit[start] = true;

    deque<int> q;
    q.push_back(start);

    int level = 1;
    while (!q.empty()) {
        int levelSize = q.size();
        
        for (int i = 0; i != levelSize; ++i) {
            int cur = q.front();

            for (int j = 0; j != g[cur].size(); j++) {
                int nxtV = g[cur][j];
                if (!visit[nxtV]) {
                    q.push_back(nxtV);
                    dist[nxtV] = level;
                    visit[nxtV] = true;
                }
            }

            q.pop_front();
        }

        level++;
    }

    if (!visit[end]) {
        cout << -1 << endl;
    } else {
        cout << dist[end] << endl;
    }
    
    return 0;
}