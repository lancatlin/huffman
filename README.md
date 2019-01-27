# 霍夫曼編碼器 Huffman Coding Tool
> 利用霍夫曼編碼樹演算法來壓縮資料的 CLI 工具

## 操作
壓縮： `./huffman -c file`，預設輸出檔名 `file.huf`  
壓縮並指定檔名： `./huffman -c -f file.huf file`  
解壓縮： `./huffman -x file.huf` 預設輸出檔名 `file.out`  
解壓縮並指定檔名： `./huffman -x -f file.out file.huf`

# 程式碼導航
| Feature | Source |
| ------- | ------ |
| 演算法核心 | huffman.go |
| 排序演算法 | sort.go |
| Struct, Interface 定義 | type.go |
| 計算字元數量 | count.go |
| 編碼器 | coding.go |
| 資料讀寫、編碼 | data.go |
| 主程式 | main.go|
