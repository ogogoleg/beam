// beam-playground:
//   name: IntroductionConcepts
//   description: Assignment for Introduction/Basic Concepts unit.
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

	// implement missing code 

}