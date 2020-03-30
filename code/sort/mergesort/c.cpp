 #include <iostream>
 #include <cstdlib>
 using namespace std;
 void print(int *arr, int start, int end)
 {
	for (int i = start; i <= end; ++i)
		cout << arr[i] << ' ';
	cout << endl;
 }
 void randData(int *arr, int start, int end)
 {
	for (int i = start; i <= end; ++i)
		arr[i] = rand() % 20;
	print(arr, start, end);
 }

void merge(int *arr, int start, int mid, int end)
{
	int i, j, k, key;
	i = start;
	j = mid;
	while (i < j && j <= end)  //当i等于j或者j到达末尾时终止
	{
		if (arr[i] > arr[j])
		{
			k = j;
			key = arr[j];
			while (k > i && arr[k - 1] > key)
			{
				arr[k] = arr[k - 1];
				--k;

			}
			arr[k] = key;
			++j;
		}
		++i;
	}
}
void mergeSort(int *arr, int start, int end)
 {
	if(start < end)
	{
		int mid = (end + start) / 2;
		mergeSort(arr, start, mid);
		mergeSort(arr, mid + 1, end);
		merge(arr, start, mid + 1,end);
		print(arr, start, end);
	}
 }
 /*11 4 2 13 12 2 1 16 18 15*/
 int main()
 {
	bool bIsContinue = true;
	char ch = 'n';
	const int Len = 10;
	int arr[Len];
	print(arr, 0, Len - 1);

	while (true == bIsContinue)
	{
		randData(arr, 0, Len - 1);
		mergeSort(arr, 0, Len - 1);
		cout << "the new array: ";
		print(arr, 0, Len - 1);
		cout << "please input yes or no" << endl;
		cin >> ch;
		if (ch == 'y' || ch == 'Y')
			bIsContinue = true;
		else
			bIsContinue = false;
	}
    return 0;
 }