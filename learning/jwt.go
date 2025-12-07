package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySecretKey = []byte("my-super-secret-password-123456")

type MyCustomClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	// 嵌入标准声明 (Standard Claims)，帮我们自动处理 exp, iss, iat 等字段
	jwt.RegisteredClaims
}

// -------------------------------------------------------
// 核心函数 1: 生成 Token
// -------------------------------------------------------
func GenerateJWT() (string, error) {
	// 1. 准备数据 (Claims)
	claims := MyCustomClaims{
		UserID:   1001,
		Username: "GeminiUser",
		Role:     "Admin",
		RegisteredClaims: jwt.RegisteredClaims{
			// 设置过期时间：24小时后过期
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			// 设置签发者
			Issuer: "MyGoApp",
			// 设置签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	// 2. 创建 Token 对象
	// 指定签名算法：HS256 (HMAC-SHA256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 3. 签名 (Sign)
	// 使用我们的秘钥生成最终的字符串
	// 这一步就是生成 Header.Payload.Signature 的过程
	tokenStr, err := token.SignedString(mySecretKey)
	return tokenStr, err
}

// -------------------------------------------------------
// 核心函数 2: 解析与验证 Token
// -------------------------------------------------------

func PareseAndVerifyJWT(tokenStr string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return mySecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 3. 提取数据
	// 如果 token.Valid 为 true，说明签名一致且没过期
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}

}

func main() {
	fmt.Println("--- 第一步：制造 JWT (签发) ---")
	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Printf("生成失败: %v\n", err)
		return
	}
	fmt.Printf("生成的 Token (发给客户端):\n%s\n\n", tokenString)

	// ---------------------------------------------------------

	fmt.Println("--- 第二步：验证 JWT (解析) ---")
	// 模拟服务端收到 Token 后进行解析
	// 我们试着解析刚才生成的 tokenString
	claims, err := PareseAndVerifyJWT(tokenString)
	if err != nil {
		fmt.Printf("❌ 验证失败: %v\n", err)
	} else {
		fmt.Println("✅ 验证成功！数据如下：")
		fmt.Printf("ID: %d\n", claims.UserID)
		fmt.Printf("用户: %s\n", claims.Username)
		fmt.Printf("角色: %s\n", claims.Role)
		fmt.Printf("过期时间: %v\n", claims.ExpiresAt)
	}

	// ---------------------------------------------------------

	fmt.Println("\n--- 第三步：模拟黑客篡改 ---")
	// 假如黑客拿到了 Token，修改了最后一位字符，试图绕过验证
	fakeToken := tokenString[0:len(tokenString)-1] + "X"
	fmt.Println("黑客修改后的 Token:", fakeToken)

	_, err = PareseAndVerifyJWT(fakeToken)
	if err != nil {
		fmt.Printf("❌ 篡改检测成功 (验证应该失败): %v\n", err)
	}
}
