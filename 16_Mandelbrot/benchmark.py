import subprocess
import codecs

size = [60, 100, 500, 1000, 10000]
accuracy = [100, 200, 500, 1000, 10000]
cpuCount = [1, 2, 4, 6, 8, 10, 12]

with open("results.txt", "w+") as f:
    for s in size:
        for a in accuracy:
            for c in cpuCount:
                print(f"Running: size={s} accuracy={a} cpuCount={c}")
                f.write(codecs.decode((subprocess.check_output(
                    f"./main -benchmark -size={s} -accuracy={a} -cpuCount={c}", shell=True)), 'UTF-8'))
            f.flush()
