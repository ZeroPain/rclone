// Package vfsflags implements command line flags to set up a vfs
package vfsflags

import (
	"github.com/ncw/rclone/fs/config"
	"github.com/ncw/rclone/vfs"
	"github.com/spf13/pflag"
)

// Options set by command line flags
var (
	Opt = vfs.DefaultOpt
)

// AddFlags adds the non filing system specific flags to the command
func AddFlags(flags *pflag.FlagSet) {
	config.BoolVarP(flags, &Opt.NoModTime, "no-modtime", "", Opt.NoModTime, "Don't read/write the modification time (can speed things up).")
	config.BoolVarP(flags, &Opt.NoChecksum, "no-checksum", "", Opt.NoChecksum, "Don't compare checksums on up/download.")
	config.BoolVarP(flags, &Opt.NoSeek, "no-seek", "", Opt.NoSeek, "Don't allow seeking in files.")
	config.DurationVarP(flags, &Opt.DirCacheTime, "dir-cache-time", "", Opt.DirCacheTime, "Time to cache directory entries for.")
	config.DurationVarP(flags, &Opt.PollInterval, "poll-interval", "", Opt.PollInterval, "Time to wait between polling for changes. Must be smaller than dir-cache-time. Only on supported remotes. Set to 0 to disable.")
	config.BoolVarP(flags, &Opt.ReadOnly, "read-only", "", Opt.ReadOnly, "Mount read-only.")
	config.FlagsVarP(flags, &Opt.CacheMode, "vfs-cache-mode", "", "Cache mode off|minimal|writes|full")
	config.DurationVarP(flags, &Opt.CachePollInterval, "vfs-cache-poll-interval", "", Opt.CachePollInterval, "Interval to poll the cache for stale objects.")
	config.DurationVarP(flags, &Opt.CacheMaxAge, "vfs-cache-max-age", "", Opt.CacheMaxAge, "Max age of objects in the cache.")
	platformFlags(flags)
}
