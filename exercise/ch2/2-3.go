package main

var pc [256]byte

/*
 0000: 0
 0001: 1
 0010: 2
 0011: 3
 0100: 4
 0101: 5
 0110: 6
 0111: 7

 1000: 8
 1001: 9
 1010: 10
 1011: 11
 1100: 12
 1101: 13
 1110: 14
 1111: 15

	pc : from 0 ~ 255
	i/2: 0 ~ 127
	i&1: 0 or 1

	i == 0
		pc[0] = pc[0] + 0
		=> 0

	i == 1
		pc[1] = pc[0] + 1
		=> 1

	i == 2
		pc[2] = pc[1] + 0
		=> 1

	i == 3
		pc[3] = pc[1] + 1
		=> 2

	i == 4
		pc[4] = pc[2] + 0
		=> 1

	i == 5
		pc[5] = pc[2] + 1
		=> 2

	i == 6
		pc[6] = pc[3] + 0
		=> 2

	i == 7
		pc[7] = pc[3] + 1
		=> 3

	i == 8
		pc[8] = pc[4] + 0
		=> 1

	i == 10
		pc[10] = pc[5] + 0
		=> 2

	i == 12
		pc[12] = pc[6] + 0
		=> 2

	i == 14
		pc[14] = pc[7] + 0
		=> 3

	i == 16
		pc[16] = pc[8] + 0
		=> 1
*/

func init() {
	for i := range pc {
		// 2進法だと
		// 4 => 8 は 1bit だけ左シフトした数でビットの数は等しい
		// 4 => 9 は上記の左シフトの後に +1 したもの

		// 1 => 2: 1を 1bit だけ左シフト
		// 1 => 3: 1を 1bit だけ左シフト + 一番右のビットを立てる
		// 2 => 4: 2を 1bit だけ左シフト
		// 2 => 5: 2を 1bit だけ左シフト + 一番右のビットを立てる
		// 3 => 6: 3を 1bit だけ左シフト
		// 3 => 7: 3を 1bit だけ左シフト + 一番右のビットを立てる
		// ...
		// つまり求めたい数の BitCount =  1bit 右シフトした数の BitCount + (奇数なら1・偶数なら0)

		// pc[i/2]  : 求めたい数を 1bit だけ右シフトした数の BitCount (計算済み)
		// byte(i&1): 求めたい数が奇数なら +1 をして偶数なら何もしない
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCountByTableLookUp(x uint64) int {
	// 1byte (8bit) の BitCount のテーブルしか持っていないので
	// 64bit を8分割してそれぞれの BitCount を計算して足し合わせる
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// 余計なループを使うパターン
func PopCountByTableLookUp2(x uint64) int {
	n := 0
	for i := uint64(0); i < 8; i++ {
		n += int(pc[byte(x>>(i*8))])
	}
	return n
}

func PopCountByShift(x uint64) int {
	n := 0
	for i := uint64(0); i < 64; i++ {
		//if (x & (1 << i)) != 0 {
		if ((x >> i) & 1) != 0 {
			n++
		}
	}
	return n
}

func PopCountByClear(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}
