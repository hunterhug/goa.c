/*自底向上，迭代
  Array：待排序的数组首地址
  low：待排序的范围的下界
  high：待排序的范围的上界的后一个位置
  比如你要对数组Array[0]~Array[5]进行排序，那么low=0，high=6*/
void mergeSort(int* Array, int low, int high)
{
    int step = 1;
    while (step < high - low) {
        for (int i = low; i < high; i += step << 1) {
            int lo = i, hi = (i + (step << 1)) <= high ? (i + (step << 1)) : high; //定义二路归并的上界与下界
            int mid = i + step <= high ? (i + step) : high;
            merge(Array, lo, mid, hi);
        }

        //将i和i+step这两个有序序列进行合并
        //序列长度为step
        //当i以后的长度小于或者等于step时，退出
        step <<= 1;//在按某一步长归并序列之后，步长加倍
    }
}