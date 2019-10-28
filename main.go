package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(config.ServerName)

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	ev, _ := clientset.CoreV1().Events("").List(metav1.ListOptions{})
	for {
		for _, even := range ev.Items {

			probnode := Findprob(even.Source.Host, even.Message)
			fmt.Println(even.Message)
			if probnode == "" {
				fmt.Println("empty")
			}
			fmt.Println(probnode)

		}
		time.Sleep(10 * time.Second)
	}

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

//Findprob finds the problematic node
func Findprob(node string, message string) (nodename string) {

	b := strings.Contains(message, "mount")
	if b == true {
		if node == "" {
			fmt.Println("*")
		}
		return node

	}
	return
}


func Workonprob(node string ) (ok string ) {

 // we will use the identity from podidentity it will be easier 

 // do stuff with this node -like evict this node and reimage it 
 // then watch this node to become ready

 //return node is ready


}
