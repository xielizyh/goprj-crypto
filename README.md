# goprj-crypto
基于go语言实现的常见密码算法运算命令行工具

## utf8转换

以在linux上进行举例。

1. utf8转换为hex

   ```bash
   ./crypto utf8 -f hex -m 'hello world'
   ```

2. hex转换为utf8

   ```bash
   ./crypto utf8 -f hex -r -m 68656c6c6f20776f726c64
   ```

3. utf8转换为base64

   ```bash
   ./crypto utf8 -f base64 -m 'hello world'
   ```

4. base64转换为utf8

   ```bash
   ./crypto utf8 -f base64 -r -m aGVsbG8gd29ybGQ=
   ```

   

## 哈希算法

以在linux上进行举例。

1. 计算终端输入消息摘要

   ```bash
   ./crypto hash -a sha256 -m 'hello world'
   ```

2. 计算文件的消息摘要

   ```bash
   ./crypto hash -a sha256 -f -f README.md
   ```

## 消息认证码算法

以在linux上进行举例。

1. 计算终端输入消息认证码

   ```bash
   ./crypto mac -a hmac-sha256 -k 000102030405060708090A0B0C0D0E0F -m 'hello world'
   ```

2. 计算文件的消息认证码

   ```bash
   ./crypto mac -a hmac-sha256 -k 000102030405060708090A0B0C0D0E0F -f README.md
   ```

## 对称加解密算法

以在linux上进行举例。

1. 加密消息

   ```bash
   ./crypto enc -a aes-128-ecb -p e -k 000102030405060708090A0B0C0D0E0F -m '0102030405'
   ```

2. 解密消息

   ```bash
   ./crypto enc -a aes-128-ecb -p d -k 000102030405060708090A0B0C0D0E0F -m '47b6b76d59a92e8d3ab9abd0e287571d'
   ```

3. 加密文件

   ```bash
   ./crypto enc -a aes-128-ecb -p e -k 000102030405060708090A0B0C0D0E0F -i hello.txt -o cipher.txt
   ```

4. 解密文件

   ```bash
   ./crypto enc -a aes-128-ecb -p d -k 000102030405060708090A0B0C0D0E0F -i cipher.txt -o plain.txt
   ```

## 总结

1. 选项参数带空格

   添加引号表示作为选项m的选项参数，否则只会截取到"hello"

   ```bash
   -m 'hello world'
   -m "hello world"
   ```


2. delve带参数调试

   [添加--](https://ask.csdn.net/questions/1031161)

   ```bash
   dlv debug . -- hash -a sha256 -m 'hello world'
   ```

   

