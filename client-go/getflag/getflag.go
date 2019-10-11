package getflag

import "flag"

//ConfigDir is pwd of kubeconfig dir
var ConfigDir string

//LogDir figures out the pwd of log files
var LogDir string

//GetFlag get command line flags
func GetFlag() {
	flag.StringVar(&LogDir, "logdir", "", "absolute path to log dir")
	flag.StringVar(&ConfigDir, "configdir", "", "absolute path to kubeconfig dir")
	flag.Parse()
}
