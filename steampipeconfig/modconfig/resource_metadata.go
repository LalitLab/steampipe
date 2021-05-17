package modconfig

// ResourceMetadata ius a struct containing additional data
// about each resource, used to populate the reflection tables
type ResourceMetadata struct {
	ResourceName string
	// mod name in the format mod.<modName>@<version?
	ModName          string
	FileName         string
	StartLineNumber  int
	EndLineNumber    int
	IsAutoGenerated  bool
	SourceDefinition string

	// mod short name
	ModShortName string
}

// SetMod sets the mod name and mod short name
func (m *ResourceMetadata) SetMod(mod *Mod) {
	// if the mod is the auto-generated default workspace mod, do not save in metadata
	if mod.IsDefaultMod() {
		return
	}
	m.ModShortName = mod.ShortName
	m.ModName = mod.FullName
}

// TODO ADD PATH ltree
