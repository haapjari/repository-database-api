package util

import "log/slog"

func Error(err error) error {
	slog.Error(err.Error())
	return err
}
