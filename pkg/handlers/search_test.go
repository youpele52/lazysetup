package handlers

import (
	"testing"
)

// P2 - Search filtering correctness
// Tests that FilterTools correctly filters tools based on query
func TestFilterTools_CaseInsensitive(t *testing.T) {
	tools := []string{"git", "docker", "kubectl", "k9s", "lazydocker"}

	t.Run("lowercase query matches", func(t *testing.T) {
		result := FilterTools(tools, "kube")
		if len(result) != 1 || result[0] != "kubectl" {
			t.Errorf("Expected [kubectl], got %v", result)
		}
	})

	t.Run("uppercase query matches", func(t *testing.T) {
		result := FilterTools(tools, "KUBE")
		if len(result) != 1 || result[0] != "kubectl" {
			t.Errorf("Expected [kubectl], got %v", result)
		}
	})

	t.Run("mixed case query matches", func(t *testing.T) {
		result := FilterTools(tools, "KuBe")
		if len(result) != 1 || result[0] != "kubectl" {
			t.Errorf("Expected [kubectl], got %v", result)
		}
	})

	t.Run("empty query returns all", func(t *testing.T) {
		result := FilterTools(tools, "")
		if len(result) != len(tools) {
			t.Errorf("Expected all %d tools, got %d", len(tools), len(result))
		}
	})

	t.Run("no matches returns empty", func(t *testing.T) {
		result := FilterTools(tools, "xyz")
		if len(result) != 0 {
			t.Errorf("Expected empty result, got %v", result)
		}
	})

	t.Run("partial match works", func(t *testing.T) {
		result := FilterTools(tools, "k")
		// Should match kubectl, k9s, docker, and lazydocker (via display names)
		if len(result) != 4 {
			t.Errorf("Expected 4 results, got %v", result)
		}
		// Verify kubectl and k9s are in results
		found := make(map[string]bool)
		for _, tool := range result {
			found[tool] = true
		}
		if !found["kubectl"] || !found["k9s"] {
			t.Errorf("Expected kubectl and k9s to be in results, got %v", result)
		}
	})

	t.Run("multiple matches", func(t *testing.T) {
		result := FilterTools(tools, "docker")
		// Should match docker and lazydocker
		if len(result) != 2 {
			t.Errorf("Expected 2 results, got %v", result)
		}
		// Verify both docker and lazydocker are in results
		found := make(map[string]bool)
		for _, tool := range result {
			found[tool] = true
		}
		if !found["docker"] || !found["lazydocker"] {
			t.Errorf("Expected docker and lazydocker, got %v", result)
		}
	})
}

func TestFilterTools_SpecialCharacters(t *testing.T) {
	tools := []string{"git", "docker-compose", "k8s-tools"}

	t.Run("dash in query", func(t *testing.T) {
		result := FilterTools(tools, "docker-")
		if len(result) != 1 || result[0] != "docker-compose" {
			t.Errorf("Expected [docker-compose], got %v", result)
		}
	})

	t.Run("number in query", func(t *testing.T) {
		result := FilterTools(tools, "8s")
		if len(result) != 1 || result[0] != "k8s-tools" {
			t.Errorf("Expected [k8s-tools], got %v", result)
		}
	})
}

func TestFilterTools_EmptyToolsList(t *testing.T) {
	tools := []string{}

	t.Run("empty tools list returns empty", func(t *testing.T) {
		result := FilterTools(tools, "anything")
		if len(result) != 0 {
			t.Errorf("Expected empty result, got %v", result)
		}
	})
}

func TestFilterTools_SingleTool(t *testing.T) {
	tools := []string{"git"}

	t.Run("match returns tool", func(t *testing.T) {
		result := FilterTools(tools, "git")
		if len(result) != 1 || result[0] != "git" {
			t.Errorf("Expected [git], got %v", result)
		}
	})

	t.Run("no match returns empty", func(t *testing.T) {
		result := FilterTools(tools, "docker")
		if len(result) != 0 {
			t.Errorf("Expected empty result, got %v", result)
		}
	})
}
