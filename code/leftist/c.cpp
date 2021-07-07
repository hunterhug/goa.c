/* 
与平衡树不同（平衡树是具有非常小的深度的，这也意味着到达任何一个节点所经过的边数很少），左偏树并不是为了 
快速访问所有的节点而设计的，它的目的是快速访问最小节点以及在对树修改后快速的恢复堆性质。 
左偏树是一种可合并堆，常用于优先级队列。 
 
左偏树有两个性质： 
1）堆的性质（注意：一般堆是完全二叉树，但这里不是）： 
A[parent(i)]>=A[i] or A[parent(i)]<=A[i]，即父节点大于（或者小于）子节点的值（这个“值”可以看成Key或者优先级）。 
 
2）左偏性质： 
节点的左子节点的距离不小于右子节点的距离，即dist(left(i))≥dist(right(i)) 。 
其中节点的距离等于它的右子节点的距离加1，即 dist(i) = dist( right(i) ) + 1 ；如果一个节点的右子节点为空节点， 
则其距离为0；为了满足性质，故规定空节点的距离为-1。 
 
这里左偏树提供一下功能： 
1）插入节点Insert， 时间复杂度O(logn) 
2）获得根节点（即最小节点（如果是最小堆）或者最大节点（如果是最大堆）），时间复杂度O(1) 
3）删除根节点DeleteRoot，时间复杂度O(logn) 
4）合并两个左偏树Merge，时间复杂度O(logn) 
*/  
#include <iostream>  
#include <functional>  
#include <cassert>  
using namespace std;  
  
//节点数据结构  
template<class Type>  
struct Node  
{  
    Type data;//数据  
    int dist;//距离  
    Node<Type> *left;//左儿子指针  
    Node<Type> *right;//右儿子指针  
  
    Node(Type val)  
    {  
        data = val;  
        dist = 0;  
        left = right = NULL;  
    }  
  
    ~Node()      
    {  
    }  
};  
  
//左偏树  
template <class Type,class Compare = less<Type>/*默认是小堆*/ >

class LeftistTree  
{  
private:  
    Node<Type> *root;  
  
    //删除以指定节点为根的左偏树  
    void Delete(Node<Type> *node)  
    {  
        if(NULL != node)  
        {  
            Delete(node->left);  
            Delete(node->right);  
            delete node;  
            //node = NULL;  
        }  
    }  
  
    //判断指定节点是否有左儿子  
    //static bool HasLeft(Node<Type>* node)  
    //{  
    //  if(!node) return false;//空节点没有左儿子  
    //  return node->left;  
    //}  
  
    //获取指定节点的左儿子  
    static Node<Type>*& Left(Node<Type>*& node)  
    {     
        assert(NULL != node);  
        return node->left;  
    }  
  
    //判断指定节点是否有右儿子  
    //static bool HasRight(Node<Type>* node)  
    //{  
    //  if(!node) return false;//空节点没有右儿子  
    //  return node->right;  
    //}  
  
    //获取指定节点的右儿子  
    static Node<Type>*& Right(Node<Type>*& node)  
    {  
        assert(NULL != node);  
        return node->right;  
    }  
      
    //获取指定节点的距离  
    static int Dist(Node<Type>* node)  
    {  
        if(NULL == node) return -1;//规定空节点的距离为-1  
        return node->dist;  
    }  
  
    //交换left指针与right指针  
    static void Swap(Node<Type>*& left,Node<Type>*& right)  
    {    
        Node<Type> *temp = left;     
        left = right;  
        right = temp;  
    }  
  
    //合并两棵左偏树（返回合并之后左偏树的根结点）
    static Node<Type>*& Merge(Node<Type>*& t1, Node<Type>*& t2)  
    {  
  
        if(NULL == t1)  
            return t2;  
        else if(NULL == t2)  
            return t1;  
  
        //确定t1，t2两个节点谁作为合并后的根节点，以满足左偏树的堆的性质（A[parent(i)]>=A[i] or A[parent(i)]<=A[i]）。  
        //但这里的“堆”未必是完全二叉树  
        //注意是个递归过程  
        if( Compare()(t2->data, t1->data) ) //Compare是比较规则，它确定堆是大堆还是小堆  
            Swap(t1,t2);                  
        Right(t1) = Merge(Right(t1), t2);  
  
        //为左右儿子排序，以满足左偏树的左偏性质，即节点的左子节点的距离不小于右子节点的距离  
        if( Dist(Right(t1)) > Dist(Left(t1)) )  
            Swap( Left(t1),Right(t1) );  
        //调整距离  
        if( NULL == Right(t1) )  
            t1->dist = 0;  
        else t1->dist = Dist(Right(t1)) + 1;  
  
  
        return t1;  
    }  
  
    //输出以指定节点为根的左偏树(递归方式的先序遍历)  
    void Print(Node<Type> *node)  
    {  
        if(NULL != node)  
        {  
  
            cout << node->data <<endl;  
              
            cout<<"节点"<<node->data<<"的左儿子为：";  
            if(NULL != node->left)         
                Print(node->left);  
            else  
                cout<<"NULL";  
            cout<<endl;  
  
            cout<<"节点"<<node->data<<"的右儿子为：";  
            if(NULL != node->right)                
                Print(node->right);  
            else  
                cout<<"NULL";  
            cout<<endl;  
        }  
    }  
public:  
  
    LeftistTree(): root(NULL)  
    {  
    }  
  
    ~LeftistTree()  
    {  
        Delete(root);  
    }  
  
    //向左偏树插入元素  
    void Insert(Type val)  
    {  
        Node<Type> *newNode;  
        newNode = new Node<Type>(val);  
        root = Merge(root, newNode);  
    }  
  
    //删除左偏树的根结点  
    void DeleteRoot()  
    {  
        if(NULL == root)  
        {  
            cout << "警告：左偏树为空" << endl;  
            return;  
        }  
        Node<Type> *p = root;  
        root = Merge(p->left, p->right);  
        delete p;  
    }  
  
    //合并两棵左偏树  
    //合并后t2就为空树了。  
    void Merge(LeftistTree<Type,Compare>& t2)  
    {  
        root = Merge(root, t2.root);  
        t2.root = NULL;  
    }  
  
    //获取根结点的值（相当于优先级）  
    Type Root()  
    {  
        if(NULL == root)  
        {  
            cout << "左偏树为空" << endl;  
            return NULL;  
        }  
        return root->data;  
    }  
  
    //输出整棵左偏树中的元素  
    void Print()  
    {  
        Print(root);  
        cout << endl;  
    }  
};  
  
  
int main()  
{  
    //建立小堆的左偏树  
    {  
    LeftistTree<double> tree;  
    tree.Insert(5);  
    tree.Insert(4);  
    tree.Insert(3);  
    tree.Insert(8);  
    tree.Insert(1);  
    tree.Insert(2);  
    tree.Insert(9);  
    tree.Insert(7);  
    tree.Insert(6);  
    tree.Insert(0);  
    cout<< "第一棵左偏树：" << endl;
    tree.Print();  
  
    tree.DeleteRoot();  
    tree.DeleteRoot();  
    cout<< "新的左偏树：" << endl;  
    tree.Print();  
  
    cout <<"新的左偏树的根节点："<< tree.Root() << endl;  
  
    LeftistTree<double> tr;  
    tr.Insert(4.3);  
    tr.Insert(9.6);  
    tr.Insert(5.5);  
    tr.Insert(6.6);  
  
    cout<<"第二棵左偏树:"<<endl;
    tr.Print();  
    cout <<"第二棵左偏树的根节点："<<tr.Root() << endl;
  
    tree.Merge(tr);  
    cout <<"合并后的左偏树为：" << endl;  
    tree.Print();  
    }  
  
    cout<<"-----------------------------------"<<endl;  
  
    //建立大堆的左偏树  
    {  
    LeftistTree< double, greater<double> > tree2;  
    tree2.Insert(5);  
    tree2.Insert(4);  
    tree2.Insert(3);  
    tree2.Insert(8);  
    tree2.Insert(1);  
    tree2.Insert(2);  
    tree2.Insert(9);  
    tree2.Insert(7);  
    tree2.Insert(6);  
    tree2.Insert(0);  
    cout<< "第一棵左偏树：" << endl;
    tree2.Print();  
  
    tree2.DeleteRoot();  
    tree2.DeleteRoot();  
    cout<< "新的左偏树：" << endl;  
    tree2.Print();  
  
    cout <<"新的左偏树的根节点："<< tree2.Root() << endl;  
  
    LeftistTree< double, greater<double> > tr2;  
    tr2.Insert(4.3);  
    tr2.Insert(9.6);  
    tr2.Insert(5.5);  
    tr2.Insert(6.6);  
  
    cout<<"第二棵左偏树:"<<endl;
    tr2.Print();  
    cout <<"第二棵左偏树的根节点："<<tr2.Root() << endl;
  
    tree2.Merge(tr2);  
    cout <<"合并后的左偏树为：" << endl;  
    tree2.Print();  
    }  
  
    return 0;  
}  