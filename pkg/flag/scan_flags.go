package flag

import (
	"github.com/aquasecurity/trivy/pkg/types"
	xstrings "github.com/aquasecurity/trivy/pkg/x/strings"
)

var (
	SkipDirsFlag = Flag{
		Name:       "skip-dirs",
		ConfigName: "scan.skip-dirs",
		Default:    []string{},
		Usage:      "specify the directories or glob patterns to skip",
	}
	SkipFilesFlag = Flag{
		Name:       "skip-files",
		ConfigName: "scan.skip-files",
		Default:    []string{},
		Usage:      "specify the files or glob patterns to skip",
	}
	OfflineScanFlag = Flag{
		Name:       "offline-scan",
		ConfigName: "scan.offline",
		Default:    false,
		Usage:      "do not issue API requests to identify dependencies",
	}
	ScannersFlag = Flag{
		Name:       "scanners",
		ConfigName: "scan.scanners",
		Default: xstrings.ToStringSlice(types.Scanners{
			types.VulnerabilityScanner,
			types.SecretScanner,
		}),
		Values: xstrings.ToStringSlice(types.Scanners{
			types.VulnerabilityScanner,
			types.MisconfigScanner,
			types.SecretScanner,
			types.LicenseScanner,
		}),
		Aliases: []Alias{
			{
				Name:       "security-checks",
				ConfigName: "scan.security-checks",
				Deprecated: true, // --security-checks was renamed to --scanners
			},
		},
		Usage: "comma-separated list of what security issues to detect",
	}
	FilePatternsFlag = Flag{
		Name:       "file-patterns",
		ConfigName: "scan.file-patterns",
		Default:    []string{},
		Usage:      "specify config file patterns",
	}
	SlowFlag = Flag{
		Name:       "slow",
		ConfigName: "scan.slow",
		Default:    false,
		Usage:      "scan over time with lower CPU and memory utilization",
	}
	SBOMSourcesFlag = Flag{
		Name:       "sbom-sources",
		ConfigName: "scan.sbom-sources",
		Default:    []string{},
		Values:     []string{"oci", "rekor"},
		Usage:      "[EXPERIMENTAL] try to retrieve SBOM from the specified sources",
	}
	RekorURLFlag = Flag{
		Name:       "rekor-url",
		ConfigName: "scan.rekor-url",
		Default:    "https://rekor.sigstore.dev",
		Usage:      "[EXPERIMENTAL] address of rekor STL server",
	}
	IncludeDevDepsFlag = Flag{
		Name:       "include-dev-deps",
		ConfigName: "include-dev-deps",
		Default:    false,
		Usage:      "include development dependencies in the report (supported: npm, yarn)",
	}
)

type ScanFlagGroup struct {
	SkipDirs       *Flag
	SkipFiles      *Flag
	OfflineScan    *Flag
	Scanners       *Flag
	FilePatterns   *Flag
	Slow           *Flag
	SBOMSources    *Flag
	RekorURL       *Flag
	IncludeDevDeps *Flag
}

type ScanOptions struct {
	Target         string
	SkipDirs       []string
	SkipFiles      []string
	OfflineScan    bool
	Scanners       types.Scanners
	FilePatterns   []string
	Slow           bool
	SBOMSources    []string
	RekorURL       string
	IncludeDevDeps bool
}

func NewScanFlagGroup() *ScanFlagGroup {
	return &ScanFlagGroup{
		SkipDirs:       &SkipDirsFlag,
		SkipFiles:      &SkipFilesFlag,
		OfflineScan:    &OfflineScanFlag,
		Scanners:       &ScannersFlag,
		FilePatterns:   &FilePatternsFlag,
		Slow:           &SlowFlag,
		SBOMSources:    &SBOMSourcesFlag,
		RekorURL:       &RekorURLFlag,
		IncludeDevDeps: &IncludeDevDepsFlag,
	}
}

func (f *ScanFlagGroup) Name() string {
	return "Scan"
}

func (f *ScanFlagGroup) Flags() []*Flag {
	return []*Flag{
		f.SkipDirs,
		f.SkipFiles,
		f.OfflineScan,
		f.Scanners,
		f.FilePatterns,
		f.Slow,
		f.SBOMSources,
		f.RekorURL,
		f.IncludeDevDeps,
	}
}

func (f *ScanFlagGroup) ToOptions(args []string) (ScanOptions, error) {
	var target string
	if len(args) == 1 {
		target = args[0]
	}

	return ScanOptions{
		Target:         target,
		SkipDirs:       getStringSlice(f.SkipDirs),
		SkipFiles:      getStringSlice(f.SkipFiles),
		OfflineScan:    getBool(f.OfflineScan),
		Scanners:       getUnderlyingStringSlice[types.Scanner](f.Scanners),
		FilePatterns:   getStringSlice(f.FilePatterns),
		Slow:           getBool(f.Slow),
		SBOMSources:    getStringSlice(f.SBOMSources),
		RekorURL:       getString(f.RekorURL),
		IncludeDevDeps: getBool(f.IncludeDevDeps),
	}, nil
}
