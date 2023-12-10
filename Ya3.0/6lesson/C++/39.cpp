#include <iostream>

#include <vector>
#include <deque>
#include <string>

using namespace std;

struct point {
    int x, y, z;
};

vector<int> dx = {-1, 0, 0, 0, 0, 1};
vector<int> dy = {0, 1, -1, 0, 0, 0};
vector<int> dz = {0, 0, 0, 1, -1, 0};


int main()
{
    ios::sync_with_stdio(NULL), cin.tie(0), cout.tie(0);
    int n;
    string line;
    cin >> n;

    vector<vector<vector<bool>>> cube(n, vector<vector<bool>>(n, vector<bool>(n, false)));
    vector<vector<vector<bool>>> visit(n, vector<vector<bool>>(n, vector<bool>(n, false)));
    point start;

    for (int z = 0; z != n; ++z) {
        getline(cin, line);
        getline(cin, line);

        for (int y = 0; y != n; ++y) {
            for (int x = 0; x != n; ++x) {
                char ch;
                cin >> ch;
                if (ch == '.') {
                    cube[x][y][z] = true;
                }
                if (ch == 'S') {
                    start.x = x;
                    start.y = y;
                    start.z = z;
                }
            }
        }
    }
    // блять...

    int level = 0;
    visit[start.x][start.y][start.z] = true;

    deque<point> q;
    q.push_back(start);

    int len = 1;

    while (!q.empty()) {
        int oldS = q.size();

        for (int i = 0; i != oldS; ++i) {
            point cur = q.front();

            for (int j = 0; j != 6; ++j) {
                point newP = {cur.x + dx[j], cur.y + dy[j], cur.z + dz[j]};

                if (newP.x >= n || newP.x < 0 || newP.y >= n || newP.y < 0 || newP.z >= n || newP.z < 0) {
                    continue;
                }

                if (cube[newP.x][newP.y][newP.z] && !visit[newP.x][newP.y][newP.z]) {
                    if (newP.z == 0) {
                        cout << len << endl;
                        return 0;
                    }

                    q.push_back(newP);
                    visit[newP.x][newP.y][newP.z] = true;
                }

            }

            q.pop_front();
        }

        len++;
    }


    return 0;
}