import numpy as np
from timeit import default_timer as timer
import matplotlib
import matplotlib.pyplot as plt
import subprocess


def time_up():
    start = timer()
    process = subprocess.run(['./network.sh', 'up'])    
    end = timer()
    print(process.returncode)
    return end - start


def time_channel():
    start = timer()
    process = subprocess.run(['./network.sh', 'createChannel'])    
    end = timer()
    print(process.returncode) 
    endProcess = subprocess.run(['./network.sh', 'down'])
    print(endProcess.returncode) 
    return end - start


def main():
    f = open('networkGen.txt', 'w')
    f.write('Network Deploy\n')
    f.write('QS-Fabric  ')
    f.write(' Fabric\n')

    g = open('channelcreation.txt', 'w')
    g.write('Channel Creation\n')
    g.write('QS-Fabric  ')
    g.write(' Fabric\n')
    '''
    h = open('certificateGen.txt', 'w')
    h.write('Certificate Generation\n')
    h.write('QS-Fabric  ')
    h.write(' Fabric\n')
    '''
    official = False
    tempoQSFNetGen = []
    tempoFabricNetGen = []
    tempoQSFCh = []
    tempoFabricCh = []
    tempoQSFCG = []
    tempoFabricCG = []
    timeQSNG = 0
    timeFNG = 0
    timeQSCh = 0
    timeFCh = 0
    timeQSCG = 0
    timeFabricCG = 0
    for j in range(1, 50, 1):
        if not official: 
            timeQSNG = time_up()
            timeQSCh = time_channel()
            if timeQSCh > 50:
                timeQSCh = 50
        else:
            timeFNG = time_up()
            timeFCh = time_channel()
 
        #timeQSCG
        #timeFabricCG
 
        tempoQSFNetGen.insert(j, timeQSNG)
        tempoFabricNetGen.insert(j, timeFNG)
        f.write(str(j))
        f.write('  & ')
        f.write(str(round(timeQSNG, 4)))
        f.write('  & ')
        f.write(str(round(timeFNG, 4)))
        f.write('\ \ \hline\n')
        tempoQSFCh.insert(j, timeQSCh)
        tempoFabricCh.insert(j, timeFCh)
        g.write(str(j))
        g.write(' &  ')
        g.write(str(round(timeQSCh, 4)))
        g.write(' &  ')
        g.write(str(round(timeFCh, 4)))
        g.write('\ \ \hline\n')
        '''
        tempoFabricCG.insert(j, timeFabricCG)
        tempoQSFCG.insert(j, timeQSCG)
        print()
        h.write(str(j))
        h.write(' &  ')
        h.write(str(round(timeQSCG, 4)))
        h.write(' &  ')
        h.write(str(round(timeFabricCG, 4)))
        h.write('\ \ \hline\n')
        '''

    x = np.arange(1, 50, 1)
    plt.plot(x, tempoQSFNetGen, label="Quantum-safe Fabric")
    plt.plot(x, tempoFabricNetGen, label="Fabric")
    plt.legend()
    plt.savefig('networkGeneration.png', bbox_inches='tight')
    plt.close()
    plt.plot(x, tempoQSFCh, label="Quantum-safe Fabric")
    plt.plot(x, tempoFabricCh, label="Fabric")
    plt.legend()
    plt.savefig('channelGeneration.png')
    plt.close()
    '''
    plt.plot(x, tempoQSFCG, label="Fabric")
    plt.plot(x, tempoFabricCG, label="Quantum-safe Fabric")
    plt.legend()
    plt.savefig('certificateGeneration.png', bbox_inches='tight')
    '''
    f.write(str(tempoQSFNetGen))
    f.write(str(tempoQSFCh))
    f.close()
    g.close()
    #h.close()


if __name__ == '__main__':
    main()