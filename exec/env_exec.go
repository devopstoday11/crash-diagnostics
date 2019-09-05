package exec

import "gitlab.eng.vmware.com/vivienv/flare/script"

// execEnvs rertrieves all specified environment variables
// and return as a slice of []string{key=value} pair.
func exeEnvs(src *script.Script) []string {
	envCmds := src.Preambles[script.CmdEnv]
	var envPairs []string
	for _, envCmd := range envCmds {
		env := envCmd.(*script.EnvCommand)
		if len(env.Envs()) > 0 {
			for _, arg := range env.Envs() {
				envPairs = append(envPairs, arg)
			}
		}
	}
	return envPairs
}