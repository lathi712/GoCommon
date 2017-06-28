package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

/*Function to invoke Environments depending on the User Input
  lle - low level environment,hle - high level environment and prod - production*/
func kubecom(automataEnv string) {
	switch automataEnv {
	case "lle":
		lle()
	case "hle":
		hle()
	case "prod":
		prod()
	default:
		break
	}

}

/*
This function will execute couple of Unix and Kubernetes commands to Automate things.
We are automating the creation of Deployments,ConfigMaps and Secrets using kubernetes commands.

*/
func commandexec(env string) {

	var SEDREP = "s/ENV_NAME/" + env + "/g"
	var ZONE = "us-central1-a"
	var CLUSTER = "mean-cluster"
	var KUBERNETES_PATH = "./kubernetes"
	var SECRETS = "/secret.json"
	var DEPLOYS = "/deployment.yaml"
	var DEP_PATH = KUBERNETES_PATH + DEPLOYS
	var SECRETS_PATH = KUBERNETES_PATH + SECRETS

	gcloudComputeSet := exec.Command("gcloud", "config", "set", "compute/zone", ZONE)
	printCmd(gcloudComputeSet)

	gcloudContainerSet := exec.Command("gcloud", "container", "clusters", "get-credentials", CLUSTER)
	printCmd(gcloudContainerSet)

	createNamespace := exec.Command("kubectl", "create", "namespace", "ingress-"+env)
	printCmd(createNamespace)

	deploymentReplace := exec.Command("sed", "-i", SEDREP, DEP_PATH)
	printCmd(deploymentReplace)

	secretReplace := exec.Command("sed", "-i", SEDREP, SECRETS_PATH)
	printCmd(secretReplace)

	createConfigMap := exec.Command("kubectl", "create", "configmap", "sites-config-"+env, "--from-file=./kubernetes/configmap", "--namespace", "ingress-"+env)
	printCmd(createConfigMap)

	createSecret := exec.Command("kubectl", "create", "-f", SECRETS_PATH, "--namespace", "ingress-"+env)
	printCmd(createSecret)

	createdeploy := exec.Command("kubectl", "create", "-f", DEP_PATH)
	printCmd(createdeploy)

	revertsecret := exec.Command("cp", "secret.json", KUBERNETES_PATH)
	printCmd(revertsecret)

	revertdeploy := exec.Command("cp", "deployment.yaml", KUBERNETES_PATH)
	printCmd(revertdeploy)
}

//Iterating HLE Environments
func hle() {
	hle := [3]string{"relqa", "perf", "uat"}
	for j := 0; j < len(hle); j++ {
		commandexec(hle[j])
	}
}

//Iterating LLE Environments
func lle() {
	lle := [3]string{"dev", "qa", "reldev"}
	for j := 0; j < len(lle); j++ {
		commandexec(lle[j])
	}
}

//Iterating PROD Environments
func prod() {
	prod := [2]string{"preprod", "prod"}
	for j := 0; j < len(prod); j++ {
		commandexec(prod[j])
	}
}

//Here we are printing the Standard Output and Standard Error
func printCmd(cmd *exec.Cmd) {
	//cmd := exec.Command("go", "version")

	cmdOutput := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Stderr = stdErr
	err := cmd.Run()
	if err != nil {
		//os.Stderr.WriteString(err.Error())
		fmt.Println(fmt.Sprint(err) + ": " + stdErr.String())
		return
	}
	fmt.Print(string(cmdOutput.Bytes()))
}

/*Start of the program - Entry Point*/
func main() {
	//arg:= os.Args[1]
	var automataEnv = "hle"
	kubecom(automataEnv)
}
