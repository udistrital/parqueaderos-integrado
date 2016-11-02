#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# Este script busca en un directorio todos los archivos punto terminados en .#
# por ejemplo .1 .2 .3 ... .100 y los reemplaza por el texto indicado
# en el archivo del mismo nombre pero sin punto.
# ex: si cambié el archivo output.go que está en el directorio
# github.com/astaxie/beego/context/output.go
# debo crear un archivo punto con el mismo nombre .output.go.1
# poner en este la línea que quiero cambiar y luego de esto crear un archivo
# output.go.1 (nótese que solo se le quita el punto) y en este agregar el
# código con todo lo que he modificado.
import glob

directorio_a_reemplazar = '/home/jorge2/workspace/parqueaderos/src/' #con /

archivos = glob.glob('github.com/**/.*', recursive=True)
archivos = sorted(archivos)
print('Archivos encontrados:')
print(archivos)

class bcolors:
    HEADER = '\033[95m'
    OKBLUE = '\033[94m'
    OKGREEN = '\033[92m'
    WARNING = '\033[93m'
    FAIL = '\033[91m'
    ENDC = '\033[0m'
    BOLD = '\033[1m'
    UNDERLINE = '\033[4m'

for archivo in archivos:
    index_ult_punto = archivo.rfind('.')
    arch_sin_num = archivo[:index_ult_punto]
    index_punto = archivo.rfind('/.') + 1
    arch_sin_punto = archivo[:index_punto] + archivo[index_punto + 1:]
    index_punto = arch_sin_num.rfind('/.') + 1
    arch_sin_num_sin_punto = arch_sin_num[:index_punto] + arch_sin_num[index_punto + 1:]

    texto_busqueda = open(archivo, 'r').read()
    texto_reemplazo = open(arch_sin_punto, 'r').read()

    filedata = None
    archivo_original = directorio_a_reemplazar + arch_sin_num_sin_punto
    print(bcolors.WARNING)
    print('Archivo modificado: ' + archivo)
    print(bcolors.OKBLUE)
    print('Archivo original: ' + archivo_original)
    print(bcolors.ENDC)
    # Abrir archivo original para buscar
    with open(archivo_original, 'r') as file :
        filedata = file.read()

    # Reemplaza los strings
    print('Buscando:')
    print(bcolors.FAIL)
    print(texto_busqueda)
    print(bcolors.ENDC)
    print('Reemplazando:')
    print(bcolors.OKGREEN)
    print(texto_reemplazo)
    print(bcolors.ENDC)
    filedata = filedata.replace(texto_busqueda, texto_reemplazo)

    # Guarda en un nuevo archivo el restulado
    with open(archivo_original + '.new', 'w') as file:
        file.write(filedata)
