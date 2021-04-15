#include <bits/stdc++.h>
using namespace std;
#define N 5
int graph[5][5] = {
    {1, 0, 1, 1, 0},
    {0, 1, 0, 1, 1},
    {1, 1, 1, 1, 1},
    {0, 0, 1, 1, 1}
};
int parent[N] = {0};
int main(){
    int s = 0;
    for(int i = 0; i < N; i++)
    {
        for(int j = 0; j < N; j++){
            if(graph[i][j] == 1)

        }
    }

    return 0;
}