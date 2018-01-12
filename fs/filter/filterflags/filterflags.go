// Package filterflags implements command line flags to set up a filter
package filterflags

import (
	"github.com/ncw/rclone/fs/config"
	"github.com/ncw/rclone/fs/filter"
	"github.com/spf13/pflag"
)

// Options set by command line flags
var (
	Opt = filter.DefaultOpt
)

// AddFlags adds the non filing system specific flags to the command
func AddFlags(flags *pflag.FlagSet) {
	config.BoolVarP(flags, &Opt.DeleteExcluded, "delete-excluded", "", false, "Delete files on dest excluded from sync")
	config.StringArrayVarP(flags, &Opt.FilterRule, "filter", "f", nil, "Add a file-filtering rule")
	config.StringArrayVarP(flags, &Opt.FilterFrom, "filter-from", "", nil, "Read filtering patterns from a file")
	config.StringArrayVarP(flags, &Opt.ExcludeRule, "exclude", "", nil, "Exclude files matching pattern")
	config.StringArrayVarP(flags, &Opt.ExcludeFrom, "exclude-from", "", nil, "Read exclude patterns from file")
	config.StringVarP(flags, &Opt.ExcludeFile, "exclude-if-present", "", "", "Exclude directories if filename is present")
	config.StringArrayVarP(flags, &Opt.IncludeRule, "include", "", nil, "Include files matching pattern")
	config.StringArrayVarP(flags, &Opt.IncludeFrom, "include-from", "", nil, "Read include patterns from file")
	config.StringArrayVarP(flags, &Opt.FilesFrom, "files-from", "", nil, "Read list of source-file names from file")
	config.FlagsVarP(flags, &Opt.MinAge, "min-age", "", "Don't transfer any file younger than this in s or suffix ms|s|m|h|d|w|M|y")
	config.FlagsVarP(flags, &Opt.MaxAge, "max-age", "", "Don't transfer any file older than this in s or suffix ms|s|m|h|d|w|M|y")
	config.FlagsVarP(flags, &Opt.MinSize, "min-size", "", "Don't transfer any file smaller than this in k or suffix b|k|M|G")
	config.FlagsVarP(flags, &Opt.MaxSize, "max-size", "", "Don't transfer any file larger than this in k or suffix b|k|M|G")
	//cvsExclude     = BoolP("cvs-exclude", "C", false, "Exclude files in the same way CVS does")
}
