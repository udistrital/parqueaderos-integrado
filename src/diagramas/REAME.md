# DOT (GraphViz) graph description language
Proveé una forma simple de describir gráficas que sean entendibles por humanos y computadoras. (https://es.wikipedia.org/wiki/DOT)

Para generar salida en formato PS (PostScript):
```bash
dot -Tps ingreso_automovil.gv > ingreso_automovil.ps
dot -Tps salida_automovil.gv > salida_automovil.ps
```

Para generar salida en formato PDF (Portable Document Format):
```bash
dot -Tpdf ingreso_automovil.gv -o ingreso_automovil.pdf
dot -Tpdf salida_automovil.gv -o salida_automovil.pdf
```
