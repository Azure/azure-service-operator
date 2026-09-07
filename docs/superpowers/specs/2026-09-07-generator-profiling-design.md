# ASO Generator Profiling

## Why this matters

The ASO code generator can consume substantial CPU and memory during a run, but it currently provides no direct way to capture profiles for offline investigation. Developers need opt-in profiles that represent the command's complete workload and that can be inspected with standard `go tool pprof` commands.

## Command-line interface

The generator root command will expose two persistent options:

- `--cpu-prof <file>` writes a CPU profile.
- `--memory-prof <file>` writes a cumulative allocation profile.

Because these are persistent options, they will work with every generator subcommand. Either option may be used alone, or both may be used in the same invocation.

The memory profile will use Go's `allocs` profile rather than a live-heap snapshot. This records sampled allocations made throughout the command, including objects that were released before the command finished.

## Profiling lifecycle

A small profiling lifecycle component will own the output files and the interaction with `runtime/pprof`.

After Cobra has parsed and validated the command line, but before the selected command starts its work, the lifecycle will:

1. Create each requested output file.
2. Start CPU profiling when `--cpu-prof` is present.

After command execution returns, whether it succeeded or failed, the lifecycle will:

1. Stop CPU profiling so buffered samples are flushed.
2. Write the binary `allocs` profile when `--memory-prof` is present.
3. Close all output files.

Starting the profiler from the root command's persistent pre-run hook applies it consistently to all subcommands. Finalising it outside Cobra's post-run hooks is important because Cobra does not run post-run hooks after a command error.

The resulting profiles cover the complete execution of the selected command. Cobra construction, argument parsing, and profile initialisation necessarily happen before profiling starts.

## Error handling

If a requested output file cannot be created, or CPU profiling cannot start, command execution will not begin. Any files already opened during a partially successful start will be closed.

Errors encountered while writing or closing profiles will be returned to the caller. If command execution and profile finalisation both fail, both errors will be retained so profiling does not hide the generator's original failure.

## Testing

Focused tests will verify:

- the new options are persistent root flags;
- CPU profiling produces a pprof-compatible file;
- memory profiling produces a pprof-compatible cumulative allocation profile;
- both profiles can be enabled concurrently;
- profiles are finalised when the selected command returns an error;
- invalid output paths prevent command execution and return a useful error.

The tests will use short in-process Cobra commands and temporary output files. They will parse the generated data with the pprof profile parser already available through the repository's dependencies, avoiding fragile assertions about binary file size alone.
