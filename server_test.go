package socketio

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestFiberRoute_EmptyPath_Fix(t *testing.T) {
	app := fiber.New()
	io := New()

	// Mount the socket.io handler
	app.Route("/socket.io", io.FiberRoute)

	// Send a request to the root of the socket.io path
	// This previously caused os.Open("") which resulted in a 500 error (or panic handled by Fiber)
	req := httptest.NewRequest("GET", "/socket.io/", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}

	// We expect the request to be handled gracefully (likely 426 Upgrade Required or 400 Bad Request from websocket handler)
	// But definitely NOT 500 Internal Server Error
	if resp.StatusCode == 500 {
		t.Errorf("Test failed: Received 500 Internal Server Error. The fix might not be working.")
	} else {
		t.Logf("Test passed: Received status code %d", resp.StatusCode)
	}
}
