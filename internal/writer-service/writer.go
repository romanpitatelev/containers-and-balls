package writerservice

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

type Writer struct{}

func New() *Writer {
	return &Writer{}
}

func (w *Writer) Write(result string) error {
	output := bufio.NewWriter(os.Stdout)

	defer func() {
		err := output.Flush()
		if err != nil {
			log.Warn().Err(err).Msg("failed to write buffered data")
		}
	}()

	_, err := fmt.Fprintln(output, result)
	if err != nil {
		return fmt.Errorf("failed to write: %w", err)
	}

	return nil
}
