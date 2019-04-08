# Thuật toán sắp xếp Bubble Sort. Giả định sắp xếp từ bé --> lớn

## Duyệt qua các phần tử trong mảng. So sánh từng phần tử với phần tử liền kề ngay sau. Nếu lớn hơn --> đổi chỗ. Lặp lại cho đến khi mảng được sắp xếp từ bé --> lớn

### Triển khai với Golang

```go
func bubbleSort(list []int) []int {
    var firstNumber, secondNumber int
    checkResult := 0
    for {
        for i := 0; i < len(list) - 1; i++ {
            firstNumber = list[i]
            secondNumber = list[i+1]
            if firstNumber > secondNumber {
                list[i] = secondNumber
                list[i+1] = firstNumber
                checkResult++
            }
        }

        if checkResult == 0 {
            return list;
        }

        checkResult = 0
    }
}
```

1. 2 biến firstNumber và secondNumber dùng để lưu giá trị lấy ra từ phần tử i và i+1 trong mảng

2. Biến checkResult để lưu kết quả sắp xếp sau khi duyệt qua tất cả các phần tử trong mảng:

3. Chạy vòng lặp for. Bên trong vòng lặp for tạo một vòng lặp for khác, vòng lặp for này chạy qua từng phần tử trong mảng và thực hiện so sánh giá trị phần tử thứ i và phần tử thứ i+1:

- Giá trị phần tử thứ i được lưu ở biến firstNumber

- Giá trị phần tử thứ i+1 được lưu ở biến secondNumber

- Nếu firstNumber > secondNumber --> đổi chỗ: gán secondNumber cho biến i và firstNumber cho biến i+1, đồng thời tăng giá trị biến checkResult thêm 1 đơn vị

- Nếu firstNumber < secondNumber --> không làm gì cả

4. Sau khi duyệt hết phần tử trong mảng, tiến hành kiếm tra giá trị của biến checkResult:

- Nếu sau khi duyệt qua hết các phần tử trong mảng mà checkResult vẫn bằng 0 --> chứng tỏ mảng đã được sắp xếp từ bé đến lớn --> return ngay, thoát khỏi vòng lặp

- Còn nếu checkResult <> 0 --> reset lại giá trị của checkResult về0 và tiếp tục chạy. Nếu không reset giá trị checkResult về 0 thì biến checkResult sẽ không bao giờ = 0 --> vòng lặp for vô tận

### Một cách triển khai khác

```go
func bubbleSort_2(list []int) []int {
    for i := 0; i < len(list); i++ {
        if !compare(list, i) {
            return list
        }
    }
    return list
}

func compare(list []int, index int) bool {
    firstIndex := 0
    secondIndex := 1
    numberOfItems := len(list)

    var firstNumber, secondNumber int
    isSwap := false

    for secondIndex < (numberOfItems - index) {
        firstNumber = list[firstIndex]
        secondNumber = list[secondIndex]

        if firstNumber > secondNumber {
            list[firstIndex] = secondNumber
            list[secondIndex] = firstNumber
            isSwap = true
        }

        firstIndex++
        secondIndex++
    }

    return isSwap
}
```