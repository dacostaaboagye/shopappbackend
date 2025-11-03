package messages

import (
	"net/http"

	"github.com/Aboagye-Dacosta/shopBackend/internal/codes"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/entities"
)

var successRegistry = map[string]map[int]Message{
	entities.USER: {
		http.StatusCreated: {
			UserMessage: "User registered successfully.",
			DevMessage:  "User entity persisted to database.",
		},
		http.StatusOK: {
			UserMessage: "User details fetched successfully.",
			DevMessage:  "User entity retrieved from DB.",
		},
		http.StatusAccepted: {
			UserMessage: "User updated successfully.",
			DevMessage:  "User entity updated in database.",
		},
		http.StatusNoContent: {
			UserMessage: "User deleted successfully.",
			DevMessage:  "User entity deleted from database.",
		},
		codes.LOGIN_SUCCESS: {
			UserMessage: "Login successful. Welcome back!",
			DevMessage:  "User authenticated successfully, JWT token generated.",
		},
	},
	entities.PRODUCT: {
		http.StatusCreated: {
			UserMessage: "Product created successfully.",
			DevMessage:  "Product entity persisted to database.",
		},
		http.StatusOK: {
			UserMessage: "Product details retrieved successfully.",
			DevMessage:  "Product entity retrieved from DB.",
		},
		http.StatusAccepted: {
			UserMessage: "Product updated successfully.",
			DevMessage:  "Product entity updated in database.",
		},
		http.StatusNoContent: {
			UserMessage: "Product deleted successfully.",
			DevMessage:  "Product entity deleted from database.",
		},
	},
	entities.ORDER: {
		http.StatusCreated: {
			UserMessage: "Order placed successfully.",
			DevMessage:  "Order persisted and confirmation sent.",
		},
		http.StatusOK: {
			UserMessage: "Order details retrieved successfully.",
			DevMessage:  "Order details loaded.",
		},
		http.StatusAccepted: {
			UserMessage: "Order updated successfully.",
			DevMessage:  "Order updated in database.",
		},
		http.StatusNoContent: {
			UserMessage: "Order deleted successfully.",
			DevMessage:  "Order deleted from database.",
		},
	},
	entities.PAYMENT: {
		http.StatusCreated: {
			UserMessage: "Payment processed successfully.",
			DevMessage:  "Payment recorded and processed.",
		},
		http.StatusOK: {
			UserMessage: "Payment details retrieved successfully.",
			DevMessage:  "Payment details loaded.",
		},
		http.StatusAccepted: {
			UserMessage: "Payment updated successfully.",
			DevMessage:  "Payment record updated in database.",
		},
		http.StatusNoContent: {
			UserMessage: "Payment deleted successfully.",
			DevMessage:  "Payment record deleted from database.",
		},
	},
}

func Success(entity string, status int) string {
	if entityMap, ok := successRegistry[entity]; ok {
		if msg, ok := entityMap[status]; ok {
			return msg.UserMessage
		}
	}
	return "Operation completed successfully."
}

func SuccessDev(entity string, status int) string {
	if entityMap, ok := successRegistry[entity]; ok {
		if msg, ok := entityMap[status]; ok {
			return msg.DevMessage
		}
	}
	return "Operation succeeded with unspecified details."
}
