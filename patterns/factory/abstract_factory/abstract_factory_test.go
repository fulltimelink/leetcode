package abstract_factory

import (
	"testing"
)

// TestGetSportsFactory 测试 GetSportsFactory 函数.
func TestGetSportsFactory(t *testing.T) {
	// 测试 Adidas 品牌.
	if factory, err := GetSportsFactory("adidas"); err != nil || factory == nil {
		t.Errorf("Failed to get Adidas factory: %v", err)
	} else {
		// 你可以在这里添加对 Adidas 工厂的特定行为的测试.

	}

	// 测试 Nike 品牌.
	if factory, err := GetSportsFactory("nike"); err != nil || factory == nil {
		t.Errorf("Failed to get Nike factory: %v", err)
	} else {
		// 你可以在这里添加对 Nike 工厂的特定行为的测试.
	}

	// 测试一个不存在的品牌.
	if factory, err := GetSportsFactory("puma"); err == nil || factory != nil {
		t.Errorf("Expected an error for unknown brand, got factory: %v", factory)
	}
}
