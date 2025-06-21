package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/workloom/auth-services/internal/config"
	"github.com/workloom/auth-services/internal/repository"
	"github.com/workloom/auth-services/internal/utils"
	"github.com/workloom/shared/db"
	"github.com/workloom/shared/models"
	"golang.org/x/crypto/bcrypt"
)

func GoogleLogin(c *gin.Context) {
    url := config.GoogleOauthConfig.AuthCodeURL("randomstate")
    c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(c *gin.Context) {
    code := c.Query("code")
    db.Init()

    token, err := config.GoogleOauthConfig.Exchange(context.Background(), code)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Token exchange failed"})
        return
    }

    fmt.Println("TOKEN: ====>>>>> ", token)

    client := config.GoogleOauthConfig.Client(context.Background(), token)
    resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

    fmt.Println("RESP: ====>>>>> ", resp)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user info"})
        return
    }
    defer resp.Body.Close()

	fmt.Println(resp)

    type GoogleUserInfo struct {
	    ID            string `json:"id"`
	    Email         string `json:"email"`
	    VerifiedEmail bool   `json:"verified_email"`
	    GivenName     string `json:"given_name"`
	    FamilyName    string `json:"family_name"`
	    Name          string `json:"name"`
	    Picture       string `json:"picture"`
    }

    var userInfo GoogleUserInfo

    if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user info response"})
        return
    }

    fmt.Println("USER INFO: ====>>>>> ", userInfo)
    fmt.Println("User access token: ", token.AccessToken)
    fmt.Println("User refresh token: ", token.RefreshToken)

    user := models.User{
        Provider:    models.ProviderGoogle,
        ProviderID:  userInfo.ID,
        Name:   userInfo.GivenName + " " + userInfo.FamilyName,
        Email:       userInfo.Email,
        AvatarURL:   userInfo.Picture,
        AccessToken: token.AccessToken,
        RefreshToken: token.RefreshToken,
        LastLogin:   time.Now(),
    }


    if err := db.DB.FirstOrCreate(&user, models.User{
        Provider:   models.ProviderGoogle,
        ProviderID: user.ProviderID,
    }).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
        return
    }

    c.JSON(http.StatusOK, user)
}

func Register(c *gin.Context) {
    type RegisterRequest struct {
        Email           string  `json:"email" binding:"required,email"`
        Password        string  `json:"password" binding:"required,min=6"`
        Name            string  `json:"name" binding:"required"`
        ConfirmPassword string  `json:"confirmPassword" binding:"required"`
    }

    var req RegisterRequest

    body, err := io.ReadAll(c.Request.Body)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
        return
    }
    
    // UnPack the Json body into a User struct
    if err := json.Unmarshal(body, &req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON structure"})
        return
    }
    
    if req.Password != req.ConfirmPassword {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
        return
    }

    user := models.User{
        Name:         req.Name,
        Email:        req.Email,
        PasswordHash: string(hashedPassword),
        Provider:     models.ProviderLocal,
        LastLogin:    time.Now(),
    }
    
    db.Init()
    existingUser, err := repositories.NewUserRepository(db.DB).FindUserByEmail(req.Email)
    if err == nil && existingUser != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
        return
    }

    if err := repositories.NewUserRepository(db.DB).CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    response := map[string]string{"message": "Registration successful"}
    jsonResp, _ := json.Marshal(response)
    c.Data(http.StatusCreated, "application/json", jsonResp)
}

func Login(c *gin.Context) {
    type LoginRequest struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
    }

    body, err := io.ReadAll(c.Request.Body)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
    }
    
    var req LoginRequest
    if err := json.Unmarshal(body, &req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON structure"})
    }

    db.Init()
    user, err := repositories.NewUserRepository(db.DB).FindUserByEmail(req.Email)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
        return
    }

    // Update last login time
    user.LastLogin = time.Now()
    if err := repositories.NewUserRepository(db.DB).SaveUser(user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update last login time"})
        return
    }

    // Generate a session token or JWT here if needed
    // For simplicity, we are just returning the user data
    token, err := utils.GenerateJWTFromPayload(map[string]interface{}{
        "id":    user.ID,
        "email": user.Email,
        "name":  user.Name,
        "provider": user.Provider,
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    // Here you would typically check the user's credentials against the database
    // For simplicity, we are just returning a success message
    c.SetCookie(
        "token",       // name
        token,         // value
        3600*24,       // maxAge in seconds (1 day)
        "/",           // path
        "",            // domain (use your domain in production)
        false,         // secure
        true,          // httpOnly
    )

    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func LogOut(c *gin.Context) {
    c.SetCookie(
        "token",     // name
        "",          // value
        -1,          // maxAge (expires immediately)
        "/",         // path
        "",          // domain (empty = current domain)
        false,       // secure (true if using HTTPS)
        true,        // httpOnly
    )

    c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
