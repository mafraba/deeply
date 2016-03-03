package deeply

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	a := map[string]interface{}{
		"z": map[string]interface{}{
			"one":   "az1",
			"three": "az3",
		},
		"x": "ax",
	}
	b := map[string]interface{}{
		"z": map[string]interface{}{
			"one": "az1!",
			"two": "az2",
		},
		"y": "ay",
	}
	c := Merge(a, b)
	z1 := c["z"].(map[string]interface{})["one"]
	if z1 != "az1!" {
		t.Errorf("Expected value for 'z.one' was 'az1!', but found '%v'", z1)
	}
	z2 := c["z"].(map[string]interface{})["two"]
	if z2 != "az2" {
		t.Errorf("Expected value for 'z.two' was 'az2', but found '%v'", z2)
	}
	z3 := c["z"].(map[string]interface{})["three"]
	if z3 != "az3" {
		t.Errorf("Expected value for 'z.three' was 'az3', but found '%v'", z3)
	}
	y := c["y"]
	if y != "ay" {
		t.Errorf("Expected value for 'y' was 'ay', but found '%v'", y)
	}
	x := c["x"]
	if x != "ax" {
		t.Errorf("Expected value for 'x' was 'ax', but found '%v'", x)
	}
}

func TestCopy(t *testing.T) {
	a := map[string]interface{}{
		"z": map[string]interface{}{
			"one": "az1",
			"two": "az2",
		},
		"y": "ay",
	}
	b := Copy(a)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Copy not equal than original: '%v' '%v'", a, b)
	}
	b["y"] = "!!"
	if a["y"] == "!!" {
		t.Errorf("Modification on copy changed the original")
	}
	b["z"].(map[string]interface{})["one"] = "!!"
	if a["z"].(map[string]interface{})["one"] == "!!" {
		t.Errorf("Modification on copy changed the original")
	}
}
