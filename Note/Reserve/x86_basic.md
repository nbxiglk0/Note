# 函数传参方式
## 栈
1.参数进栈的顺序  
2.栈平衡
### 调用约定
类型 | _cdecl(C) |  pascal | stdcall | Fastcall| 
-|-|-|-|-|
参数传递顺序 | 右->左 | 左->右 | 右->左 | 寄存器,栈
平衡栈 | 调用者 | 子程序 | 子程序 | 子程序
VARARG | 1 | 0 | 1 |
VARARG:表示参数个数不确定,stdcall使用VAR则是调用者平衡,否则子程序  
C/C++: _cdecl
WIN32 API: stdcall  
ret X == add esp X ,X=参数个数x4h  
  
**enter xxx,0**:(自动为被调用程序建立堆栈,arv1:保留的堆栈空间，arv2:嵌套层次0)  
```
push ebp  
move ebp esp
sub esp xxx
```
**leave**:(释放堆栈)  
```
add esp xxx
pop ebp
```
### 优化编译模式:
直接使用esp进行寻址,节省ebp寄存器,加快速率  
```
mov eax dword ptr [esp+04]
mov ecx dword ptr [esp+08]
```
## 寄存器
遵循Fastcall规范  
MV C++:左边不大于4字节的参数放在ecx,edx寄存器，寄存器用完后使用栈，右->左,子程序平衡,浮点值,远指针,__int64用栈传递  
Borland c++:左边三个不大于4字节的参数使用eax,edx,ecx寄存器,用完后使用PASCAL方式入栈  
WATCOM C:为每一个参数分配一个寄存器,寄存器使用完使用栈,可以自行指定任意一个寄存器传递参数  
thiscall:非静态类成员函数的调用方式,使用exc寄存器传递this指针  
```
...
lea ecx [edp-04]
...
```
### 名称修饰约定
pascal:输出函数不能有任何修饰,且全为大写  
C:  
* stdcall:  函数名称前_,@后跟参数字节数 _functionname@number  
* __cdecl:  函数名称前_, _functionname  
* Fastcall:  函数名称前@,后@跟参数字节数  

C++：  
* stdcall : ?开始,函数后面跟@@YG和参数表,参数表依次为返回值类型,参数数据类型.后面以@Z表明程序结束,没有参数则使用Z结束 ?functionname@@YG****(@)Z  
* __cdecl : 与stdcall类似.@@YG->@@YA
* Fastcall : 与stdcall类似.@@YG->@@YI
## 函数返回
### return
函数返回值一般放在eax寄存器,超过容量则放在edx寄存器  
#### 按传引用方式返回值 
##### 传值
建立一份副本,修改参数值不会影响到原值
##### 传引用
直接传递变量的地址,在函数中使用间接引用运算符修改内存单元的值
# 数据结构
## 局部变量
通常使用栈和寄存器
### 栈
```
sub esp n   add esp -n   push reg(寄存器)
...
add esp n   sub esp -n   pop reg
```
### 寄存器
除了esp和ebp之外的6个通用寄存器
## 全局变量
通常位于.data的固定地址处  
`mov dword ptr[z] 7`  
硬编码直接寻址
## 数组
访问方式:基址+变址
## 虚函数
通过指向虚函数表的指针间接地调用,利用两次简介寻址得到虚函数的地址
## 控制语句
### if-then-else
整数用cmp比较,浮点值用fcom、fcomp等指令  
test(or)指令`test eax eax(逻辑与)`如果eax为0则结果为0,ZF->1 否则ZF->0
### switch-case
```
cpm ... ...
je ...
优化:
cpm->dec
```
如果为算术级:
`jmp dword ptr [4*eax+基址]` eax则为算术级索引
### 转移指令机器码计算
短转移:无条件转移机器码为2字节(-128~127)  
长转移:无条件转移为5字节(1字节转移类型,4字节偏移量),条件转移为6字节(2个字节表示转移类型,4个字节表示转移偏移量)  
call:直接call 地址类似与长转移还有call 寄存器，栈等
#### 短转移
无条件转移机器码为EBXX,EB00h~EB7Fh为向后转移,EB80h~EBFFh向前转移  
位移量 = 目的地址 - 起始地址 - 跳转指令本身长度(2,5,6)  
转移指令机器码 = 转移类别机器码 + 位移量
```
00401000 jmp 00401005
...
00401005 test eax,eax
位移量=00401005-00401000-00000002=00000003
机器码=EB+03->EB03
```
#### 长转移
长转移机器码为E9
```
00401000 jmp 00402398
...
00402398 test eax,eax
位移量=00402398h-00401000h-5h=00001393h,低位字节存入内存低位->从上到下93 13 00 00
机器码=E9 93 13 00 00
```
向前转移
```
00401000 test eax,eax
...
00402398 jmp 00401000
位移量=00401000h-00402398-5h=FFFFEC63h
机器码=E9 63 EC FF FF
```
#### 条件设置指令
