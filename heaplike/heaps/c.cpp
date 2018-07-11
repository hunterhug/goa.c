// 模拟最小堆
// 最小堆是二叉堆的一种，其特点是父节点的键值总是小于或者等于子节点。
// 实现细节：
//          push：向堆中插入数据时，首先在堆的末尾插入数据，然后不断向上提升，直到没有大小颠倒时。
//          pop：从堆中删除最小值时首先把最后一个值复制到根节点上，并且删除最后一个数值。然后不断向下交换
//               直到没有大小颠倒为止。在向下交换过程中，如果有两个子儿子都小于自己，就选择较小的
#include <iostream>
using namespace std;
const int MAX_N = 1005;
int heap[MAX_N], sz = 0;
void push(int x);
void display();
int pop();
int main()
{
    // 测试
    int x;
    int cmd;
    do
    {
        cout << "请输入命令:1.push\t2.pop\t3.display\t0.quit\n";
        cin >> cmd;
        switch(cmd)
        {
        case 1:
            cout << "Input X:";
            cin >> x;
            push(x);
            break;
        case 2:
            x = pop();
            cout << x << "已取出!\n";
            break;
        case 3:
            display();
            break;
        }
    }while(cmd);

    return 0;
}

void push(int x)
{
    // i是要插入节点的下标
    int i = sz++;
    while(i > 0)
    {
        // p为父亲节点的下标
        int p = (i - 1) / 2;
        // 如果父亲节点小于等于插入的值，则说明大小没有跌倒，可以退出
        if(heap[p] <= x)
            break;
        // 互换当前父亲节点与要插入的值
        heap[i] = heap[p];
        i = p;
    }

    heap[i] = x;
    cout << "数据插入成功!\n";
}

int pop()
{
    // 取出根节点
    int ret = heap[0];
    // 将最后一个节点的值提到根节点上
    int x = heap[--sz];
    int i = 0;
    while(i * 2 + 1 < sz)
    {
        // a，b为左右两个子节点的下标
        int a = 2 * i + 1, b = 2 * i + 2;
        // 去两个子节点中较小的值
        if(b < sz && heap[b] < heap[a])
            a = b;
        // 如果已经没有大小颠倒的话则退出循环
        if(heap[a] >= x)
            break;
        // 将父亲节点与子节点互换
        heap[i] = heap[a];
        i = a;
    }
    heap[i] = x;

    return ret;
}

void display()
{
    for(int i = 0; i < sz; i++)
        cout << heap[i] << "\t";
    cout << endl;
