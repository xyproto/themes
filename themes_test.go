package themes

import "testing"

func TestBuiltinThemes(t *testing.T) {
	all := BuiltinThemes()
	if len(all) == 0 {
		t.Fatal("BuiltinThemes returned no themes")
	}
	seen := make(map[string]bool)
	for _, ti := range all {
		if ti.Name == "" {
			t.Error("ThemeInfo has empty Name")
		}
		if ti.EnvValue == "" {
			t.Errorf("ThemeInfo %q has empty EnvValue", ti.Name)
		}
		if seen[ti.EnvValue] {
			t.Errorf("duplicate EnvValue %q", ti.EnvValue)
		}
		seen[ti.EnvValue] = true
		if ti.New == nil {
			t.Errorf("ThemeInfo %q has nil New function", ti.Name)
		}
		if ti.Register == nil {
			t.Errorf("ThemeInfo %q has nil Register function", ti.Name)
		}
	}
}

func TestThemeByEnvValue(t *testing.T) {
	ti, ok := ThemeByEnvValue("gruvbox")
	if !ok {
		t.Fatal("expected to find gruvbox theme")
	}
	if ti.Name != "Gruvbox" {
		t.Errorf("expected Name \"Gruvbox\", got %q", ti.Name)
	}
	if !ti.NeedsRegistration {
		t.Error("expected Gruvbox to need registration")
	}
	if !ti.Needs256Colors {
		t.Error("expected Gruvbox to need 256 colors")
	}
	if ti.Fallback16 != "gruvbox16" {
		t.Errorf("expected Fallback16 \"gruvbox16\", got %q", ti.Fallback16)
	}

	_, ok = ThemeByEnvValue("nonexistent")
	if ok {
		t.Error("expected not to find nonexistent theme")
	}
}

func TestNewDefaultTheme(t *testing.T) {
	theme := NewDefaultTheme()
	if theme.Name != "Default" {
		t.Errorf("expected Name \"Default\", got %q", theme.Name)
	}
	if theme.Light {
		t.Error("expected default theme to be dark")
	}
}

func TestTextConfig(t *testing.T) {
	theme := NewOrbTheme()
	tc := theme.TextConfig()
	if tc == nil {
		t.Fatal("TextConfig returned nil")
	}
	if tc.Keyword == "" {
		t.Error("expected non-empty Keyword in TextConfig")
	}
	if tc.String == "" {
		t.Error("expected non-empty String in TextConfig")
	}
	if tc.Comment == "" {
		t.Error("expected non-empty Comment in TextConfig")
	}
}

func TestAllThemesConstruct(t *testing.T) {
	for _, ti := range BuiltinThemes() {
		theme := ti.New()
		if theme.Name == "" {
			t.Errorf("theme from %q constructor has empty Name", ti.EnvValue)
		}
		tc := theme.TextConfig()
		if tc == nil {
			t.Errorf("theme %q returned nil TextConfig", ti.Name)
		}
	}
}
