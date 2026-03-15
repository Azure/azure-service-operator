# plot-metrics.gp — Generate CPU and memory charts from a perf metrics CSV.
#
# Usage:
#   gnuplot -c plot-metrics.gp <csv_file> [output_dir]
#
# The CSV must have the header:
#   timestamp,elapsed_seconds,pod,container,cpu_millicores,memory_mb
#
# Outputs two PNG files:
#   <output_dir>/<basename>_cpu.png
#   <output_dir>/<basename>_memory.png
#
# If output_dir is omitted, PNGs are written next to the input CSV.

if (ARGC < 1) {
    print "Usage: gnuplot -c plot-metrics.gp <csv_file> [output_dir]"
    exit
}

csv_file = ARG1

# Derive base name by stripping directory and .csv extension
bare = system(sprintf("basename '%s' .csv", csv_file))

if (ARGC >= 2) {
    out_dir = ARG2
} else {
    # Default: same directory as the CSV
    out_dir = system(sprintf("dirname '%s'", csv_file))
}

cpu_out  = out_dir . "/" . bare . "_cpu.png"
mem_out  = out_dir . "/" . bare . "_memory.png"

# Discover unique pod names from the CSV (skip header)
pods = system(sprintf("tail -n +2 '%s' | cut -d',' -f3 | sort -u | tr '\\n' ' '", csv_file))

# Shorten pod name for legend: take last dash-delimited segment
shorten(name) = system(sprintf("echo '%s' | rev | cut -d'-' -f1 | rev | tr -d '\\n'", name))

set datafile separator ','
set key outside right
set grid
set style data linespoints

# ── CPU chart ──────────────────────────────────────────────────────
set terminal pngcairo size 1200,600 enhanced font 'Arial,12'
set output cpu_out

set title 'ASO Controller CPU Usage'
set xlabel 'Elapsed (s)'
set ylabel 'CPU (millicores)'

plot for [pod in pods] \
    sprintf("< grep '%s' '%s'", pod, csv_file) \
    using 2:5 title shorten(pod)

# ── Memory chart ───────────────────────────────────────────────────
set output mem_out

set title 'ASO Controller Memory Usage'
set xlabel 'Elapsed (s)'
set ylabel 'Memory (MB)'

plot for [pod in pods] \
    sprintf("< grep '%s' '%s'", pod, csv_file) \
    using 2:6 title shorten(pod)

print sprintf("Wrote %s", cpu_out)
print sprintf("Wrote %s", mem_out)
