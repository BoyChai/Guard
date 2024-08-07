# Guard
## 数据通信
软件和验证端可以选择HTTP、Socket进行链接, 数据校验端在构建时会生成私钥和公钥，客户端需要把公钥写死，通过各种保护方式保护私钥不被篡改。后端这里仅提供公钥，客户端需要通过公钥验证服务端身份。http和socket的卡密校验如果成功则返回的是卡密的密文，如果失败则返回错误信息。

## 公私钥生成
```bash
openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048 && openssl rsa -pubout -in private.pem -out public.pem
```
## 加解密算法
PKCS1v15 + base64