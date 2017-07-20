package walg

import ()

type UnsetEnvVarError struct {
	names []string
}

func (e UnsetEnvVarError) Error() string {
	msg := "Did not set the following environment variables:\n"
	for _, v := range e.names {
		msg = msg + v + "\n"
	}

	return msg
}