// tour-of-beam:
//   name: IntroductionConceptsSolution
//   description: An example that analyzes traffic sensor data using SlidingWindows.
//     For each window, it calculates the average speed over the window for some small set
//     of predefined 'routes', and looks for 'slowdowns' in those routes. It writes its
//     results to a BigQuery table.
//   multifile: true
//   context_line: 97
//   categories:
//     - Combiners
//     - Streaming
//     - Options
//     - Windowing

package main

import (
	"beam.apache.org/learning/katas/introduction/hello_beam/hello_beam/pkg/task"
	"context"
	"github.com/apache/beam/sdks/go/pkg/beam"
	"github.com/apache/beam/sdks/go/pkg/beam/log"
	"github.com/apache/beam/sdks/go/pkg/beam/x/beamx"
	"github.com/apache/beam/sdks/go/pkg/beam/x/debug"
)

func main() {
	p, s := beam.NewPipelineWithRoot()

	hello := task.HelloBeam(s)

	debug.Print(s, hello)

	err := beamx.Run(context.Background(), p)
	if err != nil {
		log.Exitf(context.Background(), "Failed to execute job: %v", err)
	}
}