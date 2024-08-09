# Guard
## 数据通信
软件和验证端可以选择HTTP、Socket进行链接, 数据校验端在构建时会生成私钥和公钥，客户端需要把公钥写死，通过各种保护方式保护私钥不被篡改。后端生成密钥，客户端需要通过http和socket的卡密校验，发送的卡密需使用PKCS1v15+Base64加密然后发送时需要把配置文件指定的截断符号`Socket_EndSymbol`加上。， 如果成功则返回的是服务端使用私钥签名的卡密+base64，这是为了防止被代理服务器绕过，如果失败则返回错误信息，为了方便验证错误，返回的第一个字符是错误码。


```bash
openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048 && openssl rsa -pubout -in private.pem -out public.pem
```
## 加解密算法
PKCS1v15 + base64

# Socket错误码
1 发送过来的数据解密失败，即私钥被篡改，或者当前端口正在被测试。
2 发送过来的数据有效期已结束，或者说是就不存在这卡密
# Http错误码
Http的没有错误码，而是直接返回中文的错误，如果成功则直接返回服务端使用私钥签名的卡密+base64。