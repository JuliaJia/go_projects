package string_func

import (
	"fmt"
	"unsafe"
)

type mapextra struct {
}

type hmap struct {
	count      int //map中的元素个数，必须放在struct的第一个位置，因为内置的len函数会从这里读取
	flags      uint8
	B          uint8          //说明包含2^B个bucket
	noverflow  uint16         //溢出的bucket的个数
	hash0      uint32         //hash种子
	buckets    unsafe.Pointer //buckets的数组指针
	oldbuckets unsafe.Pointer //结构扩容的时候用于赋值bucketss数组
	nevacuate  uintptr        //搬迁进度（已经搬迁的buckets数量）
	extra      *mapextra
}

const bucketCnt = 10

type bmap struct {
	tophash [bucketCnt]uint8 //tophash是hash值的高8位
}

func GetStringBySlice(s []byte) string {
	return string(s)
}

func GetSliceByString(s string) []byte {
	return []byte(s)
}

func WrongAddSlice(s []byte) {
	s = append(s, 'a', 'a', 'a')
	fmt.Println(s)
}

func RightAddSlice(s *([]byte)) {
	*s = append(*s, 'a', 'a', 'a')
	fmt.Println(*s)
}
