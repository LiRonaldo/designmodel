package factory

import "testing"

func TestNewRestaurant(t *testing.T) {
	NewRestaurant("Q").GetFood()
	NewRestaurant("D").GetFood()
}
