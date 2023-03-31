import subprocess
import random
import time
import resource
import sys
import os

my_env = os.environ.copy()
my_env['FABRIC_CFG_PATH'] = os.environ["PWD"] + "/../config/"


colors = ['blue', 'red', 'green', 'yellow', 'black', 'purple', 'white', 'violet', 'indigo', 'brown']
makes = ['Toyota', 'Ford', 'Hyundai', 'Volkswagen', 'Tesla', 'Peugeot', 'Chery', 'Fiat', 'Tata', 'Holden']
models = ['Prius', 'Mustang', 'Tucson', 'Passat', 'S', '205', 'S22L', 'Punto', 'Nano', 'Barina']
owners = ['Tomoko', 'Brad', 'Jin Soo', 'Max', 'Adrianna', 'Michel', 'Aarav', 'Pari', 'Valeria', 'Shotaro']
channel = 'mychannel'
chaincode = 'fabcar'
failed = 0

# Funzione per eseguire una transazione
def run_transaction(command, max_tries=5):
    tries = 0
    global failed
    while tries < max_tries:
        start_time = time.time()
        try:
            subprocess.check_output(command, env=my_env, shell=True)
        except subprocess.CalledProcessError:
            tries += 1
            failed += 1
            continue
        end_time = time.time()
        elapsed_time = end_time - start_time
        rusage = resource.getrusage(resource.RUSAGE_CHILDREN)
        '''
        print("CPU time: {} seconds".format(rusage.ru_utime))
        print("Memory usage: {} bytes".format(rusage.ru_maxrss))
        print("Elapsed time: {} seconds".format(elapsed_time))
        '''

        return elapsed_time



# Funzione per testare una transazione
def test_transaction(function_name, num_executions, org):
    total_time = 0
    if function_name == "CreateCar" or function_name == "ChangeCarOwner":
        orderer_tls_host = 'orderer.example.com'
        if org == 'org1':
            orderer_address = 'localhost:7050'
            peer_addresses = ['localhost:7051', 'localhost:9051']
            peer_tls_files = [f'{os.environ["PWD"]}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt',
                              f'{os.environ["PWD"]}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt']
        elif org == 'org2':
            orderer_address = 'localhost:9050'
            peer_addresses = ['localhost:9051', 'localhost:7051']
            peer_tls_files = [f'{os.environ["PWD"]}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt',
                              f'{os.environ["PWD"]}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt']
        else:
            raise ValueError(f'Invalid org: {org}')
        
        '''invoke_command = f'peer chaincode invoke -o {orderer_address} --ordererTLSHostnameOverride {orderer_tls_host} --tls --cafile "{os.environ["PWD"]}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C {channel} -n {chaincode} --peerAddresses localhost:7051 --tlsRootCertFiles "{os.environ["PWD"]}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "{os.environ["PWD"]}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"'
        '''
        invoke_command = f'peer chaincode invoke -o {orderer_address} --ordererTLSHostnameOverride {orderer_tls_host} --tls --cafile "{os.environ["PWD"]}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C {channel} -n {chaincode}'
        for i in range(len(peer_addresses)):
            invoke_command += f' --peerAddresses {peer_addresses[i]} --tlsRootCertFiles "{peer_tls_files[i]}"'
            
        if function_name == "ChangeCarOwner":
            for i in range(num_executions):
                command = invoke_command + f' -c \'{{"function":"{function_name}","Args":["CAR3","{random.choice(owners)}"]}}\''
                total_time += run_transaction(command)
                time.sleep(1.5)
            
        else:
            for i in range(num_executions):
                command = invoke_command + f' -c \'{{"function":"{function_name}","Args":["CAR{i+10}","{random.choice(makes)}","{random.choice(models)}","{random.choice(colors)}","{random.choice(owners)}"]}}\''
                total_time += run_transaction(command)
                time.sleep(1.5)

                
    
    elif function_name == "QueryAllCars" or function_name == "QueryCar":
        if function_name == "QueryAllCars":
            command = f'peer chaincode query -C {channel} -n {chaincode} -c \'{{"Args":["{function_name}"]}}\''
            for i in range(num_executions):
                    total_time += run_transaction(command)
        else:
            command = f'peer chaincode query -C {channel} -n {chaincode} -c \'{{"function":"{function_name}","Args":["CAR0"]}}\''
            for i in range(num_executions):
                    total_time += run_transaction(command)
    else:
        print(f"Error: {function_name} is not a valid function in chaincode {chaincode}.")
        sys.exit(1)
    
    
    tps = num_executions / total_time
    print(f"{function_name}:\niteration: {num_executions} \ntotal-time: ({total_time} s) \ntransactions per second: ({tps} tps) \naverage time for execution: ({total_time/num_executions} s)\nfailed transactions: {failed}")

# Get the command line arguments
function_name = sys.argv[1]
num_executions = int(sys.argv[2])
org_name = sys.argv[3]
# Set the environment variables for the specified organization
os.environ['CORE_PEER_TLS_ENABLED'] = 'true'

if org_name == "org1":
    my_env['CORE_PEER_LOCALMSPID'] = 'Org1MSP'
    my_env['CORE_PEER_TLS_ROOTCERT_FILE'] = os.environ["PWD"] + "/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt"
    my_env['CORE_PEER_MSPCONFIGPATH'] = os.environ["PWD"] + "/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp"
    my_env['CORE_PEER_ADDRESS'] = 'localhost:7051'
elif org_name == "org2":
    my_env['CORE_PEER_LOCALMSPID'] = 'Org2MSP'
    my_env['CORE_PEER_TLS_ROOTCERT_FILE'] = os.environ["PWD"] + "/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"
    my_env['CORE_PEER_MSPCONFIGPATH'] = os.environ["PWD"] + "/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp"
    my_env['CORE_PEER_ADDRESS'] = 'localhost:9051'
else:
    print("Invalid organization name")
    sys.exit(1)

test_transaction(function_name, num_executions, org_name)
