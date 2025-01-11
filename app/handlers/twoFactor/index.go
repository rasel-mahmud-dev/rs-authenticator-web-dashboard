package twoFactor

import (
	"encoding/base64"
	"fmt"
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
	"net/http"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
)

func Generate2FASecret(w http.ResponseWriter, r *http.Request) {

	authSession := (*r).Context().Value("authSession").(*models.AuthSession)

	if authSession == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Generate the TOTP secret
	secret, err := totp.Generate(totp.GenerateOpts{
		Issuer:      fmt.Sprintf("RsAuth (%s)", authSession.Email),
		AccountName: authSession.Email,
	})
	if err != nil {
		http.Error(w, "Failed to generate secret", http.StatusInternalServerError)
		return
	}

	// Create QR code
	qrCodeData, err := qrcode.Encode(secret.URL(), qrcode.Medium, 256)
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}

	// Optionally save the secret to a database (simulated here)
	//twoFASecret := TwoFASecret{
	//	UserID: user.ID,
	//	Secret: secret.Secret(),
	//}
	// Simulate saving the secret (replace with actual database logic)
	//fmt.Printf("Saving 2FA secret for user %s: %v\n", user.ID, twoFASecret)
	
	w.Header().Set("Content-Type", "application/json")
	response.Respond(w, statusCode.OK, "Success", map[string]interface{}{
		"message": "2FA secret generated successfully",
		"qrCode":  fmt.Sprintf("data:image/png;base64,%s", toBase64(qrCodeData)),
		"secret":  secret.Secret(),
	})
}

func toBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
