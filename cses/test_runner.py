import os
import subprocess
import glob
import sys

GO_FILE = "missing-number.go"
BINARY_NAME = "solucao_temp"

def run_tests():
    print(f"--- Iniciando Testes para {GO_FILE} ---")
    
    # 1. Compila o código Go
    print("Compilando...", end=" ")
    if os.name == 'nt':
        BINARY_NAME_FULL = BINARY_NAME + ".exe"
    else:
        BINARY_NAME_FULL = "./" + BINARY_NAME

    compile_cmd = ["go", "build", "-o", BINARY_NAME, GO_FILE]
    if subprocess.call(compile_cmd) != 0:
        print("❌ Erro de compilação!")
        return
    print("Sucesso!\n")

    inputs = sorted(glob.glob("./testes/*.in"))
    
    if not inputs:
        print("Nenhum arquivo .in encontrado.")
        return

    for input_file in inputs:
        expected_output_file = input_file.replace(".in", ".out")
        
        if not os.path.exists(expected_output_file):
            print(f"⚠️  {input_file}: Arquivo .out correspondente não encontrado. Pulando.")
            continue
            
        print(f"Testando {input_file}...", end=" ")
        
        with open(input_file, 'r') as infile:
            try:
                result = subprocess.run(
                    [BINARY_NAME_FULL], 
                    stdin=infile, 
                    capture_output=True, 
                    text=True, 
                    timeout=2
                )
                
                user_output = result.stdout.strip()
                
                with open(expected_output_file, 'r') as outfile:
                    expected_output = outfile.read().strip()
                
                if user_output == expected_output:
                    print("✅ PASSOU")
                else:
                    print("❌ FALHOU")
                    print(f"   Esperado (início): {expected_output[:40]}...")
                    print(f"   Obtido   (início): {user_output[:40]}...")
                    
            except subprocess.TimeoutExpired:
                print("⏱️  TEMPO EXCEDIDO (TLE)")
            except Exception as e:
                print(f"⚠️  ERRO DE EXECUÇÃO: {e}")

    try:
        os.remove(BINARY_NAME_FULL)
    except:
        pass

if __name__ == "__main__":
    run_tests()
