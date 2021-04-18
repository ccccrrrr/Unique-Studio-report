#include <bits/stdc++.h>
#include <stack>
using namespace std;
#define N 10
struct TreeNode{
    int value;
    struct TreeNode * left;
    struct TreeNode * right;
};

void CreateTree(struct TreeNode * root);

void preOrder_r(struct TreeNode * root);
void preOrder__(struct TreeNode * root);
void preOrder_morris(struct TreeNode * root);

void inOrder_r(struct TreeNode * root);
void inOrder__(struct TreeNode * root);
void inOrder_morris(struct TreeNode * root);

void postOrder_r(struct TreeNode * root);
void postOrder__(struct TreeNode * root);
void postOrder_morris(struct TreeNode * root);

void levelOrder_r(struct TreeNode * root);
void levelOrder__(struct TreeNode * root);

void bfsTree(struct TreeNode * root);
void dfsTree(struct TreeNode * root);

void addPath(vector<int>& res, struct TreeNode * root);

int main(){

    struct TreeNode * root = (struct TreeNode *)malloc(sizeof(struct TreeNode));
    root->value = 0;
    root->left = root->right = nullptr;

    CreateTree(root);
    // 先序
    preOrder_r(root);
    printf("\n");
    preOrder__(root);
    printf("\n");
    preOrder_morris(root);
    printf("\n");

    //中序
    printf("in order using recursion: ");
    inOrder_r(root);
    printf("\n");
    printf("in order using queue: ");
    inOrder__(root);
    printf("\n");
    inOrder_morris(root);
    printf("\n");

    //后序
    printf("post order recursion: ");
    postOrder_r(root);
    printf("\n");
    printf("post order iteration: ");
    postOrder__(root);
    printf("\n");
    printf("post order morris: ");
    postOrder_morris(root);
    printf("\n");

    //层序遍历
    levelOrder__(root);
    printf("\n");

    dfsTree(root);
    printf("\n");
    bfsTree(root);
}

void preOrder_r(struct TreeNode * root){
    if(root == NULL) return;
    //if(root->left == NULL)
    //    printf("%d ", root->value);
    printf("%d ", root->value);
    if(root->left)
        preOrder_r(root->left);
    if(root->right)
        preOrder_r(root->right);
    
}
void preOrder__(struct TreeNode * root){
    struct TreeNode * t = root;
    stack<struct TreeNode *> s;
    while(t || s.size() != 0){
        while(t){
            printf("%d ", t->value);
            s.push(t);
            t = t->left;
        }
        if(s.size() != 0){
            t = s.top();
            s.pop();
            t = t->right;
        }
    }

}
void preOrder_morris(struct TreeNode * root){
    if(root == nullptr){
        printf("error input...\n");
        return;
    }
    printf("preorder morris: ");
    struct TreeNode * cur = root;
    struct TreeNode * mostright;
    while(!(cur->left == nullptr && cur->right == nullptr)){
        //printf("%d ", cur->value);
        if(cur->left == nullptr){
            printf("%d ", cur->value);
            cur = cur->right;
        }
        else{
            mostright = cur->left;
            while(mostright->right != nullptr && mostright->right != cur){
                mostright = mostright->right;
            }
            if(mostright->right == nullptr){
                printf("%d ", cur->value);
                mostright->right = cur; 
                cur = cur->left;
            }
            else{
                mostright->right = nullptr;
                cur = cur->right;
            }
        }
    }
}

void inOrder_r(struct TreeNode * root){
    if(root == nullptr) return;
    if(root->left)
        inOrder_r(root->left);
    printf("%d ", root->value);
    if(root->right)
        inOrder_r(root->right);
}
void inOrder__(struct TreeNode * root){
    struct TreeNode * t = root;
    stack<struct TreeNode *> s;
    while(t || s.size() != 0){
        while(t){
            s.push(t);
            t = t->left;
        }
        t = s.top();
        printf("%d ", t->value);
        s.pop();
        t = t->right;
    }
}
void inOrder_morris(struct TreeNode * root){
    if(root == nullptr){
        printf("error input...\n");
        return;
    }
    printf("in order morris: ");
    struct TreeNode * cur = root;
    struct TreeNode * mostright;
    while(cur){
        if(cur->left == nullptr){
            printf("%d ", cur->value);
            cur = cur->right;
        }
        else{
            mostright = cur->left;
            while(mostright->right != nullptr && mostright->right != cur){
                mostright = mostright->right;
            }
            if(mostright->right == nullptr){
                mostright->right = cur;
                cur = cur->left;
                continue;
            }
            else{
                mostright->right = nullptr;
                printf("%d ", cur->value);
                cur = cur->right;
            }
        }
    }
}

void postOrder_r(struct TreeNode * root){//O(n)
    if(root == nullptr) return;
    if(root->left)
        postOrder_r(root->left);
    if(root->right)
        postOrder_r(root->right);
    printf("%d ", root->value);
}
void postOrder__(struct TreeNode * root){//O(n)
    stack<struct TreeNode *> s;
    struct TreeNode * t = root;
    struct TreeNode * ptr;
    while(t || s.size() != 0){
        while(t){
            s.push(t);
            t = t->left;
        }
        t = s.top();
        s.pop();
        if(t->right == nullptr || ptr == t->right){
            printf("%d ", t->value);
            ptr = t;
            t = nullptr;
        }
        else{
            s.push(t);
            t = t->right;
        }
    }
}
void postOrder_morris(struct TreeNode * root){//O(n)
    //problem reserved
    struct TreeNode * cur = root;
    struct TreeNode * mostright;
    stack<struct TreeNode *> s;
    vector<int> res;
    while(cur != nullptr){
        if(cur->left == nullptr){
            cur = cur->right;
        }else{
            mostright = cur->left;
            while(mostright->right != nullptr && mostright->right != cur){
                mostright = mostright->right;
            }
            if(mostright->right == nullptr){
                mostright->right = cur;
                cur = cur->left;
                continue;
            }else {
                mostright->right = nullptr;
                addPath(res, cur->left);
                cur = cur->right;
            }
        }
        
    }
    addPath(res, root);
    for(vector<int>::iterator it = res.begin(); it != res.end(); it++)
        printf("%d ", *it);
}

void addPath(vector<int>& res, struct TreeNode * root){
    int count = 0;
    while(root != nullptr){
        ++count;
        res.emplace_back(root->value);
        root = root->right;
    }
    reverse(res.end() - count, res.end());
}

void levelOrder_r(struct TreeNode * root){
    printf("is there any recursive solution...?\n");
}
void levelOrder__(struct TreeNode * root){
    if(root == nullptr) {
        printf("root is null pointer...\n");
        return;
    }
    printf("level order: ");
    queue<struct TreeNode *> q;
    struct TreeNode * t = root;
    q.push(root);
    while(q.size() != 0){
        t = q.front();
        if(t->left) q.push(t->left);
        if(t->right) q.push(t->right);
        printf("%d ", t->value);
        q.pop(); 
    }
}

void dfsTree(struct TreeNode * root){
    if(root == nullptr){
        printf("this is a null pointer...\n");
        return;
    }
    printf("dfs: ");
    stack<struct TreeNode *> s;
    s.push(root);
    while(s.size() != 0){
        printf("%d ", s.top()->value);
        struct TreeNode * temp = s.top();
        s.pop();
        if(temp->right) s.push(temp->right);
        if(temp->left) s.push(temp->left);
    }
}
void bfsTree(struct TreeNode * root){
    if(root == nullptr){
        printf("this is a null pointer...\n");
        return;
    }
    printf("bfs: ");
    list<struct TreeNode *> l;
    l.push_back(root);
    struct TreeNode * temp;
    while(l.size() != 0){
        temp = l.front();
        if(temp->left) l.push_back(temp->left);
        if(temp->right) l.push_back(temp->right);
        printf("%d ", temp->value);
        l.pop_front();
    }
}

void CreateTree(struct TreeNode * root){
    printf("tree node number is %d, you can change tree size in #define N\n", N);

    static struct TreeNode * list[N];
    list[0] = root;
    for(int i = 1; i < N; i++){
        list[i] = (struct TreeNode *)malloc(sizeof(struct TreeNode));
        list[i]->left = list[i]->right = nullptr;
        list[i]->value = i;
    }
    int parent = 0;
    while(2 * parent + 1 < N){
        //printf("parent: %d\n", parent);
        list[parent]->left = list[2 * parent + 1];
        if(2 * parent + 2 < N){
            list[parent]->right = list[2*parent + 2];
            //printf("parent! = %d\n", parent);
        }
        parent++;
    }
    
}