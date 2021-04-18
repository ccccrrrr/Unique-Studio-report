#include <bits/stdc++.h>
#include <algorithm>
using namespace std;
#define N 5
#define M 7
struct node {
    int u;
    int v;
    int w;
};
bool func(const node* a, const node* b){
    return a->w < b->w;
}
struct node * vertex[N];
queue<struct node *> q;
//int u[N] = {0, 0, 0, 1, 2, 2, 3};
//int v[N] = {1, 3, 4, 2, 3, 4, 5};
//int w[N] = {2, 5, 9, 6, 3, 7, 4};
void CreateGraph(void);
int findParent(int v);
void addTree(int u, int v, int w);
int parent[N] = {0};
int main(){

    CreateGraph();

    // use Krustal algorithm
    // when value > 0 its value is parent index
    // while value <= 0, it means it is a single node or it is a parent node
    // and its total weight is its value
    while(q.size() != 0){
        struct node * temp = q.front();
        if(findParent(temp->u) != findParent(temp->v) || (findParent(temp->u) == 0 && findParent(temp->v) == 0)){
            printf("%d %d %d\n", temp->u, temp->v, temp->w);
            addTree(temp->u, temp->v, temp->w);
        }
        q.pop();
    }
    printf("\n");
    for(int i = 0; i < N; i++){
        printf("%d ", parent[i]);
    }
    

    return 0;
}

int findParent(int p){
    while(parent[p] > 0){
        if(parent[parent[p]] > 0)
            parent[p] = parent[parent[p]];
        p = parent[p];
    }
    return p;
}

void addTree(int u, int v, int w){
    if(findParent(u) < findParent(v)){
        //value weight sum
        parent[findParent(v)] = parent[findParent(v)] + parent[findParent(u)] - w;
        parent[findParent(u)] = v;
    }else{
        parent[findParent(u)] = parent[findParent(u)] + parent[findParent(v)] - w;
        parent[findParent(v)] = u;
    }
}

void CreateGraph(void){
    for(int i = 0; i < M; i++)
        vertex[i] = new node;
    vertex[0]->u = 0;vertex[0]->v = 1;vertex[0]->w = 2;
    vertex[1]->u = 0;vertex[1]->v = 3;vertex[1]->w = 5;
    vertex[2]->u = 0;vertex[2]->v = 4;vertex[2]->w = 9;
    vertex[3]->u = 1;vertex[3]->v = 2;vertex[3]->w = 6;
    vertex[4]->u = 2;vertex[4]->v = 3;vertex[4]->w = 3;
    vertex[5]->u = 2;vertex[5]->v = 4;vertex[5]->w = 7;
    vertex[6]->u = 3;vertex[6]->v = 4;vertex[6]->w = 4;
    vector<struct node *> v;
    for(int i = 0; i < M; i++)
        v.push_back(vertex[i]);
    sort(v.begin(), v.end(), func);
    for(int i = 0; i < v.size(); i++){
        q.push(v.at(i));
    }
    printf("successfully create the graph!\n");    
}
