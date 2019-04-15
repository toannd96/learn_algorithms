# Thuật toán Binary Search

## Chỉ áp dụng được với dãy đã được sắp xếp

### Triển khai

1. Xác định vị trí chính giữa trong dãy và số ở vị trí đó

```
mediumIndex = (firstIndex + lastIndex)/2
mediumNumber = list[mediumIndex]
```

2. So sánh mediumNumber với số cần tìm

- Nếu số cần tìm = mediumNumber: return mediumIndex

- Nếu số cần tìm > mediumNumber --> Cập nhật firstIndex

```
firstIndex = mediumIndex + 1
```
(Cộng thêm 1 là bởi số ở vị trí mediumIndex không phải là số cần tìm --> không quan tâm đến mediumIndex nữa)

- Nếu số cần tìm < mediumNumber --> Cập nhật lastIndex

```
lastIndex = mediumIndex - 1
```
(Trừ đi 1 là bởi số ở vị trí mediumIndex không phải là số cần tìm --> không quan tâm đến mediumIndex nữa)

3. Lặp lại bước 1 và 2 đối với firstIndex và lastIndex mới. Quá trình lặp diễn ra với điều kiện firstIndex <= lastIndex