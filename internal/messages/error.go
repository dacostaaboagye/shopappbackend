package messages

import (
	"fmt"
	"net/http"

	"github.com/Aboagye-Dacosta/shopBackend/internal/codes"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/entities"
)

type Message struct {
	UserMessage string
	DevMessage  string
}

var errorRegistry = map[string]map[int]Message{
	entities.USER: {
		http.StatusNotFound: {
			UserMessage: "User not found.",
			DevMessage:  "User entity lookup failed in database.",
		},
		http.StatusConflict: {
			UserMessage: "User already exists.",
			DevMessage:  "Duplicate user email constraint.",
		},
	},
	entities.PRODUCT: {
		http.StatusNotFound: {
			UserMessage: "Product not found.",
			DevMessage:  "Product ID not found in DB.",
		},
		http.StatusConflict: {
			UserMessage: "Product already exists.",
			DevMessage:  "Duplicate product identifier (name/SKU).",
		},
	},
	entities.ORDER: {
		http.StatusNotFound: {
			UserMessage: "Order not found.",
			DevMessage:  "Order ID not found in DB.",
		},
	},
	entities.PAYMENT: {
		http.StatusNotFound: {
			UserMessage: "Payment not found.",
			DevMessage:  "Payment record not found in DB.",
		},
		http.StatusBadRequest: {
			UserMessage: "Payment request is invalid.",
			DevMessage:  "Invalid payment details or state.",
		},
	},
	entities.AUTHORIZATION: {
		codes.INVALID_TOKEN: {
			UserMessage: "Invalid or malformed token.",
			DevMessage:  "JWT verification failed: invalid signature or structure.",
		},
		codes.EXPIRED_TOKEN: {
			UserMessage: "Session expired. Please sign in again.",
			DevMessage:  "JWT expired per exp claim.",
		},
		codes.NO_TOKEN_PROVIDED: {
			UserMessage: "Authorization token is required.",
			DevMessage:  "Missing Authorization header or token.",
		},

		codes.INVALID_EMAIL: {
			UserMessage: "Invalid email format. Please provide a valid email address.",
			DevMessage:  "Email validation failed: format does not match email pattern.",
		},
		codes.INVALID_EMAIL_OR_PASSWORD: {
			UserMessage: "Invalid email or password. Please check your credentials and try again.",
			DevMessage:  "Authentication failed: email not found or password mismatch.",
		},
		codes.INVALID_PASSWORD: {
			UserMessage: "Invalid password. Please check your password and try again.",
			DevMessage:  "Password verification failed: password hash mismatch or validation error.",
		},
	},
}

func Error(entity string, status int) string {
	if entityMap, ok := errorRegistry[entity]; ok {
		if msg, ok := entityMap[status]; ok {
			return msg.UserMessage
		}
	}
	return fmt.Sprintf("An unexpected error occurred while processing %s.", entity)
}

func ErrorDev(entity string, status int) string {
	if entityMap, ok := errorRegistry[entity]; ok {
		if msg, ok := entityMap[status]; ok {
			return msg.DevMessage
		}
	}
	return fmt.Sprintf("Unhandled error for %s, status %d", entity, status)
}
