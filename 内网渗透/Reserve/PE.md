<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
- [PE](#pe)

- [PE](#pe)
  - [Dos头](#dos%E5%A4%B4)
    - [IMAGE_DOS_HEADER结构体](#image_dos_header%E7%BB%93%E6%9E%84%E4%BD%93)
  - [Dos存根](#dos%E5%AD%98%E6%A0%B9)
  - [NT头](#nt%E5%A4%B4)
    - [IMAGE_NT_HEADERS](#image_nt_headers)
      - [文件头(IMAGE_FILE_HEADER)](#%E6%96%87%E4%BB%B6%E5%A4%B4image_file_header)
        - [Machine](#machine)
        - [NumberOfSections](#numberofsections)
        - [SizeOfOptionalHeader](#sizeofoptionalheader)
        - [Characteristics](#characteristics)
      - [可选头(IMAGE_OPTIONAL_HEADER)](#%E5%8F%AF%E9%80%89%E5%A4%B4image_optional_header)
        - [Magic](#magic)
        - [AdderssOfEntryPoint](#adderssofentrypoint)
        - [ImageBase](#imagebase)
        - [SectionAlignment,FileAlignment](#sectionalignmentfilealignment)
        - [SizeOfImage](#sizeofimage)
        - [SizeOfHeaders](#sizeofheaders)
        - [Subsystem](#subsystem)
        - [NumberOfRvaAndSizes](#numberofrvaandsizes)
        - [DataDirectory](#datadirectory)
  - [节区头](#%E8%8A%82%E5%8C%BA%E5%A4%B4)
    - [IMAGE_SECTION_HEADER](#image_section_header)
  - [RVA TO RAW](#rva-to-raw)
  - [IAT](#iat)
    - [DLL](#dll)
    - [IMAGE_IMPORT_DESCRIPTOR](#image_import_descriptor)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->


# PE
可执行系列: EXE SCR  
库系列: DLL OCX CPL DRV  
驱动程序系列: SYS VXD  
对象文件系列: OBJ  

PE头:DOS头--节区头,剩下为PE体  
文件:代码(.text),数据(.data),资源(.rsrc)  
RVA(相对虚拟地址)+ImageBase=VA(绝对虚拟地址)
## Dos头
### IMAGE_DOS_HEADER结构体
大小:64字节  
```
WORD e_magic;
...
...
...
LONG e_lfanew;
```
**e_magic**：DOS签名->4D5A->'MZ'  
**e_ifanew**：NT头偏移
## Dos存根
可选，代码与数据混合而成，大小不固定
## NT头
### IMAGE_NT_HEADERS
大小:F8
```
DWORD Signatrue;//PE签名 5045 0000 ("PE"00) 
IMAGE_FILE_HEADER FileHeader;
IMAGE_OPTIONAL_HEADER32 OptionlHeader;
```
#### 文件头(IMAGE_FILE_HEADER)
```C++
WORD Machine;
WORD NumberOfSections;
DWORD TimeDateStamp;
DWORD PointerToSymbolTable;
DWORD NumberOfSymbols;
WORD SizeOfOptionalHeader;
WORD Characteristics;
```
##### Machine  
每个CPU拥有的唯一机器码,intel x86为14C，定义为winnt.h文件中
##### NumberOfSections
文件存在的节区数量,值>0
##### SizeOfOptionalHeader
用来指出可选头(IMAGE_OPTIONAL_HEADER32)的长度
##### Characteristics
标识文件的属性
```
...
#define IMAGE_FILE_EXECUTABLE_IMAGE 0x0002//可执行
...
#define IMAGE_FILE_DLL 0x2000//DLL文件
```
#### 可选头(IMAGE_OPTIONAL_HEADER)
##### Magic
32结构体时为10B，64结构体时为64
##### AdderssOfEntryPoint
EP的RVA值,该程序最先执行的代码起始地址
##### ImageBase
文件装入内存的优先装入地址,exe->0040000 dll->10000000,EIP的值=ImageBase+AdderssOfEntryPoint
##### SectionAlignment,FileAlignment
SectionAlignment->节区在磁盘文件中的最小单位  
FileAlignment->节区在内存中的最小单位
##### SizeOfImage
PE image在虚拟内存中所占空间大小
##### SizeOfHeaders
整个PE头的大小，为FileAlignment的整数倍
##### Subsystem
区分系统驱动文件与普通可执行文件
* 1->Driver文件 系统驱动(sys)
* 2->GUI文件 窗口程序
* 3->CUI 控制台程序
##### NumberOfRvaAndSizes
指定DataDirectory数组的个数,PE通过查看该值来识别数组大小
##### DataDirectory
由IMAGE_DATA_DIRECTORY结构体组成的数组  
**IMAGE_DATA_DIRECTORY**
```
DWORD VirtualAddress;
DWORD Size;
```
## 节区头
code、data、resource,由IMAGE_SECTION_HEADER结构体组成的数据，一个结构体对应一个节区    
类别 | 访问权限 
- | -
code | 执行、读取
data | 非执行、读写
resource | 非执行、读取  
### IMAGE_SECTION_HEADER
主要成员 | 含义
- | -
VirtualSize | 内存中节区所占大小
VirtualAddress | 内存中节区的起始地址
SizeOfRawData | 磁盘文件中节去所占大小
PointerToRawData | 磁盘文件中节区起始位置
Characteristics | 节区属性
## RVA TO RAW
RAW - PointerToRawData = RVA - VirtualAddress  
RAW = RVA - VirtualAddress + PointerToRawData  
文件偏移=相对虚拟地址+磁盘节区起始位置-内存的偏移地址
## IAT
导入地址表
### DLL
* 显示链接:程序使用时才加载,使用完毕后释放内存
* 隐式链接:程序开始时就一并加载,程序结束后释放内存
### IMAGE_IMPORT_DESCRIPTOR
位置信息存储于可选头的最后一位成员DataDirectory[1].VirtualAddress  
主要成员 | 含义
- | -|
OriginalFirstThunk | INT的地址(RVA)
Name | 库名称字符串的地址(RVA)
FirstThunk | IAT的地址(RVA)
* INT:IAT的副本,不可改写,保存了间接指向函数地址的指针,当dll文件被修改后,可以根据INT重建来IAT
* IAT:根据INT中的函数地址指针,将函数的实际地址取出构成。  
**IAT的构建过程：**  
1.根据IID的Name成员找到库函数名称导入  
2.读取IID的OriginalFirstThunk值获取INT的相对虚拟地址  
3.从INT数组中一一获取IMAGE_IMPORT_BY_NAME的相对虚拟地址(INT数组每个元素都指向一个IMAGE_IMPORT_BY_NAME结构体的指针)  
4.从IMAGE_BY_NAME中的Hint成员或者Name成员来获取对应函数的起始地址  
5.读取IID的FirstThunk获取IAT地址的相对虚拟地址  
6.将4中获得的函数起始地址存入IAT中  
7.重复3-6直到遇到INT中的NULL(INT的所有函数地址都被写入IAT中)  


