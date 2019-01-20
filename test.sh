#! /bin/sh
# 測試 Huffman 主程式
./huffman $1
./huffman -x -f $1.out $1.huf
du  $1 $1.huf $1.out
