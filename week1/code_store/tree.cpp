#include <bits/stdc++.h>
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

void inOrder_r(struct TreeNode * root);
void inOrder__(struct TreeNode * root);

void postOrder_r(struct TreeNode * root);
void postOrder__(struct TreeNode * root);

int main(){

    struct TreeNode * root = (struct TreeNode *)malloc(sizeof(struct TreeNode));
    root->value = 0;
    root->left = root->right = nullptr;

    CreateTree(root);
    // 先序
    preOrder_r(root);
    printf("\n");
    preOrder__(root);

    //中序
    inOrder_r(root);
    printf("\n");
    //后序
    postOrder_r(root);
    printf("\n");
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

}

void postOrder_r(struct TreeNode * root){
    if(root == nullptr) return;
    if(root->left)
        postOrder_r(root->left);
    if(root->right)
        postOrder_r(root->right);
    printf("%d ", root->value);
}
void postOrder__(struct TreeNode * root){

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