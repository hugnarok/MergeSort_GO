import pandas as pd
import matplotlib.pyplot as plt

df = pd.read_csv("results.csv")

def plot(metric, ylabel):
    for struct in df['structure'].unique():
        subset = df[df['structure'] == struct]
        for rep in subset['representation'].unique():
            data = subset[subset['representation'] == rep]
            label = f"{struct}-{rep}"
            plt.plot(data['size'], data[metric], marker='o', label=label)

    plt.xscale('log')
    plt.yscale('log')
    plt.xlabel("Tamanho da entrada (log)")
    plt.ylabel(ylabel)
    plt.title(ylabel + " vs tamanho")
    plt.legend()
    plt.tight_layout()
    plt.savefig("grafico.png")
    plt.close()

plot('duration_ns', 'Duração (ns)')
plot('mem_bytes',  'Memória (bytes)')
