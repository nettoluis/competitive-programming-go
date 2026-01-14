import os
import subprocess
import glob
import sys

BINARY_NAME = "solucao_temp"

def run_tests():
    
    if len(sys.argv) < 2:
        print("❌ Erro: Especifique o arquivo Go.")
        print("   Uso correto: python test_runner.py nome_do_arquivo.go")
        return

    go_file = sys.argv[1] 

    if not os.path.exists(go_file):
        print(f"❌ Erro: O arquivo '{go_file}' não existe nesta pasta.")
        return
    

    print(f"--- Iniciando Testes para {go_file} ---")
    
    
    print("Compilando...", end=" ")
    
    if os.name == 'nt': 
        binary_name_full = BINARY_NAME + ".exe"
    else: 
        binary_name_full = "./" + BINARY_NAME

    
    compile_cmd = ["go", "build", "-o", BINARY_NAME, go_file]
    
    if subprocess.call(compile_cmd) != 0:
        print("❌ Erro de compilação!")
        return
    print("Sucesso!\n")

    
    inputs = sorted(glob.glob("testes/*.in"))
    
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
                    [binary_name_full], 
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
        os.remove(binary_name_full)
    except:
        pass

if __name__ == "__main__":
    run_tests()
