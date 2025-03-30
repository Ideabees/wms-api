package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"wms-app/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// Secret key for JWT signing
var jwtSecret = []byte("your_secret_key")

// OTP Store with rate limiting
var otpStore = struct {
	sync.Mutex
	data map[string]string
}{data: make(map[string]string)}

// Rate limiters per user (bucket strategy)
var rateLimiters = struct {
	sync.Mutex
	limiters map[string]*rate.Limiter
}{limiters: make(map[string]*rate.Limiter)}

func generateOTP() string {
	//rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000)) // 6-digit OTP
}

type MaskedMobileNumber struct {
	MaskedMobile string `mask:"mobileNumber"`
}

// Rate limiter setup (1 request per 10 seconds)
func getRateLimiter(userID string) *rate.Limiter {
	rateLimiters.Lock()
	defer rateLimiters.Unlock()

	if limiter, exists := rateLimiters.limiters[userID]; exists {
		return limiter
	}

	limiter := rate.NewLimiter(1, 1) // 1 request per 10 seconds
	rateLimiters.limiters[userID] = limiter
	return limiter
}

func generateOTPHandler(c *gin.Context) {
	//userID := c.Query("user")
	var req struct {
		MobileNumber string `json:"mobile_number"`
	}

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	fmt.Println(req.MobileNumber)
	if req.MobileNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "mobile number is required"})
		return
	}

	limiter := getRateLimiter(req.MobileNumber)
	if !limiter.Allow() {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests. Try again later."})
		return
	}

	otp := generateOTP()
	otpStore.Lock()
	otpStore.data[req.MobileNumber] = otp
	otpStore.Unlock()

	// mask mobile number
	/*maskedMobile := &MaskedMobileNumber{
	        MaskedMobile: req.MobileNumber,
	    }

	    m := masker.NewMaskerMarshaler()

	    maskedData, err := m.Struct(maskedMobile)
		fmt.Println(err)
		if err != nil {
			c.JSON(http.StatusFailedDependency, gin.H{
				"status":       "success",
				"mobile_numer": "",
				"msg":          "",
				"error_code":   ""})
		}
	    fmt.Println(maskedData)
	*/
	// In real applications, send OTP via SMS/Email
	c.JSON(http.StatusOK, gin.H{
		"status":       "success",
		"mobile_numer": req.MobileNumber,
		"msg":          otp,
		"error_code":   ""})
}

func verifyOTPHandler(c *gin.Context) {
	var req struct {
		MobileNumber string `json:"mobile_number"`
		OTP          string `json:"otp"`
	}
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	limiter := getRateLimiter(req.MobileNumber)
	if !limiter.Allow() {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many attempts. Try again later."})
		return
	}

	/*otpStore.Lock()
	storedOTP, exists := otpStore.data[req.MobileNumber]
	otpStore.Unlock()

	if !exists || storedOTP != req.OTP {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid OTP"})
		return
	}*/

	// OTP Verified -> Generate JWT token
	token, err := createJWT(req.MobileNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Render dashboard with JWT token
	c.JSON(http.StatusOK, gin.H{
		"access_token":   token,
		"otp_expired_at": "",
		"error_code":     "",
		"msg":            "",
		"mobile_number":  req.MobileNumber,
	})
}

// Create JWT Token
func createJWT(mobileNumber string) (string, error) {
	claims := jwt.MapClaims{
		"user": mobileNumber,
		"exp":  time.Now().Add(time.Hour * 1).Unix(), // Token valid for 1 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Middleware to protect routes
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims["user"])
		c.Next()
	}
}

func main() {
	// create db connection
	_, err := config.GetDBConnection()
	if err != nil {
		fmt.Println("DB not initialized")
	}
	fmt.Println("Connected to DB")
	r := gin.Default()
	v1 := r.Group("/api/wms/v1")
	{
		v1.POST("/generate-otp", generateOTPHandler)
		v1.POST("/verify-otp", verifyOTPHandler)

	}
	r.Run(":8080") // Start server
}
