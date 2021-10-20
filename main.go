package main

/*
#cgo CFLAGS: -Inmap
#cgo LDFLAGS: -Llib -L/usr/local/Cellar/pcre/8.45/lib -lstdc++
#cgo LDFLAGS: -lblas -ldnet -llinear -llua -lnbase -lnetutil -lnmap -lnsock -lz -lpcre -lpcap

#include <stdio.h>
#include <string.h>
#include <nmap.h>

void split(char *src,const char *separator,char **dest,int *num) {
	char *pNext;
	int count = 0;
	if (src == NULL || strlen(src) == 0) // 如果传入的地址为空或长度为0，直接终止
		return;
	if (separator == NULL || strlen(separator) == 0) // 如未指定分割的字符串，直接终止
		return;
	pNext = (char *)strtok(src,separator); // 必须使用(char *)进行强制类型转换(虽然不写有的编译器中不会出现指针错误)
	while(pNext != NULL) {
		*dest++ = pNext;
		++count;
		pNext = (char *)strtok(NULL,separator);  // 必须使用(char *)进行强制类型转换
	}
	*num = count;
}

void run_nmap(char *buf) {
	int num = 0;
	char *argv[255];
	split(buf, " ", argv, &num);
	set_program_name(argv[0]);
	nmap_main(num, argv);
}
*/
import "C"
import (
	"unsafe"
)

func main() {
	arg := []byte("nmap 127.0.0.1")
	C.run_nmap((*C.char)(unsafe.Pointer(&arg[0])))
}
