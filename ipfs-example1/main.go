package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	// El paquete shell implementa una interfaz API remota para un IPFS daemon
	// en ejecucion.
	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	//-----------------------------------------------------------------------
	// Esta linea conecta a tu IPFS daemon corriendo en segundo plano.
	//-----------------------------------------------------------------------
	sh := shell.NewShell("localhost:5001")
	//-----------------------------------------------------------------------

	//-----------------------------------------------------------------------
	// Aqui agregamos el archivo a IPFS.
	//-----------------------------------------------------------------------
	cid, err := sh.Add(strings.NewReader("hola IPFS!"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s\n", cid)
	//-----------------------------------------------------------------------

	//-----------------------------------------------------------------------
	// Aqui obtenemos la data de IPFS y guardamos el contenido en un archivo.
	//-----------------------------------------------------------------------
	out := fmt.Sprintf("%s.txt", cid)
	err = sh.Get(cid, out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	//-----------------------------------------------------------------------

	//-----------------------------------------------------------------------
	// Aqui obtenemos la data de IPFS y guardamos el contenido en la variable
	// <<data>>.
	//-----------------------------------------------------------------------
	data, err := sh.Cat(cid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	//-----------------------------------------------------------------------

	//-----------------------------------------------------------------------
	// Finalmente convertimos la variable <<data>> en string pasandola atravez
	// de un buffer y la imprimimos en pantalla.
	//-----------------------------------------------------------------------
	buf := new(bytes.Buffer)
	buf.ReadFrom(data)
	newStr := buf.String()
	fmt.Printf("data %s", newStr)
}
